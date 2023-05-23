package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"mosaic-go-interview/src/cache"
	"mosaic-go-interview/src/commands"
	apperrors "mosaic-go-interview/src/errors"
	"mosaic-go-interview/src/response"
	"mosaic-go-interview/src/services"
)

var operationMap = map[string]func(commands.BasicArythemticCommand) (int, error){
	"add":      services.Add,
	"subtract": services.Subtract,
	"multiply": services.Multiply,
	"divide":   services.Divide,
}

func BasicArythmeticHandler(cacheService cache.CacheService, w http.ResponseWriter, r *http.Request) {

	command, err := getCommand(r.URL.Query(), r.URL)
	if err != nil {
		sendErrorResponse(w, apperrors.InvalidOperandsError{}, http.StatusBadRequest)
		return
	}

	value, hit, err := cacheService.Get(command.GetCacheKey())
	if err != nil {
		sendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	if hit {
		response := response.ArthmeticResponse{
			Action: command.Action,
			X:      command.X,
			Y:      command.Y,
			Answer: value,
			Cached: true,
		}
		sendOkResponse(w, response)
		return
	}

	service, ok := operationMap[command.Action]
	if !ok {
		sendErrorResponse(w, apperrors.InvalidOperationError{}, http.StatusBadRequest)
		return
	}

	result, err := service(command)
	if err != nil {
		sendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	cacheService.Set(command.GetCacheKey(), result, 60*time.Second)

	response := response.ArthmeticResponse{
		Action: command.Action,
		X:      command.X,
		Y:      command.Y,
		Answer: result,
		Cached: false,
	}

	sendOkResponse(w, response)
}

func getCommand(params url.Values, url *url.URL) (commands.BasicArythemticCommand, error) {
	command := commands.BasicArythemticCommand{}
	x, err := strconv.Atoi(params.Get("x"))
	if err != nil {
		return command, err
	}

	y, err := strconv.Atoi(params.Get("y"))
	if err != nil {
		return command, err
	}

	command.X = x
	command.Y = y
	command.Action = url.Path[1:]

	return command, nil
}

func sendOkResponse(w http.ResponseWriter, response response.ArthmeticResponse) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func sendErrorResponse(w http.ResponseWriter, err error, httpsStatus int) {

	userMessage := "Internal server error"
	if userFriendlyErr, ok := err.(apperrors.UserFriendlyError); ok {
		userMessage = userFriendlyErr.UserMessage()
	}

	errResponse := response.ErrorResponse{
		Error: userMessage,
	}
	jsonResponse, _ := json.Marshal(errResponse)
	http.Error(w, string(jsonResponse), httpsStatus)
}

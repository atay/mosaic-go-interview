package response

type ArthmeticResponse struct {
	Action string `json:"action"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Answer int    `json:"answer"`
	Cached bool   `json:"cached"`
}

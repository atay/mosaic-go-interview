FROM golang:1.20-alpine3.18 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY .  /app
RUN cd src && go test ./...
RUN CGO_ENABLED=0 GOOS=linux cd src && go build -a -installsuffix cgo -o app .


FROM alpine:3.18.0

RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app/src/app .

CMD ["./app"]
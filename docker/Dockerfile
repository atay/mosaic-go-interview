FROM golang:1.20-alpine3.18 AS build

WORKDIR /app
COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/  ./
RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:3.18.0

RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app/app .

CMD ["./app"]
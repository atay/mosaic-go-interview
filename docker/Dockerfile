FROM golang:1.20.5-bullseye AS build

WORKDIR /app
COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/  ./
RUN go test ./...
RUN CGO_ENABLED=0 go build -o app .


FROM scratch

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
COPY --from=build /app/app /app/app

CMD ["/app/app"]
up:
	docker-compose up -d

down:
	docker-compose down

run-local:
	docker-compose up redis -d
	cd src && HTTP_PORT=8080 REDIS_URL="localhost:6379" go run main.go

test:
	cd src && go test ./...
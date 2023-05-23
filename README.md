# Mosaic Group Interview Task

That's an interview task prepared for a recruitment in Mosiac Group for a senior Developer role.

# How to run it?

## Via docker

1. Ensure Docker and Docker Compose are installed in your system, if not - install those.
2. Type `docker compose up` in your console. The project will build and you should see `Server listening on http://localhost:8080/`
3. You can use browser or curl to use the project, f.e.:
```bash
> curl "localhost:8080/add?x=5&y=7"
{"action":"add","x":5,"y":7,"answer":12,"cached":false}
```

## Via docker in the background

1. Run `make up`


## Via localhost

1. Execute `make run-local`

# How to execute tests?

1. Run `make test`

# Notes

1. The code is divided to handlers and services, so for each endpoint there's handler that checks the request, build the command that could be used by the service and sends it to proper service. This way we could provide a separation between infrastructure/app and domain layers of the app.
2. I decided to use redis as cache system, becuse of TTL requirement, so I decided it will be easier. But I also created InMemoryCache in application, as I required it for handler tests, but this one is simpler and does not support TTL. The app uses `CacheService` interface, so it could be easily switched between those two.
3. The app uses Dependency Injection, so `CacheService` is injected into handler.
4. Port number and redis address are configurable via environment variables in `docker-compose.yml` file.
5. There is a `Makefile` that contains most common commands, so it's easy to run the app, tests, etc.
6. I'm using UserFriendlyError to return errors to the user, so this way I know this is the error that should be returned to the user and not some internal error that should be never returned to the output.

# Docker notes

1. There's a `docker-compose.yml` file, so the app can be easily run in docker with one command.
2. The `Dockerfile` is multi-stage, so it builds the app in one container and then copies the binary to another container, so the final image is much smaller.
3. The `Dockerfile` uses `scratch` as base image, so it's really small and contains only the binary and nothing else.
4. The `Dockerfile` uses `CGO_ENABLED=0` to build the app, so it's static and does not require any libraries to run.
5. The `Dockerfile` uses `go test -v ./...` to run tests, so it's easy to run tests in docker.
6. Version of images in `Dockerfile` is not set to `latest`, this way you may be sure that all developers use the same version of images.
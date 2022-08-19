<pre>
██████╗  ██████╗     ████████╗███████╗███╗   ███╗██████╗ ██╗      █████╗ ████████╗███████╗
██╔════╝ ██╔═══██╗    ╚══██╔══╝██╔════╝████╗ ████║██╔══██╗██║     ██╔══██╗╚══██╔══╝██╔════╝
██║  ███╗██║   ██║       ██║   █████╗  ██╔████╔██║██████╔╝██║     ███████║   ██║   █████╗
██║   ██║██║   ██║       ██║   ██╔══╝  ██║╚██╔╝██║██╔═══╝ ██║     ██╔══██║   ██║   ██╔══╝
╚██████╔╝╚██████╔╝       ██║   ███████╗██║ ╚═╝ ██║██║     ███████╗██║  ██║   ██║   ███████╗
╚═════╝  ╚═════╝        ╚═╝   ╚══════╝╚═╝     ╚═╝╚═╝     ╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝
</pre>

<p align="center">
  <img src="resources/assets/gophere_logo.png" width="125">
</p>

# Go Template - Hexa(go)nal architecture based on DDD

Go-template is an opinionated Hexagonal backend template written in GO.

## Features

- Provides database migration with [pressly/gooose](https://github.com/pressly/goose).
- Integrates [go-swagger](https://github.com/go-swagger/go-swagger) for compile-time generation of swagger.(yml|json).
- Integrates [SwaggerUI](https://github.com/swagger-api/swagger-ui) for live-previewing Swagger v2 schema.
- Comes with an initial MariaDB database structure (
  see [migrations](https://github.com/NicklasWallgren/go-template/tree/main/resources/database/migrations))
- API endpoints for health and readiness probes.
- K8s manifest files.
- Provides support for parallel database integration tests.
- Comes with support for AMQP publisher and consumer (via RabbitMQ).
- Various convenience scripts via `Makefile`.
- Multi-staged Dockerfile with support for BuildX for faster builds.
- CLI Layer which is easily extensible (and provided by spf13/cobra)
    - Command to handle migrations (migration).
    - Command to serve the http server.
    - Command to launch amqp consumer.
- Support for observability via [dd-trace-go](https://github.com/DataDog/dd-trace-go/)

## Prerequisites

Requires the following prerequisites for development in a local environment

- Docker CE (v20.10+)
- Docker Compose (v1.27+)
- Go 1.18

## Quickstart
Initialize a new git repository through [GitHub Template feature](https://github.com/allaboutapps/go-starter/generate)

```bash 
# Clone the repository and then start the dev environment through the Makefile
make
```

## Merge with the repository to get future updates
```bash
git merge --no-commit --no-ff --allow-unrelated-histories NicklasWallgren/go-template
```

## Set the custom project name
```bash
find . -not -path '*/\.*' -type f -exec sed -i "" "s|github.com/NicklasWallgren/go-template|<REPLACE>|g" {} \;
```

## CLI
```bash
# Start the HTTP server
go-template start
# Start the AMQP consumers
go-template start-consumers
# Create a new migration file
go-template migrate --create <NAME>
# Rollup migrations
go-template migrate --up
```

## Contributing

## License

## TODO

mockery --all --output ./tests/mocks --keeptree --case underscore --with-expecter

go fmt $(go list ./... | grep -v /test/mocks/)

golangci-lint run --fix
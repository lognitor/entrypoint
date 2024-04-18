gen.proto:
	protoc --go_out=pkg/transport --go-grpc_out=pkg/transport proto/*.proto

# IMPORTANT: COMMAND INCLUDE PUSH TO GITHUB
# Build multiarch image and push to github
# > make docker.build version=v1.0.0
docker.build:
	docker buildx build --platform=linux/amd64,linux/arm64 --push -t ghcr.io/lognitor/entrypoint:${version} -t ghcr.io/lognitor/entrypoint:latest .

# Migration create from name
# > make migrate.create name=create_table
migrate.create:
	docker run --rm -u $(shell id -u):$(shell id -g) -v ./migrations:/migrations \
		migrate/migrate create -ext sql -dir /migrations ${name}
gen.proto:
	protoc --go_out=. --go-grpc_out=. proto/*.proto

# IMPORTANT: COMMAND INCLUDE PUSH TO GITHUB
# Build multiarch image and push to github
# make docker.build version=v1.0.0
docker.build:
	docker buildx build --platform=linux/amd64,linux/arm64 --push -t ghcr.io/lognitor/entrypoint:${version} -t ghcr.io/lognitor/entrypoint:latest .
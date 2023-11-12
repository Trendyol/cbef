format:
	@gofumpt -l -w -extra .
dockerize:
	@docker build -f build/docker/Dockerfile .

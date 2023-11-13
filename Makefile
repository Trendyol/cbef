lint:
	@golangci-lint run ./...
format:
	@gofumpt -l -w -extra .
dockerize:
	@docker build -f build/docker/Dockerfile .
setup-hook:
	@chmod 755 githooks/pre-commit
	@git config --local include.path ../.gitconfig

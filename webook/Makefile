.PHONY: docker mock
docker:
	@rm -f webook || true
	@go mod tidy
	@GOOS=linux GOARCH=arm go build -tags=k8s -o main .
	@docker rmi -f sgy111222333/webook:v0.0.1
	@docker build -t sgy111222333/webook:v0.0.1 .

mock:
	# mock UserService这个interface
	@mockgen -source=./internal/service/user.go -package=svcmocks -destination=./internal/service/mocks/user.mock.go
	# mock CodeService这个interface
	@mockgen -source=./internal/service/code.go -package=svcmocks -destination=./internal/service/mocks/code.mock.go
	# mock UserRepository这个interface
	@mockgen -source=./internal/repository/user.go -package=repomocks -destination=./internal/repository/mocks/user.mock.go
	# mock CodeRepository这个interface
	@mockgen -source=./internal/repository/code.go -package=repomocks -destination=./internal/repository/mocks/code.mock.go

	@go mod tidy

.PHONY: docker
docker:
	@rm -f webook || true
	@go mod tidy
	@GOOS=linux GOARCH=arm go build -tags=k8s -o main .
	@docker rmi -f sgy111222333/webook:v0.0.1
	@docker build -t sgy111222333/webook:v0.0.1 .
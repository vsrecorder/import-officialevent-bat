.PHONY: batch
batch:
	go mod tidy
	go run cmd/batch/main.go


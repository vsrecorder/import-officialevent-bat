.PHONY: batch
batch:
	go mod download && go mod tidy
	go run cmd/batch/main.go

.PHONY: build
build:
	go mod download && go mod tidy
	go build -o bin/batch cmd/batch/main.go
	go build -o bin/daily_batch cmd/daily_batch/main.go

.PHONY: clean
clean:
	rm -f bin/batch
	rm -f bin/daily_batch

.PHONY: deploy
deploy: build
	sudo cp systemd/* /etc/systemd/system/
	sudo systemctl daemon-reload
	sudo systemctl enable --now import-officialevent-bat.timer
	sudo systemctl restart import-officialevent-bat.timer
#!/bin/bash

source /home/ubuntu/import-officialevent-bat/.env

curl -X POST -H 'Content-type: application/json' --data '{"text":"日次バッチを開始します。"}' ${SLACK_WEBHOOK_URL}
cd /home/ubuntu/import-officialevent-bat
go run cmd/daily_batch/main.go
curl -X POST -H 'Content-type: application/json' --data '{"text":"日次バッチが終了しました。"}' ${SLACK_WEBHOOK_URL}

exit 0

#!/bin/bash

source .env

curl -X POST -H 'Content-type: application/json' --data '{"text":"日次バッチを開始します。"}' ${SLACK_WEBHOOK_URL}

go run cmd/daily_batch/main.go 2> /dev/null

curl -X POST -H 'Content-type: application/json' --data '{"text":"日次バッチが終了しました。"}' ${SLACK_WEBHOOK_URL}

exit 0

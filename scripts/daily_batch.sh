#!/bin/bash

source /home/ubuntu/vsrecorder/import-officialevent-bat/.env

curl -X POST -H 'Content-type: application/json' --data '{"text":"日次バッチを開始します。"}' ${SLACK_WEBHOOK_URL}

./bin/daily_batch

curl -X POST -H 'Content-type: application/json' --data '{"text":"日次バッチが終了しました。"}' ${SLACK_WEBHOOK_URL}

exit 0
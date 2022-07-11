# SQS + Lambda demo
## Setup Lambda by SQS hook
*Need to configure AWS profile to execute `aws` command with correct profile
```
$ cat .envrc
export AWS_PROFILE=gosample
```

Create executable and create a function in lambda
(When you need other files on runtime, you can add more files/directories in zip)
```
$ make lmd-sassample
```

Confirm if the function is registered
```
$ aws lambda list-functions --endpoint-url http://localhost:4566
$ aws lambda invoke \
  --endpoint-url http://localhost:4566 \
  --function-name sqssample \
  --cli-binary-format raw-in-base64-out \
  --payload '{"Records": [{"messageId": "id1"}]}' \
  response.json
```

Register SQS hook to the function
```
$ aws lambda create-event-source-mapping \
  --endpoint-url=http://localhost:4566 \
  --function-name sqssample \
  --event-source-arn arn:aws:sqs:ap-northeast-1:000000000000:sqssample
```

Confirm if the hook is registered
```
$ aws lambda list-event-source-mappings --endpoint-url http://localhost:4566
$ aws sqs send-message \
  --queue-url http://localhost:4566/queue/sqssample \
  --message-body '{"name":"hello"}' \
  --endpoint-url=http://localhost:4566
```

## Setup S3
```
$ aws s3api list-buckets --endpoint-url=http://localhost:4566
$ aws s3api list-objects --bucket s3sample --endpoint-url=http://localhost:4566
```

# Debezium demo
Start
```
docker-compose --profile cdc --profile debezium up -d
```

## Register a connector
```
curl --location --request POST 'http://localhost:8083/connectors' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "name": "gosample2022-connector",
    "config": {
        "connector.class": "io.debezium.connector.mysql.MySqlConnector",
        "topic.creation.default.replication.factor": 1,
        "topic.creation.default.partitions": 2,
        "database.hostname": "mysql",
        "database.port": "3306",
        "database.user": "root",
        "database.password": "root",
        "database.include.list": "gosample2022_development",
        "table.include.list": "gosample2022_development.samples",
        "message.key.columns": "gosample2022_development.samples:office_id",
        "database.server.name": "gosample2022_dbserver",
        "database.history.kafka.bootstrap.servers": "kafka:29092",
        "database.history.kafka.topic": "schema-changes.gosample2022",
        "include.schema.changes": false,
        "provide.transaction.metadata": true
    }
}'
```

## Confirm the configuration
```
curl -H "Accept:application/json" http://localhost:8083/connectors/gosample2022-connector | jq
```

## Monitor the topic
```
kcat -b localhost:9092 -t dbserver1.gosample2022_development.samples
```

## Delete connector
```
curl -i -X DELETE localhost:8083/connectors/gosample2022-connector
```

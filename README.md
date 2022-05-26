# gosample 2022
## Setup Lambda by SQS hook
*Need to configure AWS profile to execute `aws` command with correct profile
```
$ cat .envrc
export AWS_PROFILE=gosample
```

Create executable and create a function in lambda
```
$ GOOS=linux CGO_ENABLED=0 go build ./cmd/sqssample/main.go && zip main.zip main
```

Register or Update function in lambda
```
$ aws lambda create-function \
  --endpoint-url=http://localhost:4566 \
  --function-name sqssample \
  --handler main \
  --runtime go1.x \
  --zip-file fileb://main.zip \
  --role sqssample
$ aws lambda update-function-code \
  --endpoint-url http://localhost:4566 \
  --function-name sqssample \
  --zip-file fileb://main.zip \
  --publish
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

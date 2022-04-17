#!/bin/bash
AWS_ACCESS_KEY_ID=gosample2022
AWS_SECRET_ACCESS_KEY=gosample2022
REGION=ap-northeast-1

# Configuration
aws configure set aws_access_key_id $AWS_ACCESS_KEY_ID
aws configure set aws_secret_access_key $AWS_SECRET_ACCESS_KEY
aws configure set region $REGION

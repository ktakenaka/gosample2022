#!/bin/bash
apt update -y && apt install jq -y

cd /data/localstack

# https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-create-cmk.html
KeyId=`jq -r '.KeyMetadata.KeyId' <<< $(awslocal kms create-key --origin EXTERNAL)`

# https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-get-public-key-and-token.html
awslocal kms get-parameters-for-import --key-id ${KeyId} --wrapping-algorithm RSAES_OAEP_SHA_1 --wrapping-key-spec RSA_2048 > get-parameters-for-import-output.json
openssl enc -d -base64 -out PublicKey.bin <<< $(jq -r '.PublicKey' get-parameters-for-import-output.json)
openssl enc -d -base64 -out ImportToken.bin <<< $(jq -r '.ImportToken' get-parameters-for-import-output.json)

# https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-encrypt-key-material.html
openssl enc -base64 -d -out PlaintextKeyMaterial.bin <<< 7qajDsDXhmA+9pUKrQvrVczPMM51WN0HIWZbUVVGQns=
openssl rsautl -encrypt -in PlaintextKeyMaterial.bin -oaep -inkey PublicKey.bin -keyform DER -pubin -out EncryptedKeyMaterial.bin

# https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-import-key-material.html
awslocal kms import-key-material --key-id ${KeyId} --encrypted-key-material fileb://EncryptedKeyMaterial.bin --import-token fileb://ImportToken.bin --expiration-model KEY_MATERIAL_DOES_NOT_EXPIRE

# Make key alias
awslocal kms create-alias --target-key-id ${KeyId} --alias-name 'alias/local-kms-key'

# You can confirm if the alias is made by this command
# awslocal kms describe-key --key-id 'alias/local-kms-key'

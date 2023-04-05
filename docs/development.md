# Getting started 
## Package the Go executable for deployment
Since AWS Lambdas use Linux.

Run the following:

### Package the Go executable for deployment
```shell
GOOS=linux GOARCH=amd64 go build -o main
```

### Zip the executable
```shell
zip deployment.zip main statuses.json
```

### Deploy the Lambda function
This assumes you have the AWS CLI installed and configured.

```shell
aws lambda create-function \
    --function-name UpdateGitHubStatus \
    --runtime go1.x \
    --handler main \
    --zip-file fileb://./deployment.zip \
    --role arn:aws:iam::YOUR_ACCOUNT_ID:role/UpdateGitHubStatusRole
```
replace `YOUR_ACCOUNT_ID` with your AWS account ID.

### Update the function with your GitHub token
```shell
aws lambda update-function-configuration \
    --function-name UpdateGitHubStatus \
    --environment "Variables={GITHUB_TOKEN
=YOUR_GITHUB_TOKEN}"
```
replace `YOUR_GITHUB_TOKEN` with your GitHub token.

### Create a CloudWatch event to trigger the Lambda function
```shell
aws events put-rule \
   --name UpdateGitHubStatusRule \
   --schedule-expression 'cron(0 5 * * ? *)'
```
By default this will run the Lambda function every day at 5am.

```shell
aws events put-targets \
    --rule UpdateGitHubStatusRule \
    --targets "Id"="1","Arn"="arn:aws:lambda:YOUR_REGION:YOUR_ACCOUNT_ID:function:UpdateGitHubStatus"


aws lambda add-permission \
    --function-name UpdateGitHubStatus \
    --statement-id UpdateGitHubStatusRule \
    --action 'lambda:InvokeFunction' \
    --principal events.amazonaws.com \
    --source-arn arn:aws:events:YOUR_REGION:YOUR_ACCOUNT_ID:rule/UpdateGitHubStatusRule
```
replace `YOUR_REGION` and `YOUR_ACCOUNT_ID` with your AWS region and account ID.

### Useful functions you can run 
Update the function with new executable code:
```shell
aws lambda update-function-code --function-name UpdateGitHubStatus --zip-file fileb://deployment.zip
```

Delete the function:
```shell
aws lambda delete-function \
    --function-name UpdateGitHubStatus 
```

## The status.json file
The `statuses.json` file is a JSON file that contains a list of statuses. Each status is a JSON object with the following fields:
- `emoji`: The emoji to use for the status. This can be any GitHub supported emoji.
- `message`: The message to use for the status.
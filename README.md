# github-status-lambda
A lambda function that updates my GitHub status with a random quote and emoji.

## Why did I make this?
I wanted to learn how to make a lambda function and I thought it would be fun to have a random quote and emoji show up as my GitHub status. Also I tend to not really use GitHub statuses all that much so why not make it fun?

## Package the Go executable for deployment
Since AWS Lambdas use Linux.

Run the following:
```shell
GOOS=linux go build -o main
```

## Zip the executable
```shell
zip deployment.zip main
```

## Deploy the Lambda function
```shell
aws lambda create-function --function-name UpdateMyGithubStatus --runtime go1.19 --handler main --role arn:aws:iam::123456789012:role/YourRole --zip-file fileb://deployment.zip
```


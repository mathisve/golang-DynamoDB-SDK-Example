# golang-DynamoDB-SDK-Example
Short example of the 5 most common operations on the DynamoDB SDK for Golang.
 - [CreateTable](https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.CreateTable)
 - [PutItem](https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.PutItem)
 - [UpdateItem](https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.UpdateItem)
 - [DeleteItem](https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.DeleteItem)
 - [GetItem](https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.GetItem)

Video in where I go over them one by one:
[![image](https://github.com/mathisve/golang-DynamoDB-SDK-Example/blob/master/img/dynamodb.png)](https://youtu.be/pr4x8KdIfDU)

## Quite a bit has changed about the code!
I noticed the code I wrote in the video was **perfect for explaining DynamoDB concepts**, but not so much for using it in a production environment. 
This is why I have slightly updated the code, but most of the pure DynamoDB related functions should be roughly the same.
When in doubt, feel free ask questions!

Furthermore, I highly recommend using DynamoDB Attributes instead of writing everything explicitly.
It will make your code a lot more portable and clean!
Fortunately for you, I made a video about this. Find it [here](https://youtu.be/dLf9A9YQ97Y). Or go directly to the [Github repository](https://github.com/mathisve/golang-DynamoDBAttributes-SDK-Example)!

**Thank you** for understanding!

## Read more documentation
Find the official DynamoDB docs for Golang [here](https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/)!

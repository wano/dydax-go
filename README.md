# dydax-go

client merged DynamoDB and DAX.

if DAX non supported API called, invoke DynamoDB Client.

if DAX supported API called, invoke DAX Client.

## usage

``` go
import "github.com/wano/dydax-go"
...
client := dydax.New(dynamodbClient, daxClient)
```

package dydax

import (
	"github.com/aws/aws-dax-go/dax"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"io"
)

type DynamoDBCloser interface {
	dynamodbiface.DynamoDBAPI
	io.Closer
}

type daxWrapper struct {
	*dax.Dax
}

func newDaxWrapper(d *dax.Dax) *daxWrapper {
	return &daxWrapper{d}
}

var _ dynamodbiface.DynamoDBAPI = (*daxWrapper)(nil)


/**
daxとdynamodbのクライアントをガッチャンコしたラッパークライアント
daxが対応してないapiのみ、dynamodbクライアントの方で実行する
*/
type Client struct {
	dynamodbiface.DynamoDBAPI
	daxClient DynamoDBCloser
}

func New(dynamoClient *dynamodb.DynamoDB, daxClient *dax.Dax) *Client {
	return &Client{
		DynamoDBAPI: dynamoClient,
		daxClient:   newDaxWrapper(daxClient),
	}
}

func (d *Client) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return d.daxClient.PutItem(input)
}

func (d *Client) PutItemWithContext(ctx aws.Context, input *dynamodb.PutItemInput, opts ...request.Option) (*dynamodb.PutItemOutput, error) {
	return d.daxClient.PutItemWithContext(ctx, input, opts...)
}

func (d *Client) PutItemRequest(input *dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput) {
	return d.daxClient.PutItemRequest(input)
}

func (d *Client) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	return d.daxClient.DeleteItem(input)
}

func (d *Client) DeleteItemWithContext(ctx aws.Context, input *dynamodb.DeleteItemInput, opts ...request.Option) (*dynamodb.DeleteItemOutput, error) {
	return d.daxClient.DeleteItemWithContext(ctx, input, opts...)
}

func (d *Client) DeleteItemRequest(input *dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput) {
	return d.daxClient.DeleteItemRequest(input)
}

func (d *Client) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	return d.daxClient.UpdateItem(input)
}

func (d *Client) UpdateItemWithContext(ctx aws.Context, input *dynamodb.UpdateItemInput, opts ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	return d.daxClient.UpdateItemWithContext(ctx, input, opts...)
}

func (d *Client) UpdateItemRequest(input *dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput) {
	return d.daxClient.UpdateItemRequest(input)
}

func (d *Client) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return d.daxClient.GetItem(input)
}

func (d *Client) GetItemWithContext(ctx aws.Context, input *dynamodb.GetItemInput, opts ...request.Option) (*dynamodb.GetItemOutput, error) {
	return d.daxClient.GetItemWithContext(ctx, input, opts...)
}

func (d *Client) GetItemRequest(input *dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput) {
	return d.daxClient.GetItemRequest(input)
}

func (d *Client) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	return d.daxClient.Scan(input)
}

func (d *Client) ScanWithContext(ctx aws.Context, input *dynamodb.ScanInput, opts ...request.Option) (*dynamodb.ScanOutput, error) {
	return d.daxClient.ScanWithContext(ctx, input, opts...)
}

func (d *Client) ScanRequest(input *dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput) {
	return d.daxClient.ScanRequest(input)
}

func (d *Client) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return d.daxClient.Query(input)
}

func (d *Client) QueryWithContext(ctx aws.Context, input *dynamodb.QueryInput, opts ...request.Option) (*dynamodb.QueryOutput, error) {
	return d.daxClient.QueryWithContext(ctx, input, opts...)
}

func (d *Client) QueryRequest(input *dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput) {
	return d.daxClient.QueryRequest(input)
}

func (d *Client) BatchWriteItem(input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	return d.daxClient.BatchWriteItem(input)
}

func (d *Client) BatchWriteItemWithContext(ctx aws.Context, input *dynamodb.BatchWriteItemInput, opts ...request.Option) (*dynamodb.BatchWriteItemOutput, error) {
	return d.daxClient.BatchWriteItemWithContext(ctx, input, opts...)
}

func (d *Client) BatchWriteItemRequest(input *dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput) {
	return d.daxClient.BatchWriteItemRequest(input)
}

func (d *Client) BatchGetItem(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	return d.daxClient.BatchGetItem(input)
}

func (d *Client) BatchGetItemWithContext(ctx aws.Context, input *dynamodb.BatchGetItemInput, opts ...request.Option) (*dynamodb.BatchGetItemOutput, error) {
	return d.daxClient.BatchGetItemWithContext(ctx, input, opts...)
}

func (d *Client) BatchGetItemRequest(input *dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput) {
	return d.daxClient.BatchGetItemRequest(input)
}

func (d *Client) TransactWriteItems(input *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
	return d.daxClient.TransactWriteItems(input)
}

func (d *Client) TransactWriteItemsWithContext(ctx aws.Context, input *dynamodb.TransactWriteItemsInput, opts ...request.Option) (*dynamodb.TransactWriteItemsOutput, error) {
	return d.daxClient.TransactWriteItemsWithContext(ctx, input, opts...)
}

func (d *Client) TransactWriteItemsRequest(input *dynamodb.TransactWriteItemsInput) (*request.Request, *dynamodb.TransactWriteItemsOutput) {
	return d.daxClient.TransactWriteItemsRequest(input)
}

func (d *Client) TransactGetItems(input *dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error) {
	return d.daxClient.TransactGetItems(input)
}

func (d *Client) TransactGetItemsWithContext(ctx aws.Context, input *dynamodb.TransactGetItemsInput, opts ...request.Option) (*dynamodb.TransactGetItemsOutput, error) {
	return d.daxClient.TransactGetItemsWithContext(ctx, input, opts...)
}

func (d *Client) TransactGetItemsRequest(input *dynamodb.TransactGetItemsInput) (*request.Request, *dynamodb.TransactGetItemsOutput) {
	return d.daxClient.TransactGetItemsRequest(input)
}

func (d *Client) Close() error {
	return d.daxClient.Close()
}

var _ dynamodbiface.DynamoDBAPI = (*Client)(nil)


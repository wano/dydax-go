package dydax

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"testing"
)

type mockDynamoDBAPI struct {
	dynamodbiface.DynamoDBAPI
}

type mockDynamoCloser struct {
	DynamoDBCloser
}

func (m *mockDynamoCloser) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) PutItemWithContext(ctx aws.Context, input *dynamodb.PutItemInput, opts ...request.Option) (*dynamodb.PutItemOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) PutItemRequest(input *dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput) {
	return nil, nil
}

func (m *mockDynamoCloser) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) DeleteItemWithContext(ctx aws.Context, input *dynamodb.DeleteItemInput, opts ...request.Option) (*dynamodb.DeleteItemOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) DeleteItemRequest(input *dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput) {
	return nil, nil
}

func (m *mockDynamoCloser) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) UpdateItemWithContext(ctx aws.Context, input *dynamodb.UpdateItemInput, opts ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) UpdateItemRequest(input *dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput) {
	return nil, nil
}

func (m *mockDynamoCloser) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) GetItemWithContext(ctx aws.Context, input *dynamodb.GetItemInput, opts ...request.Option) (*dynamodb.GetItemOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) GetItemRequest(input *dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput) {
	return nil, nil
}

func (m *mockDynamoCloser) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) ScanWithContext(ctx aws.Context, input *dynamodb.ScanInput, opts ...request.Option) (*dynamodb.ScanOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) ScanRequest(input *dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput) {
	return nil, nil
}

func (m *mockDynamoCloser) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) QueryWithContext(ctx aws.Context, input *dynamodb.QueryInput, opts ...request.Option) (*dynamodb.QueryOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) QueryRequest(input *dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput) {
	return nil, nil
}

func (m *mockDynamoCloser) BatchWriteItem(input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) BatchWriteItemWithContext(ctx aws.Context, input *dynamodb.BatchWriteItemInput, opts ...request.Option) (*dynamodb.BatchWriteItemOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) BatchWriteItemRequest(input *dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput) {
	return nil, nil
}

func (m *mockDynamoCloser) BatchGetItem(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) BatchGetItemWithContext(ctx aws.Context, input *dynamodb.BatchGetItemInput, opts ...request.Option) (*dynamodb.BatchGetItemOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) BatchGetItemRequest(input *dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput) {
	return nil, nil
}

func (m *mockDynamoCloser) TransactWriteItems(input *dynamodb.TransactWriteItemsInput) (*dynamodb.TransactWriteItemsOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) TransactWriteItemsWithContext(ctx aws.Context, input *dynamodb.TransactWriteItemsInput, opts ...request.Option) (*dynamodb.TransactWriteItemsOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) TransactWriteItemsRequest(input *dynamodb.TransactWriteItemsInput) (*request.Request, *dynamodb.TransactWriteItemsOutput) {
	return nil, nil
}

func (m *mockDynamoCloser) TransactGetItems(input *dynamodb.TransactGetItemsInput) (*dynamodb.TransactGetItemsOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) TransactGetItemsWithContext(ctx aws.Context, input *dynamodb.TransactGetItemsInput, opts ...request.Option) (*dynamodb.TransactGetItemsOutput, error) {
	return nil, nil
}

func (m *mockDynamoCloser) TransactGetItemsRequest(input *dynamodb.TransactGetItemsInput) (*request.Request, *dynamodb.TransactGetItemsOutput) {
	return nil, nil
}

func (m *mockDynamoCloser) Close() error {
	return nil
}

func TestClinet_DaxSupportedAPICall(t *testing.T) {
	m := &mockDynamoDBAPI{}
	m2 := &mockDynamoCloser{}
	client := &Client{daxClient: m2, DynamoDBAPI: m}

	tests := []struct {
		f func() (interface{}, interface{})
	}{
		{
			f: func() (interface{}, interface{}) {
				return client.PutItem(nil)
			},
		},
		{
			f: func() (interface{}, interface{}) {
				return client.PutItemRequest(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.PutItemWithContext(nil, nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.GetItem(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.GetItemRequest(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.GetItemWithContext(nil, nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.BatchGetItemRequest(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.BatchGetItem(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.BatchGetItemWithContext(nil, nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.BatchWriteItemRequest(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.BatchWriteItem(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.BatchWriteItemWithContext(nil, nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.DeleteItemRequest(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.DeleteItem(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.UpdateItemRequest(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.UpdateItem(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.UpdateItemWithContext(nil, nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.TransactGetItemsRequest(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.TransactGetItems(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.TransactGetItemsWithContext(nil, nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.TransactWriteItemsRequest(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.TransactWriteItems(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.TransactWriteItemsWithContext(nil, nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.Scan(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.ScanRequest(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.ScanWithContext(nil, nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.Query(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.QueryRequest(nil)
			},
		},
		{
			f: func() (i interface{}, i2 interface{}) {
				return client.QueryWithContext(nil, nil)
			},
		},
	}

	for _, test := range tests {
		test.f() // if mockDynamoDBAPI called, panic occur. should be called mockDynamoCloser
	}

}

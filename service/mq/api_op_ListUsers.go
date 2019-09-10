// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mq

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/ListUsersRequest
type ListUsersInput struct {
	_ struct{} `type:"structure"`

	// BrokerId is a required field
	BrokerId *string `location:"uri" locationName:"broker-id" type:"string" required:"true"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListUsersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListUsersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListUsersInput"}

	if s.BrokerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BrokerId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListUsersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BrokerId != nil {
		v := *s.BrokerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "broker-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/ListUsersResponse
type ListUsersOutput struct {
	_ struct{} `type:"structure"`

	BrokerId *string `locationName:"brokerId" type:"string"`

	MaxResults *int64 `locationName:"maxResults" min:"5" type:"integer"`

	NextToken *string `locationName:"nextToken" type:"string"`

	Users []UserSummary `locationName:"users" type:"list"`
}

// String returns the string representation
func (s ListUsersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListUsersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BrokerId != nil {
		v := *s.BrokerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "brokerId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Users != nil {
		v := s.Users

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "users", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListUsers = "ListUsers"

// ListUsersRequest returns a request value for making API operation for
// AmazonMQ.
//
// Returns a list of all ActiveMQ users.
//
//    // Example sending a request using ListUsersRequest.
//    req := client.ListUsersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/ListUsers
func (c *Client) ListUsersRequest(input *ListUsersInput) ListUsersRequest {
	op := &aws.Operation{
		Name:       opListUsers,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/brokers/{broker-id}/users",
	}

	if input == nil {
		input = &ListUsersInput{}
	}

	req := c.newRequest(op, input, &ListUsersOutput{})
	return ListUsersRequest{Request: req, Input: input, Copy: c.ListUsersRequest}
}

// ListUsersRequest is the request type for the
// ListUsers API operation.
type ListUsersRequest struct {
	*aws.Request
	Input *ListUsersInput
	Copy  func(*ListUsersInput) ListUsersRequest
}

// Send marshals and sends the ListUsers API request.
func (r ListUsersRequest) Send(ctx context.Context) (*ListUsersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListUsersResponse{
		ListUsersOutput: r.Request.Data.(*ListUsersOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListUsersResponse is the response type for the
// ListUsers API operation.
type ListUsersResponse struct {
	*ListUsersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListUsers request.
func (r *ListUsersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
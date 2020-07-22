// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sesv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to remove an email address from the suppression list for your account.
type DeleteSuppressedDestinationInput struct {
	_ struct{} `type:"structure"`

	// The suppressed email destination to remove from the account suppression list.
	//
	// EmailAddress is a required field
	EmailAddress *string `location:"uri" locationName:"EmailAddress" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSuppressedDestinationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSuppressedDestinationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSuppressedDestinationInput"}

	if s.EmailAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("EmailAddress"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSuppressedDestinationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.EmailAddress != nil {
		v := *s.EmailAddress

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "EmailAddress", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type DeleteSuppressedDestinationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteSuppressedDestinationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSuppressedDestinationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteSuppressedDestination = "DeleteSuppressedDestination"

// DeleteSuppressedDestinationRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Removes an email address from the suppression list for your account.
//
//    // Example sending a request using DeleteSuppressedDestinationRequest.
//    req := client.DeleteSuppressedDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sesv2-2019-09-27/DeleteSuppressedDestination
func (c *Client) DeleteSuppressedDestinationRequest(input *DeleteSuppressedDestinationInput) DeleteSuppressedDestinationRequest {
	op := &aws.Operation{
		Name:       opDeleteSuppressedDestination,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v2/email/suppression/addresses/{EmailAddress}",
	}

	if input == nil {
		input = &DeleteSuppressedDestinationInput{}
	}

	req := c.newRequest(op, input, &DeleteSuppressedDestinationOutput{})

	return DeleteSuppressedDestinationRequest{Request: req, Input: input, Copy: c.DeleteSuppressedDestinationRequest}
}

// DeleteSuppressedDestinationRequest is the request type for the
// DeleteSuppressedDestination API operation.
type DeleteSuppressedDestinationRequest struct {
	*aws.Request
	Input *DeleteSuppressedDestinationInput
	Copy  func(*DeleteSuppressedDestinationInput) DeleteSuppressedDestinationRequest
}

// Send marshals and sends the DeleteSuppressedDestination API request.
func (r DeleteSuppressedDestinationRequest) Send(ctx context.Context) (*DeleteSuppressedDestinationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSuppressedDestinationResponse{
		DeleteSuppressedDestinationOutput: r.Request.Data.(*DeleteSuppressedDestinationOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSuppressedDestinationResponse is the response type for the
// DeleteSuppressedDestination API operation.
type DeleteSuppressedDestinationResponse struct {
	*DeleteSuppressedDestinationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSuppressedDestination request.
func (r *DeleteSuppressedDestinationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
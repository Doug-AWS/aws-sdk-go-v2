// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the input of a DeregisterOnPremisesInstance operation.
type DeregisterOnPremisesInstanceInput struct {
	_ struct{} `type:"structure"`

	// The name of the on-premises instance to deregister.
	//
	// InstanceName is a required field
	InstanceName *string `locationName:"instanceName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeregisterOnPremisesInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterOnPremisesInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterOnPremisesInstanceInput"}

	if s.InstanceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeregisterOnPremisesInstanceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeregisterOnPremisesInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeregisterOnPremisesInstance = "DeregisterOnPremisesInstance"

// DeregisterOnPremisesInstanceRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Deregisters an on-premises instance.
//
//    // Example sending a request using DeregisterOnPremisesInstanceRequest.
//    req := client.DeregisterOnPremisesInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/DeregisterOnPremisesInstance
func (c *Client) DeregisterOnPremisesInstanceRequest(input *DeregisterOnPremisesInstanceInput) DeregisterOnPremisesInstanceRequest {
	op := &aws.Operation{
		Name:       opDeregisterOnPremisesInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterOnPremisesInstanceInput{}
	}

	req := c.newRequest(op, input, &DeregisterOnPremisesInstanceOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeregisterOnPremisesInstanceRequest{Request: req, Input: input, Copy: c.DeregisterOnPremisesInstanceRequest}
}

// DeregisterOnPremisesInstanceRequest is the request type for the
// DeregisterOnPremisesInstance API operation.
type DeregisterOnPremisesInstanceRequest struct {
	*aws.Request
	Input *DeregisterOnPremisesInstanceInput
	Copy  func(*DeregisterOnPremisesInstanceInput) DeregisterOnPremisesInstanceRequest
}

// Send marshals and sends the DeregisterOnPremisesInstance API request.
func (r DeregisterOnPremisesInstanceRequest) Send(ctx context.Context) (*DeregisterOnPremisesInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterOnPremisesInstanceResponse{
		DeregisterOnPremisesInstanceOutput: r.Request.Data.(*DeregisterOnPremisesInstanceOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterOnPremisesInstanceResponse is the response type for the
// DeregisterOnPremisesInstance API operation.
type DeregisterOnPremisesInstanceResponse struct {
	*DeregisterOnPremisesInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterOnPremisesInstance request.
func (r *DeregisterOnPremisesInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
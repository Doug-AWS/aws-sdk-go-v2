// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListVPCEConfigurationsInput struct {
	_ struct{} `type:"structure"`

	// An integer specifying the maximum number of items you want to return in the
	// API response.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`
}

// String returns the string representation
func (s ListVPCEConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListVPCEConfigurationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListVPCEConfigurationsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListVPCEConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`

	// An array of VPCEConfiguration objects containing information about your VPC
	// endpoint configuration.
	VpceConfigurations []VPCEConfiguration `locationName:"vpceConfigurations" type:"list"`
}

// String returns the string representation
func (s ListVPCEConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListVPCEConfigurations = "ListVPCEConfigurations"

// ListVPCEConfigurationsRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Returns information about all Amazon Virtual Private Cloud (VPC) endpoint
// configurations in the AWS account.
//
//    // Example sending a request using ListVPCEConfigurationsRequest.
//    req := client.ListVPCEConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListVPCEConfigurations
func (c *Client) ListVPCEConfigurationsRequest(input *ListVPCEConfigurationsInput) ListVPCEConfigurationsRequest {
	op := &aws.Operation{
		Name:       opListVPCEConfigurations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListVPCEConfigurationsInput{}
	}

	req := c.newRequest(op, input, &ListVPCEConfigurationsOutput{})
	return ListVPCEConfigurationsRequest{Request: req, Input: input, Copy: c.ListVPCEConfigurationsRequest}
}

// ListVPCEConfigurationsRequest is the request type for the
// ListVPCEConfigurations API operation.
type ListVPCEConfigurationsRequest struct {
	*aws.Request
	Input *ListVPCEConfigurationsInput
	Copy  func(*ListVPCEConfigurationsInput) ListVPCEConfigurationsRequest
}

// Send marshals and sends the ListVPCEConfigurations API request.
func (r ListVPCEConfigurationsRequest) Send(ctx context.Context) (*ListVPCEConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListVPCEConfigurationsResponse{
		ListVPCEConfigurationsOutput: r.Request.Data.(*ListVPCEConfigurationsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListVPCEConfigurationsResponse is the response type for the
// ListVPCEConfigurations API operation.
type ListVPCEConfigurationsResponse struct {
	*ListVPCEConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListVPCEConfigurations request.
func (r *ListVPCEConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

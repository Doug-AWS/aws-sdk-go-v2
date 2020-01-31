// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetVPCEConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the VPC endpoint configuration you want
	// to describe.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s GetVPCEConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetVPCEConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetVPCEConfigurationInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetVPCEConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// An object containing information about your VPC endpoint configuration.
	VpceConfiguration *VPCEConfiguration `locationName:"vpceConfiguration" type:"structure"`
}

// String returns the string representation
func (s GetVPCEConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetVPCEConfiguration = "GetVPCEConfiguration"

// GetVPCEConfigurationRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Returns information about the configuration settings for your Amazon Virtual
// Private Cloud (VPC) endpoint.
//
//    // Example sending a request using GetVPCEConfigurationRequest.
//    req := client.GetVPCEConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetVPCEConfiguration
func (c *Client) GetVPCEConfigurationRequest(input *GetVPCEConfigurationInput) GetVPCEConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetVPCEConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetVPCEConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetVPCEConfigurationOutput{})
	return GetVPCEConfigurationRequest{Request: req, Input: input, Copy: c.GetVPCEConfigurationRequest}
}

// GetVPCEConfigurationRequest is the request type for the
// GetVPCEConfiguration API operation.
type GetVPCEConfigurationRequest struct {
	*aws.Request
	Input *GetVPCEConfigurationInput
	Copy  func(*GetVPCEConfigurationInput) GetVPCEConfigurationRequest
}

// Send marshals and sends the GetVPCEConfiguration API request.
func (r GetVPCEConfigurationRequest) Send(ctx context.Context) (*GetVPCEConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetVPCEConfigurationResponse{
		GetVPCEConfigurationOutput: r.Request.Data.(*GetVPCEConfigurationOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetVPCEConfigurationResponse is the response type for the
// GetVPCEConfiguration API operation.
type GetVPCEConfigurationResponse struct {
	*GetVPCEConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetVPCEConfiguration request.
func (r *GetVPCEConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

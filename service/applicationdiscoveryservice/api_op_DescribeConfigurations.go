// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeConfigurationsInput struct {
	_ struct{} `type:"structure"`

	// One or more configuration IDs.
	//
	// ConfigurationIds is a required field
	ConfigurationIds []string `locationName:"configurationIds" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeConfigurationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeConfigurationsInput"}

	if s.ConfigurationIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	// A key in the response map. The value is an array of data.
	Configurations []map[string]string `locationName:"configurations" type:"list"`
}

// String returns the string representation
func (s DescribeConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeConfigurations = "DescribeConfigurations"

// DescribeConfigurationsRequest returns a request value for making API operation for
// AWS Application Discovery Service.
//
// Retrieves attributes for a list of configuration item IDs.
//
// All of the supplied IDs must be for the same asset type from one of the following:
//
//    * server
//
//    * application
//
//    * process
//
//    * connection
//
// Output fields are specific to the asset type specified. For example, the
// output for a server configuration item includes a list of attributes about
// the server, such as host name, operating system, number of network cards,
// etc.
//
// For a complete list of outputs for each asset type, see Using the DescribeConfigurations
// Action (http://docs.aws.amazon.com/application-discovery/latest/APIReference/discovery-api-queries.html#DescribeConfigurations).
//
//    // Example sending a request using DescribeConfigurationsRequest.
//    req := client.DescribeConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/DescribeConfigurations
func (c *Client) DescribeConfigurationsRequest(input *DescribeConfigurationsInput) DescribeConfigurationsRequest {
	op := &aws.Operation{
		Name:       opDescribeConfigurations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeConfigurationsInput{}
	}

	req := c.newRequest(op, input, &DescribeConfigurationsOutput{})
	return DescribeConfigurationsRequest{Request: req, Input: input, Copy: c.DescribeConfigurationsRequest}
}

// DescribeConfigurationsRequest is the request type for the
// DescribeConfigurations API operation.
type DescribeConfigurationsRequest struct {
	*aws.Request
	Input *DescribeConfigurationsInput
	Copy  func(*DescribeConfigurationsInput) DescribeConfigurationsRequest
}

// Send marshals and sends the DescribeConfigurations API request.
func (r DescribeConfigurationsRequest) Send(ctx context.Context) (*DescribeConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeConfigurationsResponse{
		DescribeConfigurationsOutput: r.Request.Data.(*DescribeConfigurationsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeConfigurationsResponse is the response type for the
// DescribeConfigurations API operation.
type DescribeConfigurationsResponse struct {
	*DescribeConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeConfigurations request.
func (r *DescribeConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

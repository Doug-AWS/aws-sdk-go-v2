// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure for an Amplify App delete request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/DeleteAppRequest
type DeleteAppInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAppInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAppInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAppInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAppInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure for an Amplify App delete request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/DeleteAppResult
type DeleteAppOutput struct {
	_ struct{} `type:"structure"`

	// Amplify App represents different branches of a repository for building, deploying,
	// and hosting.
	//
	// App is a required field
	App *App `locationName:"app" type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteAppOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAppOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.App != nil {
		v := s.App

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "app", v, metadata)
	}
	return nil
}

const opDeleteApp = "DeleteApp"

// DeleteAppRequest returns a request value for making API operation for
// AWS Amplify.
//
// Delete an existing Amplify App by appId.
//
//    // Example sending a request using DeleteAppRequest.
//    req := client.DeleteAppRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/DeleteApp
func (c *Client) DeleteAppRequest(input *DeleteAppInput) DeleteAppRequest {
	op := &aws.Operation{
		Name:       opDeleteApp,
		HTTPMethod: "DELETE",
		HTTPPath:   "/apps/{appId}",
	}

	if input == nil {
		input = &DeleteAppInput{}
	}

	req := c.newRequest(op, input, &DeleteAppOutput{})
	return DeleteAppRequest{Request: req, Input: input, Copy: c.DeleteAppRequest}
}

// DeleteAppRequest is the request type for the
// DeleteApp API operation.
type DeleteAppRequest struct {
	*aws.Request
	Input *DeleteAppInput
	Copy  func(*DeleteAppInput) DeleteAppRequest
}

// Send marshals and sends the DeleteApp API request.
func (r DeleteAppRequest) Send(ctx context.Context) (*DeleteAppResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAppResponse{
		DeleteAppOutput: r.Request.Data.(*DeleteAppOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAppResponse is the response type for the
// DeleteApp API operation.
type DeleteAppResponse struct {
	*DeleteAppOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApp request.
func (r *DeleteAppResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
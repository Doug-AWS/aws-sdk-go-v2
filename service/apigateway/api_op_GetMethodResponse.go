// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes a MethodResponse () resource.
func (c *Client) GetMethodResponse(ctx context.Context, params *GetMethodResponseInput, optFns ...func(*Options)) (*GetMethodResponseOutput, error) {
	stack := middleware.NewStack("GetMethodResponse", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetMethodResponseMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetMethodResponseValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMethodResponse(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "GetMethodResponse",
			Err:           err,
		}
	}
	out := result.(*GetMethodResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to describe a MethodResponse () resource.
type GetMethodResponseInput struct {

	// [Required] The HTTP verb of the Method () resource.
	//
	// This member is required.
	HttpMethod *string

	// [Required] The Resource () identifier for the MethodResponse () resource.
	//
	// This member is required.
	ResourceId *string

	// [Required] The string identifier of the associated RestApi ().
	//
	// This member is required.
	RestApiId *string

	// [Required] The status code for the MethodResponse () resource.
	//
	// This member is required.
	StatusCode *string

	Name *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// Represents a method response of a given HTTP status code returned to the client.
// The method response is passed from the back end through the associated
// integration response that can be transformed using a mapping template.
// Example:
// A MethodResponse instance of an API
//
// Request
//
// The example request retrieves a
// MethodResponse of the 200 status code. GET
// /restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200 HTTP/1.1
// Content-Type: application/json Host: apigateway.us-east-1.amazonaws.com
// X-Amz-Date: 20160603T222952Z Authorization: AWS4-HMAC-SHA256
// Credential={access_key_ID}/20160603/us-east-1/apigateway/aws4_request,
// SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
// Response
//
// The
// successful response returns 200 OK status and a payload as follows: { "_links":
// { "curies": { "href":
// "https://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-method-response-{rel}.html",
// "name": "methodresponse", "templated": true }, "self": { "href":
// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200", "title":
// "200" }, "methodresponse:delete": { "href":
// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200" },
// "methodresponse:update": { "href":
// "/restapis/fugvjdxtri/resources/3kzxbg5sa2/methods/GET/responses/200" } },
// "responseModels": { "application/json": "Empty" }, "responseParameters": {
// "method.response.header.Content-Type": false }, "statusCode": "200" }Method (),
// IntegrationResponse (), Integration ()Creating an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type GetMethodResponseOutput struct {

	// Specifies the Model () resources used for the response's content-type. Response
	// models are represented as a key/value map, with a content-type as the key and a
	// Model () name as the value.
	ResponseModels map[string]*string

	// A key-value map specifying required or optional response parameters that API
	// Gateway can send back to the caller. A key defines a method response header and
	// the value specifies whether the associated method response header is required or
	// not. The expression of the key must match the pattern
	// method.response.header.{name}, where name is a valid and unique header name. API
	// Gateway passes certain integration response data to the method response headers
	// specified here according to the mapping you prescribe in the API's
	// IntegrationResponse (). The integration response data that can be mapped include
	// an integration response header expressed in integration.response.header.{name},
	// a static value enclosed within a pair of single quotes (e.g.,
	// 'application/json'), or a JSON expression from the back-end response payload in
	// the form of integration.response.body.{JSON-expression}, where JSON-expression
	// is a valid JSON expression without the $ prefix.)
	ResponseParameters map[string]*bool

	// The method response's status code.
	StatusCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetMethodResponseMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetMethodResponse{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMethodResponse{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetMethodResponse(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetMethodResponse",
	}
}

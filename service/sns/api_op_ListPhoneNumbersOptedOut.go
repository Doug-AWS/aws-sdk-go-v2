// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the ListPhoneNumbersOptedOut action.
type ListPhoneNumbersOptedOutInput struct {
	_ struct{} `type:"structure"`

	// A NextToken string is used when you call the ListPhoneNumbersOptedOut action
	// to retrieve additional records that are available after the first page of
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListPhoneNumbersOptedOutInput) String() string {
	return awsutil.Prettify(s)
}

// The response from the ListPhoneNumbersOptedOut action.
type ListPhoneNumbersOptedOutOutput struct {
	_ struct{} `type:"structure"`

	// A NextToken string is returned when you call the ListPhoneNumbersOptedOut
	// action if additional records are available after the first page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A list of phone numbers that are opted out of receiving SMS messages. The
	// list is paginated, and each page can contain up to 100 phone numbers.
	PhoneNumbers []string `locationName:"phoneNumbers" type:"list"`
}

// String returns the string representation
func (s ListPhoneNumbersOptedOutOutput) String() string {
	return awsutil.Prettify(s)
}

const opListPhoneNumbersOptedOut = "ListPhoneNumbersOptedOut"

// ListPhoneNumbersOptedOutRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Returns a list of phone numbers that are opted out, meaning you cannot send
// SMS messages to them.
//
// The results for ListPhoneNumbersOptedOut are paginated, and each page returns
// up to 100 phone numbers. If additional phone numbers are available after
// the first page of results, then a NextToken string will be returned. To receive
// the next page, you call ListPhoneNumbersOptedOut again using the NextToken
// string received from the previous call. When there are no more records to
// return, NextToken will be null.
//
//    // Example sending a request using ListPhoneNumbersOptedOutRequest.
//    req := client.ListPhoneNumbersOptedOutRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/ListPhoneNumbersOptedOut
func (c *Client) ListPhoneNumbersOptedOutRequest(input *ListPhoneNumbersOptedOutInput) ListPhoneNumbersOptedOutRequest {
	op := &aws.Operation{
		Name:       opListPhoneNumbersOptedOut,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListPhoneNumbersOptedOutInput{}
	}

	req := c.newRequest(op, input, &ListPhoneNumbersOptedOutOutput{})

	return ListPhoneNumbersOptedOutRequest{Request: req, Input: input, Copy: c.ListPhoneNumbersOptedOutRequest}
}

// ListPhoneNumbersOptedOutRequest is the request type for the
// ListPhoneNumbersOptedOut API operation.
type ListPhoneNumbersOptedOutRequest struct {
	*aws.Request
	Input *ListPhoneNumbersOptedOutInput
	Copy  func(*ListPhoneNumbersOptedOutInput) ListPhoneNumbersOptedOutRequest
}

// Send marshals and sends the ListPhoneNumbersOptedOut API request.
func (r ListPhoneNumbersOptedOutRequest) Send(ctx context.Context) (*ListPhoneNumbersOptedOutResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPhoneNumbersOptedOutResponse{
		ListPhoneNumbersOptedOutOutput: r.Request.Data.(*ListPhoneNumbersOptedOutOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListPhoneNumbersOptedOutResponse is the response type for the
// ListPhoneNumbersOptedOut API operation.
type ListPhoneNumbersOptedOutResponse struct {
	*ListPhoneNumbersOptedOutOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPhoneNumbersOptedOut request.
func (r *ListPhoneNumbersOptedOutResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
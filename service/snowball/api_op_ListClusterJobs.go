// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package snowball

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListClusterJobsInput struct {
	_ struct{} `type:"structure"`

	// The 39-character ID for the cluster that you want to list, for example CID123e4567-e89b-12d3-a456-426655440000.
	//
	// ClusterId is a required field
	ClusterId *string `min:"39" type:"string" required:"true"`

	// The number of JobListEntry objects to return.
	MaxResults *int64 `type:"integer"`

	// HTTP requests are stateless. To identify what object comes "next" in the
	// list of JobListEntry objects, you have the option of specifying NextToken
	// as the starting point for your returned list.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListClusterJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListClusterJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListClusterJobsInput"}

	if s.ClusterId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterId"))
	}
	if s.ClusterId != nil && len(*s.ClusterId) < 39 {
		invalidParams.Add(aws.NewErrParamMinLen("ClusterId", 39))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListClusterJobsOutput struct {
	_ struct{} `type:"structure"`

	// Each JobListEntry object contains a job's state, a job's ID, and a value
	// that indicates whether the job is a job part, in the case of export jobs.
	JobListEntries []JobListEntry `type:"list"`

	// HTTP requests are stateless. If you use the automatically generated NextToken
	// value in your next ListClusterJobsResult call, your list of returned jobs
	// will start from this point in the array.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListClusterJobsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListClusterJobs = "ListClusterJobs"

// ListClusterJobsRequest returns a request value for making API operation for
// Amazon Import/Export Snowball.
//
// Returns an array of JobListEntry objects of the specified length. Each JobListEntry
// object is for a job in the specified cluster and contains a job's state,
// a job's ID, and other information.
//
//    // Example sending a request using ListClusterJobsRequest.
//    req := client.ListClusterJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/snowball-2016-06-30/ListClusterJobs
func (c *Client) ListClusterJobsRequest(input *ListClusterJobsInput) ListClusterJobsRequest {
	op := &aws.Operation{
		Name:       opListClusterJobs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListClusterJobsInput{}
	}

	req := c.newRequest(op, input, &ListClusterJobsOutput{})

	return ListClusterJobsRequest{Request: req, Input: input, Copy: c.ListClusterJobsRequest}
}

// ListClusterJobsRequest is the request type for the
// ListClusterJobs API operation.
type ListClusterJobsRequest struct {
	*aws.Request
	Input *ListClusterJobsInput
	Copy  func(*ListClusterJobsInput) ListClusterJobsRequest
}

// Send marshals and sends the ListClusterJobs API request.
func (r ListClusterJobsRequest) Send(ctx context.Context) (*ListClusterJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListClusterJobsResponse{
		ListClusterJobsOutput: r.Request.Data.(*ListClusterJobsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListClusterJobsResponse is the response type for the
// ListClusterJobs API operation.
type ListClusterJobsResponse struct {
	*ListClusterJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListClusterJobs request.
func (r *ListClusterJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
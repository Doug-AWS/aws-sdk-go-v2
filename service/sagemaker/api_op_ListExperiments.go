// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListExperimentsInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns only experiments created after the specified time.
	CreatedAfter *time.Time `type:"timestamp"`

	// A filter that returns only experiments created before the specified time.
	CreatedBefore *time.Time `type:"timestamp"`

	// The maximum number of experiments to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the previous call to ListExperiments didn't return the full set of experiments,
	// the call returns a token for getting the next set of experiments.
	NextToken *string `type:"string"`

	// The property used to sort results. The default value is CreationTime.
	SortBy SortExperimentsBy `type:"string" enum:"true"`

	// The sort order. The default value is Descending.
	SortOrder SortOrder `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListExperimentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListExperimentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListExperimentsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListExperimentsOutput struct {
	_ struct{} `type:"structure"`

	// A list of the summaries of your experiments.
	ExperimentSummaries []ExperimentSummary `type:"list"`

	// A token for getting the next set of experiments, if there are any.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListExperimentsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListExperiments = "ListExperiments"

// ListExperimentsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Lists all the experiments in your account. The list can be filtered to show
// only experiments that were created in a specific time range. The list can
// be sorted by experiment name or creation time.
//
//    // Example sending a request using ListExperimentsRequest.
//    req := client.ListExperimentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListExperiments
func (c *Client) ListExperimentsRequest(input *ListExperimentsInput) ListExperimentsRequest {
	op := &aws.Operation{
		Name:       opListExperiments,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListExperimentsInput{}
	}

	req := c.newRequest(op, input, &ListExperimentsOutput{})
	return ListExperimentsRequest{Request: req, Input: input, Copy: c.ListExperimentsRequest}
}

// ListExperimentsRequest is the request type for the
// ListExperiments API operation.
type ListExperimentsRequest struct {
	*aws.Request
	Input *ListExperimentsInput
	Copy  func(*ListExperimentsInput) ListExperimentsRequest
}

// Send marshals and sends the ListExperiments API request.
func (r ListExperimentsRequest) Send(ctx context.Context) (*ListExperimentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListExperimentsResponse{
		ListExperimentsOutput: r.Request.Data.(*ListExperimentsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListExperimentsRequestPaginator returns a paginator for ListExperiments.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListExperimentsRequest(input)
//   p := sagemaker.NewListExperimentsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListExperimentsPaginator(req ListExperimentsRequest) ListExperimentsPaginator {
	return ListExperimentsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListExperimentsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListExperimentsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListExperimentsPaginator struct {
	aws.Pager
}

func (p *ListExperimentsPaginator) CurrentPage() *ListExperimentsOutput {
	return p.Pager.CurrentPage().(*ListExperimentsOutput)
}

// ListExperimentsResponse is the response type for the
// ListExperiments API operation.
type ListExperimentsResponse struct {
	*ListExperimentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListExperiments request.
func (r *ListExperimentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

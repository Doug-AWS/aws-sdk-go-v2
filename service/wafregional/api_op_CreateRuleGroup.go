// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/waf"
)

type CreateRuleGroupInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// A friendly name or description for the metrics for this RuleGroup. The name
	// can contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length
	// 128 and minimum length one. It can't contain whitespace or metric names reserved
	// for AWS WAF, including "All" and "Default_Action." You can't change the name
	// of the metric after you create the RuleGroup.
	//
	// MetricName is a required field
	MetricName *string `type:"string" required:"true"`

	// A friendly name or description of the RuleGroup. You can't change Name after
	// you create a RuleGroup.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	Tags []waf.Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s CreateRuleGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRuleGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRuleGroupInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.MetricName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MetricName"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateRuleGroupOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the CreateRuleGroup request. You
	// can also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`

	// An empty RuleGroup.
	RuleGroup *waf.RuleGroup `type:"structure"`
}

// String returns the string representation
func (s CreateRuleGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateRuleGroup = "CreateRuleGroup"

// CreateRuleGroupRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Creates a RuleGroup. A rule group is a collection of predefined rules that
// you add to a web ACL. You use UpdateRuleGroup to add rules to the rule group.
//
// Rule groups are subject to the following limits:
//
//    * Three rule groups per account. You can request an increase to this limit
//    by contacting customer support.
//
//    * One rule group per web ACL.
//
//    * Ten rules per rule group.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using CreateRuleGroupRequest.
//    req := client.CreateRuleGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/CreateRuleGroup
func (c *Client) CreateRuleGroupRequest(input *CreateRuleGroupInput) CreateRuleGroupRequest {
	op := &aws.Operation{
		Name:       opCreateRuleGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateRuleGroupInput{}
	}

	req := c.newRequest(op, input, &CreateRuleGroupOutput{})
	return CreateRuleGroupRequest{Request: req, Input: input, Copy: c.CreateRuleGroupRequest}
}

// CreateRuleGroupRequest is the request type for the
// CreateRuleGroup API operation.
type CreateRuleGroupRequest struct {
	*aws.Request
	Input *CreateRuleGroupInput
	Copy  func(*CreateRuleGroupInput) CreateRuleGroupRequest
}

// Send marshals and sends the CreateRuleGroup API request.
func (r CreateRuleGroupRequest) Send(ctx context.Context) (*CreateRuleGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRuleGroupResponse{
		CreateRuleGroupOutput: r.Request.Data.(*CreateRuleGroupOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRuleGroupResponse is the response type for the
// CreateRuleGroup API operation.
type CreateRuleGroupResponse struct {
	*CreateRuleGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRuleGroup request.
func (r *CreateRuleGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

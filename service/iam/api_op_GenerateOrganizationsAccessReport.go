// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Generates a report for service last accessed data for AWS Organizations. You can
// generate a report for any entities (organization root, organizational unit, or
// account) or policies in your organization. To call this operation, you must be
// signed in using your AWS Organizations master account credentials. You can use
// your long-term IAM user or root user credentials, or temporary credentials from
// assuming an IAM role. SCPs must be enabled for your organization root. You must
// have the required IAM and AWS Organizations permissions. For more information,
// see Refining Permissions Using Service Last Accessed Data
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html)
// in the IAM User Guide. You can generate a service last accessed data report for
// entities by specifying only the entity's path. This data includes a list of
// services that are allowed by any service control policies (SCPs) that apply to
// the entity. You can generate a service last accessed data report for a policy by
// specifying an entity's path and an optional AWS Organizations policy ID. This
// data includes a list of services that are allowed by the specified SCP. For each
// service in both report types, the data includes the most recent account activity
// that the policy allows to account principals in the entity or the entity's
// children. For important information about the data, reporting period,
// permissions required, troubleshooting, and supported Regions see Reducing
// Permissions Using Service Last Accessed Data
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html)
// in the IAM User Guide. The data includes all attempts to access AWS, not just
// the successful ones. This includes all attempts that were made using the AWS
// Management Console, the AWS API through any of the SDKs, or any of the command
// line tools. An unexpected entry in the service last accessed data does not mean
// that an account has been compromised, because the request might have been
// denied. Refer to your CloudTrail logs as the authoritative source for
// information about all API calls and whether they were successful or denied
// access. For more information, see Logging IAM Events with CloudTrail
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/cloudtrail-integration.html)
// in the IAM User Guide. This operation returns a JobId. Use this parameter in the
// GetOrganizationsAccessReport () operation to check the status of the report
// generation. To check the status of this request, use the JobId parameter in the
// GetOrganizationsAccessReport () operation and test the JobStatus response
// parameter. When the job is complete, you can retrieve the report. To generate a
// service last accessed data report for entities, specify an entity path without
// specifying the optional AWS Organizations policy ID. The type of entity that you
// specify determines the data returned in the report.
//
//     * Root – When you
// specify the organizations root as the entity, the resulting report lists all of
// the services allowed by SCPs that are attached to your root. For each service,
// the report includes data for all accounts in your organization except the master
// account, because the master account is not limited by SCPs.
//
//     * OU – When you
// specify an organizational unit (OU) as the entity, the resulting report lists
// all of the services allowed by SCPs that are attached to the OU and its parents.
// For each service, the report includes data for all accounts in the OU or its
// children. This data excludes the master account, because the master account is
// not limited by SCPs.
//
//     * Master account – When you specify the master
// account, the resulting report lists all AWS services, because the master account
// is not limited by SCPs. For each service, the report includes data for only the
// master account.
//
//     * Account – When you specify another account as the entity,
// the resulting report lists all of the services allowed by SCPs that are attached
// to the account and its parents. For each service, the report includes data for
// only the specified account.
//
// To generate a service last accessed data report for
// policies, specify an entity path and the optional AWS Organizations policy ID.
// The type of entity that you specify determines the data returned for each
// service.
//
//     * Root – When you specify the root entity and a policy ID, the
// resulting report lists all of the services that are allowed by the specified
// SCP. For each service, the report includes data for all accounts in your
// organization to which the SCP applies. This data excludes the master account,
// because the master account is not limited by SCPs. If the SCP is not attached to
// any entities in the organization, then the report will return a list of services
// with no data.
//
//     * OU – When you specify an OU entity and a policy ID, the
// resulting report lists all of the services that are allowed by the specified
// SCP. For each service, the report includes data for all accounts in the OU or
// its children to which the SCP applies. This means that other accounts outside
// the OU that are affected by the SCP might not be included in the data. This data
// excludes the master account, because the master account is not limited by SCPs.
// If the SCP is not attached to the OU or one of its children, the report will
// return a list of services with no data.
//
//     * Master account – When you specify
// the master account, the resulting report lists all AWS services, because the
// master account is not limited by SCPs. If you specify a policy ID in the CLI or
// API, the policy is ignored. For each service, the report includes data for only
// the master account.
//
//     * Account – When you specify another account entity and
// a policy ID, the resulting report lists all of the services that are allowed by
// the specified SCP. For each service, the report includes data for only the
// specified account. This means that other accounts in the organization that are
// affected by the SCP might not be included in the data. If the SCP is not
// attached to the account, the report will return a list of services with no
// data.
//
// Service last accessed data does not use other policy types when
// determining whether a principal could access a service. These other policy types
// include identity-based policies, resource-based policies, access control lists,
// IAM permissions boundaries, and STS assume role policies. It only applies SCP
// logic. For more about the evaluation of policy types, see Evaluating Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-basics)
// in the IAM User Guide. For more information about service last accessed data,
// see Reducing Policy Scope by Viewing User Activity
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html)
// in the IAM User Guide.
func (c *Client) GenerateOrganizationsAccessReport(ctx context.Context, params *GenerateOrganizationsAccessReportInput, optFns ...func(*Options)) (*GenerateOrganizationsAccessReportOutput, error) {
	stack := middleware.NewStack("GenerateOrganizationsAccessReport", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGenerateOrganizationsAccessReportMiddlewares(stack)
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
	addOpGenerateOrganizationsAccessReportValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateOrganizationsAccessReport(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

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
			OperationName: "GenerateOrganizationsAccessReport",
			Err:           err,
		}
	}
	out := result.(*GenerateOrganizationsAccessReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateOrganizationsAccessReportInput struct {

	// The path of the AWS Organizations entity (root, OU, or account). You can build
	// an entity path using the known structure of your organization. For example,
	// assume that your account ID is 123456789012 and its parent OU ID is
	// ou-rge0-awsabcde. The organization root ID is r-f6g7h8i9j0example and your
	// organization ID is o-a1b2c3d4e5. Your entity path is
	// o-a1b2c3d4e5/r-f6g7h8i9j0example/ou-rge0-awsabcde/123456789012.
	//
	// This member is required.
	EntityPath *string

	// The identifier of the AWS Organizations service control policy (SCP). This
	// parameter is optional. This ID is used to generate information about when an
	// account principal that is limited by the SCP attempted to access an AWS service.
	OrganizationsPolicyId *string
}

type GenerateOrganizationsAccessReportOutput struct {

	// The job identifier that you can use in the GetOrganizationsAccessReport ()
	// operation.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGenerateOrganizationsAccessReportMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGenerateOrganizationsAccessReport{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGenerateOrganizationsAccessReport{}, middleware.After)
}

func newServiceMetadataMiddleware_opGenerateOrganizationsAccessReport(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GenerateOrganizationsAccessReport",
	}
}

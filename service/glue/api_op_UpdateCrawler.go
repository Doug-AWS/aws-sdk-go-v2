// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a crawler. If a crawler is running, you must stop it using StopCrawler
// before updating it.
func (c *Client) UpdateCrawler(ctx context.Context, params *UpdateCrawlerInput, optFns ...func(*Options)) (*UpdateCrawlerOutput, error) {
	stack := middleware.NewStack("UpdateCrawler", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateCrawlerMiddlewares(stack)
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
	addOpUpdateCrawlerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCrawler(options.Region), middleware.Before)
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
			OperationName: "UpdateCrawler",
			Err:           err,
		}
	}
	out := result.(*UpdateCrawlerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCrawlerInput struct {

	// Name of the new crawler.
	//
	// This member is required.
	Name *string

	// A list of custom classifiers that the user has registered. By default, all
	// built-in classifiers are included in a crawl, but these custom classifiers
	// always override the default classifiers for a given classification.
	Classifiers []*string

	// Crawler configuration information. This versioned JSON string allows users to
	// specify aspects of a crawler's behavior. For more information, see Configuring a
	// Crawler (https://docs.aws.amazon.com/glue/latest/dg/crawler-configuration.html).
	Configuration *string

	// The name of the SecurityConfiguration structure to be used by this crawler.
	CrawlerSecurityConfiguration *string

	// The AWS Glue database where results are stored, such as:
	// arn:aws:daylight:us-east-1::database/sometable/*.
	DatabaseName *string

	// A description of the new crawler.
	Description *string

	// The IAM role or Amazon Resource Name (ARN) of an IAM role that is used by the
	// new crawler to access customer resources.
	Role *string

	// A cron expression used to specify the schedule (see Time-Based Schedules for
	// Jobs and Crawlers
	// (https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html).
	// For example, to run something every day at 12:15 UTC, you would specify: cron(15
	// 12 * * ? *).
	Schedule *string

	// The policy for the crawler's update and deletion behavior.
	SchemaChangePolicy *types.SchemaChangePolicy

	// The table prefix used for catalog tables that are created.
	TablePrefix *string

	// A list of targets to crawl.
	Targets *types.CrawlerTargets
}

type UpdateCrawlerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateCrawlerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCrawler{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCrawler{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateCrawler(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "UpdateCrawler",
	}
}

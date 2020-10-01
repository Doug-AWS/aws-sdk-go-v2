// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Assigns new properties to a user. Parameters you pass modify any or all of the
// following: the home directory, role, and policy for the UserName and ServerId
// you specify.  <p>The response returns the <code>ServerId</code> and the
// <code>UserName</code> for the updated user.</p>
func (c *Client) UpdateUser(ctx context.Context, params *UpdateUserInput, optFns ...func(*Options)) (*UpdateUserOutput, error) {
	stack := middleware.NewStack("UpdateUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateUserMiddlewares(stack)
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
	addOpUpdateUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUser(options.Region), middleware.Before)
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
			OperationName: "UpdateUser",
			Err:           err,
		}
	}
	out := result.(*UpdateUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateUserInput struct {

	// A system-assigned unique identifier for a file transfer protocol-enabled server
	// instance that the user account is assigned to.
	//
	// This member is required.
	ServerId *string

	// A unique string that identifies a user and is associated with a file transfer
	// protocol-enabled server as specified by the ServerId. This is the string that
	// will be used by your user when they log in to your server. This user name is a
	// minimum of 3 and a maximum of 32 characters long. The following are valid
	// characters: a-z, A-Z, 0-9, underscore, and hyphen. The user name can't start
	// with a hyphen.
	//
	// This member is required.
	UserName *string

	// Specifies the landing directory (folder) for a user when they log in to the file
	// transfer protocol-enabled server using their file transfer protocol client.
	// <p>An example is <code>your-Amazon-S3-bucket-name>/home/username</code>.</p>
	HomeDirectory *string

	// Logical directory mappings that specify what Amazon S3 paths and keys should be
	// visible to your user and how you want to make them visible. You will need to
	// specify the "Entry" and "Target" pair, where Entry shows how the path is made
	// visible and Target is the actual Amazon S3 path. If you only specify a target,
	// it will be displayed as is. You will need to also make sure that your IAM role
	// provides access to paths in Target. The following is an example.  <p> <code>'[
	// "/bucket2/documentation", { "Entry": "your-personal-report.pdf", "Target":
	// "/bucket3/customized-reports/${transfer:UserName}.pdf" } ]'</code> </p> <p>In
	// most cases, you can use this value instead of the scope-down policy to lock your
	// user down to the designated home directory ("chroot"). To do this, you can set
	// <code>Entry</code> to '/' and set <code>Target</code> to the HomeDirectory
	// parameter value.</p> <note> <p>If the target of a logical directory entry does
	// not exist in Amazon S3, the entry will be ignored. As a workaround, you can use
	// the Amazon S3 api to create 0 byte objects as place holders for your directory.
	// If using the CLI, use the <code>s3api</code> call instead of <code>s3</code> so
	// you can use the put-object operation. For example, you use the following:
	// <code>aws s3api put-object --bucket bucketname --key path/to/folder/</code>.
	// Make sure that the end of the key name ends in a / for it to be considered a
	// folder.</p> </note>
	HomeDirectoryMappings []*types.HomeDirectoryMapEntry

	// The type of landing directory (folder) you want your users' home directory to be
	// when they log into the file transfer protocol-enabled server. If you set it to
	// PATH, the user will see the absolute Amazon S3 bucket paths as is in their file
	// transfer protocol clients. If you set it LOGICAL, you will need to provide
	// mappings in the HomeDirectoryMappings for how you want to make Amazon S3 paths
	// visible to your users.
	HomeDirectoryType types.HomeDirectoryType

	// Allows you to supply a scope-down policy for your user so you can use the same
	// IAM role across multiple users. The policy scopes down user access to portions
	// of your Amazon S3 bucket. Variables you can use inside this policy include
	// ${Transfer:UserName}, ${Transfer:HomeDirectory}, and ${Transfer:HomeBucket}. For
	// scope-down policies, AWS Transfer Family stores the policy as a JSON blob,
	// instead of the Amazon Resource Name (ARN) of the policy. You save the policy as
	// a JSON blob and pass it in the Policy argument.  <p>For an example of a
	// scope-down policy, see <a
	// href="https://docs.aws.amazon.com/transfer/latest/userguide/users.html#users-policies-scope-down">Creating
	// a scope-down policy</a>.</p> <p>For more information, see <a
	// href="https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html">AssumeRole</a>
	// in the <i>AWS Security Token Service API Reference</i>.</p> </note>
	Policy *string

	// The IAM role that controls your users' access to your Amazon S3 bucket. The
	// policies attached to this role will determine the level of access you want to
	// provide your users when transferring files into and out of your Amazon S3 bucket
	// or buckets. The IAM role should also contain a trust relationship that allows
	// the file transfer protocol-enabled server to access your resources when
	// servicing your users' transfer requests.
	Role *string
}

// UpdateUserResponse returns the user name and file transfer protocol-enabled
// server identifier for the request to update a user's properties.
type UpdateUserOutput struct {

	// A system-assigned unique identifier for a file transfer protocol-enabled server
	// instance that the user account is assigned to.
	//
	// This member is required.
	ServerId *string

	// The unique identifier for a user that is assigned to a file transfer
	// protocol-enabled server instance that was specified in the request.
	//
	// This member is required.
	UserName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateUser{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "UpdateUser",
	}
}

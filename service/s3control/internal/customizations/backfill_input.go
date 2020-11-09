package customizations

import (
	"context"
	"fmt"

	"github.com/awslabs/smithy-go/middleware"

	"github.com/aws/aws-sdk-go-v2/service/internal/s3shared"
)

// BackfillInputMiddleware validates and backfill's values from ARN into request serializable input.
// This middleware must be executed after `ARNLookup` and before `inputValidationMiddleware`.
type BackfillInputMiddleware struct {

	// CopyInput creates a copy of input to be modified, this ensures the original input is not modified.
	CopyInput func(interface{}) (interface{}, error)

	// BackfillAccountID points to a function that validates the input for accountID. If absent, it populates the
	// accountID and returns a copy. If present, but different than passed in accountID value throws an error
	BackfillAccountID func(interface{}, string) error
}

// ID representing the middleware
func (m *BackfillInputMiddleware) ID() string {
	return "S3Control:BackfillInputMiddleware"
}

// HandleInitialize handles the middleware behavior in an Initialize step.
func (m *BackfillInputMiddleware) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	// fetch arn from context
	av, ok := s3shared.GetARNResourceFromContext(ctx)
	if !ok {
		return next.HandleInitialize(ctx, in)
	}

	// create a copy of input, and assign it on params
	in.Parameters, err = m.CopyInput(in.Parameters)
	if err != nil {
		return out, metadata, fmt.Errorf("error creating a copy of input")
	}

	// if not supported, move to next
	if m.BackfillAccountID == nil {
		return next.HandleInitialize(ctx, in)
	}

	err = m.BackfillAccountID(in.Parameters, av.AccountID)
	if err != nil {
		return out, metadata, fmt.Errorf("invalid ARN, %w", err)
	}

	return next.HandleInitialize(ctx, in)
}

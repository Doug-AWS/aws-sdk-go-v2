// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpCreateAccessPoint struct {
}

func (*validateOpCreateAccessPoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateAccessPoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateAccessPointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateAccessPointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateJob struct {
}

func (*validateOpCreateJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAccessPoint struct {
}

func (*validateOpDeleteAccessPoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAccessPoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAccessPointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAccessPointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAccessPointPolicy struct {
}

func (*validateOpDeleteAccessPointPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAccessPointPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAccessPointPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAccessPointPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteJobTagging struct {
}

func (*validateOpDeleteJobTagging) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteJobTagging) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteJobTaggingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteJobTaggingInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeletePublicAccessBlock struct {
}

func (*validateOpDeletePublicAccessBlock) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeletePublicAccessBlock) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeletePublicAccessBlockInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeletePublicAccessBlockInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeJob struct {
}

func (*validateOpDescribeJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetAccessPoint struct {
}

func (*validateOpGetAccessPoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAccessPoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetAccessPointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetAccessPointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetAccessPointPolicy struct {
}

func (*validateOpGetAccessPointPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAccessPointPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetAccessPointPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetAccessPointPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetAccessPointPolicyStatus struct {
}

func (*validateOpGetAccessPointPolicyStatus) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAccessPointPolicyStatus) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetAccessPointPolicyStatusInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetAccessPointPolicyStatusInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetJobTagging struct {
}

func (*validateOpGetJobTagging) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetJobTagging) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetJobTaggingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetJobTaggingInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetPublicAccessBlock struct {
}

func (*validateOpGetPublicAccessBlock) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetPublicAccessBlock) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetPublicAccessBlockInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetPublicAccessBlockInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListAccessPoints struct {
}

func (*validateOpListAccessPoints) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListAccessPoints) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListAccessPointsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListAccessPointsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListJobs struct {
}

func (*validateOpListJobs) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListJobs) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListJobsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListJobsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutAccessPointPolicy struct {
}

func (*validateOpPutAccessPointPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutAccessPointPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutAccessPointPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutAccessPointPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutJobTagging struct {
}

func (*validateOpPutJobTagging) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutJobTagging) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutJobTaggingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutJobTaggingInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutPublicAccessBlock struct {
}

func (*validateOpPutPublicAccessBlock) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutPublicAccessBlock) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutPublicAccessBlockInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutPublicAccessBlockInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateJobPriority struct {
}

func (*validateOpUpdateJobPriority) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateJobPriority) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateJobPriorityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateJobPriorityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateJobStatus struct {
}

func (*validateOpUpdateJobStatus) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateJobStatus) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateJobStatusInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateJobStatusInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateAccessPointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateAccessPoint{}, middleware.After)
}

func addOpCreateJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateJob{}, middleware.After)
}

func addOpDeleteAccessPointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAccessPoint{}, middleware.After)
}

func addOpDeleteAccessPointPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAccessPointPolicy{}, middleware.After)
}

func addOpDeleteJobTaggingValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteJobTagging{}, middleware.After)
}

func addOpDeletePublicAccessBlockValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeletePublicAccessBlock{}, middleware.After)
}

func addOpDescribeJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeJob{}, middleware.After)
}

func addOpGetAccessPointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAccessPoint{}, middleware.After)
}

func addOpGetAccessPointPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAccessPointPolicy{}, middleware.After)
}

func addOpGetAccessPointPolicyStatusValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAccessPointPolicyStatus{}, middleware.After)
}

func addOpGetJobTaggingValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetJobTagging{}, middleware.After)
}

func addOpGetPublicAccessBlockValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetPublicAccessBlock{}, middleware.After)
}

func addOpListAccessPointsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListAccessPoints{}, middleware.After)
}

func addOpListJobsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListJobs{}, middleware.After)
}

func addOpPutAccessPointPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutAccessPointPolicy{}, middleware.After)
}

func addOpPutJobTaggingValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutJobTagging{}, middleware.After)
}

func addOpPutPublicAccessBlockValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutPublicAccessBlock{}, middleware.After)
}

func addOpUpdateJobPriorityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateJobPriority{}, middleware.After)
}

func addOpUpdateJobStatusValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateJobStatus{}, middleware.After)
}

func validateJobManifest(v *types.JobManifest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "JobManifest"}
	if v.Location == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Location"))
	} else if v.Location != nil {
		if err := validateJobManifestLocation(v.Location); err != nil {
			invalidParams.AddNested("Location", err.(smithy.InvalidParamsError))
		}
	}
	if v.Spec == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Spec"))
	} else if v.Spec != nil {
		if err := validateJobManifestSpec(v.Spec); err != nil {
			invalidParams.AddNested("Spec", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateJobManifestLocation(v *types.JobManifestLocation) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "JobManifestLocation"}
	if v.ObjectArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ObjectArn"))
	}
	if v.ETag == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ETag"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateJobManifestSpec(v *types.JobManifestSpec) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "JobManifestSpec"}
	if len(v.Format) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Format"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateJobOperation(v *types.JobOperation) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "JobOperation"}
	if v.S3PutObjectLegalHold != nil {
		if err := validateS3SetObjectLegalHoldOperation(v.S3PutObjectLegalHold); err != nil {
			invalidParams.AddNested("S3PutObjectLegalHold", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3PutObjectTagging != nil {
		if err := validateS3SetObjectTaggingOperation(v.S3PutObjectTagging); err != nil {
			invalidParams.AddNested("S3PutObjectTagging", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3PutObjectRetention != nil {
		if err := validateS3SetObjectRetentionOperation(v.S3PutObjectRetention); err != nil {
			invalidParams.AddNested("S3PutObjectRetention", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3PutObjectCopy != nil {
		if err := validateS3CopyObjectOperation(v.S3PutObjectCopy); err != nil {
			invalidParams.AddNested("S3PutObjectCopy", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3PutObjectAcl != nil {
		if err := validateS3SetObjectAclOperation(v.S3PutObjectAcl); err != nil {
			invalidParams.AddNested("S3PutObjectAcl", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateJobReport(v *types.JobReport) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "JobReport"}
	if v.Enabled == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Enabled"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3AccessControlList(v *types.S3AccessControlList) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3AccessControlList"}
	if v.Owner == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Owner"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3AccessControlPolicy(v *types.S3AccessControlPolicy) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3AccessControlPolicy"}
	if v.AccessControlList != nil {
		if err := validateS3AccessControlList(v.AccessControlList); err != nil {
			invalidParams.AddNested("AccessControlList", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3CopyObjectOperation(v *types.S3CopyObjectOperation) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3CopyObjectOperation"}
	if v.NewObjectTagging != nil {
		if err := validateS3TagSet(v.NewObjectTagging); err != nil {
			invalidParams.AddNested("NewObjectTagging", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3ObjectLockLegalHold(v *types.S3ObjectLockLegalHold) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3ObjectLockLegalHold"}
	if len(v.Status) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Status"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3SetObjectAclOperation(v *types.S3SetObjectAclOperation) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3SetObjectAclOperation"}
	if v.AccessControlPolicy != nil {
		if err := validateS3AccessControlPolicy(v.AccessControlPolicy); err != nil {
			invalidParams.AddNested("AccessControlPolicy", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3SetObjectLegalHoldOperation(v *types.S3SetObjectLegalHoldOperation) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3SetObjectLegalHoldOperation"}
	if v.LegalHold == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LegalHold"))
	} else if v.LegalHold != nil {
		if err := validateS3ObjectLockLegalHold(v.LegalHold); err != nil {
			invalidParams.AddNested("LegalHold", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3SetObjectRetentionOperation(v *types.S3SetObjectRetentionOperation) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3SetObjectRetentionOperation"}
	if v.Retention == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Retention"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3SetObjectTaggingOperation(v *types.S3SetObjectTaggingOperation) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3SetObjectTaggingOperation"}
	if v.TagSet != nil {
		if err := validateS3TagSet(v.TagSet); err != nil {
			invalidParams.AddNested("TagSet", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3Tag(v *types.S3Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3Tag"}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3TagSet(v []*types.S3Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3TagSet"}
	for i := range v {
		if err := validateS3Tag(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateVpcConfiguration(v *types.VpcConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VpcConfiguration"}
	if v.VpcId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VpcId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateAccessPointInput(v *CreateAccessPointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAccessPointInput"}
	if v.Bucket == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Bucket"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if v.VpcConfiguration != nil {
		if err := validateVpcConfiguration(v.VpcConfiguration); err != nil {
			invalidParams.AddNested("VpcConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateJobInput(v *CreateJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateJobInput"}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if v.Report == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Report"))
	} else if v.Report != nil {
		if err := validateJobReport(v.Report); err != nil {
			invalidParams.AddNested("Report", err.(smithy.InvalidParamsError))
		}
	}
	if v.Priority == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Priority"))
	}
	if v.ClientRequestToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientRequestToken"))
	}
	if v.Operation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Operation"))
	} else if v.Operation != nil {
		if err := validateJobOperation(v.Operation); err != nil {
			invalidParams.AddNested("Operation", err.(smithy.InvalidParamsError))
		}
	}
	if v.Manifest == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Manifest"))
	} else if v.Manifest != nil {
		if err := validateJobManifest(v.Manifest); err != nil {
			invalidParams.AddNested("Manifest", err.(smithy.InvalidParamsError))
		}
	}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if v.Tags != nil {
		if err := validateS3TagSet(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAccessPointInput(v *DeleteAccessPointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAccessPointInput"}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAccessPointPolicyInput(v *DeleteAccessPointPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAccessPointPolicyInput"}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteJobTaggingInput(v *DeleteJobTaggingInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteJobTaggingInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeletePublicAccessBlockInput(v *DeletePublicAccessBlockInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeletePublicAccessBlockInput"}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeJobInput(v *DescribeJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeJobInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetAccessPointInput(v *GetAccessPointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetAccessPointInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetAccessPointPolicyInput(v *GetAccessPointPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetAccessPointPolicyInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetAccessPointPolicyStatusInput(v *GetAccessPointPolicyStatusInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetAccessPointPolicyStatusInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetJobTaggingInput(v *GetJobTaggingInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetJobTaggingInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetPublicAccessBlockInput(v *GetPublicAccessBlockInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetPublicAccessBlockInput"}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListAccessPointsInput(v *ListAccessPointsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListAccessPointsInput"}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListJobsInput(v *ListJobsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListJobsInput"}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutAccessPointPolicyInput(v *PutAccessPointPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutAccessPointPolicyInput"}
	if v.Policy == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Policy"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutJobTaggingInput(v *PutJobTaggingInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutJobTaggingInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateS3TagSet(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutPublicAccessBlockInput(v *PutPublicAccessBlockInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutPublicAccessBlockInput"}
	if v.PublicAccessBlockConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PublicAccessBlockConfiguration"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateJobPriorityInput(v *UpdateJobPriorityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateJobPriorityInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if v.Priority == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Priority"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateJobStatusInput(v *UpdateJobStatusInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateJobStatusInput"}
	if v.JobId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if len(v.RequestedJobStatus) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("RequestedJobStatus"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Contains details about the analyzed resource.
type AnalyzedResource struct {

	// The time at which the resource was analyzed.
	//
	// This member is required.
	AnalyzedAt *time.Time

	// The time at which the finding was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// Indicates whether the policy that generated the finding grants public access to
	// the resource.
	//
	// This member is required.
	IsPublic *bool

	// The ARN of the resource that was analyzed.
	//
	// This member is required.
	ResourceArn *string

	// The AWS account ID that owns the resource.
	//
	// This member is required.
	ResourceOwnerAccount *string

	// The type of the resource that was analyzed.
	//
	// This member is required.
	ResourceType ResourceType

	// The time at which the finding was updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The actions that an external principal is granted permission to use by the
	// policy that generated the finding.
	Actions []*string

	// An error message.
	Error *string

	// Indicates how the access that generated the finding is granted. This is
	// populated for Amazon S3 bucket findings.
	SharedVia []*string

	// The current status of the finding generated from the analyzed resource.
	Status FindingStatus
}

// Contains the ARN of the analyzed resource.
type AnalyzedResourceSummary struct {

	// The ARN of the analyzed resource.
	//
	// This member is required.
	ResourceArn *string

	// The AWS account ID that owns the resource.
	//
	// This member is required.
	ResourceOwnerAccount *string

	// The type of resource that was analyzed.
	//
	// This member is required.
	ResourceType ResourceType
}

// Contains information about the analyzer.
type AnalyzerSummary struct {

	// The ARN of the analyzer.
	//
	// This member is required.
	Arn *string

	// A timestamp for the time at which the analyzer was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The name of the analyzer.
	//
	// This member is required.
	Name *string

	// The status of the analyzer. An Active analyzer successfully monitors supported
	// resources and generates new findings. The analyzer is Disabled when a user
	// action, such as removing trusted access for IAM Access Analyzer from AWS
	// Organizations, causes the analyzer to stop generating new findings. The status
	// is Creating when the analyzer creation is in progress and Failed when the
	// analyzer creation has failed.
	//
	// This member is required.
	Status AnalyzerStatus

	// The type of analyzer, which corresponds to the zone of trust chosen for the
	// analyzer.
	//
	// This member is required.
	Type Type

	// The resource that was most recently analyzed by the analyzer.
	LastResourceAnalyzed *string

	// The time at which the most recently analyzed resource was analyzed.
	LastResourceAnalyzedAt *time.Time

	// The statusReason provides more details about the current status of the analyzer.
	// For example, if the creation for the analyzer fails, a Failed status is
	// displayed. For an analyzer with organization as the type, this failure can be
	// due to an issue with creating the service-linked roles required in the member
	// accounts of the AWS organization.
	StatusReason *StatusReason

	// The tags added to the analyzer.
	Tags map[string]*string
}

// Contains information about an archive rule.
type ArchiveRuleSummary struct {

	// The time at which the archive rule was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// A filter used to define the archive rule.
	//
	// This member is required.
	Filter map[string]*Criterion

	// The name of the archive rule.
	//
	// This member is required.
	RuleName *string

	// The time at which the archive rule was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time
}

// The criteria to use in the filter that defines the archive rule.
type Criterion struct {

	// A "contains" operator to match for the filter used to create the rule.
	Contains []*string

	// An "equals" operator to match for the filter used to create the rule.
	Eq []*string

	// An "exists" operator to match for the filter used to create the rule.
	Exists *bool

	// A "not equals" operator to match for the filter used to create the rule.
	Neq []*string
}

// Contains information about a finding.
type Finding struct {

	// The time at which the resource was analyzed.
	//
	// This member is required.
	AnalyzedAt *time.Time

	// The condition in the analyzed policy statement that resulted in a finding.
	//
	// This member is required.
	Condition map[string]*string

	// The time at which the finding was generated.
	//
	// This member is required.
	CreatedAt *time.Time

	// The ID of the finding.
	//
	// This member is required.
	Id *string

	// The AWS account ID that owns the resource.
	//
	// This member is required.
	ResourceOwnerAccount *string

	// The type of the resource reported in the finding.
	//
	// This member is required.
	ResourceType ResourceType

	// The current status of the finding.
	//
	// This member is required.
	Status FindingStatus

	// The time at which the finding was updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The action in the analyzed policy statement that an external principal has
	// permission to use.
	Action []*string

	// An error.
	Error *string

	// Indicates whether the policy that generated the finding allows public access to
	// the resource.
	IsPublic *bool

	// The external principal that access to a resource within the zone of trust.
	Principal map[string]*string

	// The resource that an external principal has access to.
	Resource *string

	// The sources of the finding. This indicates how the access that generated the
	// finding is granted. It is populated for Amazon S3 bucket findings.
	Sources []*FindingSource
}

// The source of the finding. This indicates how the access that generated the
// finding is granted. It is populated for Amazon S3 bucket findings.
type FindingSource struct {

	// Indicates the type of access that generated the finding.
	//
	// This member is required.
	Type FindingSourceType

	// Includes details about how the access that generated the finding is granted.
	// This is populated for Amazon S3 bucket findings.
	Detail *FindingSourceDetail
}

// Includes details about how the access that generated the finding is granted.
// This is populated for Amazon S3 bucket findings.
type FindingSourceDetail struct {

	// The ARN of the access point that generated the finding.
	AccessPointArn *string
}

// Contains information about a finding.
type FindingSummary struct {

	// The time at which the resource-based policy that generated the finding was
	// analyzed.
	//
	// This member is required.
	AnalyzedAt *time.Time

	// The condition in the analyzed policy statement that resulted in a finding.
	//
	// This member is required.
	Condition map[string]*string

	// The time at which the finding was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The ID of the finding.
	//
	// This member is required.
	Id *string

	// The AWS account ID that owns the resource.
	//
	// This member is required.
	ResourceOwnerAccount *string

	// The type of the resource that the external principal has access to.
	//
	// This member is required.
	ResourceType ResourceType

	// The status of the finding.
	//
	// This member is required.
	Status FindingStatus

	// The time at which the finding was most recently updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The action in the analyzed policy statement that an external principal has
	// permission to use.
	Action []*string

	// The error that resulted in an Error finding.
	Error *string

	// Indicates whether the finding reports a resource that has a policy that allows
	// public access.
	IsPublic *bool

	// The external principal that has access to a resource within the zone of trust.
	Principal map[string]*string

	// The resource that the external principal has access to.
	Resource *string

	// The sources of the finding. This indicates how the access that generated the
	// finding is granted. It is populated for Amazon S3 bucket findings.
	Sources []*FindingSource
}

// An criterion statement in an archive rule. Each archive rule may have multiple
// criteria.
type InlineArchiveRule struct {

	// The condition and values for a criterion.
	//
	// This member is required.
	Filter map[string]*Criterion

	// The name of the rule.
	//
	// This member is required.
	RuleName *string
}

// The criteria used to sort.
type SortCriteria struct {

	// The name of the attribute to sort on.
	AttributeName *string

	// The sort order, ascending or descending.
	OrderBy OrderBy
}

// Provides more details about the current status of the analyzer. For example, if
// the creation for the analyzer fails, a Failed status is displayed. For an
// analyzer with organization as the type, this failure can be due to an issue with
// creating the service-linked roles required in the member accounts of the AWS
// organization.
type StatusReason struct {

	// The reason code for the current status of the analyzer.
	//
	// This member is required.
	Code ReasonCode
}

// Contains information about a validation exception.
type ValidationExceptionField struct {

	// A message about the validation exception.
	//
	// This member is required.
	Message *string

	// The name of the validation exception.
	//
	// This member is required.
	Name *string
}

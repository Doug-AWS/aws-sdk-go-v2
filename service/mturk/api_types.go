// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// The Assignment data structure represents a single assignment of a HIT to
// a Worker. The assignment tracks the Worker's efforts to complete the HIT,
// and contains the results for later retrieval.
type Assignment struct {
	_ struct{} `type:"structure"`

	// The date and time the Worker accepted the assignment.
	AcceptTime *time.Time `type:"timestamp"`

	// The Worker's answers submitted for the HIT contained in a QuestionFormAnswers
	// document, if the Worker provides an answer. If the Worker does not provide
	// any answers, Answer may contain a QuestionFormAnswers document, or Answer
	// may be empty.
	Answer *string `type:"string"`

	// If the Worker has submitted results and the Requester has approved the results,
	// ApprovalTime is the date and time the Requester approved the results. This
	// value is omitted from the assignment if the Requester has not yet approved
	// the results.
	ApprovalTime *time.Time `type:"timestamp"`

	// A unique identifier for the assignment.
	AssignmentId *string `min:"1" type:"string"`

	// The status of the assignment.
	AssignmentStatus AssignmentStatus `type:"string" enum:"true"`

	// If results have been submitted, AutoApprovalTime is the date and time the
	// results of the assignment results are considered Approved automatically if
	// they have not already been explicitly approved or rejected by the Requester.
	// This value is derived from the auto-approval delay specified by the Requester
	// in the HIT. This value is omitted from the assignment if the Worker has not
	// yet submitted results.
	AutoApprovalTime *time.Time `type:"timestamp"`

	// The date and time of the deadline for the assignment. This value is derived
	// from the deadline specification for the HIT and the date and time the Worker
	// accepted the HIT.
	Deadline *time.Time `type:"timestamp"`

	// The ID of the HIT.
	HITId *string `min:"1" type:"string"`

	// If the Worker has submitted results and the Requester has rejected the results,
	// RejectionTime is the date and time the Requester rejected the results.
	RejectionTime *time.Time `type:"timestamp"`

	// The feedback string included with the call to the ApproveAssignment operation
	// or the RejectAssignment operation, if the Requester approved or rejected
	// the assignment and specified feedback.
	RequesterFeedback *string `type:"string"`

	// If the Worker has submitted results, SubmitTime is the date and time the
	// assignment was submitted. This value is omitted from the assignment if the
	// Worker has not yet submitted results.
	SubmitTime *time.Time `type:"timestamp"`

	// The ID of the Worker who accepted the HIT.
	WorkerId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s Assignment) String() string {
	return awsutil.Prettify(s)
}

// An object representing a Bonus payment paid to a Worker.
type BonusPayment struct {
	_ struct{} `type:"structure"`

	// The ID of the assignment associated with this bonus payment.
	AssignmentId *string `min:"1" type:"string"`

	// A string representing a currency amount.
	BonusAmount *string `type:"string"`

	// The date and time of when the bonus was granted.
	GrantTime *time.Time `type:"timestamp"`

	// The Reason text given when the bonus was granted, if any.
	Reason *string `type:"string"`

	// The ID of the Worker to whom the bonus was paid.
	WorkerId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s BonusPayment) String() string {
	return awsutil.Prettify(s)
}

// The HIT data structure represents a single HIT, including all the information
// necessary for a Worker to accept and complete the HIT.
type HIT struct {
	_ struct{} `type:"structure"`

	// The length of time, in seconds, that a Worker has to complete the HIT after
	// accepting it.
	AssignmentDurationInSeconds *int64 `type:"long"`

	// The amount of time, in seconds, after the Worker submits an assignment for
	// the HIT that the results are automatically approved by Amazon Mechanical
	// Turk. This is the amount of time the Requester has to reject an assignment
	// submitted by a Worker before the assignment is auto-approved and the Worker
	// is paid.
	AutoApprovalDelayInSeconds *int64 `type:"long"`

	// The date and time the HIT was created.
	CreationTime *time.Time `type:"timestamp"`

	// A general description of the HIT.
	Description *string `type:"string"`

	// The date and time the HIT expires.
	Expiration *time.Time `type:"timestamp"`

	// The ID of the HIT Group of this HIT.
	HITGroupId *string `min:"1" type:"string"`

	// A unique identifier for the HIT.
	HITId *string `min:"1" type:"string"`

	// The ID of the HIT Layout of this HIT.
	HITLayoutId *string `min:"1" type:"string"`

	// Indicates the review status of the HIT. Valid Values are NotReviewed | MarkedForReview
	// | ReviewedAppropriate | ReviewedInappropriate.
	HITReviewStatus HITReviewStatus `type:"string" enum:"true"`

	// The status of the HIT and its assignments. Valid Values are Assignable |
	// Unassignable | Reviewable | Reviewing | Disposed.
	HITStatus HITStatus `type:"string" enum:"true"`

	// The ID of the HIT type of this HIT
	HITTypeId *string `min:"1" type:"string"`

	// One or more words or phrases that describe the HIT, separated by commas.
	// Search terms similar to the keywords of a HIT are more likely to have the
	// HIT in the search results.
	Keywords *string `type:"string"`

	// The number of times the HIT can be accepted and completed before the HIT
	// becomes unavailable.
	MaxAssignments *int64 `type:"integer"`

	// The number of assignments for this HIT that are available for Workers to
	// accept.
	NumberOfAssignmentsAvailable *int64 `type:"integer"`

	// The number of assignments for this HIT that have been approved or rejected.
	NumberOfAssignmentsCompleted *int64 `type:"integer"`

	// The number of assignments for this HIT that are being previewed or have been
	// accepted by Workers, but have not yet been submitted, returned, or abandoned.
	NumberOfAssignmentsPending *int64 `type:"integer"`

	// Conditions that a Worker's Qualifications must meet in order to accept the
	// HIT. A HIT can have between zero and ten Qualification requirements. All
	// requirements must be met in order for a Worker to accept the HIT. Additionally,
	// other actions can be restricted using the ActionsGuarded field on each QualificationRequirement
	// structure.
	QualificationRequirements []QualificationRequirement `type:"list"`

	// The data the Worker completing the HIT uses produce the results. This is
	// either either a QuestionForm, HTMLQuestion or an ExternalQuestion data structure.
	Question *string `type:"string"`

	// An arbitrary data field the Requester who created the HIT can use. This field
	// is visible only to the creator of the HIT.
	RequesterAnnotation *string `type:"string"`

	// A string representing a currency amount.
	Reward *string `type:"string"`

	// The title of the HIT.
	Title *string `type:"string"`
}

// String returns the string representation
func (s HIT) String() string {
	return awsutil.Prettify(s)
}

// The HITLayoutParameter data structure defines parameter values used with
// a HITLayout. A HITLayout is a reusable Amazon Mechanical Turk project template
// used to provide Human Intelligence Task (HIT) question data for CreateHIT.
type HITLayoutParameter struct {
	_ struct{} `type:"structure"`

	// The name of the parameter in the HITLayout.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The value substituted for the parameter referenced in the HITLayout.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s HITLayoutParameter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *HITLayoutParameter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "HITLayoutParameter"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The Locale data structure represents a geographical region or location.
type Locale struct {
	_ struct{} `type:"structure"`

	// The country of the locale. Must be a valid ISO 3166 country code. For example,
	// the code US refers to the United States of America.
	//
	// Country is a required field
	Country *string `min:"2" type:"string" required:"true"`

	// The state or subdivision of the locale. A valid ISO 3166-2 subdivision code.
	// For example, the code WA refers to the state of Washington.
	Subdivision *string `min:"2" type:"string"`
}

// String returns the string representation
func (s Locale) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Locale) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Locale"}

	if s.Country == nil {
		invalidParams.Add(aws.NewErrParamRequired("Country"))
	}
	if s.Country != nil && len(*s.Country) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Country", 2))
	}
	if s.Subdivision != nil && len(*s.Subdivision) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Subdivision", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The NotificationSpecification data structure describes a HIT event notification
// for a HIT type.
type NotificationSpecification struct {
	_ struct{} `type:"structure"`

	// The target for notification messages. The Destination’s format is determined
	// by the specified Transport:
	//
	//    * When Transport is Email, the Destination is your email address.
	//
	//    * When Transport is SQS, the Destination is your queue URL.
	//
	//    * When Transport is SNS, the Destination is the ARN of your topic.
	//
	// Destination is a required field
	Destination *string `type:"string" required:"true"`

	// The list of events that should cause notifications to be sent. Valid Values:
	// AssignmentAccepted | AssignmentAbandoned | AssignmentReturned | AssignmentSubmitted
	// | AssignmentRejected | AssignmentApproved | HITCreated | HITExtended | HITDisposed
	// | HITReviewable | HITExpired | Ping. The Ping event is only valid for the
	// SendTestEventNotification operation.
	//
	// EventTypes is a required field
	EventTypes []EventType `type:"list" required:"true"`

	// The method Amazon Mechanical Turk uses to send the notification. Valid Values:
	// Email | SQS | SNS.
	//
	// Transport is a required field
	Transport NotificationTransport `type:"string" required:"true" enum:"true"`

	// The version of the Notification API to use. Valid value is 2006-05-05.
	//
	// Version is a required field
	Version *string `type:"string" required:"true"`
}

// String returns the string representation
func (s NotificationSpecification) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *NotificationSpecification) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "NotificationSpecification"}

	if s.Destination == nil {
		invalidParams.Add(aws.NewErrParamRequired("Destination"))
	}

	if s.EventTypes == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventTypes"))
	}
	if len(s.Transport) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Transport"))
	}

	if s.Version == nil {
		invalidParams.Add(aws.NewErrParamRequired("Version"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// When MTurk encounters an issue with notifying the Workers you specified,
// it returns back this object with failure details.
type NotifyWorkersFailureStatus struct {
	_ struct{} `type:"structure"`

	// Encoded value for the failure type.
	NotifyWorkersFailureCode NotifyWorkersFailureCode `type:"string" enum:"true"`

	// A message detailing the reason the Worker could not be notified.
	NotifyWorkersFailureMessage *string `type:"string"`

	// The ID of the Worker.
	WorkerId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s NotifyWorkersFailureStatus) String() string {
	return awsutil.Prettify(s)
}

// This data structure is the data type for the AnswerKey parameter of the ScoreMyKnownAnswers/2011-09-01
// Review Policy.
type ParameterMapEntry struct {
	_ struct{} `type:"structure"`

	// The QuestionID from the HIT that is used to identify which question requires
	// Mechanical Turk to score as part of the ScoreMyKnownAnswers/2011-09-01 Review
	// Policy.
	Key *string `type:"string"`

	// The list of answers to the question specified in the MapEntry Key element.
	// The Worker must match all values in order for the answer to be scored correctly.
	Values []string `type:"list"`
}

// String returns the string representation
func (s ParameterMapEntry) String() string {
	return awsutil.Prettify(s)
}

// Name of the parameter from the Review policy.
type PolicyParameter struct {
	_ struct{} `type:"structure"`

	// Name of the parameter from the list of Review Polices.
	Key *string `type:"string"`

	// List of ParameterMapEntry objects.
	MapEntries []ParameterMapEntry `type:"list"`

	// The list of values of the Parameter
	Values []string `type:"list"`
}

// String returns the string representation
func (s PolicyParameter) String() string {
	return awsutil.Prettify(s)
}

// The Qualification data structure represents a Qualification assigned to a
// user, including the Qualification type and the value (score).
type Qualification struct {
	_ struct{} `type:"structure"`

	// The date and time the Qualification was granted to the Worker. If the Worker's
	// Qualification was revoked, and then re-granted based on a new Qualification
	// request, GrantTime is the date and time of the last call to the AcceptQualificationRequest
	// operation.
	GrantTime *time.Time `type:"timestamp"`

	// The value (score) of the Qualification, if the Qualification has an integer
	// value.
	IntegerValue *int64 `type:"integer"`

	// The Locale data structure represents a geographical region or location.
	LocaleValue *Locale `type:"structure"`

	// The ID of the Qualification type for the Qualification.
	QualificationTypeId *string `min:"1" type:"string"`

	// The status of the Qualification. Valid values are Granted | Revoked.
	Status QualificationStatus `type:"string" enum:"true"`

	// The ID of the Worker who possesses the Qualification.
	WorkerId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s Qualification) String() string {
	return awsutil.Prettify(s)
}

// The QualificationRequest data structure represents a request a Worker has
// made for a Qualification.
type QualificationRequest struct {
	_ struct{} `type:"structure"`

	// The Worker's answers for the Qualification type's test contained in a QuestionFormAnswers
	// document, if the type has a test and the Worker has submitted answers. If
	// the Worker does not provide any answers, Answer may be empty.
	Answer *string `type:"string"`

	// The ID of the Qualification request, a unique identifier generated when the
	// request was submitted.
	QualificationRequestId *string `type:"string"`

	// The ID of the Qualification type the Worker is requesting, as returned by
	// the CreateQualificationType operation.
	QualificationTypeId *string `min:"1" type:"string"`

	// The date and time the Qualification request had a status of Submitted. This
	// is either the time the Worker submitted answers for a Qualification test,
	// or the time the Worker requested the Qualification if the Qualification type
	// does not have a test.
	SubmitTime *time.Time `type:"timestamp"`

	// The contents of the Qualification test that was presented to the Worker,
	// if the type has a test and the Worker has submitted answers. This value is
	// identical to the QuestionForm associated with the Qualification type at the
	// time the Worker requests the Qualification.
	Test *string `type:"string"`

	// The ID of the Worker requesting the Qualification.
	WorkerId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s QualificationRequest) String() string {
	return awsutil.Prettify(s)
}

// The QualificationRequirement data structure describes a Qualification that
// a Worker must have before the Worker is allowed to accept a HIT. A requirement
// may optionally state that a Worker must have the Qualification in order to
// preview the HIT, or see the HIT in search results.
type QualificationRequirement struct {
	_ struct{} `type:"structure"`

	// Setting this attribute prevents Workers whose Qualifications do not meet
	// this QualificationRequirement from taking the specified action. Valid arguments
	// include "Accept" (Worker cannot accept the HIT, but can preview the HIT and
	// see it in their search results), "PreviewAndAccept" (Worker cannot accept
	// or preview the HIT, but can see the HIT in their search results), and "DiscoverPreviewAndAccept"
	// (Worker cannot accept, preview, or see the HIT in their search results).
	// It's possible for you to create a HIT with multiple QualificationRequirements
	// (which can have different values for the ActionGuarded attribute). In this
	// case, the Worker is only permitted to perform an action when they have met
	// all QualificationRequirements guarding the action. The actions in the order
	// of least restrictive to most restrictive are Discover, Preview and Accept.
	// For example, if a Worker meets all QualificationRequirements that are set
	// to DiscoverPreviewAndAccept, but do not meet all requirements that are set
	// with PreviewAndAccept, then the Worker will be able to Discover, i.e. see
	// the HIT in their search result, but will not be able to Preview or Accept
	// the HIT. ActionsGuarded should not be used in combination with the RequiredToPreview
	// field.
	ActionsGuarded HITAccessActions `type:"string" enum:"true"`

	// The kind of comparison to make against a Qualification's value. You can compare
	// a Qualification's value to an IntegerValue to see if it is LessThan, LessThanOrEqualTo,
	// GreaterThan, GreaterThanOrEqualTo, EqualTo, or NotEqualTo the IntegerValue.
	// You can compare it to a LocaleValue to see if it is EqualTo, or NotEqualTo
	// the LocaleValue. You can check to see if the value is In or NotIn a set of
	// IntegerValue or LocaleValue values. Lastly, a Qualification requirement can
	// also test if a Qualification Exists or DoesNotExist in the user's profile,
	// regardless of its value.
	//
	// Comparator is a required field
	Comparator Comparator `type:"string" required:"true" enum:"true"`

	// The integer value to compare against the Qualification's value. IntegerValue
	// must not be present if Comparator is Exists or DoesNotExist. IntegerValue
	// can only be used if the Qualification type has an integer value; it cannot
	// be used with the Worker_Locale QualificationType ID. When performing a set
	// comparison by using the In or the NotIn comparator, you can use up to 15
	// IntegerValue elements in a QualificationRequirement data structure.
	IntegerValues []int64 `type:"list"`

	// The locale value to compare against the Qualification's value. The local
	// value must be a valid ISO 3166 country code or supports ISO 3166-2 subdivisions.
	// LocaleValue can only be used with a Worker_Locale QualificationType ID. LocaleValue
	// can only be used with the EqualTo, NotEqualTo, In, and NotIn comparators.
	// You must only use a single LocaleValue element when using the EqualTo or
	// NotEqualTo comparators. When performing a set comparison by using the In
	// or the NotIn comparator, you can use up to 30 LocaleValue elements in a QualificationRequirement
	// data structure.
	LocaleValues []Locale `type:"list"`

	// The ID of the Qualification type for the requirement.
	//
	// QualificationTypeId is a required field
	QualificationTypeId *string `type:"string" required:"true"`

	// DEPRECATED: Use the ActionsGuarded field instead. If RequiredToPreview is
	// true, the question data for the HIT will not be shown when a Worker whose
	// Qualifications do not meet this requirement tries to preview the HIT. That
	// is, a Worker's Qualifications must meet all of the requirements for which
	// RequiredToPreview is true in order to preview the HIT. If a Worker meets
	// all of the requirements where RequiredToPreview is true (or if there are
	// no such requirements), but does not meet all of the requirements for the
	// HIT, the Worker will be allowed to preview the HIT's question data, but will
	// not be allowed to accept and complete the HIT. The default is false. This
	// should not be used in combination with the ActionsGuarded field.
	RequiredToPreview *bool `deprecated:"true" type:"boolean"`
}

// String returns the string representation
func (s QualificationRequirement) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *QualificationRequirement) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "QualificationRequirement"}
	if len(s.Comparator) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Comparator"))
	}

	if s.QualificationTypeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("QualificationTypeId"))
	}
	if s.LocaleValues != nil {
		for i, v := range s.LocaleValues {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "LocaleValues", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The QualificationType data structure represents a Qualification type, a description
// of a property of a Worker that must match the requirements of a HIT for the
// Worker to be able to accept the HIT. The type also describes how a Worker
// can obtain a Qualification of that type, such as through a Qualification
// test.
type QualificationType struct {
	_ struct{} `type:"structure"`

	// The answers to the Qualification test specified in the Test parameter.
	AnswerKey *string `type:"string"`

	// Specifies that requests for the Qualification type are granted immediately,
	// without prompting the Worker with a Qualification test. Valid values are
	// True | False.
	AutoGranted *bool `type:"boolean"`

	// The Qualification integer value to use for automatically granted Qualifications,
	// if AutoGranted is true. This is 1 by default.
	AutoGrantedValue *int64 `type:"integer"`

	// The date and time the Qualification type was created.
	CreationTime *time.Time `type:"timestamp"`

	// A long description for the Qualification type.
	Description *string `type:"string"`

	// Specifies whether the Qualification type is one that a user can request through
	// the Amazon Mechanical Turk web site, such as by taking a Qualification test.
	// This value is False for Qualifications assigned automatically by the system.
	// Valid values are True | False.
	IsRequestable *bool `type:"boolean"`

	// One or more words or phrases that describe theQualification type, separated
	// by commas. The Keywords make the type easier to find using a search.
	Keywords *string `type:"string"`

	// The name of the Qualification type. The type name is used to identify the
	// type, and to find the type using a Qualification type search.
	Name *string `type:"string"`

	// A unique identifier for the Qualification type. A Qualification type is given
	// a Qualification type ID when you call the CreateQualificationType operation.
	QualificationTypeId *string `min:"1" type:"string"`

	// The status of the Qualification type. A Qualification type's status determines
	// if users can apply to receive a Qualification of this type, and if HITs can
	// be created with requirements based on this type. Valid values are Active
	// | Inactive.
	QualificationTypeStatus QualificationTypeStatus `type:"string" enum:"true"`

	// The amount of time, in seconds, Workers must wait after taking the Qualification
	// test before they can take it again. Workers can take a Qualification test
	// multiple times if they were not granted the Qualification from a previous
	// attempt, or if the test offers a gradient score and they want a better score.
	// If not specified, retries are disabled and Workers can request a Qualification
	// only once.
	RetryDelayInSeconds *int64 `type:"long"`

	// The questions for a Qualification test associated with this Qualification
	// type that a user can take to obtain a Qualification of this type. This parameter
	// must be specified if AnswerKey is present. A Qualification type cannot have
	// both a specified Test parameter and an AutoGranted value of true.
	Test *string `type:"string"`

	// The amount of time, in seconds, given to a Worker to complete the Qualification
	// test, beginning from the time the Worker requests the Qualification.
	TestDurationInSeconds *int64 `type:"long"`
}

// String returns the string representation
func (s QualificationType) String() string {
	return awsutil.Prettify(s)
}

// Both the AssignmentReviewReport and the HITReviewReport elements contains
// the ReviewActionDetail data structure. This structure is returned multiple
// times for each action specified in the Review Policy.
type ReviewActionDetail struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the action.
	ActionId *string `min:"1" type:"string"`

	// The nature of the action itself. The Review Policy is responsible for examining
	// the HIT and Assignments, emitting results, and deciding which other actions
	// will be necessary.
	ActionName *string `type:"string"`

	// The date when the action was completed.
	CompleteTime *time.Time `type:"timestamp"`

	// Present only when the Results have a FAILED Status.
	ErrorCode *string `type:"string"`

	// A description of the outcome of the review.
	Result *string `type:"string"`

	// The current disposition of the action: INTENDED, SUCCEEDED, FAILED, or CANCELLED.
	Status ReviewActionStatus `type:"string" enum:"true"`

	// The specific HITId or AssignmentID targeted by the action.
	TargetId *string `min:"1" type:"string"`

	// The type of object in TargetId.
	TargetType *string `type:"string"`
}

// String returns the string representation
func (s ReviewActionDetail) String() string {
	return awsutil.Prettify(s)
}

// HIT Review Policy data structures represent HIT review policies, which you
// specify when you create a HIT.
type ReviewPolicy struct {
	_ struct{} `type:"structure"`

	// Name of the parameter from the Review policy.
	Parameters []PolicyParameter `type:"list"`

	// Name of a Review Policy: SimplePlurality/2011-09-01 or ScoreMyKnownAnswers/2011-09-01
	//
	// PolicyName is a required field
	PolicyName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ReviewPolicy) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReviewPolicy) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReviewPolicy"}

	if s.PolicyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains both ReviewResult and ReviewAction elements for a particular HIT.
type ReviewReport struct {
	_ struct{} `type:"structure"`

	// A list of ReviewAction objects for each action specified in the Review Policy.
	ReviewActions []ReviewActionDetail `type:"list"`

	// A list of ReviewResults objects for each action specified in the Review Policy.
	ReviewResults []ReviewResultDetail `type:"list"`
}

// String returns the string representation
func (s ReviewReport) String() string {
	return awsutil.Prettify(s)
}

// This data structure is returned multiple times for each result specified
// in the Review Policy.
type ReviewResultDetail struct {
	_ struct{} `type:"structure"`

	// A unique identifier of the Review action result.
	ActionId *string `min:"1" type:"string"`

	// Key identifies the particular piece of reviewed information.
	Key *string `type:"string"`

	// Specifies the QuestionId the result is describing. Depending on whether the
	// TargetType is a HIT or Assignment this results could specify multiple values.
	// If TargetType is HIT and QuestionId is absent, then the result describes
	// results of the HIT, including the HIT agreement score. If ObjectType is Assignment
	// and QuestionId is absent, then the result describes the Worker's performance
	// on the HIT.
	QuestionId *string `min:"1" type:"string"`

	// The HITID or AssignmentId about which this result was taken. Note that HIT-level
	// Review Policies will often emit results about both the HIT itself and its
	// Assignments, while Assignment-level review policies generally only emit results
	// about the Assignment itself.
	SubjectId *string `min:"1" type:"string"`

	// The type of the object from the SubjectId field.
	SubjectType *string `type:"string"`

	// The values of Key provided by the review policies you have selected.
	Value *string `type:"string"`
}

// String returns the string representation
func (s ReviewResultDetail) String() string {
	return awsutil.Prettify(s)
}

// The WorkerBlock data structure represents a Worker who has been blocked.
// It has two elements: the WorkerId and the Reason for the block.
type WorkerBlock struct {
	_ struct{} `type:"structure"`

	// A message explaining the reason the Worker was blocked.
	Reason *string `type:"string"`

	// The ID of the Worker who accepted the HIT.
	WorkerId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s WorkerBlock) String() string {
	return awsutil.Prettify(s)
}
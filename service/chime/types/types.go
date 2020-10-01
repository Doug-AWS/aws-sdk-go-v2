// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// The Amazon Chime account details. An AWS account can have multiple Amazon Chime
// accounts.
type Account struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The AWS account ID.
	//
	// This member is required.
	AwsAccountId *string

	// The Amazon Chime account name.
	//
	// This member is required.
	Name *string

	// The Amazon Chime account type. For more information about different account
	// types, see Managing Your Amazon Chime Accounts
	// (https://docs.aws.amazon.com/chime/latest/ag/manage-chime-account.html) in the
	// Amazon Chime Administration Guide.
	AccountType AccountType

	// The Amazon Chime account creation timestamp, in ISO 8601 format.
	CreatedTimestamp *time.Time

	// The default license for the Amazon Chime account.
	DefaultLicense License

	// The sign-in delegate groups associated with the account.
	SigninDelegateGroups []*SigninDelegateGroup

	// Supported licenses for the Amazon Chime account.
	SupportedLicenses []License
}

// Settings related to the Amazon Chime account. This includes settings that start
// or stop remote control of shared screens, or start or stop the dial-out option
// in the Amazon Chime web application. For more information about these settings,
// see Use the Policies Page
// (https://docs.aws.amazon.com/chime/latest/ag/policies.html) in the Amazon Chime
// Administration Guide.
type AccountSettings struct {

	// Setting that stops or starts remote control of shared screens during meetings.
	DisableRemoteControl *bool

	// Setting that allows meeting participants to choose the Call me at a phone number
	// option. For more information, see Join a Meeting without the Amazon Chime App
	// (https://docs.aws.amazon.com/chime/latest/ug/chime-join-meeting.html).
	EnableDialOut *bool
}

// The Alexa for Business metadata associated with an Amazon Chime user, used to
// integrate Alexa for Business with a device.
type AlexaForBusinessMetadata struct {

	// The ARN of the room resource.
	AlexaForBusinessRoomArn *string

	// Starts or stops Alexa for Business.
	IsAlexaForBusinessEnabled *bool
}

// An Amazon Chime SDK meeting attendee. Includes a unique AttendeeId and
// JoinToken. The JoinToken allows a client to authenticate and join as the
// specified attendee. The JoinToken expires when the meeting ends or when
// DeleteAttendee () is called. After that, the attendee is unable to join the
// meeting. We recommend securely transferring each JoinToken from your server
// application to the client so that no other client has access to the token except
// for the one authorized to represent the attendee.
type Attendee struct {

	// The Amazon Chime SDK attendee ID.
	AttendeeId *string

	// The Amazon Chime SDK external user ID. Links the attendee to an identity managed
	// by a builder application.
	ExternalUserId *string

	// The join token used by the Amazon Chime SDK attendee.
	JoinToken *string
}

// A resource that allows Enterprise account administrators to configure an
// interface to receive events from Amazon Chime.
type Bot struct {

	// The bot email address.
	BotEmail *string

	// The bot ID.
	BotId *string

	// The bot type.
	BotType BotType

	// The bot creation timestamp, in ISO 8601 format.
	CreatedTimestamp *time.Time

	// When true, the bot is stopped from running in your account.
	Disabled *bool

	// The bot display name.
	DisplayName *string

	// The security token used to authenticate Amazon Chime with the outgoing event
	// endpoint.
	SecurityToken *string

	// The updated bot timestamp, in ISO 8601 format.
	UpdatedTimestamp *time.Time

	// The unique ID for the bot user.
	UserId *string
}

// The Amazon Chime Business Calling settings for the administrator's AWS account.
// Includes any Amazon S3 buckets designated for storing call detail records.
type BusinessCallingSettings struct {

	// The Amazon S3 bucket designated for call detail record storage.
	CdrBucket *string
}

// The retention settings that determine how long to retain chat conversation
// messages for an Amazon Chime Enterprise account.
type ConversationRetentionSettings struct {

	// The number of days for which to retain chat conversation messages.
	RetentionDays *int32
}

// The list of errors returned when errors are encountered during the
// BatchCreateAttendee and CreateAttendee actions. This includes external user IDs,
// error codes, and error messages.
type CreateAttendeeError struct {

	// The error code.
	ErrorCode *string

	// The error message.
	ErrorMessage *string

	// The Amazon Chime SDK external user ID. Links the attendee to an identity managed
	// by a builder application.
	ExternalUserId *string
}

// The Amazon Chime SDK attendee fields to create, used with the
// BatchCreateAttendee action.
type CreateAttendeeRequestItem struct {

	// The Amazon Chime SDK external user ID. Links the attendee to an identity managed
	// by a builder application.
	//
	// This member is required.
	ExternalUserId *string

	// The tag key-value pairs.
	Tags []*Tag
}

// The SIP credentials used to authenticate requests to your Amazon Chime Voice
// Connector.
type Credential struct {

	// The RFC2617 compliant password associated with the SIP credentials, in US-ASCII
	// format.
	Password *string

	// The RFC2617 compliant user name associated with the SIP credentials, in US-ASCII
	// format.
	Username *string
}

// The Dialed Number Identification Service (DNIS) emergency calling configuration
// details associated with an Amazon Chime Voice Connector's emergency calling
// configuration.
type DNISEmergencyCallingConfiguration struct {

	// The country from which emergency calls are allowed, in ISO 3166-1 alpha-2
	// format.
	//
	// This member is required.
	CallingCountry *string

	// The DNIS phone number to route emergency calls to, in E.164 format.
	//
	// This member is required.
	EmergencyPhoneNumber *string

	// The DNIS phone number to route test emergency calls to, in E.164 format.
	TestPhoneNumber *string
}

// The emergency calling configuration details associated with an Amazon Chime
// Voice Connector.
type EmergencyCallingConfiguration struct {

	// The Dialed Number Identification Service (DNIS) emergency calling configuration
	// details.
	DNIS []*DNISEmergencyCallingConfiguration
}

// The configuration that allows a bot to receive outgoing events. Can be either an
// HTTPS endpoint or a Lambda function ARN.
type EventsConfiguration struct {

	// The bot ID.
	BotId *string

	// Lambda function ARN that allows a bot to receive outgoing events.
	LambdaFunctionArn *string

	// HTTPS endpoint that allows a bot to receive outgoing events.
	OutboundEventsHTTPSEndpoint *string
}

// The country and area code for a proxy phone number in a proxy phone session.
type GeoMatchParams struct {

	// The area code.
	//
	// This member is required.
	AreaCode *string

	// The country.
	//
	// This member is required.
	Country *string
}

// Invitation object returned after emailing users to invite them to join the
// Amazon Chime Team account.
type Invite struct {

	// The email address to which the invite is sent.
	EmailAddress *string

	// The status of the invite email.
	EmailStatus EmailStatus

	// The invite ID.
	InviteId *string

	// The status of the invite.
	Status InviteStatus
}

// The logging configuration associated with an Amazon Chime Voice Connector.
// Specifies whether SIP message logs are enabled for sending to Amazon CloudWatch
// Logs.
type LoggingConfiguration struct {

	// When true, enables SIP message logs for sending to Amazon CloudWatch Logs.
	EnableSIPLogs *bool
}

// A set of endpoints used by clients to connect to the media service group for a
// Amazon Chime SDK meeting.
type MediaPlacement struct {

	// The audio fallback URL.
	AudioFallbackUrl *string

	// The audio host URL.
	AudioHostUrl *string

	// The screen data URL.
	ScreenDataUrl *string

	// The screen sharing URL.
	ScreenSharingUrl *string

	// The screen viewing URL.
	ScreenViewingUrl *string

	// The signaling URL.
	SignalingUrl *string

	// The turn control URL.
	TurnControlUrl *string
}

// A meeting created using the Amazon Chime SDK.
type Meeting struct {

	// The external meeting ID.
	ExternalMeetingId *string

	// The media placement for the meeting.
	MediaPlacement *MediaPlacement

	// The Region in which to create the meeting. Available values: ap-northeast-1,
	// ap-southeast-1, ap-southeast-2, ca-central-1, eu-central-1, eu-north-1,
	// eu-west-1, eu-west-2, eu-west-3, sa-east-1, us-east-1, us-east-2, us-west-1,
	// us-west-2.
	MediaRegion *string

	// The Amazon Chime SDK meeting ID.
	MeetingId *string
}

// The configuration for resource targets to receive notifications when Amazon
// Chime SDK meeting and attendee events occur. The Amazon Chime SDK supports
// resource targets located in the US East (N. Virginia) AWS Region (us-east-1).
type MeetingNotificationConfiguration struct {

	// The SNS topic ARN.
	SnsTopicArn *string

	// The SQS queue ARN.
	SqsQueueArn *string
}

// The member details, such as email address, name, member ID, and member type.
type Member struct {

	// The Amazon Chime account ID.
	AccountId *string

	// The member email address.
	Email *string

	// The member name.
	FullName *string

	// The member ID (user ID or bot ID).
	MemberId *string

	// The member type.
	MemberType MemberType
}

// The list of errors returned when a member action results in an error.
type MemberError struct {

	// The error code.
	ErrorCode ErrorCode

	// The error message.
	ErrorMessage *string

	// The member ID.
	MemberId *string
}

// Membership details, such as member ID and member role.
type MembershipItem struct {

	// The member ID.
	MemberId *string

	// The member role.
	Role RoomMembershipRole
}

// A phone number for which an order has been placed.
type OrderedPhoneNumber struct {

	// The phone number, in E.164 format.
	E164PhoneNumber *string

	// The phone number status.
	Status OrderedPhoneNumberStatus
}

// Origination settings enable your SIP hosts to receive inbound calls using your
// Amazon Chime Voice Connector.
type Origination struct {

	// When origination settings are disabled, inbound calls are not enabled for your
	// Amazon Chime Voice Connector.
	Disabled *bool

	// The call distribution properties defined for your SIP hosts. Valid range:
	// Minimum value of 1. Maximum value of 20.
	Routes []*OriginationRoute
}

// Origination routes define call distribution properties for your SIP hosts to
// receive inbound calls using your Amazon Chime Voice Connector. Limit: Ten
// origination routes for each Amazon Chime Voice Connector.
type OriginationRoute struct {

	// The FQDN or IP address to contact for origination traffic.
	Host *string

	// The designated origination route port. Defaults to 5060.
	Port *int32

	// The priority associated with the host, with 1 being the highest priority. Higher
	// priority hosts are attempted first.
	Priority *int32

	// The protocol to use for the origination route. Encryption-enabled Amazon Chime
	// Voice Connectors use TCP protocol by default.
	Protocol OriginationRouteProtocol

	// The weight associated with the host. If hosts are equal in priority, calls are
	// distributed among them based on their relative weight.
	Weight *int32
}

// The phone number and proxy phone number for a participant in an Amazon Chime
// Voice Connector proxy session.
type Participant struct {

	// The participant's phone number.
	PhoneNumber *string

	// The participant's proxy phone number.
	ProxyPhoneNumber *string
}

// A phone number used for Amazon Chime Business Calling or an Amazon Chime Voice
// Connector.
type PhoneNumber struct {

	// The phone number associations.
	Associations []*PhoneNumberAssociation

	// The outbound calling name associated with the phone number.
	CallingName *string

	// The outbound calling name status.
	CallingNameStatus CallingNameStatus

	// The phone number capabilities.
	Capabilities *PhoneNumberCapabilities

	// The phone number creation timestamp, in ISO 8601 format.
	CreatedTimestamp *time.Time

	// The deleted phone number timestamp, in ISO 8601 format.
	DeletionTimestamp *time.Time

	// The phone number, in E.164 format.
	E164PhoneNumber *string

	// The phone number ID.
	PhoneNumberId *string

	// The phone number product type.
	ProductType PhoneNumberProductType

	// The phone number status.
	Status PhoneNumberStatus

	// The phone number type.
	Type PhoneNumberType

	// The updated phone number timestamp, in ISO 8601 format.
	UpdatedTimestamp *time.Time
}

// The phone number associations, such as Amazon Chime account ID, Amazon Chime
// user ID, Amazon Chime Voice Connector ID, or Amazon Chime Voice Connector group
// ID.
type PhoneNumberAssociation struct {

	// The timestamp of the phone number association, in ISO 8601 format.
	AssociatedTimestamp *time.Time

	// Defines the association with an Amazon Chime account ID, user ID, Amazon Chime
	// Voice Connector ID, or Amazon Chime Voice Connector group ID.
	Name PhoneNumberAssociationName

	// Contains the ID for the entity specified in Name.
	Value *string
}

// The phone number capabilities for Amazon Chime Business Calling phone numbers,
// such as enabled inbound and outbound calling and text messaging.
type PhoneNumberCapabilities struct {

	// Allows or denies inbound calling for the specified phone number.
	InboundCall *bool

	// Allows or denies inbound MMS messaging for the specified phone number.
	InboundMMS *bool

	// Allows or denies inbound SMS messaging for the specified phone number.
	InboundSMS *bool

	// Allows or denies outbound calling for the specified phone number.
	OutboundCall *bool

	// Allows or denies outbound MMS messaging for the specified phone number.
	OutboundMMS *bool

	// Allows or denies outbound SMS messaging for the specified phone number.
	OutboundSMS *bool
}

// If the phone number action fails for one or more of the phone numbers in the
// request, a list of the phone numbers is returned, along with error codes and
// error messages.
type PhoneNumberError struct {

	// The error code.
	ErrorCode ErrorCode

	// The error message.
	ErrorMessage *string

	// The phone number ID for which the action failed.
	PhoneNumberId *string
}

// The details of a phone number order created for Amazon Chime.
type PhoneNumberOrder struct {

	// The phone number order creation timestamp, in ISO 8601 format.
	CreatedTimestamp *time.Time

	// The ordered phone number details, such as the phone number in E.164 format and
	// the phone number status.
	OrderedPhoneNumbers []*OrderedPhoneNumber

	// The phone number order ID.
	PhoneNumberOrderId *string

	// The phone number order product type.
	ProductType PhoneNumberProductType

	// The status of the phone number order.
	Status PhoneNumberOrderStatus

	// The updated phone number order timestamp, in ISO 8601 format.
	UpdatedTimestamp *time.Time
}

// The proxy configuration for an Amazon Chime Voice Connector.
type Proxy struct {

	// The default number of minutes allowed for proxy sessions.
	DefaultSessionExpiryMinutes *int32

	// When true, stops proxy sessions from being created on the specified Amazon Chime
	// Voice Connector.
	Disabled *bool

	// The phone number to route calls to after a proxy session expires.
	FallBackPhoneNumber *string

	// The countries for proxy phone numbers to be selected from.
	PhoneNumberCountries []*string
}

// The proxy session for an Amazon Chime Voice Connector.
type ProxySession struct {

	// The proxy session capabilities.
	Capabilities []Capability

	// The created timestamp, in ISO 8601 format.
	CreatedTimestamp *time.Time

	// The ended timestamp, in ISO 8601 format.
	EndedTimestamp *time.Time

	// The number of minutes allowed for the proxy session.
	ExpiryMinutes *int32

	// The preference for matching the country or area code of the proxy phone number
	// with that of the first participant.
	GeoMatchLevel GeoMatchLevel

	// The country and area code for the proxy phone number.
	GeoMatchParams *GeoMatchParams

	// The name of the proxy session.
	Name *string

	// The preference for proxy phone number reuse, or stickiness, between the same
	// participants across sessions.
	NumberSelectionBehavior NumberSelectionBehavior

	// The proxy session participants.
	Participants []*Participant

	// The proxy session ID.
	ProxySessionId *string

	// The status of the proxy session.
	Status ProxySessionStatus

	// The updated timestamp, in ISO 8601 format.
	UpdatedTimestamp *time.Time

	// The Amazon Chime voice connector ID.
	VoiceConnectorId *string
}

// The retention settings for an Amazon Chime Enterprise account that determine how
// long to retain items such as chat room messages and chat conversation messages.
type RetentionSettings struct {

	// The chat conversation retention settings.
	ConversationRetentionSettings *ConversationRetentionSettings

	// The chat room retention settings.
	RoomRetentionSettings *RoomRetentionSettings
}

// The Amazon Chime chat room details.
type Room struct {

	// The Amazon Chime account ID.
	AccountId *string

	// The identifier of the room creator.
	CreatedBy *string

	// The room creation timestamp, in ISO 8601 format.
	CreatedTimestamp *time.Time

	// The room name.
	Name *string

	// The room ID.
	RoomId *string

	// The room update timestamp, in ISO 8601 format.
	UpdatedTimestamp *time.Time
}

// The room membership details.
type RoomMembership struct {

	// The identifier of the user that invited the room member.
	InvitedBy *string

	// The member details, such as email address, name, member ID, and member type.
	Member *Member

	// The membership role.
	Role RoomMembershipRole

	// The room ID.
	RoomId *string

	// The room membership update timestamp, in ISO 8601 format.
	UpdatedTimestamp *time.Time
}

// The retention settings that determine how long to retain chat room messages for
// an Amazon Chime Enterprise account.
type RoomRetentionSettings struct {

	// The number of days for which to retain chat room messages.
	RetentionDays *int32
}

// An Active Directory (AD) group whose members are granted permission to act as
// delegates.
type SigninDelegateGroup struct {

	// The group name.
	GroupName *string
}

// The streaming configuration associated with an Amazon Chime Voice Connector.
// Specifies whether media streaming is enabled for sending to Amazon Kinesis, and
// shows the retention period for the Amazon Kinesis data, in hours.
type StreamingConfiguration struct {

	// The retention period, in hours, for the Amazon Kinesis data.
	//
	// This member is required.
	DataRetentionInHours *int32

	// When true, media streaming to Amazon Kinesis is turned off.
	Disabled *bool

	// The streaming notification targets.
	StreamingNotificationTargets []*StreamingNotificationTarget
}

// The targeted recipient for a streaming configuration notification.
type StreamingNotificationTarget struct {

	// The streaming notification target.
	//
	// This member is required.
	NotificationTarget NotificationTarget
}

// Describes a tag applied to a resource.
type Tag struct {

	// The key of the tag.
	//
	// This member is required.
	Key *string

	// The value of the tag.
	//
	// This member is required.
	Value *string
}

// Settings that allow management of telephony permissions for an Amazon Chime
// user, such as inbound and outbound calling and text messaging.
type TelephonySettings struct {

	// Allows or denies inbound calling.
	//
	// This member is required.
	InboundCalling *bool

	// Allows or denies outbound calling.
	//
	// This member is required.
	OutboundCalling *bool

	// Allows or denies SMS messaging.
	//
	// This member is required.
	SMS *bool
}

// Termination settings enable your SIP hosts to make outbound calls using your
// Amazon Chime Voice Connector.
type Termination struct {

	// The countries to which calls are allowed, in ISO 3166-1 alpha-2 format.
	// Required.
	CallingRegions []*string

	// The IP addresses allowed to make calls, in CIDR format. Required.
	CidrAllowedList []*string

	// The limit on calls per second. Max value based on account service quota. Default
	// value of 1.
	CpsLimit *int32

	// The default caller ID phone number.
	DefaultPhoneNumber *string

	// When termination settings are disabled, outbound calls can not be made.
	Disabled *bool
}

// The termination health details, including the source IP address and timestamp of
// the last successful SIP OPTIONS message from your SIP infrastructure.
type TerminationHealth struct {

	// The source IP address.
	Source *string

	// The timestamp, in ISO 8601 format.
	Timestamp *time.Time
}

// The phone number ID, product type, or calling name fields to update, used with
// the BatchUpdatePhoneNumber () and UpdatePhoneNumber () actions.
type UpdatePhoneNumberRequestItem struct {

	// The phone number ID to update.
	//
	// This member is required.
	PhoneNumberId *string

	// The outbound calling name to update.
	CallingName *string

	// The product type to update.
	ProductType PhoneNumberProductType
}

// The user ID and user fields to update, used with the BatchUpdateUser () action.
type UpdateUserRequestItem struct {

	// The user ID.
	//
	// This member is required.
	UserId *string

	// The Alexa for Business metadata.
	AlexaForBusinessMetadata *AlexaForBusinessMetadata

	// The user license type.
	LicenseType License

	// The user type.
	UserType UserType
}

// The user on the Amazon Chime account.
type User struct {

	// The user ID.
	//
	// This member is required.
	UserId *string

	// The Amazon Chime account ID.
	AccountId *string

	// The Alexa for Business metadata.
	AlexaForBusinessMetadata *AlexaForBusinessMetadata

	// The display name of the user.
	DisplayName *string

	// Date and time when the user is invited to the Amazon Chime account, in ISO 8601
	// format.
	InvitedOn *time.Time

	// The license type for the user.
	LicenseType License

	// The user's personal meeting PIN.
	PersonalPIN *string

	// The primary email address of the user.
	PrimaryEmail *string

	// The primary phone number associated with the user.
	PrimaryProvisionedNumber *string

	// Date and time when the user is registered, in ISO 8601 format.
	RegisteredOn *time.Time

	// The user invite status.
	UserInvitationStatus InviteStatus

	// The user registration status.
	UserRegistrationStatus RegistrationStatus

	// The user type.
	UserType UserType
}

// The list of errors returned when errors are encountered during the
// BatchSuspendUser (), BatchUnsuspendUser (), or BatchUpdateUser () actions. This
// includes user IDs, error codes, and error messages.
type UserError struct {

	// The error code.
	ErrorCode ErrorCode

	// The error message.
	ErrorMessage *string

	// The user ID for which the action failed.
	UserId *string
}

// Settings associated with an Amazon Chime user, including inbound and outbound
// calling and text messaging.
type UserSettings struct {

	// The telephony settings associated with the user.
	//
	// This member is required.
	Telephony *TelephonySettings
}

// The Amazon Chime Voice Connector configuration, including outbound host name and
// encryption settings.
type VoiceConnector struct {

	// The AWS Region in which the Amazon Chime Voice Connector is created. Default:
	// us-east-1.
	AwsRegion VoiceConnectorAwsRegion

	// The Amazon Chime Voice Connector creation timestamp, in ISO 8601 format.
	CreatedTimestamp *time.Time

	// The name of the Amazon Chime Voice Connector.
	Name *string

	// The outbound host name for the Amazon Chime Voice Connector.
	OutboundHostName *string

	// Designates whether encryption is required for the Amazon Chime Voice Connector.
	RequireEncryption *bool

	// The updated Amazon Chime Voice Connector timestamp, in ISO 8601 format.
	UpdatedTimestamp *time.Time

	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId *string
}

// The Amazon Chime Voice Connector group configuration, including associated
// Amazon Chime Voice Connectors. You can include Amazon Chime Voice Connectors
// from different AWS Regions in your group. This creates a fault tolerant
// mechanism for fallback in case of availability events.
type VoiceConnectorGroup struct {

	// The Amazon Chime Voice Connector group creation timestamp, in ISO 8601 format.
	CreatedTimestamp *time.Time

	// The name of the Amazon Chime Voice Connector group.
	Name *string

	// The updated Amazon Chime Voice Connector group timestamp, in ISO 8601 format.
	UpdatedTimestamp *time.Time

	// The Amazon Chime Voice Connector group ID.
	VoiceConnectorGroupId *string

	// The Amazon Chime Voice Connectors to which to route inbound calls.
	VoiceConnectorItems []*VoiceConnectorItem
}

// For Amazon Chime Voice Connector groups, the Amazon Chime Voice Connectors to
// which to route inbound calls. Includes priority configuration settings. Limit: 3
// VoiceConnectorItems per Amazon Chime Voice Connector group.
type VoiceConnectorItem struct {

	// The priority associated with the Amazon Chime Voice Connector, with 1 being the
	// highest priority. Higher priority Amazon Chime Voice Connectors are attempted
	// first.
	//
	// This member is required.
	Priority *int32

	// The Amazon Chime Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string
}

// The Amazon Chime Voice Connector settings. Includes any Amazon S3 buckets
// designated for storing call detail records.
type VoiceConnectorSettings struct {

	// The Amazon S3 bucket designated for call detail record storage.
	CdrBucket *string
}

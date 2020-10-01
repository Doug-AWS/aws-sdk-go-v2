// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// CDN Authorization credentials
type Authorization struct {

	// The Amazon Resource Name (ARN) for the secret in Secrets Manager that your
	// Content Distribution Network (CDN) uses for authorization to access your
	// endpoint.
	//
	// This member is required.
	CdnIdentifierSecret *string

	// The Amazon Resource Name (ARN) for the IAM role that allows MediaPackage to
	// communicate with AWS Secrets Manager.
	//
	// This member is required.
	SecretsRoleArn *string
}

// A Channel resource configuration.
type Channel struct {

	// The Amazon Resource Name (ARN) assigned to the Channel.
	Arn *string

	// A short text description of the Channel.
	Description *string

	// An HTTP Live Streaming (HLS) ingest resource configuration.
	HlsIngest *HlsIngest

	// The ID of the Channel.
	Id *string

	// A collection of tags associated with a resource
	Tags map[string]*string
}

// A Common Media Application Format (CMAF) encryption configuration.
type CmafEncryption struct {

	// A configuration for accessing an external Secure Packager and Encoder Key
	// Exchange (SPEKE) service that will provide encryption keys.
	//
	// This member is required.
	SpekeKeyProvider *SpekeKeyProvider

	// Time (in seconds) between each encryption key rotation.
	KeyRotationIntervalSeconds *int32
}

// A Common Media Application Format (CMAF) packaging configuration.
type CmafPackage struct {

	// A Common Media Application Format (CMAF) encryption configuration.
	Encryption *CmafEncryption

	// A list of HLS manifest configurations
	HlsManifests []*HlsManifest

	// Duration (in seconds) of each segment. Actual segments will be rounded to the
	// nearest multiple of the source segment duration.
	SegmentDurationSeconds *int32

	// An optional custom string that is prepended to the name of each segment. If not
	// specified, it defaults to the ChannelId.
	SegmentPrefix *string

	// A StreamSelection configuration.
	StreamSelection *StreamSelection
}

// A Common Media Application Format (CMAF) packaging configuration.
type CmafPackageCreateOrUpdateParameters struct {

	// A Common Media Application Format (CMAF) encryption configuration.
	Encryption *CmafEncryption

	// A list of HLS manifest configurations
	HlsManifests []*HlsManifestCreateOrUpdateParameters

	// Duration (in seconds) of each segment. Actual segments will be rounded to the
	// nearest multiple of the source segment duration.
	SegmentDurationSeconds *int32

	// An optional custom string that is prepended to the name of each segment. If not
	// specified, it defaults to the ChannelId.
	SegmentPrefix *string

	// A StreamSelection configuration.
	StreamSelection *StreamSelection
}

// A Dynamic Adaptive Streaming over HTTP (DASH) encryption configuration.
type DashEncryption struct {

	// A configuration for accessing an external Secure Packager and Encoder Key
	// Exchange (SPEKE) service that will provide encryption keys.
	//
	// This member is required.
	SpekeKeyProvider *SpekeKeyProvider

	// Time (in seconds) between each encryption key rotation.
	KeyRotationIntervalSeconds *int32
}

// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
type DashPackage struct {

	// A list of SCTE-35 message types that are treated as ad markers in the output. If
	// empty, no ad markers are output. Specify multiple items to create ad markers for
	// all of the included message types.
	AdTriggers []AdTriggersElement

	// This setting allows the delivery restriction flags on SCTE-35 segmentation
	// descriptors to determine whether a message signals an ad. Choosing "NONE" means
	// no SCTE-35 messages become ads. Choosing "RESTRICTED" means SCTE-35 messages of
	// the types specified in AdTriggers that contain delivery restrictions will be
	// treated as ads. Choosing "UNRESTRICTED" means SCTE-35 messages of the types
	// specified in AdTriggers that do not contain delivery restrictions will be
	// treated as ads. Choosing "BOTH" means all SCTE-35 messages of the types
	// specified in AdTriggers will be treated as ads. Note that Splice Insert messages
	// do not have these flags and are always treated as ads if specified in
	// AdTriggers.
	AdsOnDeliveryRestrictions AdsOnDeliveryRestrictions

	// A Dynamic Adaptive Streaming over HTTP (DASH) encryption configuration.
	Encryption *DashEncryption

	// Determines the position of some tags in the Media Presentation Description
	// (MPD). When set to FULL, elements like SegmentTemplate and ContentProtection are
	// included in each Representation. When set to COMPACT, duplicate elements are
	// combined and presented at the AdaptationSet level.
	ManifestLayout ManifestLayout

	// Time window (in seconds) contained in each manifest.
	ManifestWindowSeconds *int32

	// Minimum duration (in seconds) that a player will buffer media before starting
	// the presentation.
	MinBufferTimeSeconds *int32

	// Minimum duration (in seconds) between potential changes to the Dynamic Adaptive
	// Streaming over HTTP (DASH) Media Presentation Description (MPD).
	MinUpdatePeriodSeconds *int32

	// A list of triggers that controls when the outgoing Dynamic Adaptive Streaming
	// over HTTP (DASH) Media Presentation Description (MPD) will be partitioned into
	// multiple periods. If empty, the content will not be partitioned into more than
	// one period. If the list contains "ADS", new periods will be created where the
	// Channel source contains SCTE-35 ad markers.
	PeriodTriggers []PeriodTriggersElement

	// The Dynamic Adaptive Streaming over HTTP (DASH) profile type. When set to
	// "HBBTV_1_5", HbbTV 1.5 compliant output is enabled.
	Profile Profile

	// Duration (in seconds) of each segment. Actual segments will be rounded to the
	// nearest multiple of the source segment duration.
	SegmentDurationSeconds *int32

	// Determines the type of SegmentTemplate included in the Media Presentation
	// Description (MPD). When set to NUMBER_WITH_TIMELINE, a full timeline is
	// presented in each SegmentTemplate, with $Number$ media URLs. When set to
	// TIME_WITH_TIMELINE, a full timeline is presented in each SegmentTemplate, with
	// $Time$ media URLs. When set to NUMBER_WITH_DURATION, only a duration is included
	// in each SegmentTemplate, with $Number$ media URLs.
	SegmentTemplateFormat SegmentTemplateFormat

	// A StreamSelection configuration.
	StreamSelection *StreamSelection

	// Duration (in seconds) to delay live content before presentation.
	SuggestedPresentationDelaySeconds *int32
}

// A HarvestJob resource configuration
type HarvestJob struct {

	// The Amazon Resource Name (ARN) assigned to the HarvestJob.
	Arn *string

	// The ID of the Channel that the HarvestJob will harvest from.
	ChannelId *string

	// The time the HarvestJob was submitted
	CreatedAt *string

	// The end of the time-window which will be harvested.
	EndTime *string

	// The ID of the HarvestJob. The ID must be unique within the region and it cannot
	// be changed after the HarvestJob is submitted.
	Id *string

	// The ID of the OriginEndpoint that the HarvestJob will harvest from. This cannot
	// be changed after the HarvestJob is submitted.
	OriginEndpointId *string

	// Configuration parameters for where in an S3 bucket to place the harvested
	// content
	S3Destination *S3Destination

	// The start of the time-window which will be harvested.
	StartTime *string

	// The current status of the HarvestJob. Consider setting up a CloudWatch Event to
	// listen for HarvestJobs as they succeed or fail. In the event of failure, the
	// CloudWatch Event will include an explanation of why the HarvestJob failed.
	Status Status
}

// An HTTP Live Streaming (HLS) encryption configuration.
type HlsEncryption struct {

	// A configuration for accessing an external Secure Packager and Encoder Key
	// Exchange (SPEKE) service that will provide encryption keys.
	//
	// This member is required.
	SpekeKeyProvider *SpekeKeyProvider

	// A constant initialization vector for encryption (optional). When not specified
	// the initialization vector will be periodically rotated.
	ConstantInitializationVector *string

	// The encryption method to use.
	EncryptionMethod EncryptionMethod

	// Interval (in seconds) between each encryption key rotation.
	KeyRotationIntervalSeconds *int32

	// When enabled, the EXT-X-KEY tag will be repeated in output manifests.
	RepeatExtXKey *bool
}

// An HTTP Live Streaming (HLS) ingest resource configuration.
type HlsIngest struct {

	// A list of endpoints to which the source stream should be sent.
	IngestEndpoints []*IngestEndpoint
}

// A HTTP Live Streaming (HLS) manifest configuration.
type HlsManifest struct {

	// The ID of the manifest. The ID must be unique within the OriginEndpoint and it
	// cannot be changed after it is created.
	//
	// This member is required.
	Id *string

	// This setting controls how ad markers are included in the packaged
	// OriginEndpoint. "NONE" will omit all SCTE-35 ad markers from the output.
	// "PASSTHROUGH" causes the manifest to contain a copy of the SCTE-35 ad markers
	// (comments) taken directly from the input HTTP Live Streaming (HLS) manifest.
	// "SCTE35_ENHANCED" generates ad markers and blackout tags based on SCTE-35
	// messages in the input source. "DATERANGE" inserts EXT-X-DATERANGE tags to signal
	// ad and program transition events in HLS and CMAF manifests. For this option, you
	// must set a programDateTimeIntervalSeconds value that is greater than 0.
	AdMarkers AdMarkers

	// When enabled, an I-Frame only stream will be included in the output.
	IncludeIframeOnlyStream *bool

	// An optional short string appended to the end of the OriginEndpoint URL. If not
	// specified, defaults to the manifestName for the OriginEndpoint.
	ManifestName *string

	// The HTTP Live Streaming (HLS) playlist type. When either "EVENT" or "VOD" is
	// specified, a corresponding EXT-X-PLAYLIST-TYPE entry will be included in the
	// media playlist.
	PlaylistType PlaylistType

	// Time window (in seconds) contained in each parent manifest.
	PlaylistWindowSeconds *int32

	// The interval (in seconds) between each EXT-X-PROGRAM-DATE-TIME tag inserted into
	// manifests. Additionally, when an interval is specified ID3Timed Metadata
	// messages will be generated every 5 seconds using the ingest time of the content.
	// If the interval is not specified, or set to 0, then no EXT-X-PROGRAM-DATE-TIME
	// tags will be inserted into manifests and no ID3Timed Metadata messages will be
	// generated. Note that irrespective of this parameter, if any ID3 Timed Metadata
	// is found in HTTP Live Streaming (HLS) input, it will be passed through to HLS
	// output.
	ProgramDateTimeIntervalSeconds *int32

	// The URL of the packaged OriginEndpoint for consumption.
	Url *string
}

// A HTTP Live Streaming (HLS) manifest configuration.
type HlsManifestCreateOrUpdateParameters struct {

	// The ID of the manifest. The ID must be unique within the OriginEndpoint and it
	// cannot be changed after it is created.
	//
	// This member is required.
	Id *string

	// This setting controls how ad markers are included in the packaged
	// OriginEndpoint. "NONE" will omit all SCTE-35 ad markers from the output.
	// "PASSTHROUGH" causes the manifest to contain a copy of the SCTE-35 ad markers
	// (comments) taken directly from the input HTTP Live Streaming (HLS) manifest.
	// "SCTE35_ENHANCED" generates ad markers and blackout tags based on SCTE-35
	// messages in the input source. "DATERANGE" inserts EXT-X-DATERANGE tags to signal
	// ad and program transition events in HLS and CMAF manifests. For this option, you
	// must set a programDateTimeIntervalSeconds value that is greater than 0.
	AdMarkers AdMarkers

	// A list of SCTE-35 message types that are treated as ad markers in the output. If
	// empty, no ad markers are output. Specify multiple items to create ad markers for
	// all of the included message types.
	AdTriggers []AdTriggersElement

	// This setting allows the delivery restriction flags on SCTE-35 segmentation
	// descriptors to determine whether a message signals an ad. Choosing "NONE" means
	// no SCTE-35 messages become ads. Choosing "RESTRICTED" means SCTE-35 messages of
	// the types specified in AdTriggers that contain delivery restrictions will be
	// treated as ads. Choosing "UNRESTRICTED" means SCTE-35 messages of the types
	// specified in AdTriggers that do not contain delivery restrictions will be
	// treated as ads. Choosing "BOTH" means all SCTE-35 messages of the types
	// specified in AdTriggers will be treated as ads. Note that Splice Insert messages
	// do not have these flags and are always treated as ads if specified in
	// AdTriggers.
	AdsOnDeliveryRestrictions AdsOnDeliveryRestrictions

	// When enabled, an I-Frame only stream will be included in the output.
	IncludeIframeOnlyStream *bool

	// An optional short string appended to the end of the OriginEndpoint URL. If not
	// specified, defaults to the manifestName for the OriginEndpoint.
	ManifestName *string

	// The HTTP Live Streaming (HLS) playlist type. When either "EVENT" or "VOD" is
	// specified, a corresponding EXT-X-PLAYLIST-TYPE entry will be included in the
	// media playlist.
	PlaylistType PlaylistType

	// Time window (in seconds) contained in each parent manifest.
	PlaylistWindowSeconds *int32

	// The interval (in seconds) between each EXT-X-PROGRAM-DATE-TIME tag inserted into
	// manifests. Additionally, when an interval is specified ID3Timed Metadata
	// messages will be generated every 5 seconds using the ingest time of the content.
	// If the interval is not specified, or set to 0, then no EXT-X-PROGRAM-DATE-TIME
	// tags will be inserted into manifests and no ID3Timed Metadata messages will be
	// generated. Note that irrespective of this parameter, if any ID3 Timed Metadata
	// is found in HTTP Live Streaming (HLS) input, it will be passed through to HLS
	// output.
	ProgramDateTimeIntervalSeconds *int32
}

// An HTTP Live Streaming (HLS) packaging configuration.
type HlsPackage struct {

	// This setting controls how ad markers are included in the packaged
	// OriginEndpoint. "NONE" will omit all SCTE-35 ad markers from the output.
	// "PASSTHROUGH" causes the manifest to contain a copy of the SCTE-35 ad markers
	// (comments) taken directly from the input HTTP Live Streaming (HLS) manifest.
	// "SCTE35_ENHANCED" generates ad markers and blackout tags based on SCTE-35
	// messages in the input source. "DATERANGE" inserts EXT-X-DATERANGE tags to signal
	// ad and program transition events in HLS and CMAF manifests. For this option, you
	// must set a programDateTimeIntervalSeconds value that is greater than 0.
	AdMarkers AdMarkers

	// A list of SCTE-35 message types that are treated as ad markers in the output. If
	// empty, no ad markers are output. Specify multiple items to create ad markers for
	// all of the included message types.
	AdTriggers []AdTriggersElement

	// This setting allows the delivery restriction flags on SCTE-35 segmentation
	// descriptors to determine whether a message signals an ad. Choosing "NONE" means
	// no SCTE-35 messages become ads. Choosing "RESTRICTED" means SCTE-35 messages of
	// the types specified in AdTriggers that contain delivery restrictions will be
	// treated as ads. Choosing "UNRESTRICTED" means SCTE-35 messages of the types
	// specified in AdTriggers that do not contain delivery restrictions will be
	// treated as ads. Choosing "BOTH" means all SCTE-35 messages of the types
	// specified in AdTriggers will be treated as ads. Note that Splice Insert messages
	// do not have these flags and are always treated as ads if specified in
	// AdTriggers.
	AdsOnDeliveryRestrictions AdsOnDeliveryRestrictions

	// An HTTP Live Streaming (HLS) encryption configuration.
	Encryption *HlsEncryption

	// When enabled, an I-Frame only stream will be included in the output.
	IncludeIframeOnlyStream *bool

	// The HTTP Live Streaming (HLS) playlist type. When either "EVENT" or "VOD" is
	// specified, a corresponding EXT-X-PLAYLIST-TYPE entry will be included in the
	// media playlist.
	PlaylistType PlaylistType

	// Time window (in seconds) contained in each parent manifest.
	PlaylistWindowSeconds *int32

	// The interval (in seconds) between each EXT-X-PROGRAM-DATE-TIME tag inserted into
	// manifests. Additionally, when an interval is specified ID3Timed Metadata
	// messages will be generated every 5 seconds using the ingest time of the content.
	// If the interval is not specified, or set to 0, then no EXT-X-PROGRAM-DATE-TIME
	// tags will be inserted into manifests and no ID3Timed Metadata messages will be
	// generated. Note that irrespective of this parameter, if any ID3 Timed Metadata
	// is found in HTTP Live Streaming (HLS) input, it will be passed through to HLS
	// output.
	ProgramDateTimeIntervalSeconds *int32

	// Duration (in seconds) of each fragment. Actual fragments will be rounded to the
	// nearest multiple of the source fragment duration.
	SegmentDurationSeconds *int32

	// A StreamSelection configuration.
	StreamSelection *StreamSelection

	// When enabled, audio streams will be placed in rendition groups in the output.
	UseAudioRenditionGroup *bool
}

// An endpoint for ingesting source content for a Channel.
type IngestEndpoint struct {

	// The system generated unique identifier for the IngestEndpoint
	Id *string

	// The system generated password for ingest authentication.
	Password *string

	// The ingest URL to which the source stream should be sent.
	Url *string

	// The system generated username for ingest authentication.
	Username *string
}

// A Microsoft Smooth Streaming (MSS) encryption configuration.
type MssEncryption struct {

	// A configuration for accessing an external Secure Packager and Encoder Key
	// Exchange (SPEKE) service that will provide encryption keys.
	//
	// This member is required.
	SpekeKeyProvider *SpekeKeyProvider
}

// A Microsoft Smooth Streaming (MSS) packaging configuration.
type MssPackage struct {

	// A Microsoft Smooth Streaming (MSS) encryption configuration.
	Encryption *MssEncryption

	// The time window (in seconds) contained in each manifest.
	ManifestWindowSeconds *int32

	// The duration (in seconds) of each segment.
	SegmentDurationSeconds *int32

	// A StreamSelection configuration.
	StreamSelection *StreamSelection
}

// An OriginEndpoint resource configuration.
type OriginEndpoint struct {

	// The Amazon Resource Name (ARN) assigned to the OriginEndpoint.
	Arn *string

	// CDN Authorization credentials
	Authorization *Authorization

	// The ID of the Channel the OriginEndpoint is associated with.
	ChannelId *string

	// A Common Media Application Format (CMAF) packaging configuration.
	CmafPackage *CmafPackage

	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *DashPackage

	// A short text description of the OriginEndpoint.
	Description *string

	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *HlsPackage

	// The ID of the OriginEndpoint.
	Id *string

	// A short string appended to the end of the OriginEndpoint URL.
	ManifestName *string

	// A Microsoft Smooth Streaming (MSS) packaging configuration.
	MssPackage *MssPackage

	// Control whether origination of video is allowed for this OriginEndpoint. If set
	// to ALLOW, the OriginEndpoint may by requested, pursuant to any other form of
	// access control. If set to DENY, the OriginEndpoint may not be requested. This
	// can be helpful for Live to VOD harvesting, or for temporarily disabling
	// origination
	Origination Origination

	// Maximum duration (seconds) of content to retain for startover playback. If not
	// specified, startover playback will be disabled for the OriginEndpoint.
	StartoverWindowSeconds *int32

	// A collection of tags associated with a resource
	Tags map[string]*string

	// Amount of delay (seconds) to enforce on the playback of live content. If not
	// specified, there will be no time delay in effect for the OriginEndpoint.
	TimeDelaySeconds *int32

	// The URL of the packaged OriginEndpoint for consumption.
	Url *string

	// A list of source IP CIDR blocks that will be allowed to access the
	// OriginEndpoint.
	Whitelist []*string
}

// Configuration parameters for where in an S3 bucket to place the harvested
// content
type S3Destination struct {

	// The name of an S3 bucket within which harvested content will be exported
	//
	// This member is required.
	BucketName *string

	// The key in the specified S3 bucket where the harvested top-level manifest will
	// be placed.
	//
	// This member is required.
	ManifestKey *string

	// The IAM role used to write to the specified S3 bucket
	//
	// This member is required.
	RoleArn *string
}

// A configuration for accessing an external Secure Packager and Encoder Key
// Exchange (SPEKE) service that will provide encryption keys.
type SpekeKeyProvider struct {

	// The resource ID to include in key requests.
	//
	// This member is required.
	ResourceId *string

	// An Amazon Resource Name (ARN) of an IAM role that AWS Elemental MediaPackage
	// will assume when accessing the key provider service.
	//
	// This member is required.
	RoleArn *string

	// The system IDs to include in key requests.
	//
	// This member is required.
	SystemIds []*string

	// The URL of the external key provider service.
	//
	// This member is required.
	Url *string

	// An Amazon Resource Name (ARN) of a Certificate Manager certificate that
	// MediaPackage will use for enforcing secure end-to-end data transfer with the key
	// provider service.
	CertificateArn *string
}

// A StreamSelection configuration.
type StreamSelection struct {

	// The maximum video bitrate (bps) to include in output.
	MaxVideoBitsPerSecond *int32

	// The minimum video bitrate (bps) to include in output.
	MinVideoBitsPerSecond *int32

	// A directive that determines the order of streams in the output.
	StreamOrder StreamOrder
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// Represents a cross-account destination that receives subscription log events.
type Destination struct {

	// An IAM policy document that governs which AWS accounts can create subscription
	// filters against this destination.
	AccessPolicy *string

	// The ARN of this destination.
	Arn *string

	// The creation time of the destination, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC.
	CreationTime *int64

	// The name of the destination.
	DestinationName *string

	// A role for impersonation, used when delivering log events to the target.
	RoleArn *string

	// The Amazon Resource Name (ARN) of the physical target to where the log events
	// are delivered (for example, a Kinesis stream).
	TargetArn *string
}

// Represents an export task.
type ExportTask struct {

	// The name of Amazon S3 bucket to which the log data was exported.
	Destination *string

	// The prefix that was used as the start of Amazon S3 key for every object
	// exported.
	DestinationPrefix *string

	// Execution info about the export task.
	ExecutionInfo *ExportTaskExecutionInfo

	// The start time, expressed as the number of milliseconds after Jan 1, 1970
	// 00:00:00 UTC. Events with a timestamp before this time are not exported.
	From *int64

	// The name of the log group from which logs data was exported.
	LogGroupName *string

	// The status of the export task.
	Status *ExportTaskStatus

	// The ID of the export task.
	TaskId *string

	// The name of the export task.
	TaskName *string

	// The end time, expressed as the number of milliseconds after Jan 1, 1970 00:00:00
	// UTC. Events with a timestamp later than this time are not exported.
	To *int64
}

// Represents the status of an export task.
type ExportTaskExecutionInfo struct {

	// The completion time of the export task, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC.
	CompletionTime *int64

	// The creation time of the export task, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC.
	CreationTime *int64
}

// Represents the status of an export task.
type ExportTaskStatus struct {

	// The status code of the export task.
	Code ExportTaskStatusCode

	// The status message related to the status code.
	Message *string
}

// Represents a matched event.
type FilteredLogEvent struct {

	// The ID of the event.
	EventId *string

	// The time the event was ingested, expressed as the number of milliseconds after
	// Jan 1, 1970 00:00:00 UTC.
	IngestionTime *int64

	// The name of the log stream to which this event belongs.
	LogStreamName *string

	// The data contained in the log event.
	Message *string

	// The time the event occurred, expressed as the number of milliseconds after Jan
	// 1, 1970 00:00:00 UTC.
	Timestamp *int64
}

// Represents a log event, which is a record of activity that was recorded by the
// application or resource being monitored.
type InputLogEvent struct {

	// The raw event message.
	//
	// This member is required.
	Message *string

	// The time the event occurred, expressed as the number of milliseconds after Jan
	// 1, 1970 00:00:00 UTC.
	//
	// This member is required.
	Timestamp *int64
}

// Represents a log group.
type LogGroup struct {

	// The Amazon Resource Name (ARN) of the log group.
	Arn *string

	// The creation time of the log group, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC.
	CreationTime *int64

	// The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
	KmsKeyId *string

	// The name of the log group.
	LogGroupName *string

	// The number of metric filters.
	MetricFilterCount *int32

	// The number of days to retain the log events in the specified log group. Possible
	// values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827,
	// and 3653.
	RetentionInDays *int32

	// The number of bytes stored.
	StoredBytes *int64
}

// The fields contained in log events found by a GetLogGroupFields operation, along
// with the percentage of queried log events in which each field appears.
type LogGroupField struct {

	// The name of a log field.
	Name *string

	// The percentage of log events queried that contained the field.
	Percent *int32
}

// Represents a log stream, which is a sequence of log events from a single emitter
// of logs.
type LogStream struct {

	// The Amazon Resource Name (ARN) of the log stream.
	Arn *string

	// The creation time of the stream, expressed as the number of milliseconds after
	// Jan 1, 1970 00:00:00 UTC.
	CreationTime *int64

	// The time of the first event, expressed as the number of milliseconds after Jan
	// 1, 1970 00:00:00 UTC.
	FirstEventTimestamp *int64

	// The time of the most recent log event in the log stream in CloudWatch Logs. This
	// number is expressed as the number of milliseconds after Jan 1, 1970 00:00:00
	// UTC. The lastEventTime value updates on an eventual consistency basis. It
	// typically updates in less than an hour from ingestion, but may take longer in
	// some rare situations.
	LastEventTimestamp *int64

	// The ingestion time, expressed as the number of milliseconds after Jan 1, 1970
	// 00:00:00 UTC.
	LastIngestionTime *int64

	// The name of the log stream.
	LogStreamName *string

	// The number of bytes stored. IMPORTANT:On June 17, 2019, this parameter was
	// deprecated for log streams, and is always reported as zero. This change applies
	// only to log streams. The storedBytes parameter for log groups is not affected.
	StoredBytes *int64

	// The sequence token.
	UploadSequenceToken *string
}

// Metric filters express how CloudWatch Logs would extract metric observations
// from ingested log events and transform them into metric data in a CloudWatch
// metric.
type MetricFilter struct {

	// The creation time of the metric filter, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC.
	CreationTime *int64

	// The name of the metric filter.
	FilterName *string

	// A symbolic description of how CloudWatch Logs should interpret the data in each
	// log event. For example, a log event may contain timestamps, IP addresses,
	// strings, and so on. You use the filter pattern to specify what to look for in
	// the log event message.
	FilterPattern *string

	// The name of the log group.
	LogGroupName *string

	// The metric transformations.
	MetricTransformations []*MetricTransformation
}

// Represents a matched event.
type MetricFilterMatchRecord struct {

	// The raw event data.
	EventMessage *string

	// The event number.
	EventNumber *int64

	// The values extracted from the event data by the filter.
	ExtractedValues map[string]*string
}

// Indicates how to transform ingested log events to metric data in a CloudWatch
// metric.
type MetricTransformation struct {

	// The name of the CloudWatch metric.
	//
	// This member is required.
	MetricName *string

	// A custom namespace to contain your metric in CloudWatch. Use namespaces to group
	// together metrics that are similar. For more information, see Namespaces
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html#Namespace).
	//
	// This member is required.
	MetricNamespace *string

	// The value to publish to the CloudWatch metric when a filter pattern matches a
	// log event.
	//
	// This member is required.
	MetricValue *string

	// (Optional) The value to emit when a filter pattern does not match a log event.
	// This value can be null.
	DefaultValue *float64
}

// Represents a log event.
type OutputLogEvent struct {

	// The time the event was ingested, expressed as the number of milliseconds after
	// Jan 1, 1970 00:00:00 UTC.
	IngestionTime *int64

	// The data contained in the log event.
	Message *string

	// The time the event occurred, expressed as the number of milliseconds after Jan
	// 1, 1970 00:00:00 UTC.
	Timestamp *int64
}

// Reserved.
type QueryCompileError struct {

	// Reserved.
	Location *QueryCompileErrorLocation

	// Reserved.
	Message *string
}

// Reserved.
type QueryCompileErrorLocation struct {

	// Reserved.
	EndCharOffset *int32

	// Reserved.
	StartCharOffset *int32
}

type QueryDefinition struct {
	LastModified *int64

	LogGroupNames []*string

	Name *string

	QueryDefinitionId *string

	QueryString *string
}

// Information about one CloudWatch Logs Insights query that matches the request in
// a DescribeQueries operation.
type QueryInfo struct {

	// The date and time that this query was created.
	CreateTime *int64

	// The name of the log group scanned by this query.
	LogGroupName *string

	// The unique ID number of this query.
	QueryId *string

	// The query string used in this query.
	QueryString *string

	// The status of this query. Possible values are Cancelled, Complete, Failed,
	// Running, Scheduled, and Unknown.
	Status QueryStatus
}

// Contains the number of log events scanned by the query, the number of log events
// that matched the query criteria, and the total number of bytes in the log events
// that were scanned.
type QueryStatistics struct {

	// The total number of bytes in the log events scanned during the query.
	BytesScanned *float64

	// The number of log events that matched the query string.
	RecordsMatched *float64

	// The total number of log events scanned during the query.
	RecordsScanned *float64
}

// Represents the rejected events.
type RejectedLogEventsInfo struct {

	// The expired log events.
	ExpiredLogEventEndIndex *int32

	// The log events that are too new.
	TooNewLogEventStartIndex *int32

	// The log events that are too old.
	TooOldLogEventEndIndex *int32
}

// A policy enabling one or more entities to put logs to a log group in this
// account.
type ResourcePolicy struct {

	// Timestamp showing when this policy was last updated, expressed as the number of
	// milliseconds after Jan 1, 1970 00:00:00 UTC.
	LastUpdatedTime *int64

	// The details of the policy.
	PolicyDocument *string

	// The name of the resource policy.
	PolicyName *string
}

// Contains one field from one log event returned by a CloudWatch Logs Insights
// query, along with the value of that field. For more information about the fields
// that are generated by CloudWatch logs, see Supported Logs and Discovered Fields
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_AnalyzeLogData-discoverable-fields.html).
type ResultField struct {

	// The log event field.
	Field *string

	// The value of this field.
	Value *string
}

// Represents the search status of a log stream.
type SearchedLogStream struct {

	// The name of the log stream.
	LogStreamName *string

	// Indicates whether all the events in this log stream were searched.
	SearchedCompletely *bool
}

// Represents a subscription filter.
type SubscriptionFilter struct {

	// The creation time of the subscription filter, expressed as the number of
	// milliseconds after Jan 1, 1970 00:00:00 UTC.
	CreationTime *int64

	// The Amazon Resource Name (ARN) of the destination.
	DestinationArn *string

	// The method used to distribute log data to the destination, which can be either
	// random or grouped by log stream.
	Distribution Distribution

	// The name of the subscription filter.
	FilterName *string

	// A symbolic description of how CloudWatch Logs should interpret the data in each
	// log event. For example, a log event may contain timestamps, IP addresses,
	// strings, and so on. You use the filter pattern to specify what to look for in
	// the log event message.
	FilterPattern *string

	// The name of the log group.
	LogGroupName *string

	//
	RoleArn *string
}

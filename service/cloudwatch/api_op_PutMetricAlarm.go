// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates or updates an alarm and associates it with the specified metric, metric
// math expression, or anomaly detection model. Alarms based on anomaly detection
// models cannot have Auto Scaling actions. When this operation creates an alarm,
// the alarm state is immediately set to INSUFFICIENT_DATA. The alarm is then
// evaluated and its state is set appropriately. Any actions associated with the
// new state are then executed. When you update an existing alarm, its state is
// left unchanged, but the update completely overwrites the previous configuration
// of the alarm.  <p>If you are an IAM user, you must have Amazon EC2 permissions
// for some alarm operations:</p> <ul> <li> <p>
// <code>iam:CreateServiceLinkedRole</code> for all alarms with EC2 actions</p>
// </li> <li> <p> <code>ec2:DescribeInstanceStatus</code> and
// <code>ec2:DescribeInstances</code> for all alarms on EC2 instance status
// metrics</p> </li> <li> <p> <code>ec2:StopInstances</code> for alarms with stop
// actions</p> </li> <li> <p> <code>ec2:TerminateInstances</code> for alarms with
// terminate actions</p> </li> <li> <p>No specific permissions are needed for
// alarms with recover actions</p> </li> </ul> <p>If you have read/write
// permissions for Amazon CloudWatch but not for Amazon EC2, you can still create
// an alarm, but the stop or terminate actions are not performed. However, if you
// are later granted the required permissions, the alarm actions that you created
// earlier are performed.</p> <p>If you are using an IAM role (for example, an EC2
// instance profile), you cannot stop or terminate the instance using alarm
// actions. However, you can still see the alarm state and perform any other
// actions such as Amazon SNS notifications or Auto Scaling policies.</p> <p>If you
// are using temporary security credentials granted using AWS STS, you cannot stop
// or terminate an EC2 instance using alarm actions.</p> <p>The first time you
// create an alarm in the AWS Management Console, the CLI, or by using the
// PutMetricAlarm API, CloudWatch creates the necessary service-linked role for
// you. The service-linked role is called
// <code>AWSServiceRoleForCloudWatchEvents</code>. For more information, see <a
// href="https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html#iam-term-service-linked-role">AWS
// service-linked role</a>.</p>
func (c *Client) PutMetricAlarm(ctx context.Context, params *PutMetricAlarmInput, optFns ...func(*Options)) (*PutMetricAlarmOutput, error) {
	stack := middleware.NewStack("PutMetricAlarm", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpPutMetricAlarmMiddlewares(stack)
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
	addOpPutMetricAlarmValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutMetricAlarm(options.Region), middleware.Before)
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
			OperationName: "PutMetricAlarm",
			Err:           err,
		}
	}
	out := result.(*PutMetricAlarmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutMetricAlarmInput struct {

	// The name for the alarm. This name must be unique within the Region.
	//
	// This member is required.
	AlarmName *string

	// The arithmetic operation to use when comparing the specified statistic and
	// threshold. The specified statistic value is used as the first operand. The
	// values LessThanLowerOrGreaterThanUpperThreshold, LessThanLowerThreshold, and
	// GreaterThanUpperThreshold are used only for alarms based on anomaly detection
	// models.
	//
	// This member is required.
	ComparisonOperator types.ComparisonOperator

	// The number of periods over which data is compared to the specified threshold. If
	// you are setting an alarm that requires that a number of consecutive data points
	// be breaching to trigger the alarm, this value specifies that number. If you are
	// setting an "M out of N" alarm, this value is the N. An alarm's total current
	// evaluation period can be no longer than one day, so this number multiplied by
	// Period cannot be more than 86,400 seconds.
	//
	// This member is required.
	EvaluationPeriods *int32

	// Indicates whether actions should be executed during any changes to the alarm
	// state. The default is TRUE.
	ActionsEnabled *bool

	// The actions to execute when this alarm transitions to the ALARM state from any
	// other state. Each action is specified as an Amazon Resource Name (ARN).
	// <p>Valid Values: <code>arn:aws:automate:<i>region</i>:ec2:stop</code> |
	// <code>arn:aws:automate:<i>region</i>:ec2:terminate</code> |
	// <code>arn:aws:automate:<i>region</i>:ec2:recover</code> |
	// <code>arn:aws:automate:<i>region</i>:ec2:reboot</code> |
	// <code>arn:aws:sns:<i>region</i>:<i>account-id</i>:<i>sns-topic-name</i> </code>
	// |
	// <code>arn:aws:autoscaling:<i>region</i>:<i>account-id</i>:scalingPolicy:<i>policy-id</i>:autoScalingGroupName/<i>group-friendly-name</i>:policyName/<i>policy-friendly-name</i>
	// </code> </p> <p>Valid Values (for use with IAM roles):
	// <code>arn:aws:swf:<i>region</i>:<i>account-id</i>:action/actions/AWS_EC2.InstanceId.Stop/1.0</code>
	// |
	// <code>arn:aws:swf:<i>region</i>:<i>account-id</i>:action/actions/AWS_EC2.InstanceId.Terminate/1.0</code>
	// |
	// <code>arn:aws:swf:<i>region</i>:<i>account-id</i>:action/actions/AWS_EC2.InstanceId.Reboot/1.0</code>
	// </p>
	AlarmActions []*string

	// The description for the alarm.
	AlarmDescription *string

	// The number of data points that must be breaching to trigger the alarm. This is
	// used only if you are setting an "M out of N" alarm. In that case, this value is
	// the M. For more information, see Evaluating an Alarm
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation)
	// in the Amazon CloudWatch User Guide.
	DatapointsToAlarm *int32

	// The dimensions for the metric specified in MetricName.
	Dimensions []*types.Dimension

	// Used only for alarms based on percentiles. If you specify ignore, the alarm
	// state does not change during periods with too few data points to be
	// statistically significant. If you specify evaluate or omit this parameter, the
	// alarm is always evaluated and possibly changes state no matter how many data
	// points are available. For more information, see Percentile-Based CloudWatch
	// Alarms and Low Data Samples
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#percentiles-with-low-samples).
	// Valid Values: evaluate | ignore
	EvaluateLowSampleCountPercentile *string

	// The percentile statistic for the metric specified in MetricName. Specify a value
	// between p0.0 and p100. When you call PutMetricAlarm and specify a MetricName,
	// you must specify either Statistic or ExtendedStatistic, but not both.
	ExtendedStatistic *string

	// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA
	// state from any other state. Each action is specified as an Amazon Resource Name
	// (ARN).  <p>Valid Values: <code>arn:aws:automate:<i>region</i>:ec2:stop</code> |
	// <code>arn:aws:automate:<i>region</i>:ec2:terminate</code> |
	// <code>arn:aws:automate:<i>region</i>:ec2:recover</code> |
	// <code>arn:aws:automate:<i>region</i>:ec2:reboot</code> |
	// <code>arn:aws:sns:<i>region</i>:<i>account-id</i>:<i>sns-topic-name</i> </code>
	// |
	// <code>arn:aws:autoscaling:<i>region</i>:<i>account-id</i>:scalingPolicy:<i>policy-id</i>:autoScalingGroupName/<i>group-friendly-name</i>:policyName/<i>policy-friendly-name</i>
	// </code> </p> <p>Valid Values (for use with IAM roles):
	// <code>>arn:aws:swf:<i>region</i>:<i>account-id</i>:action/actions/AWS_EC2.InstanceId.Stop/1.0</code>
	// |
	// <code>arn:aws:swf:<i>region</i>:<i>account-id</i>:action/actions/AWS_EC2.InstanceId.Terminate/1.0</code>
	// |
	// <code>arn:aws:swf:<i>region</i>:<i>account-id</i>:action/actions/AWS_EC2.InstanceId.Reboot/1.0</code>
	// </p>
	InsufficientDataActions []*string

	// The name for the metric associated with the alarm. For each PutMetricAlarm
	// operation, you must specify either MetricName or a Metrics array. If you are
	// creating an alarm based on a math expression, you cannot specify this parameter,
	// or any of the Dimensions, Period, Namespace, Statistic, or ExtendedStatistic
	// parameters. Instead, you specify all this information in the Metrics array.
	MetricName *string

	// An array of MetricDataQuery structures that enable you to create an alarm based
	// on the result of a metric math expression. For each PutMetricAlarm operation,
	// you must specify either MetricName or a Metrics array. Each item in the Metrics
	// array either retrieves a metric or performs a math expression. One item in the
	// Metrics array is the expression that the alarm watches. You designate this
	// expression by setting ReturnValue to true for this object in the array. For more
	// information, see MetricDataQuery
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDataQuery.html).
	// If you use the Metrics parameter, you cannot include the MetricName, Dimensions,
	// Period, Namespace, Statistic, or ExtendedStatistic parameters of PutMetricAlarm
	// in the same operation. Instead, you retrieve the metrics you are using in your
	// math expression as part of the Metrics array.
	Metrics []*types.MetricDataQuery

	// The namespace for the metric associated specified in MetricName.
	Namespace *string

	// The actions to execute when this alarm transitions to an OK state from any other
	// state. Each action is specified as an Amazon Resource Name (ARN).  <p>Valid
	// Values: <code>arn:aws:automate:<i>region</i>:ec2:stop</code> |
	// <code>arn:aws:automate:<i>region</i>:ec2:terminate</code> |
	// <code>arn:aws:automate:<i>region</i>:ec2:recover</code> |
	// <code>arn:aws:automate:<i>region</i>:ec2:reboot</code> |
	// <code>arn:aws:sns:<i>region</i>:<i>account-id</i>:<i>sns-topic-name</i> </code>
	// |
	// <code>arn:aws:autoscaling:<i>region</i>:<i>account-id</i>:scalingPolicy:<i>policy-id</i>:autoScalingGroupName/<i>group-friendly-name</i>:policyName/<i>policy-friendly-name</i>
	// </code> </p> <p>Valid Values (for use with IAM roles):
	// <code>arn:aws:swf:<i>region</i>:<i>account-id</i>:action/actions/AWS_EC2.InstanceId.Stop/1.0</code>
	// |
	// <code>arn:aws:swf:<i>region</i>:<i>account-id</i>:action/actions/AWS_EC2.InstanceId.Terminate/1.0</code>
	// |
	// <code>arn:aws:swf:<i>region</i>:<i>account-id</i>:action/actions/AWS_EC2.InstanceId.Reboot/1.0</code>
	// </p>
	OKActions []*string

	// The length, in seconds, used each time the metric specified in MetricName is
	// evaluated. Valid values are 10, 30, and any multiple of 60. Period is required
	// for alarms based on static thresholds. If you are creating an alarm based on a
	// metric math expression, you specify the period for each metric within the
	// objects in the Metrics array. Be sure to specify 10 or 30 only for metrics that
	// are stored by a PutMetricData call with a StorageResolution of 1. If you specify
	// a period of 10 or 30 for a metric that does not have sub-minute resolution, the
	// alarm still attempts to gather data at the period rate that you specify. In this
	// case, it does not receive data for the attempts that do not correspond to a
	// one-minute data resolution, and the alarm might often lapse into
	// INSUFFICENT_DATA status. Specifying 10 or 30 also sets this alarm as a
	// high-resolution alarm, which has a higher charge than other alarms. For more
	// information about pricing, see Amazon CloudWatch Pricing
	// (https://aws.amazon.com/cloudwatch/pricing/). An alarm's total current
	// evaluation period can be no longer than one day, so Period multiplied by
	// EvaluationPeriods cannot be more than 86,400 seconds.
	Period *int32

	// The statistic for the metric specified in MetricName, other than percentile. For
	// percentile statistics, use ExtendedStatistic. When you call PutMetricAlarm and
	// specify a MetricName, you must specify either Statistic or ExtendedStatistic,
	// but not both.
	Statistic types.Statistic

	// A list of key-value pairs to associate with the alarm. You can associate as many
	// as 50 tags with an alarm. Tags can help you organize and categorize your
	// resources. You can also use them to scope user permissions by granting a user
	// permission to access or change only resources with certain tag values.
	Tags []*types.Tag

	// The value against which the specified statistic is compared. This parameter is
	// required for alarms based on static thresholds, but should not be used for
	// alarms based on anomaly detection models.
	Threshold *float64

	// If this is an alarm based on an anomaly detection model, make this value match
	// the ID of the ANOMALY_DETECTION_BAND function. For an example of how to use this
	// parameter, see the Anomaly Detection Model Alarm example on this page. If your
	// alarm uses this parameter, it cannot have Auto Scaling actions.
	ThresholdMetricId *string

	// Sets how this alarm is to handle missing data points. If TreatMissingData is
	// omitted, the default behavior of missing is used. For more information, see
	// Configuring How CloudWatch Alarms Treats Missing Data
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarms-and-missing-data).
	// Valid Values: breaching | notBreaching | ignore | missing
	TreatMissingData *string

	// The unit of measure for the statistic. For example, the units for the Amazon EC2
	// NetworkIn metric are Bytes because NetworkIn tracks the number of bytes that an
	// instance receives on all network interfaces. You can also specify a unit when
	// you create a custom metric. Units help provide conceptual meaning to your data.
	// Metric data points that specify a unit of measure, such as Percent, are
	// aggregated separately. If you don't specify Unit, CloudWatch retrieves all unit
	// types that have been published for the metric and attempts to evaluate the
	// alarm. Usually, metrics are published with only one unit, so the alarm works as
	// intended. However, if the metric is published with multiple types of units and
	// you don't specify a unit, the alarm's behavior is not defined and it behaves
	// predictably. We recommend omitting Unit so that you don't inadvertently specify
	// an incorrect unit that is not published for this metric. Doing so causes the
	// alarm to be stuck in the INSUFFICIENT DATA state.
	Unit types.StandardUnit
}

type PutMetricAlarmOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpPutMetricAlarmMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpPutMetricAlarm{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpPutMetricAlarm{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutMetricAlarm(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "PutMetricAlarm",
	}
}

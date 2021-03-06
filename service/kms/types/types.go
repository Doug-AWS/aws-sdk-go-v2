// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Contains information about an alias.
type AliasListEntry struct {

	// String that contains the key ARN.
	AliasArn *string

	// String that contains the alias. This value begins with alias/.
	AliasName *string

	// String that contains the key identifier referred to by the alias.
	TargetKeyId *string
}

// Contains information about each custom key store in the custom key store list.
type CustomKeyStoresListEntry struct {

	// A unique identifier for the AWS CloudHSM cluster that is associated with the
	// custom key store.
	CloudHsmClusterId *string

	// Describes the connection error. This field appears in the response only when the
	// ConnectionState is FAILED. For help resolving these errors, see How to Fix a
	// Connection Failure
	// (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#fix-keystore-failed)
	// in AWS Key Management Service Developer Guide. Valid values are:
	//
	// *
	// CLUSTER_NOT_FOUND - AWS KMS cannot find the AWS CloudHSM cluster with the
	// specified cluster ID.
	//
	// * INSUFFICIENT_CLOUDHSM_HSMS - The associated AWS
	// CloudHSM cluster does not contain any active HSMs. To connect a custom key store
	// to its AWS CloudHSM cluster, the cluster must contain at least one active
	// HSM.
	//
	// * INTERNAL_ERROR - AWS KMS could not complete the request due to an
	// internal error. Retry the request. For ConnectCustomKeyStore requests,
	// disconnect the custom key store before trying to connect again.
	//
	// *
	// INVALID_CREDENTIALS - AWS KMS does not have the correct password for the kmsuser
	// crypto user in the AWS CloudHSM cluster. Before you can connect your custom key
	// store to its AWS CloudHSM cluster, you must change the kmsuser account password
	// and update the key store password value for the custom key store.
	//
	// *
	// NETWORK_ERRORS - Network errors are preventing AWS KMS from connecting to the
	// custom key store.
	//
	// * SUBNET_NOT_FOUND - A subnet in the AWS CloudHSM cluster
	// configuration was deleted. If AWS KMS cannot find all of the subnets in the
	// cluster configuration, attempts to connect the custom key store to the AWS
	// CloudHSM cluster fail. To fix this error, create a cluster from a recent backup
	// and associate it with your custom key store. (This process creates a new cluster
	// configuration with a VPC and private subnets.) For details, see How to Fix a
	// Connection Failure
	// (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#fix-keystore-failed)
	// in the AWS Key Management Service Developer Guide.
	//
	// * USER_LOCKED_OUT - The
	// kmsuser CU account is locked out of the associated AWS CloudHSM cluster due to
	// too many failed password attempts. Before you can connect your custom key store
	// to its AWS CloudHSM cluster, you must change the kmsuser account password and
	// update the key store password value for the custom key store.
	//
	// * USER_LOGGED_IN
	// - The kmsuser CU account is logged into the the associated AWS CloudHSM cluster.
	// This prevents AWS KMS from rotating the kmsuser account password and logging
	// into the cluster. Before you can connect your custom key store to its AWS
	// CloudHSM cluster, you must log the kmsuser CU out of the cluster. If you changed
	// the kmsuser password to log into the cluster, you must also and update the key
	// store password value for the custom key store. For help, see How to Log Out and
	// Reconnect
	// (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#login-kmsuser-2)
	// in the AWS Key Management Service Developer Guide.
	//
	// * USER_NOT_FOUND - AWS KMS
	// cannot find a kmsuser CU account in the associated AWS CloudHSM cluster. Before
	// you can connect your custom key store to its AWS CloudHSM cluster, you must
	// create a kmsuser CU account in the cluster, and then update the key store
	// password value for the custom key store.
	ConnectionErrorCode ConnectionErrorCodeType

	// Indicates whether the custom key store is connected to its AWS CloudHSM cluster.
	// You can create and use CMKs in your custom key stores only when its connection
	// state is CONNECTED. The value is DISCONNECTED if the key store has never been
	// connected or you use the DisconnectCustomKeyStore operation to disconnect it. If
	// the value is CONNECTED but you are having trouble using the custom key store,
	// make sure that its associated AWS CloudHSM cluster is active and contains at
	// least one active HSM. A value of FAILED indicates that an attempt to connect was
	// unsuccessful. The ConnectionErrorCode field in the response indicates the cause
	// of the failure. For help resolving a connection failure, see Troubleshooting a
	// Custom Key Store
	// (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html) in the
	// AWS Key Management Service Developer Guide.
	ConnectionState ConnectionStateType

	// The date and time when the custom key store was created.
	CreationDate *time.Time

	// A unique identifier for the custom key store.
	CustomKeyStoreId *string

	// The user-specified friendly name for the custom key store.
	CustomKeyStoreName *string

	// The trust anchor certificate of the associated AWS CloudHSM cluster. When you
	// initialize the cluster
	// (https://docs.aws.amazon.com/cloudhsm/latest/userguide/initialize-cluster.html#sign-csr),
	// you create this certificate and save it in the customerCA.crt file.
	TrustAnchorCertificate *string
}

// Use this structure to allow cryptographic operations
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
// in the grant only when the operation request includes the specified encryption
// context
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context).
// AWS KMS applies the grant constraints only to cryptographic operations that
// support an encryption context, that is, all cryptographic operations with a
// symmetric CMK
// (https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-concepts.html#symmetric-cmks).
// Grant constraints are not applied to operations that do not support an
// encryption context, such as cryptographic operations with asymmetric CMKs and
// management operations, such as DescribeKey or ScheduleKeyDeletion. In a
// cryptographic operation, the encryption context in the decryption operation must
// be an exact, case-sensitive match for the keys and values in the encryption
// context of the encryption operation. Only the order of the pairs can vary.
// However, in a grant constraint, the key in each key-value pair is not case
// sensitive, but the value is case sensitive. To avoid confusion, do not use
// multiple encryption context pairs that differ only by case. To require a fully
// case-sensitive encryption context, use the kms:EncryptionContext: and
// kms:EncryptionContextKeys conditions in an IAM or key policy. For details, see
// kms:EncryptionContext:
// (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-encryption-context)
// in the AWS Key Management Service Developer Guide .
type GrantConstraints struct {

	// A list of key-value pairs that must match the encryption context in the
	// cryptographic operation
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
	// request. The grant allows the operation only when the encryption context in the
	// request is the same as the encryption context specified in this constraint.
	EncryptionContextEquals map[string]*string

	// A list of key-value pairs that must be included in the encryption context of the
	// cryptographic operation
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
	// request. The grant allows the cryptographic operation only when the encryption
	// context in the request includes the key-value pairs specified in this
	// constraint, although it can include additional key-value pairs.
	EncryptionContextSubset map[string]*string
}

// Contains information about a grant.
type GrantListEntry struct {

	// A list of key-value pairs that must be present in the encryption context of
	// certain subsequent operations that the grant allows.
	Constraints *GrantConstraints

	// The date and time when the grant was created.
	CreationDate *time.Time

	// The unique identifier for the grant.
	GrantId *string

	// The identity that gets the permissions in the grant. The GranteePrincipal field
	// in the ListGrants response usually contains the user or role designated as the
	// grantee principal in the grant. However, when the grantee principal in the grant
	// is an AWS service, the GranteePrincipal field contains the service principal
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html#principal-services),
	// which might represent several different grantee principals.
	GranteePrincipal *string

	// The AWS account under which the grant was issued.
	IssuingAccount *string

	// The unique identifier for the customer master key (CMK) to which the grant
	// applies.
	KeyId *string

	// The friendly name that identifies the grant. If a name was provided in the
	// CreateGrant request, that name is returned. Otherwise this value is null.
	Name *string

	// The list of operations permitted by the grant.
	Operations []GrantOperation

	// The principal that can retire the grant.
	RetiringPrincipal *string
}

// Contains information about each entry in the key list.
type KeyListEntry struct {

	// ARN of the key.
	KeyArn *string

	// Unique identifier of the key.
	KeyId *string
}

// Contains metadata about a customer master key (CMK). This data type is used as a
// response element for the CreateKey and DescribeKey operations.
type KeyMetadata struct {

	// The globally unique identifier for the CMK.
	//
	// This member is required.
	KeyId *string

	// The twelve-digit account ID of the AWS account that owns the CMK.
	AWSAccountId *string

	// The Amazon Resource Name (ARN) of the CMK. For examples, see AWS Key Management
	// Service (AWS KMS)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kms)
	// in the Example ARNs section of the AWS General Reference.
	Arn *string

	// The cluster ID of the AWS CloudHSM cluster that contains the key material for
	// the CMK. When you create a CMK in a custom key store
	// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html),
	// AWS KMS creates the key material for the CMK in the associated AWS CloudHSM
	// cluster. This value is present only when the CMK is created in a custom key
	// store.
	CloudHsmClusterId *string

	// The date and time when the CMK was created.
	CreationDate *time.Time

	// A unique identifier for the custom key store
	// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
	// that contains the CMK. This value is present only when the CMK is created in a
	// custom key store.
	CustomKeyStoreId *string

	// Describes the type of key material in the CMK.
	CustomerMasterKeySpec CustomerMasterKeySpec

	// The date and time after which AWS KMS deletes the CMK. This value is present
	// only when KeyState is PendingDeletion.
	DeletionDate *time.Time

	// The description of the CMK.
	Description *string

	// Specifies whether the CMK is enabled. When KeyState is Enabled this value is
	// true, otherwise it is false.
	Enabled *bool

	// The encryption algorithms that the CMK supports. You cannot use the CMK with
	// other encryption algorithms within AWS KMS. This field appears only when the
	// KeyUsage of the CMK is ENCRYPT_DECRYPT.
	EncryptionAlgorithms []EncryptionAlgorithmSpec

	// Specifies whether the CMK's key material expires. This value is present only
	// when Origin is EXTERNAL, otherwise this value is omitted.
	ExpirationModel ExpirationModelType

	// The manager of the CMK. CMKs in your AWS account are either customer managed or
	// AWS managed. For more information about the difference, see Customer Master Keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#master_keys)
	// in the AWS Key Management Service Developer Guide.
	KeyManager KeyManagerType

	// The current status of the CMK. For more information about how key state affects
	// the use of a CMK, see Key state: Effect on your CMK
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
	// AWS Key Management Service Developer Guide.
	KeyState KeyState

	// The cryptographic operations
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
	// for which you can use the CMK.
	KeyUsage KeyUsageType

	// The source of the CMK's key material. When this value is AWS_KMS, AWS KMS
	// created the key material. When this value is EXTERNAL, the key material was
	// imported from your existing key management infrastructure or the CMK lacks key
	// material. When this value is AWS_CLOUDHSM, the key material was created in the
	// AWS CloudHSM cluster associated with a custom key store.
	Origin OriginType

	// The signing algorithms that the CMK supports. You cannot use the CMK with other
	// signing algorithms within AWS KMS. This field appears only when the KeyUsage of
	// the CMK is SIGN_VERIFY.
	SigningAlgorithms []SigningAlgorithmSpec

	// The time at which the imported key material expires. When the key material
	// expires, AWS KMS deletes the key material and the CMK becomes unusable. This
	// value is present only for CMKs whose Origin is EXTERNAL and whose
	// ExpirationModel is KEY_MATERIAL_EXPIRES, otherwise this value is omitted.
	ValidTo *time.Time
}

// A key-value pair. A tag consists of a tag key and a tag value. Tag keys and tag
// values are both required, but tag values can be empty (null) strings. For
// information about the rules that apply to tag keys and tag values, see
// User-Defined Tag Restrictions
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html)
// in the AWS Billing and Cost Management User Guide.
type Tag struct {

	// The key of the tag.
	//
	// This member is required.
	TagKey *string

	// The value of the tag.
	//
	// This member is required.
	TagValue *string
}

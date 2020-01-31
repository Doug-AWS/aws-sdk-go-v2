// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateCustomKeyStoreInput struct {
	_ struct{} `type:"structure"`

	// Associates the custom key store with a related AWS CloudHSM cluster.
	//
	// Enter the cluster ID of the cluster that you used to create the custom key
	// store or a cluster that shares a backup history and has the same cluster
	// certificate as the original cluster. You cannot use this parameter to associate
	// a custom key store with an unrelated cluster. In addition, the replacement
	// cluster must fulfill the requirements (https://docs.aws.amazon.com/kms/latest/developerguide/create-keystore.html#before-keystore)
	// for a cluster associated with a custom key store. To view the cluster certificate
	// of a cluster, use the DescribeClusters (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_DescribeClusters.html)
	// operation.
	CloudHsmClusterId *string `min:"19" type:"string"`

	// Identifies the custom key store that you want to update. Enter the ID of
	// the custom key store. To find the ID of a custom key store, use the DescribeCustomKeyStores
	// operation.
	//
	// CustomKeyStoreId is a required field
	CustomKeyStoreId *string `min:"1" type:"string" required:"true"`

	// Enter the current password of the kmsuser crypto user (CU) in the AWS CloudHSM
	// cluster that is associated with the custom key store.
	//
	// This parameter tells AWS KMS the current password of the kmsuser crypto user
	// (CU). It does not set or change the password of any users in the AWS CloudHSM
	// cluster.
	KeyStorePassword *string `min:"1" type:"string" sensitive:"true"`

	// Changes the friendly name of the custom key store to the value that you specify.
	// The custom key store name must be unique in the AWS account.
	NewCustomKeyStoreName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateCustomKeyStoreInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateCustomKeyStoreInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateCustomKeyStoreInput"}
	if s.CloudHsmClusterId != nil && len(*s.CloudHsmClusterId) < 19 {
		invalidParams.Add(aws.NewErrParamMinLen("CloudHsmClusterId", 19))
	}

	if s.CustomKeyStoreId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CustomKeyStoreId"))
	}
	if s.CustomKeyStoreId != nil && len(*s.CustomKeyStoreId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CustomKeyStoreId", 1))
	}
	if s.KeyStorePassword != nil && len(*s.KeyStorePassword) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyStorePassword", 1))
	}
	if s.NewCustomKeyStoreName != nil && len(*s.NewCustomKeyStoreName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NewCustomKeyStoreName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateCustomKeyStoreOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateCustomKeyStoreOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateCustomKeyStore = "UpdateCustomKeyStore"

// UpdateCustomKeyStoreRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Changes the properties of a custom key store. Use the CustomKeyStoreId parameter
// to identify the custom key store you want to edit. Use the remaining parameters
// to change the properties of the custom key store.
//
// You can only update a custom key store that is disconnected. To disconnect
// the custom key store, use DisconnectCustomKeyStore. To reconnect the custom
// key store after the update completes, use ConnectCustomKeyStore. To find
// the connection state of a custom key store, use the DescribeCustomKeyStores
// operation.
//
// Use the parameters of UpdateCustomKeyStore to edit your keystore settings.
//
//    * Use the NewCustomKeyStoreName parameter to change the friendly name
//    of the custom key store to the value that you specify.
//
//    * Use the KeyStorePassword parameter tell AWS KMS the current password
//    of the kmsuser crypto user (CU) (https://docs.aws.amazon.com/kms/latest/developerguide/key-store-concepts.html#concept-kmsuser)
//    in the associated AWS CloudHSM cluster. You can use this parameter to
//    fix connection failures (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#fix-keystore-password)
//    that occur when AWS KMS cannot log into the associated cluster because
//    the kmsuser password has changed. This value does not change the password
//    in the AWS CloudHSM cluster.
//
//    * Use the CloudHsmClusterId parameter to associate the custom key store
//    with a different, but related, AWS CloudHSM cluster. You can use this
//    parameter to repair a custom key store if its AWS CloudHSM cluster becomes
//    corrupted or is deleted, or when you need to create or restore a cluster
//    from a backup.
//
// If the operation succeeds, it returns a JSON object with no properties.
//
// This operation is part of the Custom Key Store feature (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// feature in AWS KMS, which combines the convenience and extensive integration
// of AWS KMS with the isolation and control of a single-tenant key store.
//
//    // Example sending a request using UpdateCustomKeyStoreRequest.
//    req := client.UpdateCustomKeyStoreRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/UpdateCustomKeyStore
func (c *Client) UpdateCustomKeyStoreRequest(input *UpdateCustomKeyStoreInput) UpdateCustomKeyStoreRequest {
	op := &aws.Operation{
		Name:       opUpdateCustomKeyStore,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateCustomKeyStoreInput{}
	}

	req := c.newRequest(op, input, &UpdateCustomKeyStoreOutput{})
	return UpdateCustomKeyStoreRequest{Request: req, Input: input, Copy: c.UpdateCustomKeyStoreRequest}
}

// UpdateCustomKeyStoreRequest is the request type for the
// UpdateCustomKeyStore API operation.
type UpdateCustomKeyStoreRequest struct {
	*aws.Request
	Input *UpdateCustomKeyStoreInput
	Copy  func(*UpdateCustomKeyStoreInput) UpdateCustomKeyStoreRequest
}

// Send marshals and sends the UpdateCustomKeyStore API request.
func (r UpdateCustomKeyStoreRequest) Send(ctx context.Context) (*UpdateCustomKeyStoreResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateCustomKeyStoreResponse{
		UpdateCustomKeyStoreOutput: r.Request.Data.(*UpdateCustomKeyStoreOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateCustomKeyStoreResponse is the response type for the
// UpdateCustomKeyStore API operation.
type UpdateCustomKeyStoreResponse struct {
	*UpdateCustomKeyStoreOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateCustomKeyStore request.
func (r *UpdateCustomKeyStoreResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

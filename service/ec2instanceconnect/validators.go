// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2instanceconnect

import (
	"context"
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpSendSSHPublicKey struct {
}

func (*validateOpSendSSHPublicKey) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSendSSHPublicKey) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SendSSHPublicKeyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSendSSHPublicKeyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpSendSSHPublicKeyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSendSSHPublicKey{}, middleware.After)
}

func validateOpSendSSHPublicKeyInput(v *SendSSHPublicKeyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SendSSHPublicKeyInput"}
	if v.AvailabilityZone == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AvailabilityZone"))
	}
	if v.InstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InstanceId"))
	}
	if v.SSHPublicKey == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SSHPublicKey"))
	}
	if v.InstanceOSUser == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InstanceOSUser"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

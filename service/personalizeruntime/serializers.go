// Code generated by smithy-go-codegen DO NOT EDIT.

package personalizeruntime

import (
	"bytes"
	"context"
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/httpbinding"
	smithyjson "github.com/awslabs/smithy-go/json"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsRestjson1_serializeOpGetPersonalizedRanking struct {
}

func (*awsRestjson1_serializeOpGetPersonalizedRanking) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetPersonalizedRanking) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetPersonalizedRankingInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/personalize-ranking")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentGetPersonalizedRankingInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetPersonalizedRankingInput(v *GetPersonalizedRankingInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentGetPersonalizedRankingInput(v *GetPersonalizedRankingInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CampaignArn != nil {
		ok := object.Key("campaignArn")
		ok.String(*v.CampaignArn)
	}

	if v.Context != nil {
		ok := object.Key("context")
		if err := awsRestjson1_serializeDocumentContext(v.Context, ok); err != nil {
			return err
		}
	}

	if v.InputList != nil {
		ok := object.Key("inputList")
		if err := awsRestjson1_serializeDocumentInputList(v.InputList, ok); err != nil {
			return err
		}
	}

	if v.UserId != nil {
		ok := object.Key("userId")
		ok.String(*v.UserId)
	}

	return nil
}

type awsRestjson1_serializeOpGetRecommendations struct {
}

func (*awsRestjson1_serializeOpGetRecommendations) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetRecommendations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetRecommendationsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/recommendations")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentGetRecommendationsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetRecommendationsInput(v *GetRecommendationsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentGetRecommendationsInput(v *GetRecommendationsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CampaignArn != nil {
		ok := object.Key("campaignArn")
		ok.String(*v.CampaignArn)
	}

	if v.Context != nil {
		ok := object.Key("context")
		if err := awsRestjson1_serializeDocumentContext(v.Context, ok); err != nil {
			return err
		}
	}

	if v.FilterArn != nil {
		ok := object.Key("filterArn")
		ok.String(*v.FilterArn)
	}

	if v.ItemId != nil {
		ok := object.Key("itemId")
		ok.String(*v.ItemId)
	}

	if v.NumResults != nil {
		ok := object.Key("numResults")
		ok.Integer(*v.NumResults)
	}

	if v.UserId != nil {
		ok := object.Key("userId")
		ok.String(*v.UserId)
	}

	return nil
}

func awsRestjson1_serializeDocumentContext(v map[string]*string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		if vv := v[key]; vv == nil {
			om.Null()
			continue
		}
		om.String(*v[key])
	}
	return nil
}

func awsRestjson1_serializeDocumentInputList(v []*string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			av.Null()
			continue
		}
		av.String(*v[i])
	}
	return nil
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package trustedadvisor

import (
	"bytes"
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpGetOrganizationRecommendation struct {
}

func (*awsRestjson1_serializeOpGetOrganizationRecommendation) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetOrganizationRecommendation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetOrganizationRecommendationInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/organization-recommendations/{organizationRecommendationIdentifier}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetOrganizationRecommendationInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetOrganizationRecommendationInput(v *GetOrganizationRecommendationInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.OrganizationRecommendationIdentifier == nil || len(*v.OrganizationRecommendationIdentifier) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member organizationRecommendationIdentifier must not be empty")}
	}
	if v.OrganizationRecommendationIdentifier != nil {
		if err := encoder.SetURI("organizationRecommendationIdentifier").String(*v.OrganizationRecommendationIdentifier); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetRecommendation struct {
}

func (*awsRestjson1_serializeOpGetRecommendation) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetRecommendation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetRecommendationInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/recommendations/{recommendationIdentifier}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetRecommendationInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetRecommendationInput(v *GetRecommendationInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.RecommendationIdentifier == nil || len(*v.RecommendationIdentifier) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member recommendationIdentifier must not be empty")}
	}
	if v.RecommendationIdentifier != nil {
		if err := encoder.SetURI("recommendationIdentifier").String(*v.RecommendationIdentifier); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpListChecks struct {
}

func (*awsRestjson1_serializeOpListChecks) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListChecks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListChecksInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/checks")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListChecksInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListChecksInput(v *ListChecksInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.AwsService != nil {
		encoder.SetQuery("awsService").String(*v.AwsService)
	}

	if len(v.Language) > 0 {
		encoder.SetQuery("language").String(string(v.Language))
	}

	if v.MaxResults != nil {
		encoder.SetQuery("maxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("nextToken").String(*v.NextToken)
	}

	if len(v.Pillar) > 0 {
		encoder.SetQuery("pillar").String(string(v.Pillar))
	}

	if len(v.Source) > 0 {
		encoder.SetQuery("source").String(string(v.Source))
	}

	return nil
}

type awsRestjson1_serializeOpListOrganizationRecommendationAccounts struct {
}

func (*awsRestjson1_serializeOpListOrganizationRecommendationAccounts) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListOrganizationRecommendationAccounts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListOrganizationRecommendationAccountsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/organization-recommendations/{organizationRecommendationIdentifier}/accounts")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListOrganizationRecommendationAccountsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListOrganizationRecommendationAccountsInput(v *ListOrganizationRecommendationAccountsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.AffectedAccountId != nil {
		encoder.SetQuery("affectedAccountId").String(*v.AffectedAccountId)
	}

	if v.MaxResults != nil {
		encoder.SetQuery("maxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("nextToken").String(*v.NextToken)
	}

	if v.OrganizationRecommendationIdentifier == nil || len(*v.OrganizationRecommendationIdentifier) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member organizationRecommendationIdentifier must not be empty")}
	}
	if v.OrganizationRecommendationIdentifier != nil {
		if err := encoder.SetURI("organizationRecommendationIdentifier").String(*v.OrganizationRecommendationIdentifier); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpListOrganizationRecommendationResources struct {
}

func (*awsRestjson1_serializeOpListOrganizationRecommendationResources) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListOrganizationRecommendationResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListOrganizationRecommendationResourcesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/organization-recommendations/{organizationRecommendationIdentifier}/resources")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListOrganizationRecommendationResourcesInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListOrganizationRecommendationResourcesInput(v *ListOrganizationRecommendationResourcesInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.AffectedAccountId != nil {
		encoder.SetQuery("affectedAccountId").String(*v.AffectedAccountId)
	}

	if v.MaxResults != nil {
		encoder.SetQuery("maxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("nextToken").String(*v.NextToken)
	}

	if v.OrganizationRecommendationIdentifier == nil || len(*v.OrganizationRecommendationIdentifier) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member organizationRecommendationIdentifier must not be empty")}
	}
	if v.OrganizationRecommendationIdentifier != nil {
		if err := encoder.SetURI("organizationRecommendationIdentifier").String(*v.OrganizationRecommendationIdentifier); err != nil {
			return err
		}
	}

	if v.RegionCode != nil {
		encoder.SetQuery("regionCode").String(*v.RegionCode)
	}

	if len(v.Status) > 0 {
		encoder.SetQuery("status").String(string(v.Status))
	}

	return nil
}

type awsRestjson1_serializeOpListOrganizationRecommendations struct {
}

func (*awsRestjson1_serializeOpListOrganizationRecommendations) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListOrganizationRecommendations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListOrganizationRecommendationsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/organization-recommendations")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListOrganizationRecommendationsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListOrganizationRecommendationsInput(v *ListOrganizationRecommendationsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.AfterLastUpdatedAt != nil {
		encoder.SetQuery("afterLastUpdatedAt").String(smithytime.FormatDateTime(*v.AfterLastUpdatedAt))
	}

	if v.AwsService != nil {
		encoder.SetQuery("awsService").String(*v.AwsService)
	}

	if v.BeforeLastUpdatedAt != nil {
		encoder.SetQuery("beforeLastUpdatedAt").String(smithytime.FormatDateTime(*v.BeforeLastUpdatedAt))
	}

	if v.CheckIdentifier != nil {
		encoder.SetQuery("checkIdentifier").String(*v.CheckIdentifier)
	}

	if v.MaxResults != nil {
		encoder.SetQuery("maxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("nextToken").String(*v.NextToken)
	}

	if len(v.Pillar) > 0 {
		encoder.SetQuery("pillar").String(string(v.Pillar))
	}

	if len(v.Source) > 0 {
		encoder.SetQuery("source").String(string(v.Source))
	}

	if len(v.Status) > 0 {
		encoder.SetQuery("status").String(string(v.Status))
	}

	if len(v.Type) > 0 {
		encoder.SetQuery("type").String(string(v.Type))
	}

	return nil
}

type awsRestjson1_serializeOpListRecommendationResources struct {
}

func (*awsRestjson1_serializeOpListRecommendationResources) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListRecommendationResources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListRecommendationResourcesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/recommendations/{recommendationIdentifier}/resources")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListRecommendationResourcesInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListRecommendationResourcesInput(v *ListRecommendationResourcesInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MaxResults != nil {
		encoder.SetQuery("maxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("nextToken").String(*v.NextToken)
	}

	if v.RecommendationIdentifier == nil || len(*v.RecommendationIdentifier) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member recommendationIdentifier must not be empty")}
	}
	if v.RecommendationIdentifier != nil {
		if err := encoder.SetURI("recommendationIdentifier").String(*v.RecommendationIdentifier); err != nil {
			return err
		}
	}

	if v.RegionCode != nil {
		encoder.SetQuery("regionCode").String(*v.RegionCode)
	}

	if len(v.Status) > 0 {
		encoder.SetQuery("status").String(string(v.Status))
	}

	return nil
}

type awsRestjson1_serializeOpListRecommendations struct {
}

func (*awsRestjson1_serializeOpListRecommendations) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListRecommendations) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListRecommendationsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/recommendations")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListRecommendationsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListRecommendationsInput(v *ListRecommendationsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.AfterLastUpdatedAt != nil {
		encoder.SetQuery("afterLastUpdatedAt").String(smithytime.FormatDateTime(*v.AfterLastUpdatedAt))
	}

	if v.AwsService != nil {
		encoder.SetQuery("awsService").String(*v.AwsService)
	}

	if v.BeforeLastUpdatedAt != nil {
		encoder.SetQuery("beforeLastUpdatedAt").String(smithytime.FormatDateTime(*v.BeforeLastUpdatedAt))
	}

	if v.CheckIdentifier != nil {
		encoder.SetQuery("checkIdentifier").String(*v.CheckIdentifier)
	}

	if v.MaxResults != nil {
		encoder.SetQuery("maxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("nextToken").String(*v.NextToken)
	}

	if len(v.Pillar) > 0 {
		encoder.SetQuery("pillar").String(string(v.Pillar))
	}

	if len(v.Source) > 0 {
		encoder.SetQuery("source").String(string(v.Source))
	}

	if len(v.Status) > 0 {
		encoder.SetQuery("status").String(string(v.Status))
	}

	if len(v.Type) > 0 {
		encoder.SetQuery("type").String(string(v.Type))
	}

	return nil
}

type awsRestjson1_serializeOpUpdateOrganizationRecommendationLifecycle struct {
}

func (*awsRestjson1_serializeOpUpdateOrganizationRecommendationLifecycle) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUpdateOrganizationRecommendationLifecycle) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateOrganizationRecommendationLifecycleInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/organization-recommendations/{organizationRecommendationIdentifier}/lifecycle")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "PUT"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsUpdateOrganizationRecommendationLifecycleInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentUpdateOrganizationRecommendationLifecycleInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsUpdateOrganizationRecommendationLifecycleInput(v *UpdateOrganizationRecommendationLifecycleInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.OrganizationRecommendationIdentifier == nil || len(*v.OrganizationRecommendationIdentifier) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member organizationRecommendationIdentifier must not be empty")}
	}
	if v.OrganizationRecommendationIdentifier != nil {
		if err := encoder.SetURI("organizationRecommendationIdentifier").String(*v.OrganizationRecommendationIdentifier); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentUpdateOrganizationRecommendationLifecycleInput(v *UpdateOrganizationRecommendationLifecycleInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.LifecycleStage) > 0 {
		ok := object.Key("lifecycleStage")
		ok.String(string(v.LifecycleStage))
	}

	if v.UpdateReason != nil {
		ok := object.Key("updateReason")
		ok.String(*v.UpdateReason)
	}

	if len(v.UpdateReasonCode) > 0 {
		ok := object.Key("updateReasonCode")
		ok.String(string(v.UpdateReasonCode))
	}

	return nil
}

type awsRestjson1_serializeOpUpdateRecommendationLifecycle struct {
}

func (*awsRestjson1_serializeOpUpdateRecommendationLifecycle) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUpdateRecommendationLifecycle) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateRecommendationLifecycleInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/v1/recommendations/{recommendationIdentifier}/lifecycle")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "PUT"
	var restEncoder *httpbinding.Encoder
	if request.URL.RawPath == "" {
		restEncoder, err = httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	} else {
		request.URL.RawPath = smithyhttp.JoinPath(request.URL.RawPath, opPath)
		restEncoder, err = httpbinding.NewEncoderWithRawPath(request.URL.Path, request.URL.RawPath, request.URL.RawQuery, request.Header)
	}

	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsUpdateRecommendationLifecycleInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentUpdateRecommendationLifecycleInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsUpdateRecommendationLifecycleInput(v *UpdateRecommendationLifecycleInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.RecommendationIdentifier == nil || len(*v.RecommendationIdentifier) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member recommendationIdentifier must not be empty")}
	}
	if v.RecommendationIdentifier != nil {
		if err := encoder.SetURI("recommendationIdentifier").String(*v.RecommendationIdentifier); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentUpdateRecommendationLifecycleInput(v *UpdateRecommendationLifecycleInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.LifecycleStage) > 0 {
		ok := object.Key("lifecycleStage")
		ok.String(string(v.LifecycleStage))
	}

	if v.UpdateReason != nil {
		ok := object.Key("updateReason")
		ok.String(*v.UpdateReason)
	}

	if len(v.UpdateReasonCode) > 0 {
		ok := object.Key("updateReasonCode")
		ok.String(string(v.UpdateReasonCode))
	}

	return nil
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deploy a new environment. An Proton environment is created from an environment
// template that defines infrastructure and resources that can be shared across
// services. You can provision environments using the following methods:
//
// * Amazon
// Web Services-managed provisioning – Proton makes direct calls to provision your
// resources.
//
// * Self-managed provisioning – Proton makes pull requests on your
// repository to provide compiled infrastructure as code (IaC) files that your IaC
// engine uses to provision resources.
//
// * CodeBuild-based provisioning – Proton
// uses CodeBuild to run shell commands that you provide. Your commands can read
// inputs that Proton provides, and are responsible for provisioning or
// deprovisioning infrastructure and generating output values.
//
// For more
// information, see Environments
// (https://docs.aws.amazon.com/proton/latest/userguide/ag-environments.html) and
// Provisioning methods
// (https://docs.aws.amazon.com/proton/latest/userguide/ag-works-prov-methods.html)
// in the Proton User Guide.
func (c *Client) CreateEnvironment(ctx context.Context, params *CreateEnvironmentInput, optFns ...func(*Options)) (*CreateEnvironmentOutput, error) {
	if params == nil {
		params = &CreateEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironment", params, optFns, c.addOperationCreateEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEnvironmentInput struct {

	// The name of the environment.
	//
	// This member is required.
	Name *string

	// A YAML formatted string that provides inputs as defined in the environment
	// template bundle schema file. For more information, see Environments
	// (https://docs.aws.amazon.com/proton/latest/userguide/ag-environments.html) in
	// the Proton User Guide.
	//
	// This value conforms to the media type: application/yaml
	//
	// This member is required.
	Spec *string

	// The major version of the environment template.
	//
	// This member is required.
	TemplateMajorVersion *string

	// The name of the environment template. For more information, see Environment
	// Templates
	// (https://docs.aws.amazon.com/proton/latest/userguide/ag-templates.html) in the
	// Proton User Guide.
	//
	// This member is required.
	TemplateName *string

	// The Amazon Resource Name (ARN) of the IAM service role that allows Proton to
	// provision infrastructure using CodeBuild-based provisioning on your behalf. To
	// use CodeBuild-based provisioning for the environment or for any service instance
	// running in the environment, specify either the environmentAccountConnectionId or
	// codebuildRoleArn parameter.
	CodebuildRoleArn *string

	// The Amazon Resource Name (ARN) of the IAM service role that Proton uses when
	// provisioning directly defined components in this environment. It determines the
	// scope of infrastructure that a component can provision. You must specify
	// componentRoleArn to allow directly defined components to be associated with this
	// environment. For more information about components, see Proton components
	// (https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html) in the
	// Proton User Guide.
	ComponentRoleArn *string

	// A description of the environment that's being created and deployed.
	Description *string

	// The ID of the environment account connection that you provide if you want Proton
	// to provision infrastructure resources for your environment or for any of the
	// service instances running in it in an environment account. For more information,
	// see Environment account connections
	// (https://docs.aws.amazon.com/proton/latest/userguide/ag-env-account-connections.html)
	// in the Proton User guide. If you specify the environmentAccountConnectionId
	// parameter, don't specify protonServiceRoleArn, codebuildRoleArn, or
	// provisioningRepository.
	EnvironmentAccountConnectionId *string

	// The Amazon Resource Name (ARN) of the IAM service role that allows Proton to
	// provision infrastructure using Amazon Web Services-managed provisioning and
	// CloudFormation on your behalf. To use Amazon Web Services-managed provisioning
	// for the environment or for any service instance running in the environment,
	// specify either the environmentAccountConnectionId or protonServiceRoleArn
	// parameter.
	ProtonServiceRoleArn *string

	// The linked repository that you use to host your rendered infrastructure
	// templates for self-managed provisioning. A linked repository is a repository
	// that has been registered with Proton. For more information, see
	// CreateRepository. To use self-managed provisioning for the environment or for
	// any service instance running in the environment, specify this parameter.
	ProvisioningRepository *types.RepositoryBranchInput

	// An optional list of metadata items that you can associate with the Proton
	// environment. A tag is a key-value pair. For more information, see Proton
	// resources and tagging
	// (https://docs.aws.amazon.com/proton/latest/userguide/resources.html) in the
	// Proton User Guide.
	Tags []types.Tag

	// The minor version of the environment template.
	TemplateMinorVersion *string

	noSmithyDocumentSerde
}

type CreateEnvironmentOutput struct {

	// The environment detail data that's returned by Proton.
	//
	// This member is required.
	Environment *types.Environment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironment(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "CreateEnvironment",
	}
}

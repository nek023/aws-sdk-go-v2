// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new document object and version object. The client specifies the
// parent folder ID and name of the document to upload. The ID is optionally
// specified when creating a new version of an existing document. This is the first
// step to upload a document. Next, upload the document to the URL returned from
// the call, and then call UpdateDocumentVersion. To cancel the document upload,
// call AbortDocumentVersionUpload.
func (c *Client) InitiateDocumentVersionUpload(ctx context.Context, params *InitiateDocumentVersionUploadInput, optFns ...func(*Options)) (*InitiateDocumentVersionUploadOutput, error) {
	if params == nil {
		params = &InitiateDocumentVersionUploadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InitiateDocumentVersionUpload", params, optFns, c.addOperationInitiateDocumentVersionUploadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InitiateDocumentVersionUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InitiateDocumentVersionUploadInput struct {

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string

	// The timestamp when the content of the document was originally created.
	ContentCreatedTimestamp *time.Time

	// The timestamp when the content of the document was modified.
	ContentModifiedTimestamp *time.Time

	// The content type of the document.
	ContentType *string

	// The size of the document, in bytes.
	DocumentSizeInBytes *int64

	// The ID of the document.
	Id *string

	// The name of the document.
	Name *string

	// The ID of the parent folder.
	ParentFolderId *string

	noSmithyDocumentSerde
}

type InitiateDocumentVersionUploadOutput struct {

	// The document metadata.
	Metadata *types.DocumentMetadata

	// The upload metadata.
	UploadMetadata *types.UploadMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInitiateDocumentVersionUploadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInitiateDocumentVersionUpload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInitiateDocumentVersionUpload{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInitiateDocumentVersionUpload(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInitiateDocumentVersionUpload(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "InitiateDocumentVersionUpload",
	}
}

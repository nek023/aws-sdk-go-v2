// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// For operations that require confirmation that the email address for the
// registrant contact is valid, such as registering a new domain, this operation
// returns information about whether the registrant contact has responded. If you
// want us to resend the email, use the ResendContactReachabilityEmail operation.
func (c *Client) GetContactReachabilityStatus(ctx context.Context, params *GetContactReachabilityStatusInput, optFns ...func(*Options)) (*GetContactReachabilityStatusOutput, error) {
	if params == nil {
		params = &GetContactReachabilityStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetContactReachabilityStatus", params, optFns, c.addOperationGetContactReachabilityStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetContactReachabilityStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetContactReachabilityStatusInput struct {

	// The name of the domain for which you want to know whether the registrant
	// contact has confirmed that the email address is valid.
	DomainName *string

	noSmithyDocumentSerde
}

type GetContactReachabilityStatusOutput struct {

	// The domain name for which you requested the reachability status.
	DomainName *string

	// Whether the registrant contact has responded. Values include the following:
	// PENDING We sent the confirmation email and haven't received a response yet. DONE
	// We sent the email and got confirmation from the registrant contact. EXPIRED The
	// time limit expired before the registrant contact responded.
	Status types.ReachabilityStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetContactReachabilityStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetContactReachabilityStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetContactReachabilityStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetContactReachabilityStatus"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetContactReachabilityStatus(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetContactReachabilityStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetContactReachabilityStatus",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates member accounts of the current Amazon Web Services account by
// specifying a list of Amazon Web Services account IDs. This step is a
// prerequisite for managing the associated member accounts either by invitation or
// through an organization. As a delegated administrator, using CreateMembers will
// enable GuardDuty in the added member accounts, with the exception of the
// organization delegated administrator account. A delegated administrator must
// enable GuardDuty prior to being added as a member. If you are adding accounts by
// invitation, before using InviteMembers (https://docs.aws.amazon.com/guardduty/latest/APIReference/API_InviteMembers.html)
// , use CreateMembers after GuardDuty has been enabled in potential member
// accounts. If you disassociate a member from a GuardDuty delegated administrator,
// the member account details obtained from this API, including the associated
// email addresses, will be retained. This is done so that the delegated
// administrator can invoke the InviteMembers (https://docs.aws.amazon.com/guardduty/latest/APIReference/API_InviteMembers.html)
// API without the need to invoke the CreateMembers API again. To remove the
// details associated with a member account, the delegated administrator must
// invoke the DeleteMembers (https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DeleteMembers.html)
// API.
func (c *Client) CreateMembers(ctx context.Context, params *CreateMembersInput, optFns ...func(*Options)) (*CreateMembersOutput, error) {
	if params == nil {
		params = &CreateMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMembers", params, optFns, c.addOperationCreateMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMembersInput struct {

	// A list of account ID and email address pairs of the accounts that you want to
	// associate with the GuardDuty administrator account.
	//
	// This member is required.
	AccountDetails []types.AccountDetail

	// The unique ID of the detector of the GuardDuty account that you want to
	// associate member accounts with.
	//
	// This member is required.
	DetectorId *string

	noSmithyDocumentSerde
}

type CreateMembersOutput struct {

	// A list of objects that include the accountIds of the unprocessed accounts and a
	// result string that explains why each was unprocessed.
	//
	// This member is required.
	UnprocessedAccounts []types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMembers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateMembers"); err != nil {
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
	if err = addOpCreateMembersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMembers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateMembers",
	}
}

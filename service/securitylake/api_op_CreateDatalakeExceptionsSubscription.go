// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securitylake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates the specified notification subscription in Security Lake. Creates the
// specified subscription notifications in the specified organization.
func (c *Client) CreateDatalakeExceptionsSubscription(ctx context.Context, params *CreateDatalakeExceptionsSubscriptionInput, optFns ...func(*Options)) (*CreateDatalakeExceptionsSubscriptionOutput, error) {
	if params == nil {
		params = &CreateDatalakeExceptionsSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDatalakeExceptionsSubscription", params, optFns, c.addOperationCreateDatalakeExceptionsSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDatalakeExceptionsSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDatalakeExceptionsSubscriptionInput struct {

	// The account in which the exception notifications subscription is created.
	//
	// This member is required.
	NotificationEndpoint *string

	// The subscription protocol to which exception messages are posted.
	//
	// This member is required.
	SubscriptionProtocol types.SubscriptionProtocolType

	noSmithyDocumentSerde
}

type CreateDatalakeExceptionsSubscriptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDatalakeExceptionsSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDatalakeExceptionsSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDatalakeExceptionsSubscription{}, middleware.After)
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
	if err = addOpCreateDatalakeExceptionsSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDatalakeExceptionsSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDatalakeExceptionsSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securitylake",
		OperationName: "CreateDatalakeExceptionsSubscription",
	}
}
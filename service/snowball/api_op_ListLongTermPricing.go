// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all long-term pricing types.
func (c *Client) ListLongTermPricing(ctx context.Context, params *ListLongTermPricingInput, optFns ...func(*Options)) (*ListLongTermPricingOutput, error) {
	if params == nil {
		params = &ListLongTermPricingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLongTermPricing", params, optFns, c.addOperationListLongTermPricingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLongTermPricingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLongTermPricingInput struct {

	// The maximum number of ListLongTermPricing objects to return.
	MaxResults *int32

	// Because HTTP requests are stateless, this is the starting point for your next
	// list of ListLongTermPricing to return.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLongTermPricingOutput struct {

	// Each LongTermPricingEntry object contains a status, ID, and other information
	// about the LongTermPricing type.
	LongTermPricingEntries []types.LongTermPricingListEntry

	// Because HTTP requests are stateless, this is the starting point for your next
	// list of returned ListLongTermPricing list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLongTermPricingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListLongTermPricing{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLongTermPricing{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLongTermPricing(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListLongTermPricing(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "ListLongTermPricing",
	}
}
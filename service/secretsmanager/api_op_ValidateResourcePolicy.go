// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Validates the JSON text of the resource-based policy document attached to the
// specified secret. The JSON request string input and response output displays
// formatted code with white space and line breaks for better readability. Submit
// your input as a single line JSON string. A resource-based policy is optional.
func (c *Client) ValidateResourcePolicy(ctx context.Context, params *ValidateResourcePolicyInput, optFns ...func(*Options)) (*ValidateResourcePolicyOutput, error) {
	if params == nil {
		params = &ValidateResourcePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ValidateResourcePolicy", params, optFns, addOperationValidateResourcePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ValidateResourcePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ValidateResourcePolicyInput struct {

	// Identifies the Resource Policy attached to the secret.
	//
	// This member is required.
	ResourcePolicy *string

	// The identifier for the secret that you want to validate a resource policy. You
	// can specify either the Amazon Resource Name (ARN) or the friendly name of the
	// secret. If you specify an ARN, we generally recommend that you specify a
	// complete ARN. You can specify a partial ARN too—for example, if you don’t
	// include the final hyphen and six random characters that Secrets Manager adds at
	// the end of the ARN when you created the secret. A partial ARN match can work as
	// long as it uniquely matches only one secret. However, if your secret has a name
	// that ends in a hyphen followed by six characters (before Secrets Manager adds
	// the hyphen and six characters to the ARN) and you try to use that as a partial
	// ARN, then those characters cause Secrets Manager to assume that you’re
	// specifying a complete ARN. This confusion can cause unexpected results. To avoid
	// this situation, we recommend that you don’t create secret names ending with a
	// hyphen followed by six characters. If you specify an incomplete ARN without the
	// random suffix, and instead provide the 'friendly name', you must not include the
	// random suffix. If you do include the random suffix added by Secrets Manager, you
	// receive either a ResourceNotFoundException or an AccessDeniedException error,
	// depending on your permissions.
	SecretId *string
}

type ValidateResourcePolicyOutput struct {

	// Returns a message stating that your Reource Policy passed validation.
	PolicyValidationPassed bool

	// Returns an error message if your policy doesn't pass validatation.
	ValidationErrors []types.ValidationErrorsEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationValidateResourcePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpValidateResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpValidateResourcePolicy{}, middleware.After)
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
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpValidateResourcePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opValidateResourcePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opValidateResourcePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "ValidateResourcePolicy",
	}
}
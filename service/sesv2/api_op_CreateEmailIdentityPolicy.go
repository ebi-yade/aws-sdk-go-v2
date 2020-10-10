// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates the specified sending authorization policy for the given identity (an
// email address or a domain). This API is for the identity owner only. If you have
// not verified the identity, this API will return an error. Sending authorization
// is a feature that enables an identity owner to authorize other senders to use
// its identities. For information about using sending authorization, see the
// Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
// <p>You can execute this operation no more than once per second.</p>
func (c *Client) CreateEmailIdentityPolicy(ctx context.Context, params *CreateEmailIdentityPolicyInput, optFns ...func(*Options)) (*CreateEmailIdentityPolicyOutput, error) {
	if params == nil {
		params = &CreateEmailIdentityPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEmailIdentityPolicy", params, optFns, addOperationCreateEmailIdentityPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEmailIdentityPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to create a sending authorization policy for an identity.
// Sending authorization is an Amazon SES feature that enables you to authorize
// other senders to use your identities. For information, see the Amazon SES
// Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization-identity-owner-tasks-management.html).
type CreateEmailIdentityPolicyInput struct {

	// The email identity for which you want to create a policy.
	//
	// This member is required.
	EmailIdentity *string

	// The text of the policy in JSON format. The policy cannot exceed 4 KB. For
	// information about the syntax of sending authorization policies, see the Amazon
	// SES Developer Guide
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization-policies.html).
	//
	// This member is required.
	Policy *string

	// The name of the policy.  <p>The policy name cannot exceed 64 characters and can
	// only include alphanumeric characters, dashes, and underscores.</p>
	//
	// This member is required.
	PolicyName *string
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type CreateEmailIdentityPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateEmailIdentityPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateEmailIdentityPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateEmailIdentityPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateEmailIdentityPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEmailIdentityPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateEmailIdentityPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "CreateEmailIdentityPolicy",
	}
}
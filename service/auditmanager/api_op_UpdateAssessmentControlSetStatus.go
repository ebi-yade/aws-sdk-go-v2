// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the status of a control set in an AWS Audit Manager assessment.
func (c *Client) UpdateAssessmentControlSetStatus(ctx context.Context, params *UpdateAssessmentControlSetStatusInput, optFns ...func(*Options)) (*UpdateAssessmentControlSetStatusOutput, error) {
	if params == nil {
		params = &UpdateAssessmentControlSetStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAssessmentControlSetStatus", params, optFns, addOperationUpdateAssessmentControlSetStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAssessmentControlSetStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAssessmentControlSetStatusInput struct {

	// The identifier for the specified assessment.
	//
	// This member is required.
	AssessmentId *string

	// The comment related to the status update.
	//
	// This member is required.
	Comment *string

	// The identifier for the specified control set.
	//
	// This member is required.
	ControlSetId *string

	// The status of the control set that is being updated.
	//
	// This member is required.
	Status types.ControlSetStatus
}

type UpdateAssessmentControlSetStatusOutput struct {

	// The name of the updated control set returned by the
	// UpdateAssessmentControlSetStatus API.
	ControlSet *types.AssessmentControlSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateAssessmentControlSetStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAssessmentControlSetStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAssessmentControlSetStatus{}, middleware.After)
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
	if err = addOpUpdateAssessmentControlSetStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAssessmentControlSetStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAssessmentControlSetStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "UpdateAssessmentControlSetStatus",
	}
}
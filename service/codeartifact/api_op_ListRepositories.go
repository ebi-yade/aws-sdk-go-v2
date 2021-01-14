// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of RepositorySummary
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_RepositorySummary.html)
// objects. Each RepositorySummary contains information about a repository in the
// specified AWS account and that matches the input parameters.
func (c *Client) ListRepositories(ctx context.Context, params *ListRepositoriesInput, optFns ...func(*Options)) (*ListRepositoriesOutput, error) {
	if params == nil {
		params = &ListRepositoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRepositories", params, optFns, addOperationListRepositoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRepositoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRepositoriesInput struct {

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// A prefix used to filter returned repositories. Only repositories with names that
	// start with repositoryPrefix are returned.
	RepositoryPrefix *string
}

type ListRepositoriesOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// The returned list of RepositorySummary
	// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_RepositorySummary.html)
	// objects.
	Repositories []types.RepositorySummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListRepositoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRepositories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRepositories{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRepositories(options.Region), middleware.Before); err != nil {
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

// ListRepositoriesAPIClient is a client that implements the ListRepositories
// operation.
type ListRepositoriesAPIClient interface {
	ListRepositories(context.Context, *ListRepositoriesInput, ...func(*Options)) (*ListRepositoriesOutput, error)
}

var _ ListRepositoriesAPIClient = (*Client)(nil)

// ListRepositoriesPaginatorOptions is the paginator options for ListRepositories
type ListRepositoriesPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRepositoriesPaginator is a paginator for ListRepositories
type ListRepositoriesPaginator struct {
	options   ListRepositoriesPaginatorOptions
	client    ListRepositoriesAPIClient
	params    *ListRepositoriesInput
	nextToken *string
	firstPage bool
}

// NewListRepositoriesPaginator returns a new ListRepositoriesPaginator
func NewListRepositoriesPaginator(client ListRepositoriesAPIClient, params *ListRepositoriesInput, optFns ...func(*ListRepositoriesPaginatorOptions)) *ListRepositoriesPaginator {
	options := ListRepositoriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListRepositoriesInput{}
	}

	return &ListRepositoriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRepositoriesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListRepositories page.
func (p *ListRepositoriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRepositoriesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListRepositories(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListRepositories(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "ListRepositories",
	}
}
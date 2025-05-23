// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroups

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCancelTagSyncTask struct {
}

func (*validateOpCancelTagSyncTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelTagSyncTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelTagSyncTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelTagSyncTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateGroup struct {
}

func (*validateOpCreateGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTags struct {
}

func (*validateOpGetTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTagSyncTask struct {
}

func (*validateOpGetTagSyncTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTagSyncTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTagSyncTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTagSyncTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGroupResources struct {
}

func (*validateOpGroupResources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGroupResources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GroupResourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGroupResourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListGroupingStatuses struct {
}

func (*validateOpListGroupingStatuses) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListGroupingStatuses) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListGroupingStatusesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListGroupingStatusesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListGroupResources struct {
}

func (*validateOpListGroupResources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListGroupResources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListGroupResourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListGroupResourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListGroups struct {
}

func (*validateOpListGroups) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListGroups) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListGroupsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListGroupsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutGroupConfiguration struct {
}

func (*validateOpPutGroupConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutGroupConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutGroupConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutGroupConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSearchResources struct {
}

func (*validateOpSearchResources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSearchResources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SearchResourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSearchResourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartTagSyncTask struct {
}

func (*validateOpStartTagSyncTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartTagSyncTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartTagSyncTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartTagSyncTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTag struct {
}

func (*validateOpTag) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTag) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUngroupResources struct {
}

func (*validateOpUngroupResources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUngroupResources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UngroupResourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUngroupResourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntag struct {
}

func (*validateOpUntag) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntag) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateGroupQuery struct {
}

func (*validateOpUpdateGroupQuery) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateGroupQuery) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateGroupQueryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateGroupQueryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCancelTagSyncTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelTagSyncTask{}, middleware.After)
}

func addOpCreateGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateGroup{}, middleware.After)
}

func addOpGetTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTags{}, middleware.After)
}

func addOpGetTagSyncTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTagSyncTask{}, middleware.After)
}

func addOpGroupResourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGroupResources{}, middleware.After)
}

func addOpListGroupingStatusesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListGroupingStatuses{}, middleware.After)
}

func addOpListGroupResourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListGroupResources{}, middleware.After)
}

func addOpListGroupsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListGroups{}, middleware.After)
}

func addOpPutGroupConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutGroupConfiguration{}, middleware.After)
}

func addOpSearchResourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSearchResources{}, middleware.After)
}

func addOpStartTagSyncTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartTagSyncTask{}, middleware.After)
}

func addOpTagValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTag{}, middleware.After)
}

func addOpUngroupResourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUngroupResources{}, middleware.After)
}

func addOpUntagValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntag{}, middleware.After)
}

func addOpUpdateGroupQueryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateGroupQuery{}, middleware.After)
}

func validateGroupConfigurationItem(v *types.GroupConfigurationItem) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GroupConfigurationItem"}
	if v.Type == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.Parameters != nil {
		if err := validateGroupParameterList(v.Parameters); err != nil {
			invalidParams.AddNested("Parameters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateGroupConfigurationList(v []types.GroupConfigurationItem) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GroupConfigurationList"}
	for i := range v {
		if err := validateGroupConfigurationItem(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateGroupConfigurationParameter(v *types.GroupConfigurationParameter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GroupConfigurationParameter"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateGroupFilter(v *types.GroupFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GroupFilter"}
	if len(v.Name) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Values == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Values"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateGroupFilterList(v []types.GroupFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GroupFilterList"}
	for i := range v {
		if err := validateGroupFilter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateGroupParameterList(v []types.GroupConfigurationParameter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GroupParameterList"}
	for i := range v {
		if err := validateGroupConfigurationParameter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateListGroupingStatusesFilter(v *types.ListGroupingStatusesFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListGroupingStatusesFilter"}
	if len(v.Name) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Values == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Values"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateListGroupingStatusesFilterList(v []types.ListGroupingStatusesFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListGroupingStatusesFilterList"}
	for i := range v {
		if err := validateListGroupingStatusesFilter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateResourceFilter(v *types.ResourceFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResourceFilter"}
	if len(v.Name) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Values == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Values"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateResourceFilterList(v []types.ResourceFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResourceFilterList"}
	for i := range v {
		if err := validateResourceFilter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateResourceQuery(v *types.ResourceQuery) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResourceQuery"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.Query == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Query"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelTagSyncTaskInput(v *CancelTagSyncTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelTagSyncTaskInput"}
	if v.TaskArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateGroupInput(v *CreateGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateGroupInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.ResourceQuery != nil {
		if err := validateResourceQuery(v.ResourceQuery); err != nil {
			invalidParams.AddNested("ResourceQuery", err.(smithy.InvalidParamsError))
		}
	}
	if v.Configuration != nil {
		if err := validateGroupConfigurationList(v.Configuration); err != nil {
			invalidParams.AddNested("Configuration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTagsInput(v *GetTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTagsInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTagSyncTaskInput(v *GetTagSyncTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTagSyncTaskInput"}
	if v.TaskArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGroupResourcesInput(v *GroupResourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GroupResourcesInput"}
	if v.Group == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Group"))
	}
	if v.ResourceArns == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArns"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListGroupingStatusesInput(v *ListGroupingStatusesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListGroupingStatusesInput"}
	if v.Group == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Group"))
	}
	if v.Filters != nil {
		if err := validateListGroupingStatusesFilterList(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListGroupResourcesInput(v *ListGroupResourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListGroupResourcesInput"}
	if v.Filters != nil {
		if err := validateResourceFilterList(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListGroupsInput(v *ListGroupsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListGroupsInput"}
	if v.Filters != nil {
		if err := validateGroupFilterList(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutGroupConfigurationInput(v *PutGroupConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutGroupConfigurationInput"}
	if v.Configuration != nil {
		if err := validateGroupConfigurationList(v.Configuration); err != nil {
			invalidParams.AddNested("Configuration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSearchResourcesInput(v *SearchResourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchResourcesInput"}
	if v.ResourceQuery == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceQuery"))
	} else if v.ResourceQuery != nil {
		if err := validateResourceQuery(v.ResourceQuery); err != nil {
			invalidParams.AddNested("ResourceQuery", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartTagSyncTaskInput(v *StartTagSyncTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartTagSyncTaskInput"}
	if v.Group == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Group"))
	}
	if v.ResourceQuery != nil {
		if err := validateResourceQuery(v.ResourceQuery); err != nil {
			invalidParams.AddNested("ResourceQuery", err.(smithy.InvalidParamsError))
		}
	}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagInput(v *TagInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUngroupResourcesInput(v *UngroupResourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UngroupResourcesInput"}
	if v.Group == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Group"))
	}
	if v.ResourceArns == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArns"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagInput(v *UntagInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if v.Keys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Keys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateGroupQueryInput(v *UpdateGroupQueryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateGroupQueryInput"}
	if v.ResourceQuery == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceQuery"))
	} else if v.ResourceQuery != nil {
		if err := validateResourceQuery(v.ResourceQuery); err != nil {
			invalidParams.AddNested("ResourceQuery", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

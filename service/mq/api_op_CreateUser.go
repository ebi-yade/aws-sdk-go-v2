// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mq

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/CreateUserRequest
type CreateUserInput struct {
	_ struct{} `type:"structure"`

	// BrokerId is a required field
	BrokerId *string `location:"uri" locationName:"broker-id" type:"string" required:"true"`

	ConsoleAccess *bool `locationName:"consoleAccess" type:"boolean"`

	Groups []string `locationName:"groups" type:"list"`

	Password *string `locationName:"password" type:"string"`

	// Username is a required field
	Username *string `location:"uri" locationName:"username" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateUserInput"}

	if s.BrokerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BrokerId"))
	}

	if s.Username == nil {
		invalidParams.Add(aws.NewErrParamRequired("Username"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateUserInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ConsoleAccess != nil {
		v := *s.ConsoleAccess

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "consoleAccess", protocol.BoolValue(v), metadata)
	}
	if s.Groups != nil {
		v := s.Groups

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "groups", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Password != nil {
		v := *s.Password

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "password", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BrokerId != nil {
		v := *s.BrokerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "broker-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Username != nil {
		v := *s.Username

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "username", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/CreateUserResponse
type CreateUserOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateUserOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateUserOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCreateUser = "CreateUser"

// CreateUserRequest returns a request value for making API operation for
// AmazonMQ.
//
// Creates an ActiveMQ user.
//
//    // Example sending a request using CreateUserRequest.
//    req := client.CreateUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/CreateUser
func (c *Client) CreateUserRequest(input *CreateUserInput) CreateUserRequest {
	op := &aws.Operation{
		Name:       opCreateUser,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/brokers/{broker-id}/users/{username}",
	}

	if input == nil {
		input = &CreateUserInput{}
	}

	req := c.newRequest(op, input, &CreateUserOutput{})
	return CreateUserRequest{Request: req, Input: input, Copy: c.CreateUserRequest}
}

// CreateUserRequest is the request type for the
// CreateUser API operation.
type CreateUserRequest struct {
	*aws.Request
	Input *CreateUserInput
	Copy  func(*CreateUserInput) CreateUserRequest
}

// Send marshals and sends the CreateUser API request.
func (r CreateUserRequest) Send(ctx context.Context) (*CreateUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateUserResponse{
		CreateUserOutput: r.Request.Data.(*CreateUserOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateUserResponse is the response type for the
// CreateUser API operation.
type CreateUserResponse struct {
	*CreateUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateUser request.
func (r *CreateUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
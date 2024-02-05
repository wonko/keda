// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodbstreams

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about a stream, including the current status of the stream,
// its Amazon Resource Name (ARN), the composition of its shards, and its
// corresponding DynamoDB table. You can call DescribeStream at a maximum rate of
// 10 times per second. Each shard in the stream has a SequenceNumberRange
// associated with it. If the SequenceNumberRange has a StartingSequenceNumber but
// no EndingSequenceNumber , then the shard is still open (able to receive more
// stream records). If both StartingSequenceNumber and EndingSequenceNumber are
// present, then that shard is closed and can no longer receive more data.
func (c *Client) DescribeStream(ctx context.Context, params *DescribeStreamInput, optFns ...func(*Options)) (*DescribeStreamOutput, error) {
	if params == nil {
		params = &DescribeStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStream", params, optFns, c.addOperationDescribeStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DescribeStream operation.
type DescribeStreamInput struct {

	// The Amazon Resource Name (ARN) for the stream.
	//
	// This member is required.
	StreamArn *string

	// The shard ID of the first item that this operation will evaluate. Use the value
	// that was returned for LastEvaluatedShardId in the previous operation.
	ExclusiveStartShardId *string

	// The maximum number of shard objects to return. The upper limit is 100.
	Limit *int32

	noSmithyDocumentSerde
}

// Represents the output of a DescribeStream operation.
type DescribeStreamOutput struct {

	// A complete description of the stream, including its creation date and time, the
	// DynamoDB table associated with the stream, the shard IDs within the stream, and
	// the beginning and ending sequence numbers of stream records within the shards.
	StreamDescription *types.StreamDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeStream"); err != nil {
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
	if err = addOpDescribeStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeStream",
	}
}

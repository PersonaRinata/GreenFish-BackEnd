// Code generated by Kitex v0.6.2. DO NOT EDIT.

package chatservice

import (
	chat "GoYin/server/kitex_gen/chat"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetChatHistory(ctx context.Context, req *chat.DouyinMessageGetChatHistoryRequest, callOptions ...callopt.Option) (r *chat.DouyinMessageGetChatHistoryResponse, err error)
	SentMessage(ctx context.Context, req *chat.DouyinMessageActionRequest, callOptions ...callopt.Option) (r *chat.DouyinMessageActionResponse, err error)
	GetLatestMessage(ctx context.Context, req *chat.DouyinMessageGetLatestRequest, callOptions ...callopt.Option) (r *chat.DouyinMessageGetLatestResponse, err error)
	BatchGetLatestMessage(ctx context.Context, req *chat.DouyinMessageBatchGetLatestRequest, callOptions ...callopt.Option) (r *chat.DouyinMessageBatchGetLatestResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kChatServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kChatServiceClient struct {
	*kClient
}

func (p *kChatServiceClient) GetChatHistory(ctx context.Context, req *chat.DouyinMessageGetChatHistoryRequest, callOptions ...callopt.Option) (r *chat.DouyinMessageGetChatHistoryResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetChatHistory(ctx, req)
}

func (p *kChatServiceClient) SentMessage(ctx context.Context, req *chat.DouyinMessageActionRequest, callOptions ...callopt.Option) (r *chat.DouyinMessageActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SentMessage(ctx, req)
}

func (p *kChatServiceClient) GetLatestMessage(ctx context.Context, req *chat.DouyinMessageGetLatestRequest, callOptions ...callopt.Option) (r *chat.DouyinMessageGetLatestResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLatestMessage(ctx, req)
}

func (p *kChatServiceClient) BatchGetLatestMessage(ctx context.Context, req *chat.DouyinMessageBatchGetLatestRequest, callOptions ...callopt.Option) (r *chat.DouyinMessageBatchGetLatestResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BatchGetLatestMessage(ctx, req)
}

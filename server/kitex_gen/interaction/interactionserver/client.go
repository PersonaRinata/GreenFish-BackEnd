// Code generated by Kitex v0.6.2. DO NOT EDIT.

package interactionserver

import (
	interaction "GoYin/server/kitex_gen/interaction"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Favorite(ctx context.Context, req *interaction.DouyinFavoriteActionRequest, callOptions ...callopt.Option) (r *interaction.DouyinFavoriteActionResponse, err error)
	GetFavoriteVideoIdList(ctx context.Context, req *interaction.DouyinGetFavoriteVideoIdListRequest, callOptions ...callopt.Option) (r *interaction.DouyinGetFavoriteVideoIdListResponse, err error)
	Comment(ctx context.Context, req *interaction.DouyinCommentActionRequest, callOptions ...callopt.Option) (r *interaction.DouyinCommentActionResponse, err error)
	GetCommentList(ctx context.Context, req *interaction.DouyinGetCommentListRequest, callOptions ...callopt.Option) (r *interaction.DouyinGetCommentListResponse, err error)
	GetVideoInteractInfo(ctx context.Context, req *interaction.DouyinGetVideoInteractInfoRequest, callOptions ...callopt.Option) (r *interaction.DouyinGetVideoInteractInfoResponse, err error)
	BatchGetVideoInteractInfo(ctx context.Context, req *interaction.DouyinBatchGetVideoInteractInfoRequest, callOptions ...callopt.Option) (r *interaction.DouyinBatchGetVideoInteractInfoResponse, err error)
	GetUserInteractInfo(ctx context.Context, req *interaction.DouyinGetUserInteractInfoRequest, callOptions ...callopt.Option) (r *interaction.DouyinGetUserInteractInfoResponse, err error)
	BatchGetUserInteractInfo(ctx context.Context, req *interaction.DouyinBatchGetUserInteractInfoRequest, callOptions ...callopt.Option) (r *interaction.DouyinBatchGetUserInteractInfoResponse, err error)
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
	return &kInteractionServerClient{
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

type kInteractionServerClient struct {
	*kClient
}

func (p *kInteractionServerClient) Favorite(ctx context.Context, req *interaction.DouyinFavoriteActionRequest, callOptions ...callopt.Option) (r *interaction.DouyinFavoriteActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Favorite(ctx, req)
}

func (p *kInteractionServerClient) GetFavoriteVideoIdList(ctx context.Context, req *interaction.DouyinGetFavoriteVideoIdListRequest, callOptions ...callopt.Option) (r *interaction.DouyinGetFavoriteVideoIdListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFavoriteVideoIdList(ctx, req)
}

func (p *kInteractionServerClient) Comment(ctx context.Context, req *interaction.DouyinCommentActionRequest, callOptions ...callopt.Option) (r *interaction.DouyinCommentActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Comment(ctx, req)
}

func (p *kInteractionServerClient) GetCommentList(ctx context.Context, req *interaction.DouyinGetCommentListRequest, callOptions ...callopt.Option) (r *interaction.DouyinGetCommentListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCommentList(ctx, req)
}

func (p *kInteractionServerClient) GetVideoInteractInfo(ctx context.Context, req *interaction.DouyinGetVideoInteractInfoRequest, callOptions ...callopt.Option) (r *interaction.DouyinGetVideoInteractInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideoInteractInfo(ctx, req)
}

func (p *kInteractionServerClient) BatchGetVideoInteractInfo(ctx context.Context, req *interaction.DouyinBatchGetVideoInteractInfoRequest, callOptions ...callopt.Option) (r *interaction.DouyinBatchGetVideoInteractInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BatchGetVideoInteractInfo(ctx, req)
}

func (p *kInteractionServerClient) GetUserInteractInfo(ctx context.Context, req *interaction.DouyinGetUserInteractInfoRequest, callOptions ...callopt.Option) (r *interaction.DouyinGetUserInteractInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInteractInfo(ctx, req)
}

func (p *kInteractionServerClient) BatchGetUserInteractInfo(ctx context.Context, req *interaction.DouyinBatchGetUserInteractInfoRequest, callOptions ...callopt.Option) (r *interaction.DouyinBatchGetUserInteractInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BatchGetUserInteractInfo(ctx, req)
}
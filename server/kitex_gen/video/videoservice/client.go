// Code generated by Kitex v0.7.3. DO NOT EDIT.

package videoservice

import (
	video "GreenFish/server/kitex_gen/video"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Feed(ctx context.Context, req *video.QingyuFeedRequest, callOptions ...callopt.Option) (r *video.QingyuFeedResponse, err error)
	PublishVideo(ctx context.Context, req *video.QingyuPublishActionRequest, callOptions ...callopt.Option) (r *video.QingyuPublishActionResponse, err error)
	GetPublishedVideoList(ctx context.Context, req *video.QingyuGetPublishedListRequest, callOptions ...callopt.Option) (r *video.QingyuGetPublishedListResponse, err error)
	GetFavoriteVideoList(ctx context.Context, req *video.QingyuGetFavoriteListRequest, callOptions ...callopt.Option) (r *video.QingyuGetFavoriteListResponse, err error)
	GetPublishedVideoIdList(ctx context.Context, req *video.QingyuGetPublishedVideoIdListRequest, callOptions ...callopt.Option) (r *video.QingyuGetPublishedVideoIdListResponse, err error)
	SearchVideoList(ctx context.Context, req *video.QingyuSearchVideoRequest, callOptions ...callopt.Option) (r *video.QingyuSearchVideoResponse, err error)
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
	return &kVideoServiceClient{
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

type kVideoServiceClient struct {
	*kClient
}

func (p *kVideoServiceClient) Feed(ctx context.Context, req *video.QingyuFeedRequest, callOptions ...callopt.Option) (r *video.QingyuFeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Feed(ctx, req)
}

func (p *kVideoServiceClient) PublishVideo(ctx context.Context, req *video.QingyuPublishActionRequest, callOptions ...callopt.Option) (r *video.QingyuPublishActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishVideo(ctx, req)
}

func (p *kVideoServiceClient) GetPublishedVideoList(ctx context.Context, req *video.QingyuGetPublishedListRequest, callOptions ...callopt.Option) (r *video.QingyuGetPublishedListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPublishedVideoList(ctx, req)
}

func (p *kVideoServiceClient) GetFavoriteVideoList(ctx context.Context, req *video.QingyuGetFavoriteListRequest, callOptions ...callopt.Option) (r *video.QingyuGetFavoriteListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFavoriteVideoList(ctx, req)
}

func (p *kVideoServiceClient) GetPublishedVideoIdList(ctx context.Context, req *video.QingyuGetPublishedVideoIdListRequest, callOptions ...callopt.Option) (r *video.QingyuGetPublishedVideoIdListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPublishedVideoIdList(ctx, req)
}

func (p *kVideoServiceClient) SearchVideoList(ctx context.Context, req *video.QingyuSearchVideoRequest, callOptions ...callopt.Option) (r *video.QingyuSearchVideoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SearchVideoList(ctx, req)
}

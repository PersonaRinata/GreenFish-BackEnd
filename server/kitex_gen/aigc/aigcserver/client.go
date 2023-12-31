// Code generated by Kitex v0.7.3. DO NOT EDIT.

package aigcserver

import (
	aigc "GreenFish/server/kitex_gen/aigc"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UserAskQuestion(ctx context.Context, req *aigc.QingyuAigcQuestionRequest, callOptions ...callopt.Option) (r *aigc.QingyuAigcQuestionResponse, err error)
	AnalyseIssueList(ctx context.Context, req *aigc.QingyuAigcIssueListRequest, callOptions ...callopt.Option) (r *aigc.QingyuAigcIssueListResponse, err error)
	ChooseWord(ctx context.Context, req *aigc.QingyuAigcChooseWordRequest, callOptions ...callopt.Option) (r *aigc.QingyuAigcChooseWordResponse, err error)
	DoctorAnalyse(ctx context.Context, req *aigc.QingyuAigcDoctorAnalyseRequest, callOptions ...callopt.Option) (r *aigc.QingyuAigcDoctorAnalyseResponse, err error)
	GetAIGCHistory(ctx context.Context, req *aigc.QingyuAigcGetHistoryRequest, callOptions ...callopt.Option) (r *aigc.QingyuAigcGetHistoryResponse, err error)
	RecommendDoctor(ctx context.Context, req *aigc.QingyuAigcRecommendDocotorRequest, callOptions ...callopt.Option) (r *aigc.QingyuAigcRecommendDocotorResponse, err error)
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
	return &kAIGCServerClient{
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

type kAIGCServerClient struct {
	*kClient
}

func (p *kAIGCServerClient) UserAskQuestion(ctx context.Context, req *aigc.QingyuAigcQuestionRequest, callOptions ...callopt.Option) (r *aigc.QingyuAigcQuestionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserAskQuestion(ctx, req)
}

func (p *kAIGCServerClient) AnalyseIssueList(ctx context.Context, req *aigc.QingyuAigcIssueListRequest, callOptions ...callopt.Option) (r *aigc.QingyuAigcIssueListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AnalyseIssueList(ctx, req)
}

func (p *kAIGCServerClient) ChooseWord(ctx context.Context, req *aigc.QingyuAigcChooseWordRequest, callOptions ...callopt.Option) (r *aigc.QingyuAigcChooseWordResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChooseWord(ctx, req)
}

func (p *kAIGCServerClient) DoctorAnalyse(ctx context.Context, req *aigc.QingyuAigcDoctorAnalyseRequest, callOptions ...callopt.Option) (r *aigc.QingyuAigcDoctorAnalyseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DoctorAnalyse(ctx, req)
}

func (p *kAIGCServerClient) GetAIGCHistory(ctx context.Context, req *aigc.QingyuAigcGetHistoryRequest, callOptions ...callopt.Option) (r *aigc.QingyuAigcGetHistoryResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAIGCHistory(ctx, req)
}

func (p *kAIGCServerClient) RecommendDoctor(ctx context.Context, req *aigc.QingyuAigcRecommendDocotorRequest, callOptions ...callopt.Option) (r *aigc.QingyuAigcRecommendDocotorResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RecommendDoctor(ctx, req)
}

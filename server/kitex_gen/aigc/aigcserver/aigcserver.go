// Code generated by Kitex v0.7.3. DO NOT EDIT.

package aigcserver

import (
	aigc "GreenFish/server/kitex_gen/aigc"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return aIGCServerServiceInfo
}

var aIGCServerServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "AIGCServer"
	handlerType := (*aigc.AIGCServer)(nil)
	methods := map[string]kitex.MethodInfo{
		"UserAskQuestion":  kitex.NewMethodInfo(userAskQuestionHandler, newAIGCServerUserAskQuestionArgs, newAIGCServerUserAskQuestionResult, false),
		"AnalyseIssueList": kitex.NewMethodInfo(analyseIssueListHandler, newAIGCServerAnalyseIssueListArgs, newAIGCServerAnalyseIssueListResult, false),
		"ChooseWord":       kitex.NewMethodInfo(chooseWordHandler, newAIGCServerChooseWordArgs, newAIGCServerChooseWordResult, false),
		"DoctorAnalyse":    kitex.NewMethodInfo(doctorAnalyseHandler, newAIGCServerDoctorAnalyseArgs, newAIGCServerDoctorAnalyseResult, false),
		"GetAIGCHistory":   kitex.NewMethodInfo(getAIGCHistoryHandler, newAIGCServerGetAIGCHistoryArgs, newAIGCServerGetAIGCHistoryResult, false),
		"RecommendDoctor":  kitex.NewMethodInfo(recommendDoctorHandler, newAIGCServerRecommendDoctorArgs, newAIGCServerRecommendDoctorResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "aigc",
		"ServiceFilePath": `../../idl/aigc.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.3",
		Extra:           extra,
	}
	return svcInfo
}

func userAskQuestionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*aigc.AIGCServerUserAskQuestionArgs)
	realResult := result.(*aigc.AIGCServerUserAskQuestionResult)
	success, err := handler.(aigc.AIGCServer).UserAskQuestion(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAIGCServerUserAskQuestionArgs() interface{} {
	return aigc.NewAIGCServerUserAskQuestionArgs()
}

func newAIGCServerUserAskQuestionResult() interface{} {
	return aigc.NewAIGCServerUserAskQuestionResult()
}

func analyseIssueListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*aigc.AIGCServerAnalyseIssueListArgs)
	realResult := result.(*aigc.AIGCServerAnalyseIssueListResult)
	success, err := handler.(aigc.AIGCServer).AnalyseIssueList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAIGCServerAnalyseIssueListArgs() interface{} {
	return aigc.NewAIGCServerAnalyseIssueListArgs()
}

func newAIGCServerAnalyseIssueListResult() interface{} {
	return aigc.NewAIGCServerAnalyseIssueListResult()
}

func chooseWordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*aigc.AIGCServerChooseWordArgs)
	realResult := result.(*aigc.AIGCServerChooseWordResult)
	success, err := handler.(aigc.AIGCServer).ChooseWord(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAIGCServerChooseWordArgs() interface{} {
	return aigc.NewAIGCServerChooseWordArgs()
}

func newAIGCServerChooseWordResult() interface{} {
	return aigc.NewAIGCServerChooseWordResult()
}

func doctorAnalyseHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*aigc.AIGCServerDoctorAnalyseArgs)
	realResult := result.(*aigc.AIGCServerDoctorAnalyseResult)
	success, err := handler.(aigc.AIGCServer).DoctorAnalyse(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAIGCServerDoctorAnalyseArgs() interface{} {
	return aigc.NewAIGCServerDoctorAnalyseArgs()
}

func newAIGCServerDoctorAnalyseResult() interface{} {
	return aigc.NewAIGCServerDoctorAnalyseResult()
}

func getAIGCHistoryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*aigc.AIGCServerGetAIGCHistoryArgs)
	realResult := result.(*aigc.AIGCServerGetAIGCHistoryResult)
	success, err := handler.(aigc.AIGCServer).GetAIGCHistory(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAIGCServerGetAIGCHistoryArgs() interface{} {
	return aigc.NewAIGCServerGetAIGCHistoryArgs()
}

func newAIGCServerGetAIGCHistoryResult() interface{} {
	return aigc.NewAIGCServerGetAIGCHistoryResult()
}

func recommendDoctorHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*aigc.AIGCServerRecommendDoctorArgs)
	realResult := result.(*aigc.AIGCServerRecommendDoctorResult)
	success, err := handler.(aigc.AIGCServer).RecommendDoctor(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAIGCServerRecommendDoctorArgs() interface{} {
	return aigc.NewAIGCServerRecommendDoctorArgs()
}

func newAIGCServerRecommendDoctorResult() interface{} {
	return aigc.NewAIGCServerRecommendDoctorResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) UserAskQuestion(ctx context.Context, req *aigc.QingyuAigcQuestionRequest) (r *aigc.QingyuAigcQuestionResponse, err error) {
	var _args aigc.AIGCServerUserAskQuestionArgs
	_args.Req = req
	var _result aigc.AIGCServerUserAskQuestionResult
	if err = p.c.Call(ctx, "UserAskQuestion", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AnalyseIssueList(ctx context.Context, req *aigc.QingyuAigcIssueListRequest) (r *aigc.QingyuAigcIssueListResponse, err error) {
	var _args aigc.AIGCServerAnalyseIssueListArgs
	_args.Req = req
	var _result aigc.AIGCServerAnalyseIssueListResult
	if err = p.c.Call(ctx, "AnalyseIssueList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChooseWord(ctx context.Context, req *aigc.QingyuAigcChooseWordRequest) (r *aigc.QingyuAigcChooseWordResponse, err error) {
	var _args aigc.AIGCServerChooseWordArgs
	_args.Req = req
	var _result aigc.AIGCServerChooseWordResult
	if err = p.c.Call(ctx, "ChooseWord", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DoctorAnalyse(ctx context.Context, req *aigc.QingyuAigcDoctorAnalyseRequest) (r *aigc.QingyuAigcDoctorAnalyseResponse, err error) {
	var _args aigc.AIGCServerDoctorAnalyseArgs
	_args.Req = req
	var _result aigc.AIGCServerDoctorAnalyseResult
	if err = p.c.Call(ctx, "DoctorAnalyse", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAIGCHistory(ctx context.Context, req *aigc.QingyuAigcGetHistoryRequest) (r *aigc.QingyuAigcGetHistoryResponse, err error) {
	var _args aigc.AIGCServerGetAIGCHistoryArgs
	_args.Req = req
	var _result aigc.AIGCServerGetAIGCHistoryResult
	if err = p.c.Call(ctx, "GetAIGCHistory", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RecommendDoctor(ctx context.Context, req *aigc.QingyuAigcRecommendDocotorRequest) (r *aigc.QingyuAigcRecommendDocotorResponse, err error) {
	var _args aigc.AIGCServerRecommendDoctorArgs
	_args.Req = req
	var _result aigc.AIGCServerRecommendDoctorResult
	if err = p.c.Call(ctx, "RecommendDoctor", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

// Code generated by Kitex v0.6.2. DO NOT EDIT.

package socialityservice

import (
	sociality "GreenFish/server/kitex_gen/sociality"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return socialityServiceServiceInfo
}

var socialityServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "SocialityService"
	handlerType := (*sociality.SocialityService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Action":             kitex.NewMethodInfo(actionHandler, newSocialityServiceActionArgs, newSocialityServiceActionResult, false),
		"GetRelationIdList":  kitex.NewMethodInfo(getRelationIdListHandler, newSocialityServiceGetRelationIdListArgs, newSocialityServiceGetRelationIdListResult, false),
		"GetSocialInfo":      kitex.NewMethodInfo(getSocialInfoHandler, newSocialityServiceGetSocialInfoArgs, newSocialityServiceGetSocialInfoResult, false),
		"BatchGetSocialInfo": kitex.NewMethodInfo(batchGetSocialInfoHandler, newSocialityServiceBatchGetSocialInfoArgs, newSocialityServiceBatchGetSocialInfoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "sociality",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func actionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sociality.SocialityServiceActionArgs)
	realResult := result.(*sociality.SocialityServiceActionResult)
	success, err := handler.(sociality.SocialityService).Action(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialityServiceActionArgs() interface{} {
	return sociality.NewSocialityServiceActionArgs()
}

func newSocialityServiceActionResult() interface{} {
	return sociality.NewSocialityServiceActionResult()
}

func getRelationIdListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sociality.SocialityServiceGetRelationIdListArgs)
	realResult := result.(*sociality.SocialityServiceGetRelationIdListResult)
	success, err := handler.(sociality.SocialityService).GetRelationIdList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialityServiceGetRelationIdListArgs() interface{} {
	return sociality.NewSocialityServiceGetRelationIdListArgs()
}

func newSocialityServiceGetRelationIdListResult() interface{} {
	return sociality.NewSocialityServiceGetRelationIdListResult()
}

func getSocialInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sociality.SocialityServiceGetSocialInfoArgs)
	realResult := result.(*sociality.SocialityServiceGetSocialInfoResult)
	success, err := handler.(sociality.SocialityService).GetSocialInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialityServiceGetSocialInfoArgs() interface{} {
	return sociality.NewSocialityServiceGetSocialInfoArgs()
}

func newSocialityServiceGetSocialInfoResult() interface{} {
	return sociality.NewSocialityServiceGetSocialInfoResult()
}

func batchGetSocialInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sociality.SocialityServiceBatchGetSocialInfoArgs)
	realResult := result.(*sociality.SocialityServiceBatchGetSocialInfoResult)
	success, err := handler.(sociality.SocialityService).BatchGetSocialInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialityServiceBatchGetSocialInfoArgs() interface{} {
	return sociality.NewSocialityServiceBatchGetSocialInfoArgs()
}

func newSocialityServiceBatchGetSocialInfoResult() interface{} {
	return sociality.NewSocialityServiceBatchGetSocialInfoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Action(ctx context.Context, req *sociality.QingyuRelationActionRequest) (r *sociality.QingyuRelationActionResponse, err error) {
	var _args sociality.SocialityServiceActionArgs
	_args.Req = req
	var _result sociality.SocialityServiceActionResult
	if err = p.c.Call(ctx, "Action", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetRelationIdList(ctx context.Context, req *sociality.QingyuGetRelationIdListRequest) (r *sociality.QingyuGetRelationIdListResponse, err error) {
	var _args sociality.SocialityServiceGetRelationIdListArgs
	_args.Req = req
	var _result sociality.SocialityServiceGetRelationIdListResult
	if err = p.c.Call(ctx, "GetRelationIdList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetSocialInfo(ctx context.Context, req *sociality.QingyuGetSocialInfoRequest) (r *sociality.QingyuGetSocialInfoResponse, err error) {
	var _args sociality.SocialityServiceGetSocialInfoArgs
	_args.Req = req
	var _result sociality.SocialityServiceGetSocialInfoResult
	if err = p.c.Call(ctx, "GetSocialInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BatchGetSocialInfo(ctx context.Context, req *sociality.QingyuBatchGetSocialInfoRequest) (r *sociality.QingyuBatchGetSocialInfoResponse, err error) {
	var _args sociality.SocialityServiceBatchGetSocialInfoArgs
	_args.Req = req
	var _result sociality.SocialityServiceBatchGetSocialInfoResult
	if err = p.c.Call(ctx, "BatchGetSocialInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

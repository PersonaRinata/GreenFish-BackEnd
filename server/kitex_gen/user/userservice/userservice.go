// Code generated by Kitex v0.8.0. DO NOT EDIT.

package userservice

import (
	user "GreenFish/server/kitex_gen/user"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Register":           kitex.NewMethodInfo(registerHandler, newUserServiceRegisterArgs, newUserServiceRegisterResult, false),
		"Login":              kitex.NewMethodInfo(loginHandler, newUserServiceLoginArgs, newUserServiceLoginResult, false),
		"GetUserInfo":        kitex.NewMethodInfo(getUserInfoHandler, newUserServiceGetUserInfoArgs, newUserServiceGetUserInfoResult, false),
		"BatchGetUserInfo":   kitex.NewMethodInfo(batchGetUserInfoHandler, newUserServiceBatchGetUserInfoArgs, newUserServiceBatchGetUserInfoResult, false),
		"GetFollowList":      kitex.NewMethodInfo(getFollowListHandler, newUserServiceGetFollowListArgs, newUserServiceGetFollowListResult, false),
		"GetFollowerList":    kitex.NewMethodInfo(getFollowerListHandler, newUserServiceGetFollowerListArgs, newUserServiceGetFollowerListResult, false),
		"GetFriendList":      kitex.NewMethodInfo(getFriendListHandler, newUserServiceGetFriendListArgs, newUserServiceGetFriendListResult, false),
		"UpdateIssueList":    kitex.NewMethodInfo(updateIssueListHandler, newUserServiceUpdateIssueListArgs, newUserServiceUpdateIssueListResult, false),
		"GetIssueList":       kitex.NewMethodInfo(getIssueListHandler, newUserServiceGetIssueListArgs, newUserServiceGetIssueListResult, false),
		"SearchUserList":     kitex.NewMethodInfo(searchUserListHandler, newUserServiceSearchUserListArgs, newUserServiceSearchUserListResult, false),
		"ChangeUserAvatar":   kitex.NewMethodInfo(changeUserAvatarHandler, newUserServiceChangeUserAvatarArgs, newUserServiceChangeUserAvatarResult, false),
		"ChangeUserNickname": kitex.NewMethodInfo(changeUserNicknameHandler, newUserServiceChangeUserNicknameArgs, newUserServiceChangeUserNicknameResult, false),
		"JudgeDoctor":        kitex.NewMethodInfo(judgeDoctorHandler, newUserServiceJudgeDoctorArgs, newUserServiceJudgeDoctorResult, false),
		"AddDoctor":          kitex.NewMethodInfo(addDoctorHandler, newUserServiceAddDoctorArgs, newUserServiceAddDoctorResult, false),
		"FindDoctor":         kitex.NewMethodInfo(findDoctorHandler, newUserServiceFindDoctorArgs, newUserServiceFindDoctorResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "user",
		"ServiceFilePath": `../../idl/user.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceRegisterArgs)
	realResult := result.(*user.UserServiceRegisterResult)
	success, err := handler.(user.UserService).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceRegisterArgs() interface{} {
	return user.NewUserServiceRegisterArgs()
}

func newUserServiceRegisterResult() interface{} {
	return user.NewUserServiceRegisterResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceLoginArgs)
	realResult := result.(*user.UserServiceLoginResult)
	success, err := handler.(user.UserService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceLoginArgs() interface{} {
	return user.NewUserServiceLoginArgs()
}

func newUserServiceLoginResult() interface{} {
	return user.NewUserServiceLoginResult()
}

func getUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetUserInfoArgs)
	realResult := result.(*user.UserServiceGetUserInfoResult)
	success, err := handler.(user.UserService).GetUserInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetUserInfoArgs() interface{} {
	return user.NewUserServiceGetUserInfoArgs()
}

func newUserServiceGetUserInfoResult() interface{} {
	return user.NewUserServiceGetUserInfoResult()
}

func batchGetUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceBatchGetUserInfoArgs)
	realResult := result.(*user.UserServiceBatchGetUserInfoResult)
	success, err := handler.(user.UserService).BatchGetUserInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceBatchGetUserInfoArgs() interface{} {
	return user.NewUserServiceBatchGetUserInfoArgs()
}

func newUserServiceBatchGetUserInfoResult() interface{} {
	return user.NewUserServiceBatchGetUserInfoResult()
}

func getFollowListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetFollowListArgs)
	realResult := result.(*user.UserServiceGetFollowListResult)
	success, err := handler.(user.UserService).GetFollowList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetFollowListArgs() interface{} {
	return user.NewUserServiceGetFollowListArgs()
}

func newUserServiceGetFollowListResult() interface{} {
	return user.NewUserServiceGetFollowListResult()
}

func getFollowerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetFollowerListArgs)
	realResult := result.(*user.UserServiceGetFollowerListResult)
	success, err := handler.(user.UserService).GetFollowerList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetFollowerListArgs() interface{} {
	return user.NewUserServiceGetFollowerListArgs()
}

func newUserServiceGetFollowerListResult() interface{} {
	return user.NewUserServiceGetFollowerListResult()
}

func getFriendListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetFriendListArgs)
	realResult := result.(*user.UserServiceGetFriendListResult)
	success, err := handler.(user.UserService).GetFriendList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetFriendListArgs() interface{} {
	return user.NewUserServiceGetFriendListArgs()
}

func newUserServiceGetFriendListResult() interface{} {
	return user.NewUserServiceGetFriendListResult()
}

func updateIssueListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdateIssueListArgs)
	realResult := result.(*user.UserServiceUpdateIssueListResult)
	success, err := handler.(user.UserService).UpdateIssueList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUpdateIssueListArgs() interface{} {
	return user.NewUserServiceUpdateIssueListArgs()
}

func newUserServiceUpdateIssueListResult() interface{} {
	return user.NewUserServiceUpdateIssueListResult()
}

func getIssueListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetIssueListArgs)
	realResult := result.(*user.UserServiceGetIssueListResult)
	success, err := handler.(user.UserService).GetIssueList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetIssueListArgs() interface{} {
	return user.NewUserServiceGetIssueListArgs()
}

func newUserServiceGetIssueListResult() interface{} {
	return user.NewUserServiceGetIssueListResult()
}

func searchUserListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceSearchUserListArgs)
	realResult := result.(*user.UserServiceSearchUserListResult)
	success, err := handler.(user.UserService).SearchUserList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceSearchUserListArgs() interface{} {
	return user.NewUserServiceSearchUserListArgs()
}

func newUserServiceSearchUserListResult() interface{} {
	return user.NewUserServiceSearchUserListResult()
}

func changeUserAvatarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceChangeUserAvatarArgs)
	realResult := result.(*user.UserServiceChangeUserAvatarResult)
	success, err := handler.(user.UserService).ChangeUserAvatar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceChangeUserAvatarArgs() interface{} {
	return user.NewUserServiceChangeUserAvatarArgs()
}

func newUserServiceChangeUserAvatarResult() interface{} {
	return user.NewUserServiceChangeUserAvatarResult()
}

func changeUserNicknameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceChangeUserNicknameArgs)
	realResult := result.(*user.UserServiceChangeUserNicknameResult)
	success, err := handler.(user.UserService).ChangeUserNickname(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceChangeUserNicknameArgs() interface{} {
	return user.NewUserServiceChangeUserNicknameArgs()
}

func newUserServiceChangeUserNicknameResult() interface{} {
	return user.NewUserServiceChangeUserNicknameResult()
}

func judgeDoctorHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceJudgeDoctorArgs)
	realResult := result.(*user.UserServiceJudgeDoctorResult)
	success, err := handler.(user.UserService).JudgeDoctor(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceJudgeDoctorArgs() interface{} {
	return user.NewUserServiceJudgeDoctorArgs()
}

func newUserServiceJudgeDoctorResult() interface{} {
	return user.NewUserServiceJudgeDoctorResult()
}

func addDoctorHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceAddDoctorArgs)
	realResult := result.(*user.UserServiceAddDoctorResult)
	success, err := handler.(user.UserService).AddDoctor(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceAddDoctorArgs() interface{} {
	return user.NewUserServiceAddDoctorArgs()
}

func newUserServiceAddDoctorResult() interface{} {
	return user.NewUserServiceAddDoctorResult()
}

func findDoctorHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceFindDoctorArgs)
	realResult := result.(*user.UserServiceFindDoctorResult)
	success, err := handler.(user.UserService).FindDoctor(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceFindDoctorArgs() interface{} {
	return user.NewUserServiceFindDoctorArgs()
}

func newUserServiceFindDoctorResult() interface{} {
	return user.NewUserServiceFindDoctorResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, req *user.QingyuUserRegisterRequest) (r *user.QingyuUserRegisterResponse, err error) {
	var _args user.UserServiceRegisterArgs
	_args.Req = req
	var _result user.UserServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, req *user.QingyuUserLoginRequest) (r *user.QingyuUserLoginResponse, err error) {
	var _args user.UserServiceLoginArgs
	_args.Req = req
	var _result user.UserServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfo(ctx context.Context, req *user.QingyuGetUserRequest) (r *user.QingyuGetUserResponse, err error) {
	var _args user.UserServiceGetUserInfoArgs
	_args.Req = req
	var _result user.UserServiceGetUserInfoResult
	if err = p.c.Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BatchGetUserInfo(ctx context.Context, req *user.QingyuBatchGetUserRequest) (r *user.QingyuBatchGetUserResonse, err error) {
	var _args user.UserServiceBatchGetUserInfoArgs
	_args.Req = req
	var _result user.UserServiceBatchGetUserInfoResult
	if err = p.c.Call(ctx, "BatchGetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollowList(ctx context.Context, req *user.QingyuGetRelationFollowListRequest) (r *user.QingyuGetRelationFollowListResponse, err error) {
	var _args user.UserServiceGetFollowListArgs
	_args.Req = req
	var _result user.UserServiceGetFollowListResult
	if err = p.c.Call(ctx, "GetFollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollowerList(ctx context.Context, req *user.QingyuGetRelationFollowerListRequest) (r *user.QingyuGetRelationFollowerListResponse, err error) {
	var _args user.UserServiceGetFollowerListArgs
	_args.Req = req
	var _result user.UserServiceGetFollowerListResult
	if err = p.c.Call(ctx, "GetFollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFriendList(ctx context.Context, req *user.QingyuGetRelationFriendListRequest) (r *user.QingyuGetRelationFriendListResponse, err error) {
	var _args user.UserServiceGetFriendListArgs
	_args.Req = req
	var _result user.UserServiceGetFriendListResult
	if err = p.c.Call(ctx, "GetFriendList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateIssueList(ctx context.Context, req *user.QingyuUpdateIssueListRequest) (r *user.QingyuUpdateIssueListResponse, err error) {
	var _args user.UserServiceUpdateIssueListArgs
	_args.Req = req
	var _result user.UserServiceUpdateIssueListResult
	if err = p.c.Call(ctx, "UpdateIssueList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetIssueList(ctx context.Context, req *user.QingyuGetIssueListRequest) (r *user.QingyuGetIssueListResponse, err error) {
	var _args user.UserServiceGetIssueListArgs
	_args.Req = req
	var _result user.UserServiceGetIssueListResult
	if err = p.c.Call(ctx, "GetIssueList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SearchUserList(ctx context.Context, req *user.QingyuSearchUserRequest) (r *user.QingyuSearchUserResponse, err error) {
	var _args user.UserServiceSearchUserListArgs
	_args.Req = req
	var _result user.UserServiceSearchUserListResult
	if err = p.c.Call(ctx, "SearchUserList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChangeUserAvatar(ctx context.Context, req *user.QingyuAvatarChangeRequest) (r *user.QingyuAvatarChangeResponse, err error) {
	var _args user.UserServiceChangeUserAvatarArgs
	_args.Req = req
	var _result user.UserServiceChangeUserAvatarResult
	if err = p.c.Call(ctx, "ChangeUserAvatar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChangeUserNickname(ctx context.Context, req *user.QingyuNicknameChangeRequest) (r *user.QingyuNicknameChangeResponse, err error) {
	var _args user.UserServiceChangeUserNicknameArgs
	_args.Req = req
	var _result user.UserServiceChangeUserNicknameResult
	if err = p.c.Call(ctx, "ChangeUserNickname", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) JudgeDoctor(ctx context.Context, req *user.QingyuJudgeDoctorRequest) (r *user.QingyuJudgeDoctorResponse, err error) {
	var _args user.UserServiceJudgeDoctorArgs
	_args.Req = req
	var _result user.UserServiceJudgeDoctorResult
	if err = p.c.Call(ctx, "JudgeDoctor", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddDoctor(ctx context.Context, req *user.QingyuAddDoctorRequest) (r *user.QingyuAddDoctorResponse, err error) {
	var _args user.UserServiceAddDoctorArgs
	_args.Req = req
	var _result user.UserServiceAddDoctorResult
	if err = p.c.Call(ctx, "AddDoctor", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FindDoctor(ctx context.Context, req *user.QingyuFindDoctorRequest) (r *user.QingyuFindDoctorResponse, err error) {
	var _args user.UserServiceFindDoctorArgs
	_args.Req = req
	var _result user.UserServiceFindDoctorResult
	if err = p.c.Call(ctx, "FindDoctor", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

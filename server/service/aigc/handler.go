package main

import (
	aigc "GreenFish/server/kitex_gen/aigc"
	"GreenFish/server/kitex_gen/base"
	"GreenFish/server/service/aigc/pkg"
	"context"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
)

type UserManager interface {
	GetIssueList(ctx context.Context, userID int64) (*base.IssueList, error)
}

// AIGCServerImpl implements the last service interface defined in the IDL.
type AIGCServerImpl struct {
	UserManager
}

// UserAskQuestion implements the AIGCServerImpl interface.
func (s *AIGCServerImpl) UserAskQuestion(ctx context.Context, req *aigc.QingyuAigcQuestionRequest) (resp *aigc.QingyuAigcQuestionResponse, err error) {
	resp = new(aigc.QingyuAigcQuestionResponse)

	issueList, err := s.UserManager.GetIssueList(ctx, req.UserId)
	if err != nil {
		klog.Error("aigc get user IssueList failed: %s", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  err.Error(),
		}
		return resp, err
	}
	strIssueList, err := sonic.Marshal(issueList)
	if err != nil {
		klog.Error("aigc marshal IssueList failed: %s", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  err.Error(),
		}
		return resp, err
	}
	res := pkg.GetGptMessage(string(strIssueList), req.Content)
	resp.Msg = res
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "aigc ask question success",
	}
	return
}

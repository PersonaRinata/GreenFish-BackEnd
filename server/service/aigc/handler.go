package main

import (
	aigc "GreenFish/server/kitex_gen/aigc"
	"GreenFish/server/kitex_gen/base"
	"GreenFish/server/service/aigc/pkg"
	"context"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
)

type RedisManager interface {
	GetAIGCHistory(ctx context.Context, userID int64) ([]string, error)
	AddAIGCRecord(ctx context.Context, userID int64, msg ...string) error
}

type UserManager interface {
	GetIssueList(ctx context.Context, userID int64) (*base.IssueList, error)
}

// AIGCServerImpl implements the last service interface defined in the IDL.
type AIGCServerImpl struct {
	UserManager
	RedisManager
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
	history, err := s.RedisManager.GetAIGCHistory(ctx, req.UserId)
	if err != nil {
		klog.Error("aigc redisManager get aigc history failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  err.Error(),
		}
		return nil, err
	}
	var question []string
	question = append(question, string(strIssueList))
	history = append(history, req.Content)
	question = append(question, history...)
	res := pkg.GetGptMessage(question...)

	err = s.RedisManager.AddAIGCRecord(ctx, req.UserId, req.Content, res)
	if err != nil {
		klog.Error("aigc redisManager add aigc record failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  err.Error(),
		}
		return nil, err
	}

	resp.Msg = res
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "aigc ask question success",
	}
	return
}

// ChooseWord implements the AIGCServerImpl interface.
func (s *AIGCServerImpl) ChooseWord(ctx context.Context, req *aigc.QingyuAigcChooseWordRequest) (resp *aigc.QingyuAigcChooseWordResponse, err error) {
	resp = new(aigc.QingyuAigcChooseWordResponse)

	res := pkg.GetGptMessage(req.Content)
	resp.Msg = res
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "aigc choose word success",
	}
	return
}

// DoctorAnalyse implements the AIGCServerImpl interface.
func (s *AIGCServerImpl) DoctorAnalyse(ctx context.Context, req *aigc.QingyuAigcDoctorAnalyseRequest) (resp *aigc.QingyuAigcDoctorAnalyseResponse, err error) {
	resp = new(aigc.QingyuAigcDoctorAnalyseResponse)

	res := pkg.GetGptMessage(req.Content...)
	resp.Msg = res
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "aigc doctor analyse success",
	}
	return
}

// AnalyseIssueList implements the AIGCServerImpl interface.
func (s *AIGCServerImpl) AnalyseIssueList(ctx context.Context, req *aigc.QingyuAigcIssueListRequest) (resp *aigc.QingyuAigcIssueListResponse, err error) {
	resp = new(aigc.QingyuAigcIssueListResponse)

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
	res := pkg.AnalyseIssueList(string(strIssueList))
	resp.Msg = res
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "aigc analyse issueList success",
	}
	return
}

// GetAIGCHistory implements the AIGCServerImpl interface.
func (s *AIGCServerImpl) GetAIGCHistory(ctx context.Context, req *aigc.QingyuAigcGetHistoryRequest) (resp *aigc.QingyuAigcGetHistoryResponse, err error) {
	resp = new(aigc.QingyuAigcGetHistoryResponse)

	history, err := s.RedisManager.GetAIGCHistory(ctx, req.UserId)
	if err != nil {
		klog.Error("aigc redisManager get aigc history failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  err.Error(),
		}
		return nil, err
	}
	resp.Msg = history
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "aigc get history success",
	}
	return
}

// RecommendDoctor implements the AIGCServerImpl interface.
func (s *AIGCServerImpl) RecommendDoctor(ctx context.Context, req *aigc.QingyuAigcRecommendDocotorRequest) (resp *aigc.QingyuAigcRecommendDocotorResponse, err error) {
	resp = new(aigc.QingyuAigcRecommendDocotorResponse)

	err, department := pkg.DoctorRecommend(req.Content)
	if err != nil {
		klog.Error("aigc recommend doctor failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  err.Error(),
		}
		return nil, err
	}

	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "aigc get history success",
	}

	resp.Department = department
	return
}

package main

import (
	"GreenFish/server/common/consts"
	"GreenFish/server/kitex_gen/base"
	chat "GreenFish/server/kitex_gen/chat"
	"GreenFish/server/service/chat/dao"
	"GreenFish/server/service/chat/model"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
)

// ChatServiceImpl implements the last service interface defined in the IDL.
type ChatServiceImpl struct {
	MysqlManager
	Publisher
	Subscriber
}
type Publisher interface {
	Publish(context.Context, *chat.QingyuMessageActionRequest) error
}
type Subscriber interface {
	Subscribe(ctx context.Context, dao *dao.MysqlManager) (err error)
}
type MysqlManager interface {
	GetHistoryMessage(ctx context.Context, userId, toUserId, time int64) ([]*model.Message, error)
	GetLatestMessage(ctx context.Context, userId, toUserId int64) (*model.Message, error)
	BatchGetLatestMessage(ctx context.Context, userId int64, toUserIdList []int64) ([]*model.Message, error)
	HandleMessage(ctx context.Context, msg string, userId, toUserId, time int64) error
}

// GetChatHistory implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) GetChatHistory(ctx context.Context, req *chat.QingyuMessageGetChatHistoryRequest) (resp *chat.QingyuMessageGetChatHistoryResponse, err error) {
	resp = new(chat.QingyuMessageGetChatHistoryResponse)

	msg, err := s.MysqlManager.GetHistoryMessage(ctx, req.UserId, req.ToUserId, req.PreMsgTime)
	if err != nil {
		klog.Errorf("chat mysql get history message failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "chat mysql get history message failed",
		}
		return resp, err
	}
	for _, v := range msg {
		resp.MessageList = append(resp.MessageList, &base.Message{
			Id:         v.ID,
			ToUserId:   v.ToUserId,
			FromUserId: v.FromUserId,
			Content:    v.Content,
			CreateTime: v.CreateTime,
		})
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "chat get history message success",
	}
	return resp, nil
}

// SentMessage implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) SentMessage(ctx context.Context, req *chat.QingyuMessageActionRequest) (resp *chat.QingyuMessageActionResponse, err error) {
	resp = new(chat.QingyuMessageActionResponse)

	err = s.Publish(ctx, req)
	if err != nil {
		klog.Errorf("chat sentMessage failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "chat publisher publish failed",
		}
		return resp, err
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "chat publisher publish success",
	}
	return resp, nil
}

// GetLatestMessage implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) GetLatestMessage(ctx context.Context, req *chat.QingyuMessageGetLatestRequest) (resp *chat.QingyuMessageGetLatestResponse, err error) {
	resp = new(chat.QingyuMessageGetLatestResponse)

	msg, err := s.MysqlManager.GetLatestMessage(ctx, req.UserId, req.ToUserId)
	if err != nil {
		klog.Errorf("chat mysql get latest message failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "chat mysql get latest message failed",
		}
		return resp, err
	}
	msgType := consts.ReceiveMessage
	if msg.FromUserId == msg.ToUserId {
		msgType = consts.SentMessage
	}
	resp.LatestMsg = &base.LatestMsg{
		Message: msg.Content,
		Time:    msg.CreateTime,
		MsgType: int64(msgType),
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "chat get latest message success",
	}
	return resp, nil
}

// BatchGetLatestMessage implements the ChatServiceImpl interface.
func (s *ChatServiceImpl) BatchGetLatestMessage(ctx context.Context, req *chat.QingyuMessageBatchGetLatestRequest) (resp *chat.QingyuMessageBatchGetLatestResponse, err error) {
	resp = new(chat.QingyuMessageBatchGetLatestResponse)

	msgList, err := s.MysqlManager.BatchGetLatestMessage(ctx, req.UserId, req.ToUserIdList)
	if err != nil {
		klog.Errorf("chat mysql batch get latest message failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "chat mysql batch get latest message failed",
		}
		return resp, err
	}
	for _, v := range msgList {
		msgType := consts.ReceiveMessage
		if v.FromUserId == v.ToUserId {
			msgType = consts.SentMessage
		}
		resp.LatestMsgList = append(resp.LatestMsgList, &base.LatestMsg{
			Message: v.Content,
			Time:    v.CreateTime,
			MsgType: int64(msgType),
		})
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "chat batch get latest message success",
	}
	return resp, nil
}

package pkg

import (
	"GreenFish/server/kitex_gen/base"
	"GreenFish/server/kitex_gen/chat"
	"GreenFish/server/kitex_gen/chat/chatservice"
	"context"
)

type ChatManager struct {
	client chatservice.Client
}

func NewChatManager(client chatservice.Client) *ChatManager {
	return &ChatManager{client: client}
}

func (m *ChatManager) BatchGetLatestMessage(ctx context.Context, userId int64, toUserIdList []int64) ([]*base.LatestMsg, error) {
	resp, err := m.client.BatchGetLatestMessage(ctx, &chat.QingyuMessageBatchGetLatestRequest{
		UserId:       userId,
		ToUserIdList: toUserIdList,
	})
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, err
	}
	return resp.LatestMsgList, nil
}

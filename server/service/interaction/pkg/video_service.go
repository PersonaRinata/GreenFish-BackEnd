package pkg

import (
	"GreenFish/server/kitex_gen/video"
	"GreenFish/server/kitex_gen/video/videoservice"
	"context"
)

type VideoManager struct {
	client videoservice.Client
}

func NewVideoManager(client videoservice.Client) *VideoManager {
	return &VideoManager{client: client}
}

func (m *VideoManager) GetPublishedVideoIdList(ctx context.Context, userId int64) ([]int64, error) {
	resp, err := m.client.GetPublishedVideoIdList(ctx, &video.QingyuGetPublishedVideoIdListRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, err
	}
	return resp.VideoIdList, nil
}

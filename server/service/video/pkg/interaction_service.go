package pkg

import (
	"GreenFish/server/kitex_gen/base"
	"GreenFish/server/kitex_gen/interaction"
	"GreenFish/server/kitex_gen/interaction/interactionserver"
	"context"
)

type InteractionManager struct {
	InteractionService interactionserver.Client
}

func NewInteractionManager(client interactionserver.Client) *InteractionManager {
	return &InteractionManager{InteractionService: client}
}

// GetVideoInteractInfo get video interactInfo.
func (i *InteractionManager) GetVideoInteractInfo(ctx context.Context, videoId, viewerId int64) (*base.VideoInteractInfo, error) {
	resp, err := i.InteractionService.GetVideoInteractInfo(ctx, &interaction.QingyuGetVideoInteractInfoRequest{
		VideoId:  videoId,
		ViewerId: viewerId,
	})
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, err
	}
	return resp.InteractInfo, nil
}

// GetFavoriteVideoIdList gets the favorite video id list.
func (i *InteractionManager) GetFavoriteVideoIdList(ctx context.Context, userId int64) ([]int64, error) {
	resp, err := i.InteractionService.GetFavoriteVideoIdList(ctx, &interaction.QingyuGetFavoriteVideoIdListRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, err
	}
	return resp.VideoIdList, nil
}

// BatchGetVideoInteractInfo batch get video interactInfo.
func (i *InteractionManager) BatchGetVideoInteractInfo(ctx context.Context, videoIdList []int64, viewerId int64) ([]*base.VideoInteractInfo, error) {
	resp, err := i.InteractionService.BatchGetVideoInteractInfo(ctx, &interaction.QingyuBatchGetVideoInteractInfoRequest{
		VideoIdList: videoIdList,
		ViewerId:    viewerId,
	})
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, err
	}
	return resp.InteractInfoList, nil
}

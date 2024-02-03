package main

import (
	"GreenFish/server/common/consts"
	"GreenFish/server/kitex_gen/base"
	video "GreenFish/server/kitex_gen/video"
	"GreenFish/server/service/video/model"
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct {
	RedisManager
	MysqlManager
	UserManager
	InteractionManager
	Publisher
}
type Publisher interface {
	Publish(ctx context.Context, video *model.Video) error
}
type UserManager interface {
	BatchGetUser(ctx context.Context, list []int64, uid int64) ([]*base.User, error)
	GetUser(ctx context.Context, UserId, toUserId int64) (*base.User, error)
}

type InteractionManager interface {
	GetFavoriteVideoIdList(ctx context.Context, userId int64) ([]int64, error)
	BatchGetVideoInteractInfo(ctx context.Context, videoIdList []int64, viewerId int64) ([]*base.VideoInteractInfo, error)
}
type MysqlManager interface {
	GetBasicVideoListByLatestTime(ctx context.Context, userId, latestTime int64) ([]*model.Video, error)
	GetPublishedVideoListByUserId(ctx context.Context, userId int64) ([]*model.Video, error)
	GetFavoriteVideoListByUserId(ctx context.Context, userId int64) ([]*model.Video, error)
	GetPublishedVideoIdListByUserId(ctx context.Context, userId int64) ([]int64, error)
	PublishVideo(ctx context.Context, video *model.Video) error
	SearchVideoByTitle(ctx context.Context, content string) ([]*model.Video, error)
}
type RedisManager interface {
	GetBasicVideoListByLatestTime(ctx context.Context, userId, latestTime int64) ([]*model.Video, error)
	GetPublishedVideoListByUserId(ctx context.Context, userId int64) ([]*model.Video, error)
	GetFavoriteVideoListByUserId(ctx context.Context, userId int64) ([]*model.Video, error)
	GetPublishedVideoIdListByUserId(ctx context.Context, userId int64) ([]int64, error)
	PublishVideo(ctx context.Context, video *model.Video) error
}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.QingyuFeedRequest) (resp *video.QingyuFeedResponse, err error) {
	resp = new(video.QingyuFeedResponse)

	mv, err := s.RedisManager.GetBasicVideoListByLatestTime(ctx, req.ViewerId, req.LatestTime)
	if err != nil {
		klog.Error("video redis get basic VideoListByLatestTime failed,", err)
		mv, err = s.MysqlManager.GetBasicVideoListByLatestTime(ctx, req.ViewerId, req.LatestTime)
		if err != nil {
			klog.Error("video mysql get basic VideoListByLatestTime failed,", err)
			resp.BaseResp = &base.QingyuBaseResponse{
				StatusCode: 500,
				StatusMsg:  "video get basic VideoListByLatestTime failed",
			}
			return resp, err
		}
	}
	var videoIdList []int64
	var userIdList []int64
	for _, v := range mv {
		videoIdList = append(videoIdList, v.ID)
		userIdList = append(userIdList, v.AuthorId)
	}
	iv, err := s.InteractionManager.BatchGetVideoInteractInfo(ctx, videoIdList, req.ViewerId)
	uv, err := s.UserManager.BatchGetUser(ctx, userIdList, req.ViewerId)
	if err != nil {
		klog.Error("video InteractionManager BatchGetVideoInteractInfo failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "video InteractionManager BatchGetVideoInteractInfo failed",
		}
		return resp, err
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "get feed success",
	}
	for k, v := range videoIdList {
		resp.VideoList = append(resp.VideoList, &base.Video{
			Id:            v,
			Author:        uv[k],
			PlayUrl:       mv[k].PlayUrl,
			CoverUrl:      mv[k].CoverUrl,
			FavoriteCount: iv[k].FavoriteCount,
			CommentCount:  iv[k].CommentCount,
			IsFavorite:    iv[k].IsFavorite,
			Title:         mv[k].Title,
		})
	}

	return resp, nil
}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, req *video.QingyuPublishActionRequest) (resp *video.QingyuPublishActionResponse, err error) {
	resp = new(video.QingyuPublishActionResponse)
	sf, err := snowflake.NewNode(consts.VideoSnowFlakeNode)
	if err != nil {
		klog.Error("generate snowflake failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "generate snowflake failed",
		}
		return resp, err
	}
	videoRecord := &model.Video{
		ID:         sf.Generate().Int64(),
		AuthorId:   req.UserId,
		PlayUrl:    req.PlayUrl,
		CoverUrl:   req.CoverUrl,
		Title:      req.Title,
		CreateTime: time.Now().UnixNano(),
	}
	err = s.Publisher.Publish(ctx, videoRecord)
	if err != nil {
		klog.Error("video publish video failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "video publish video failed",
		}
		return resp, err
	}
	err = s.RedisManager.PublishVideo(ctx, videoRecord)
	if err != nil {
		klog.Error("video redis publish video failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "video redis publish video failed",
		}
		return resp, err
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "publish video success",
	}
	return resp, nil
}

// GetPublishedVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishedVideoList(ctx context.Context, req *video.QingyuGetPublishedListRequest) (resp *video.QingyuGetPublishedListResponse, err error) {
	resp = new(video.QingyuGetPublishedListResponse)

	mv, err := s.RedisManager.GetPublishedVideoListByUserId(ctx, req.OwnerId)
	if err != nil {
		klog.Error("video redis get publishedVideoList failed,", err)
		mv, err = s.MysqlManager.GetPublishedVideoListByUserId(ctx, req.OwnerId)
		if err != nil {
			klog.Error("video mysql get publishedVideoList failed,", err)
			resp.BaseResp = &base.QingyuBaseResponse{
				StatusCode: 500,
				StatusMsg:  "video get publishedVideoList failed",
			}
			return resp, err
		}
	}
	var videoIdList []int64
	var userIdList []int64
	for _, v := range mv {
		videoIdList = append(videoIdList, v.ID)
		userIdList = append(userIdList, v.AuthorId)
	}
	iv, err := s.InteractionManager.BatchGetVideoInteractInfo(ctx, videoIdList, req.ViewerId)
	uv, err := s.UserManager.BatchGetUser(ctx, userIdList, req.ViewerId)
	if err != nil {
		klog.Error("video InteractionManager BatchGetVideoInteractInfo failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "video InteractionManager BatchGetVideoInteractInfo failed",
		}
		return resp, err
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "get publishedVideoList success",
	}
	for k, v := range videoIdList {
		resp.VideoList = append(resp.VideoList, &base.Video{
			Id:            v,
			Author:        uv[k],
			PlayUrl:       mv[k].PlayUrl,
			CoverUrl:      mv[k].CoverUrl,
			FavoriteCount: iv[k].FavoriteCount,
			CommentCount:  iv[k].CommentCount,
			IsFavorite:    iv[k].IsFavorite,
			Title:         mv[k].Title,
		})
	}
	return resp, nil
}

// GetFavoriteVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFavoriteVideoList(ctx context.Context, req *video.QingyuGetFavoriteListRequest) (resp *video.QingyuGetFavoriteListResponse, err error) {
	resp = new(video.QingyuGetFavoriteListResponse)

	mv, err := s.RedisManager.GetFavoriteVideoListByUserId(ctx, req.OwnerId)
	if err != nil {
		klog.Error("video redis get favoriteVideoList failed,", err)
		mv, err = s.MysqlManager.GetFavoriteVideoListByUserId(ctx, req.OwnerId)
		if err != nil {
			klog.Error("video mysql get favoriteVideoList failed,", err)
			resp.BaseResp = &base.QingyuBaseResponse{
				StatusCode: 500,
				StatusMsg:  "video get favoriteVideoList failed",
			}
			return resp, err
		}
	}
	var videoIdList []int64
	var userIdList []int64
	for _, v := range mv {
		videoIdList = append(videoIdList, v.ID)
		userIdList = append(userIdList, v.AuthorId)
	}
	iv, err := s.InteractionManager.BatchGetVideoInteractInfo(ctx, videoIdList, req.ViewerId)
	uv, err := s.UserManager.BatchGetUser(ctx, userIdList, req.ViewerId)
	if err != nil {
		klog.Error("video InteractionManager BatchGetVideoInteractInfo failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "video InteractionManager BatchGetVideoInteractInfo failed",
		}
		return resp, err
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "get favoriteVideoList success",
	}
	for k, v := range videoIdList {
		resp.VideoList = append(resp.VideoList, &base.Video{
			Id:            v,
			Author:        uv[k],
			PlayUrl:       mv[k].PlayUrl,
			CoverUrl:      mv[k].CoverUrl,
			FavoriteCount: iv[k].FavoriteCount,
			CommentCount:  iv[k].CommentCount,
			IsFavorite:    iv[k].IsFavorite,
			Title:         mv[k].Title,
		})
	}
	return resp, nil
}

// GetPublishedVideoIdList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishedVideoIdList(ctx context.Context, req *video.QingyuGetPublishedVideoIdListRequest) (resp *video.QingyuGetPublishedVideoIdListResponse, err error) {
	resp = new(video.QingyuGetPublishedVideoIdListResponse)

	idList, err := s.RedisManager.GetPublishedVideoIdListByUserId(ctx, req.UserId)
	if err != nil {
		klog.Error("video redis get publishedVideoIdList failed,", err)
		idList, err = s.MysqlManager.GetPublishedVideoIdListByUserId(ctx, req.UserId)
		if err != nil {
			klog.Error("video mysql get publishedVideoIdList failed,", err)
			resp.BaseResp = &base.QingyuBaseResponse{
				StatusCode: 500,
				StatusMsg:  "video get publishedVideoIdList failed",
			}
			return resp, err
		}
	}
	resp.VideoIdList = idList
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "video get publishedVideoIdList success",
	}
	return resp, nil
}

// SearchVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) SearchVideoList(ctx context.Context, req *video.QingyuSearchVideoRequest) (resp *video.QingyuSearchVideoResponse, err error) {
	resp = new(video.QingyuSearchVideoResponse)

	videoList, err := s.MysqlManager.SearchVideoByTitle(ctx, req.Content)
	if err != nil {
		klog.Error("user mysqlManager get videoList failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user mysqlManager get videoList failed",
		}
		return resp, err
	}
	var res []*model.Video
	res = videoList[req.Num*req.Offset : (req.Num+1)*req.Offset]
	if int(req.Num*req.Offset+1) <= len(videoList) {
		if int((req.Offset+1)*req.Num) > len(videoList) {
			res = videoList[req.Num*req.Offset:]
		} else {
			res = videoList[req.Num*req.Offset : (req.Offset+1)*req.Num]
		}
	}

	var videoIdList []int64
	var userIdList []int64
	for _, v := range res {
		videoIdList = append(videoIdList, v.ID)
		userIdList = append(userIdList, v.AuthorId)
	}

	iv, err := s.InteractionManager.BatchGetVideoInteractInfo(ctx, videoIdList, req.ViewerId)
	uv, err := s.UserManager.BatchGetUser(ctx, userIdList, req.ViewerId)
	if err != nil {
		klog.Error("video InteractionManager BatchGetVideoInteractInfo failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "video InteractionManager BatchGetVideoInteractInfo failed",
		}
		return resp, err
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "get feed success",
	}

	for k, v := range videoIdList {
		resp.VideoList = append(resp.VideoList, &base.Video{
			Id:            v,
			Author:        uv[k],
			PlayUrl:       res[k].PlayUrl,
			CoverUrl:      res[k].CoverUrl,
			FavoriteCount: iv[k].FavoriteCount,
			CommentCount:  iv[k].CommentCount,
			IsFavorite:    iv[k].IsFavorite,
			Title:         res[k].Title,
		})
	}

	return resp, nil
}

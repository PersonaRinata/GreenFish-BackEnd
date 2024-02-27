package main

import (
	"GreenFish/server/common/consts"
	"GreenFish/server/common/middleware"
	"GreenFish/server/common/tools"
	"GreenFish/server/kitex_gen/base"
	"GreenFish/server/kitex_gen/user"
	"GreenFish/server/service/api/models"
	"GreenFish/server/service/user/config"
	"GreenFish/server/service/user/model"
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type MysqlManager interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	SearchUserByUsername(ctx context.Context, content string) ([]*model.User, error)
	ChangeAvatarByUserID(ctx context.Context, avatar string, id int64) error
	ChangeNicknameByUserID(ctx context.Context, nickname string, id int64) error
	AddDoctor(ctx context.Context, id int64, department string) error
}
type RedisManager interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, id int64) (*model.User, error)
	BatchGetUserById(ctx context.Context, id []int64) ([]*model.User, error)
	UpdateIssueListById(ctx context.Context, id int64, issueList *model.IssueList) error
	GetIssueListById(ctx context.Context, id int64) (*model.IssueList, error)
	ChangeAvatarByUserID(ctx context.Context, avatar string, id int64) error
	ChangeNicknameByUserID(ctx context.Context, nickname string, id int64) error
	AddDoctor(ctx context.Context, id int64, department string) error
	JudgeDoctor(ctx context.Context, id int64) (string, error)
	FindDoctor(ctx context.Context, department string) ([]string, error)
}
type SocialManager interface {
	GetRelationList(ctx context.Context, viewerId, ownerId int64, option int8) ([]int64, error)
	GetSocialInfo(ctx context.Context, viewerId, ownerId int64) (*base.SocialInfo, error)
	BatchGetSocialInfo(ctx context.Context, viewerId int64, ownerIdList []int64) ([]*base.SocialInfo, error)
}
type InteractionManager interface {
	GetInteractInfo(ctx context.Context, userId int64) (*base.UserInteractInfo, error)
	BatchGetInteractInfo(ctx context.Context, userIdList []int64) ([]*base.UserInteractInfo, error)
}
type ChatManager interface {
	BatchGetLatestMessage(ctx context.Context, userId int64, toUserIdList []int64) ([]*base.LatestMsg, error)
}

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	Jwt *middleware.JWT
	MysqlManager
	RedisManager
	SocialManager
	InteractionManager
	ChatManager
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.QingyuUserRegisterRequest) (resp *user.QingyuUserRegisterResponse, err error) {
	resp = new(user.QingyuUserRegisterResponse)

	sf, err := snowflake.NewNode(consts.UserSnowflakeNode)
	if err != nil {
		klog.Errorf("generate user snowflake id failed: %s", err.Error())
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  err.Error(),
		}
		return resp, nil
	}
	usr := &model.User{
		ID:              sf.Generate().Int64(),
		Username:        req.Username,
		Password:        tools.Md5Crypt(req.Password, config.GlobalServerConfig.MysqlInfo.Salt),
		Avatar:          "",
		BackGroundImage: "",
		Signature:       "default signature",
		Nickname:        req.Username,
	}
	err = s.MysqlManager.CreateUser(ctx, usr)
	if err != nil {
		if err.Error() == consts.MysqlAlreadyExists {
			resp.BaseResp = &base.QingyuBaseResponse{
				StatusCode: 500,
				StatusMsg:  "user already exists",
			}
			return resp, err
		} else {
			klog.Errorf("mysql create user failed: %s", err.Error())
			resp.BaseResp = &base.QingyuBaseResponse{
				StatusCode: 500,
				StatusMsg:  fmt.Sprintf("mysql create user failed: %s", err.Error()),
			}
			return resp, err
		}
	}
	err = s.RedisManager.CreateUser(ctx, usr)
	if err != nil {
		klog.Errorf("redis create user failed: %s", err.Error())
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  fmt.Sprintf("mysql create user failed: %s", err.Error()),
		}
		return resp, err
	}
	resp.UserId = usr.ID
	resp.Token, err = s.Jwt.CreateToken(models.CustomClaims{
		ID: usr.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60*24*30,
			Issuer:    "GreenFish",
			NotBefore: time.Now().Unix(),
		},
	})
	if err != nil {
		klog.Error("register create jwt failed", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  fmt.Sprintf("register create jwt failed,%s", err),
		}
		return resp, err
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "user register success",
	}
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.QingyuUserLoginRequest) (resp *user.QingyuUserLoginResponse, err error) {
	resp = new(user.QingyuUserLoginResponse)

	usr, err := s.MysqlManager.GetUserByUsername(ctx, req.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.BaseResp = &base.QingyuBaseResponse{
				StatusCode: 500,
				StatusMsg:  "no such user",
			}
			return resp, err
		} else {
			klog.Error("mysql get user by username failed", err)
			resp.BaseResp = &base.QingyuBaseResponse{
				StatusCode: 500,
				StatusMsg:  err.Error(),
			}
			return resp, err
		}
	}

	if usr.Password != tools.Md5Crypt(req.Password, config.GlobalServerConfig.MysqlInfo.Salt) {
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "wrong password",
		}
		return resp, nil
	}

	resp.UserId = usr.ID
	resp.Token, err = s.Jwt.CreateToken(models.CustomClaims{
		ID: usr.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60*24*30,
			Issuer:    "GreenFish",
			NotBefore: time.Now().Unix(),
		},
	})
	if err != nil {
		klog.Error("register create jwt failed", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  fmt.Sprintf("register create jwt failed,%s", err),
		}
		return resp, err
	}

	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "login success",
	}
	return resp, nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.QingyuGetUserRequest) (resp *user.QingyuGetUserResponse, err error) {
	resp = new(user.QingyuGetUserResponse)

	usr, err := s.RedisManager.GetUserById(ctx, req.OwnerId)
	if err != nil {
		klog.Error("redis get user by id failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  fmt.Sprintf("redis get user by id failed,%s", err),
		}
		return nil, err
	}

	socialInfo, err := s.SocialManager.GetSocialInfo(ctx, req.ViewerId, req.OwnerId)
	if err != nil {
		klog.Error("socialManager get socialInfo failed,", err)
	}

	interactionInfo, err := s.InteractionManager.GetInteractInfo(ctx, req.OwnerId)
	if err != nil {
		klog.Error("interactionManager get interactionInfo failed,", err)
	}

	if err != nil {
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "get userInfo failed",
		}
		return resp, err
	}

	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "get user by id success",
	}
	resp.User = &base.User{
		Id:              usr.ID,
		Name:            usr.Username,
		Nickname:        usr.Nickname,
		FollowCount:     socialInfo.FollowCount,
		FollowerCount:   socialInfo.FollowerCount,
		IsFollow:        socialInfo.IsFollow,
		Avatar:          usr.Avatar,
		BackgroundImage: usr.BackGroundImage,
		Signature:       usr.Signature,
		TotalFavorited:  interactionInfo.TotalFavorited,
		WorkCount:       interactionInfo.WorkCount,
		FavoriteCount:   interactionInfo.FavoriteCount,
		Department:      usr.Department,
	}
	return resp, nil
}

// BatchGetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) BatchGetUserInfo(ctx context.Context, req *user.QingyuBatchGetUserRequest) (resp *user.QingyuBatchGetUserResonse, err error) {
	resp = new(user.QingyuBatchGetUserResonse)

	userList, err := s.RedisManager.BatchGetUserById(ctx, req.OwnerIdList)
	if err != nil {
		klog.Error("redis batch get users by id failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  fmt.Sprintf("redis batch get users by id failed,%s", err),
		}
		return nil, err
	}
	socialList, err := s.SocialManager.BatchGetSocialInfo(ctx, req.ViewerId, req.OwnerIdList)
	if err != nil {
		klog.Error("user socialManager get socialList failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user socialManager get socialList failed",
		}
		return resp, err
	}
	interactionList, err := s.InteractionManager.BatchGetInteractInfo(ctx, req.OwnerIdList)
	if err != nil {
		klog.Error("user interactionManager get interactionList failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user interactionManager get interactionList failed",
		}
		return resp, err
	}
	for i := 0; i <= len(userList)-1; i++ {
		resp.UserList = append(resp.UserList, &base.User{
			Id:              userList[i].ID,
			Name:            userList[i].Username,
			Nickname:        userList[i].Nickname,
			FollowCount:     socialList[i].FollowCount,
			FollowerCount:   socialList[i].FollowerCount,
			IsFollow:        socialList[i].IsFollow,
			Avatar:          userList[i].Avatar,
			BackgroundImage: userList[i].BackGroundImage,
			Signature:       userList[i].Signature,
			TotalFavorited:  interactionList[i].TotalFavorited,
			WorkCount:       interactionList[i].WorkCount,
			FavoriteCount:   interactionList[i].FavoriteCount,
			Department:      userList[i].Department,
		})
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "batch get userInfo success",
	}
	return resp, nil
}

// GetFollowList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFollowList(ctx context.Context, req *user.QingyuGetRelationFollowListRequest) (resp *user.QingyuGetRelationFollowListResponse, err error) {
	resp = new(user.QingyuGetRelationFollowListResponse)

	userIdlist, err := s.SocialManager.GetRelationList(ctx, req.ViewerId, req.OwnerId, consts.FollowList)
	if err != nil {
		klog.Error("user socialManager get follow list failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user socialManager get followList failed ",
		}
		return resp, err
	}
	userList, err := s.RedisManager.BatchGetUserById(ctx, userIdlist)
	if err != nil {
		klog.Error("user redis get user failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user redis get user failed ",
		}
		return resp, err
	}
	socialList, err := s.SocialManager.BatchGetSocialInfo(ctx, req.ViewerId, userIdlist)
	if err != nil {
		klog.Error("user socialManager get socialList failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user socialManager get socialList failed",
		}
		return resp, err
	}
	interactionList, err := s.InteractionManager.BatchGetInteractInfo(ctx, userIdlist)
	if err != nil {
		klog.Error("user interactionManager get interactionList failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user interactionManager get interactionList failed",
		}
		return resp, err
	}
	for i := 0; i <= len(userList)-1; i++ {
		resp.UserList = append(resp.UserList, &base.User{
			Id:              userList[i].ID,
			Name:            userList[i].Username,
			FollowCount:     socialList[i].FollowCount,
			FollowerCount:   socialList[i].FollowerCount,
			IsFollow:        socialList[i].IsFollow,
			Avatar:          userList[i].Avatar,
			BackgroundImage: userList[i].BackGroundImage,
			Signature:       userList[i].Signature,
			TotalFavorited:  interactionList[i].TotalFavorited,
			WorkCount:       interactionList[i].WorkCount,
			FavoriteCount:   interactionList[i].FavoriteCount,
			Department:      userList[i].Department,
		})
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "batch get followList success",
	}
	return resp, nil
}

// GetFollowerList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFollowerList(ctx context.Context, req *user.QingyuGetRelationFollowerListRequest) (resp *user.QingyuGetRelationFollowerListResponse, err error) {
	resp = new(user.QingyuGetRelationFollowerListResponse)

	userIdlist, err := s.SocialManager.GetRelationList(ctx, req.ViewerId, req.OwnerId, consts.FollowerList)
	if err != nil {
		klog.Error("user socialManager get follower list failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user socialManager get followerList failed ",
		}
		return resp, err
	}
	userList, err := s.RedisManager.BatchGetUserById(ctx, userIdlist)
	if err != nil {
		klog.Error("user redis get user failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user redis get user failed ",
		}
		return resp, err
	}
	socialList, err := s.SocialManager.BatchGetSocialInfo(ctx, req.ViewerId, userIdlist)
	if err != nil {
		klog.Error("user socialManager get socialList failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user socialManager get socialList failed",
		}
		return resp, err
	}
	interactionList, err := s.InteractionManager.BatchGetInteractInfo(ctx, userIdlist)
	if err != nil {
		klog.Error("user interactionManager get interactionList failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user interactionManager get interactionList failed",
		}
		return resp, err
	}
	for i := 0; i <= len(userList)-1; i++ {
		resp.UserList = append(resp.UserList, &base.User{
			Id:              userList[i].ID,
			Name:            userList[i].Username,
			FollowCount:     socialList[i].FollowCount,
			FollowerCount:   socialList[i].FollowerCount,
			IsFollow:        socialList[i].IsFollow,
			Avatar:          userList[i].Avatar,
			BackgroundImage: userList[i].BackGroundImage,
			Signature:       userList[i].Signature,
			TotalFavorited:  interactionList[i].TotalFavorited,
			WorkCount:       interactionList[i].WorkCount,
			FavoriteCount:   interactionList[i].FavoriteCount,
			Department:      userList[i].Department,
		})
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "batch get followList success",
	}
	return resp, nil

}

// GetFriendList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFriendList(ctx context.Context, req *user.QingyuGetRelationFriendListRequest) (resp *user.QingyuGetRelationFriendListResponse, err error) {
	resp = new(user.QingyuGetRelationFriendListResponse)

	userIdlist, err := s.SocialManager.GetRelationList(ctx, req.ViewerId, req.OwnerId, consts.FriendsList)
	if err != nil {
		klog.Error("user socialManager get follow list failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user socialManager get followList failed ",
		}
		return resp, err
	}
	userList, err := s.RedisManager.BatchGetUserById(ctx, userIdlist)
	if err != nil {
		klog.Error("user redis get user failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user redis get user failed ",
		}
		return resp, err
	}
	socialList, err := s.SocialManager.BatchGetSocialInfo(ctx, req.ViewerId, userIdlist)
	if err != nil {
		klog.Error("user socialManager get socialList failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user socialManager get socialList failed",
		}
		return resp, err
	}
	interactionList, err := s.InteractionManager.BatchGetInteractInfo(ctx, userIdlist)
	if err != nil {
		klog.Error("user interactionManager get interactionList failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user interactionManager get interactionList failed",
		}
		return resp, err
	}
	chatList, err := s.ChatManager.BatchGetLatestMessage(ctx, req.ViewerId, userIdlist)
	if err != nil {
		klog.Error("user chatManager get chatList failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user chatManager get chatList failed",
		}
		return resp, err
	}
	for i := 0; i <= len(userList)-1; i++ {
		resp.UserList = append(resp.UserList, &base.FriendUser{
			Id:              userList[i].ID,
			Name:            userList[i].Username,
			FollowCount:     socialList[i].FollowCount,
			FollowerCount:   socialList[i].FollowerCount,
			IsFollow:        socialList[i].IsFollow,
			Avatar:          userList[i].Avatar,
			BackgroundImage: userList[i].BackGroundImage,
			Signature:       userList[i].Signature,
			TotalFavorited:  interactionList[i].TotalFavorited,
			WorkCount:       interactionList[i].WorkCount,
			FavoriteCount:   interactionList[i].FavoriteCount,
			MsgType:         chatList[i].MsgType,
			Message:         chatList[i].Message,
			Time:            chatList[i].Time,
		})
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "batch get followList success",
	}
	return resp, nil
}

// UpdateIssueList implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateIssueList(ctx context.Context, req *user.QingyuUpdateIssueListRequest) (resp *user.QingyuUpdateIssueListResponse, err error) {
	resp = new(user.QingyuUpdateIssueListResponse)

	issueList := model.IssueList{
		Username: req.IssueList.Username,
		Gender:   req.IssueList.Gender,
		Age:      req.IssueList.Age,
		DiseaseRelation: model.DiseaseRelation{
			DiseaseIntroduction: req.IssueList.DiseaseRelation.DiseaseIntroduction,
			FamilyDiseases:      req.IssueList.DiseaseRelation.FamilyDiseases,
		},
		BodyInfo: model.BodyInfo{
			BloodPressure: req.IssueList.BodyInfo.BloodPressure,
			HeartRate:     req.IssueList.BodyInfo.HeartRate,
			Height:        req.IssueList.BodyInfo.Height,
			Weight:        req.IssueList.BodyInfo.Weight,
			BloodSugar:    req.IssueList.BodyInfo.BloodSugar,
			UpdateTime:    time.Now().Unix(),
		},
	}
	for _, v := range req.IssueList.DiseaseRelation.HistoryDiseases {
		issueList.DiseaseRelation.HistoryDiseases = append(issueList.DiseaseRelation.HistoryDiseases, model.HistoryDisease{
			Symptom:    v.Symptom,
			Medicines:  v.Medicines,
			Department: v.Department,
			UpdateTime: v.UpdateTime,
		})
	}

	err = s.RedisManager.UpdateIssueListById(ctx, req.UserId, &issueList)
	if err != nil {
		klog.Error("user redisManager update issueList failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user redisManager update issueList failed",
		}
		return resp, err
	}

	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "update issueList success",
	}
	return resp, nil
}

// GetIssueList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetIssueList(ctx context.Context, req *user.QingyuGetIssueListRequest) (resp *user.QingyuGetIssueListResponse, err error) {
	resp = new(user.QingyuGetIssueListResponse)

	issueList, err := s.RedisManager.GetIssueListById(ctx, req.UserId)
	if err != nil {
		klog.Error("user redisManager get issueList failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user redisManager get issueList failed",
		}
		return resp, err
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "get issueList success",
	}
	if issueList == nil {
		return
	}
	resp.IssueList = &base.IssueList{
		Username: issueList.Username,
		Gender:   issueList.Gender,
		Age:      issueList.Age,
		DiseaseRelation: &base.DiseaseRelation{
			DiseaseIntroduction: issueList.DiseaseRelation.DiseaseIntroduction,
			FamilyDiseases:      issueList.DiseaseRelation.FamilyDiseases,
		},
		BodyInfo: &base.BodyInfo{
			BloodPressure: issueList.BodyInfo.BloodPressure,
			HeartRate:     issueList.BodyInfo.HeartRate,
			Height:        issueList.BodyInfo.Height,
			Weight:        issueList.BodyInfo.Weight,
			BloodSugar:    issueList.BodyInfo.BloodSugar,
			UpdateTime:    issueList.BodyInfo.UpdateTime,
		},
	}
	for _, v := range issueList.DiseaseRelation.HistoryDiseases {
		resp.IssueList.DiseaseRelation.HistoryDiseases = append(resp.IssueList.DiseaseRelation.HistoryDiseases, &base.HistoryDiseases{
			Symptom:    v.Symptom,
			Medicines:  v.Medicines,
			Department: v.Department,
			UpdateTime: v.UpdateTime,
		})
	}
	return resp, nil
}

// SearchUserList implements the UserServiceImpl interface.
func (s *UserServiceImpl) SearchUserList(ctx context.Context, req *user.QingyuSearchUserRequest) (resp *user.QingyuSearchUserResponse, err error) {
	resp = &user.QingyuSearchUserResponse{
		UserList: []*base.User{}, // 为UserList分配内存空间
	}

	userList, err := s.MysqlManager.SearchUserByUsername(ctx, req.Content)
	if err != nil {
		klog.Error("user mysqlManager get userList failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user mysqlManager get userList failed",
		}
		return resp, err
	}
	var res []*model.User
	if int(req.Num*req.Offset+1) <= len(userList) {
		if int((req.Offset+1)*req.Num) > len(userList) {
			res = userList[req.Num*req.Offset:]
		} else {
			res = userList[req.Num*req.Offset : (req.Offset+1)*req.Num]
		}
		for _, v := range res {
			if v == nil {
				break
			}
			socialInfo, err := s.SocialManager.GetSocialInfo(ctx, req.ViewerId, v.ID)
			if err != nil {
				klog.Error("user  get socialInfo failed,", err)
				resp.BaseResp = &base.QingyuBaseResponse{
					StatusCode: 500,
					StatusMsg:  "user  get socialInfo failed",
				}
				return resp, err
			}
			interactionInfo, err := s.InteractionManager.GetInteractInfo(ctx, v.ID)
			if err != nil {
				klog.Error("interactionManager get interactionInfo failed,", err)
			}
			department, err := s.JudgeDoctor(ctx, &user.QingyuJudgeDoctorRequest{UserId: v.ID})
			if err != nil {
				klog.Error("user judge IsDoctor failed,", err)
				resp.BaseResp = &base.QingyuBaseResponse{
					StatusCode: 500,
					StatusMsg:  "user judge IsDoctor failed",
				}
				return resp, err
			}
			resp.UserList = append(resp.UserList, &base.User{
				Id:              v.ID,
				Name:            v.Username,
				Nickname:        v.Nickname,
				FollowCount:     socialInfo.FollowCount,
				FollowerCount:   socialInfo.FollowerCount,
				IsFollow:        socialInfo.IsFollow,
				Avatar:          v.Avatar,
				BackgroundImage: v.BackGroundImage,
				Signature:       v.Signature,
				TotalFavorited:  interactionInfo.TotalFavorited,
				WorkCount:       interactionInfo.WorkCount,
				FavoriteCount:   interactionInfo.FavoriteCount,
				Department:      department.Department,
			})
		}
	}

	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "get userList success",
	}
	return resp, nil
}

// ChangeUserAvatar implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangeUserAvatar(ctx context.Context, req *user.QingyuAvatarChangeRequest) (resp *user.QingyuAvatarChangeResponse, err error) {
	resp = new(user.QingyuAvatarChangeResponse)
	err = s.MysqlManager.ChangeAvatarByUserID(ctx, req.Avatar, req.UserId)
	if err != nil {
		klog.Error("user mysqlManager change avatar failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user mysqlManager change avatar failed",
		}
		return resp, err
	}
	err = s.RedisManager.ChangeAvatarByUserID(ctx, req.Avatar, req.UserId)
	if err != nil {
		klog.Error("user redisManager change avatar failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user redisManager change avatar failed",
		}
		return resp, err
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "change avatar success",
	}
	return
}

// JudgeDoctor implements the UserServiceImpl interface.
func (s *UserServiceImpl) JudgeDoctor(ctx context.Context, req *user.QingyuJudgeDoctorRequest) (resp *user.QingyuJudgeDoctorResponse, err error) {
	resp = new(user.QingyuJudgeDoctorResponse)

	res, err := s.RedisManager.JudgeDoctor(ctx, req.UserId)
	if err != nil {
		klog.Error("user redisManager judge doctor failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user redisManager judge doctor failed",
		}
		return resp, err
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "judge doctor success",
	}
	resp.Department = res
	return
}

// AddDoctor implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddDoctor(ctx context.Context, req *user.QingyuAddDoctorRequest) (resp *user.QingyuAddDoctorResponse, err error) {
	resp = new(user.QingyuAddDoctorResponse)

	err = s.RedisManager.AddDoctor(ctx, req.UserId, req.Department)
	if err != nil {
		klog.Error("user redisManager add doctor failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user redisManager add doctor failed",
		}
		return resp, err
	}
	err = s.MysqlManager.AddDoctor(ctx, req.UserId, req.Department)
	if err != nil {
		klog.Error("user mysqlManager add doctor failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user mysqlManager add doctor failed",
		}
		return resp, err
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "add doctor success",
	}
	return
}

// FindDoctor implements the UserServiceImpl interface.
func (s *UserServiceImpl) FindDoctor(ctx context.Context, req *user.QingyuFindDoctorRequest) (resp *user.QingyuFindDoctorResponse, err error) {
	resp = new(user.QingyuFindDoctorResponse)

	res, err := s.RedisManager.FindDoctor(ctx, req.Department)
	if err != nil {
		klog.Error("user redisManager find doctor failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user redisManager find doctor failed",
		}
		return resp, err
	}
	for _, v := range res {
		i, _ := strconv.ParseInt(v, 10, 64)
		resp.DoctorId = append(resp.DoctorId, i)
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "add doctor success",
	}
	return
}

// ChangeUserNickname implements the UserServiceImpl interface.
func (s *UserServiceImpl) ChangeUserNickname(ctx context.Context, req *user.QingyuNicknameChangeRequest) (resp *user.QingyuNicknameChangeResponse, err error) {
	resp = new(user.QingyuNicknameChangeResponse)
	err = s.MysqlManager.ChangeNicknameByUserID(ctx, req.Nickname, req.UserId)
	if err != nil {
		klog.Error("user mysqlManager change avatar failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user mysqlManager change avatar failed",
		}
		return resp, err
	}
	err = s.RedisManager.ChangeNicknameByUserID(ctx, req.Nickname, req.UserId)
	if err != nil {
		klog.Error("user redisManager change avatar failed,", err)
		resp.BaseResp = &base.QingyuBaseResponse{
			StatusCode: 500,
			StatusMsg:  "user redisManager change avatar failed",
		}
		return resp, err
	}
	resp.BaseResp = &base.QingyuBaseResponse{
		StatusCode: 0,
		StatusMsg:  "change avatar success",
	}
	return
}

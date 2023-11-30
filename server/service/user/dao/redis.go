package dao

import (
	"GreenFish/server/service/user/model"
	"context"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type RedisManager struct {
	redisClient *redis.Client
}

func NewRedisManager(client *redis.Client) *RedisManager {
	return &RedisManager{redisClient: client}
}

func (r *RedisManager) CreateUser(ctx context.Context, user *model.User) error {
	userJson, err := sonic.Marshal(user)
	if err != nil {
		klog.Error("redis marshal user failed,", err)
		return err
	}
	err = r.redisClient.Set(ctx, "user:"+strconv.FormatInt(user.ID, 10), userJson, 0).Err()
	if err != nil {
		klog.Error("redis create user failed,", err)
		return err
	}
	return nil
}

func (r *RedisManager) GetUserById(ctx context.Context, id int64) (*model.User, error) {
	userJson, err := r.redisClient.Get(ctx, "user:"+strconv.FormatInt(id, 10)).Bytes()
	if err != nil && err != redis.Nil {
		klog.Error("redis get user by id failed,", err)
		return nil, err
	}
	var user *model.User
	err = sonic.Unmarshal(userJson, &user)
	if err != nil {
		klog.Error("redis unmarshal user failed,", err)
		return nil, err
	}
	return user, nil
}

func (r *RedisManager) BatchGetUserById(ctx context.Context, id []int64) ([]*model.User, error) {
	var userList []*model.User
	for _, v := range id {
		userJson, err := r.redisClient.Get(ctx, "user:"+strconv.FormatInt(v, 10)).Bytes()
		if err != nil && err != redis.Nil {
			klog.Error("redis get user by id failed,", err)
			return nil, err
		}
		var user *model.User
		err = sonic.Unmarshal(userJson, &user)
		if err != nil {
			klog.Error("redis unmarshal user failed,", err)
			return nil, err
		}
		userList = append(userList, user)
	}
	return userList, nil
}

func (r *RedisManager) UpdateIssueListById(ctx context.Context, id int64, issueList *model.IssueList) error {
	issueListJson, err := sonic.Marshal(issueList)
	if err != nil {
		klog.Error("redis marshal issueList failed,", err)
		return err
	}
	err = r.redisClient.Set(ctx, "issueListJson:"+strconv.FormatInt(id, 10), issueListJson, 0).Err()
	if err != nil {
		klog.Error("redis create issueList failed,", err)
		return err
	}
	return nil
}

func (r *RedisManager) GetIssueListById(ctx context.Context, id int64) (*model.IssueList, error) {
	res, err := r.redisClient.Get(ctx, "issueListJson:"+strconv.FormatInt(id, 10)).Result()
	if err != nil {
		klog.Error("redis get issueList failed,", err)
		return nil, err
	}
	var issueList *model.IssueList
	err = sonic.UnmarshalString(res, &issueList)
	if err != nil {
		klog.Error("redis unmarshal issueList failed,", err)
		return nil, err
	}
	return issueList, nil
}

func (r *RedisManager) ChangeAvatarByUserID(ctx context.Context, avatar string, id int64) error {
	user, err := r.GetUserById(ctx, id)
	if err != nil {
		klog.Error("redis get user by id failed,", err)
		return err
	}
	user.Avatar = avatar
	err = r.redisClient.Set(ctx, "user:"+strconv.FormatInt(user.ID, 10), user, 0).Err()
	if err == nil {
		klog.Error("redis update user avatar failed,", err)
		return err
	}
	return nil
}

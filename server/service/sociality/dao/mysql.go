package dao

import (
	"GreenFish/server/common/consts"
	"GreenFish/server/service/sociality/model"
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type MysqlManager struct {
	db *gorm.DB
}

// IfExist 用来检测是否存在对应表
func (m MysqlManager) isExist(userId int64, option int8) (bool, error) {
	var temp model.ConcernList
	switch option {
	case consts.FollowList:
		err := m.db.Where("follower_id = ?", userId).First(&temp).Error
		if err != nil && err == gorm.ErrRecordNotFound {
			return false, nil
		}
		if err != nil && err != gorm.ErrRecordNotFound {
			return false, err
		}
		return true, nil

	case consts.FollowerList:
		err := m.db.Where("user_id = ?", userId).First(&temp).Error
		if err != nil && err == gorm.ErrRecordNotFound {
			return false, nil
		}
		if err != nil && err != gorm.ErrRecordNotFound {
			return false, err
		}
		return true, nil
	case consts.FriendsList:
		isConcern := true
		isFollow := true
		flag := true
		err1 := m.db.Where("follower_id = ?", userId).First(&temp).Error
		if err1 != nil && err1 == gorm.ErrRecordNotFound {
			isConcern = false
		}
		if err1 != nil && err1 != gorm.ErrRecordNotFound {
			flag = false
		}
		err2 := m.db.Where("user_id = ?", userId).First(&temp).Error
		if err2 != nil && err2 == gorm.ErrRecordNotFound {
			isFollow = false
		}
		if err2 != nil && err2 != gorm.ErrRecordNotFound {
			flag = false
		}

		if isConcern && isFollow && flag == true {
			return true, nil
		}
		return false, nil

	}

	return false, errors.New("invalid option")
}

func (m MysqlManager) GetUserIdList(ctx context.Context, userId int64, option int8) ([]int64, error) {
	flag, err := m.isExist(userId, option)
	if err != nil {
		return nil, err
	}

	if !flag {
		return nil, nil
	}

	switch option {
	case consts.FollowList:
		var concernList []*model.ConcernList
		if err = m.db.Where("follower_id = ?", userId).Find(&concernList).Error; err != nil {

			return nil, err
		}
		idList := make([]int64, len(concernList))
		for _, v := range concernList {
			idList = append(idList, v.UserId)
		}

		return idList, nil
	case consts.FollowerList:
		var followerList []*model.ConcernList
		if err = m.db.Where("user_id = ?", userId).Find(&followerList).Error; err != nil {
			return nil, err
		}
		idList := make([]int64, len(followerList))
		for _, v := range followerList {
			idList = append(idList, v.FollowerId)
		}

		return idList, nil
	case consts.FriendsList:
		var results []*model.ConcernList
		err = m.db.Distinct().Select("user_id, follower_id"). //复杂查询，查找互关数据
									Where("user_id IN (?) OR follower_id IN (?)",
				m.db.Table("concern_lists").Select("user_id").Where("follower_id = ?", userId),
				m.db.Table("concern_lists").Select("follower_id").Where("user_id = ?", userId).
					Or("user_id = ? AND follower_id = ?", userId, userId)).
			Find(&results).Error
		if err != nil {
			return nil, err
		}
		idList := make([]int64, len(results)+1)
		for _, v := range results {
			if v.UserId == userId {
				idList = append(idList, v.FollowerId)
			}
		}

		return idList, nil
	}
	return nil, err

}

func (m MysqlManager) GetSocialInfo(ctx context.Context, userId int64, viewerId int64) (*model.SocialInfo, error) {
	concernIdList, err := m.GetUserIdList(ctx, userId, consts.FollowList)
	if err != nil {
		klog.Errorf("get IdList wrong")
		return nil, err
	}
	followerIdList, err := m.GetUserIdList(ctx, userId, consts.FollowerList)
	if err != nil {
		klog.Errorf("get IdList wrong")
		return nil, err
	}
	var flag bool
	for _, v := range followerIdList {
		if v == viewerId {
			flag = true
		}
	}
	return &model.SocialInfo{
		FollowCount:   int64(len(concernIdList)),
		FollowerCount: int64(len(followerIdList)),
		IsFollow:      flag,
	}, nil
}

func (m MysqlManager) BatchGetSocialInfo(ctx context.Context, userId []int64, viewerId int64) ([]*model.SocialInfo, error) {
	var res []*model.SocialInfo
	for _, v := range userId {
		socialInfo, err := m.GetSocialInfo(ctx, v, viewerId)
		if err != nil {
			return nil, err
		}
		res = append(res, socialInfo)
	}
	return res, nil
}

func (m MysqlManager) HandleSocialInfo(ctx context.Context, userId int64, toUserId int64, actionType int8) error {
	var temp model.ConcernList
	err := m.db.Where("user_id = ? AND follower_id = ? ", toUserId, userId).First(&temp).Error
	switch actionType {
	case consts.Follow:
		if err != nil && err != gorm.ErrRecordNotFound { //出错返回err
			return err
		}
		if err != nil && err == gorm.ErrRecordNotFound { //无数据则插入数据
			err = m.db.Create(&model.ConcernList{
				UserId:     userId,
				FollowerId: toUserId,
			}).Error
			if err != nil {
				return err
			}
			return nil
		}
		//没有出错则说明表中存在数据,无需额外改动
		return nil
	case consts.UnFollow:
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if err != nil && err == gorm.ErrRecordNotFound { //找不到数据则说明没有关注，不作改动
			return nil
		}
		//找到了数据则进行删除
		err = m.db.Where("user_id = ? AND follower_id = ?", toUserId, userId).Delete(&model.ConcernList{}).Error
		if err != nil {
			return err
		}

		return nil
	}

	return errors.New("invalid action_type")

}

func NewMysqlManager(db *gorm.DB) *MysqlManager {
	m := db.Migrator()
	if !m.HasTable(&model.ConcernList{}) {
		err := m.CreateTable(&model.ConcernList{})
		if err != nil {
			klog.Errorf("create mysql table failed,", err)
		}
	}
	return &MysqlManager{db: db}
}

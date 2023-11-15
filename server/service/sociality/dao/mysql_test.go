package dao

import (
	"GreenFish/server/service/sociality/config"
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkMysqlManager_GetUserIdList(b *testing.B) {
	rand.Seed(time.Now().Unix())
	m := GetMysqlManager()
	for n := 0; n < b.N; n++ {
		m.GetUserIdList(context.TODO(), 1693512896484478976, int8(rand.Intn(12)%3))
	}
}

func BenchmarkMysqlManager_GetSocialInfo(b *testing.B) {
	m := GetMysqlManager()
	for n := 0; n < b.N; n++ {
		m.GetSocialInfo(context.TODO(), 1693512896484478976, 1693526693303554048)
	}
}

func BenchmarkMysqlManager_BatchGetSocialInfo(b *testing.B) {
	m := GetMysqlManager()
	userId := []int64{1693526693303554048, 1693602654074179584, 1693605866164457472}
	for n := 0; n < b.N; n++ {
		m.BatchGetSocialInfo(context.TODO(), userId, 1693512896484478976)
	}
}

func GetMysqlManager() *MysqlManager {
	InitNacos()
	c := config.GlobalServerConfig.MysqlInfo
	db, _ := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Name)),
		&gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	m := MysqlManager{db: db}
	return &m
}

func InitNacos() {
	v := viper.New()
	v.SetConfigFile("../config.yaml")
	if err := v.ReadInConfig(); err != nil {
		klog.Fatalf("read viper config failed: %s", err)
	}
	if err := v.Unmarshal(&config.GlobalNacosConfig); err != nil {
		klog.Fatalf("unmarshal err failed: %s", err)
	}

	sc := []constant.ServerConfig{
		*constant.NewServerConfig(config.GlobalNacosConfig.Host, config.GlobalNacosConfig.Port),
	}

	cc := constant.ClientConfig{
		NamespaceId:         config.GlobalNacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		Username:            config.GlobalNacosConfig.User,
		Password:            config.GlobalNacosConfig.Password,
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		klog.Fatalf("create config client failed: %s", err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: config.GlobalNacosConfig.DataId,
		Group:  config.GlobalNacosConfig.Group,
	})
	if err != nil {
		klog.Fatalf("get config failed: %s", err.Error())
	}

	err = sonic.Unmarshal([]byte(content), &config.GlobalServerConfig)
	if err != nil {
		klog.Fatalf("nacos config failed: %s", err)
	}
}

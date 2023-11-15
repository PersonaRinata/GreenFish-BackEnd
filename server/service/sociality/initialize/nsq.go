package initialize

import (
	"GreenFish/server/common/consts"
	"GreenFish/server/service/sociality/config"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nsqio/go-nsq"
)

func InitProducer() *nsq.Producer {
	producer, err := nsq.NewProducer(config.GlobalServerConfig.NsqInfo.Host+":"+config.GlobalServerConfig.NsqInfo.Port, nsq.NewConfig())
	if err != nil {
		klog.Error("initialize nsq producer failed,", err)
		return nil
	}
	return producer
}

func InitSubscriber() *nsq.Consumer {
	subscriber, err := nsq.NewConsumer(consts.NsqSocialityTopic, consts.NsqSocialityChannel, nsq.NewConfig())
	if err != nil {
		klog.Error("nsq initialize subscriber failed,", err)
		return nil
	}
	return subscriber
}

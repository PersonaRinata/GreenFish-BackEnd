package config

import (
	"GreenFish/server/kitex_gen/chat/chatservice"
	"GreenFish/server/kitex_gen/interaction/interactionserver"
	"GreenFish/server/kitex_gen/sociality/socialityservice"
	"GreenFish/server/kitex_gen/user/userservice"
	"GreenFish/server/kitex_gen/video/videoservice"
	"github.com/minio/minio-go/v7"
)

var (
	GlobalServerConfig = &ServerConfig{}
	GlobalNacosConfig  = &NacosConfig{}

	GlobalChatClient        chatservice.Client
	GlobalUserClient        userservice.Client
	GlobalVideoClient       videoservice.Client
	GlobalSocialClient      socialityservice.Client
	GlobalInteractionClient interactionserver.Client
	GlobalMinioClient       *minio.Client
)

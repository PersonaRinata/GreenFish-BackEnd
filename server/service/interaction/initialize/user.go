package initialize

import (
	"GreenFish/server/common/consts"
	user "GreenFish/server/kitex_gen/user/userservice"
	"GreenFish/server/service/interaction/config"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	nacos "github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// InitUser init user service.
func InitUser() user.Client {
	// init resolver
	// Read configuration information from nacos
	sc := []constant.ServerConfig{
		{
			IpAddr: config.GlobalNacosConfig.Host,
			Port:   config.GlobalNacosConfig.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         config.GlobalNacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              consts.NacosLogDir,
		CacheDir:            consts.NacosCacheDir,
		LogLevel:            consts.NacosLogLevel,
	}

	nacosCli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		})
	r := nacos.NewNacosResolver(nacosCli, nacos.WithGroup("USER_GROUP"))
	if err != nil {
		klog.Fatalf("new nacos client failed: %s", err.Error())
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)

	// create a new client
	c, err := user.NewClient(
		config.GlobalServerConfig.UserSrvInfo.Name,
		client.WithResolver(r),                                     // service discovery
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()), // load balance
		client.WithMuxConnection(1),                                // multiplexing
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.UserSrvInfo.Name}),
	)
	if err != nil {
		klog.Fatalf("ERROR: cannot init client: %v\n", err)
	}
	return c
}

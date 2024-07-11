package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/spf13/viper"
	"go-cloud-storage/app/user/global"
	"go-cloud-storage/app/user/internal/cache/redis"
	"go-cloud-storage/app/user/internal/service"
	"go-cloud-storage/app/user/internal/store/mysql"
	"go-cloud-storage/pb/kitex_gen/api/v1/userservice"
	"go-cloud-storage/pkg/config"
	"go-cloud-storage/pkg/logger"
	"net"
)

type Config struct {
	Server config.Server `yaml:"server"`
	Mysql  config.MySQL  `yaml:"mysql"`
	Redis  config.Redis  `yaml:"redis"`
	Etcd   config.Etcd   `yaml:"etcd"`
}

var Conf *Config

func InitConfig() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath(global.RootDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
	logger.Infow("user-server config finished", "path", viper.ConfigFileUsed())
}

func main() {
	// 初始化配置参数
	InitConfig()
	// 数据库初始化
	mysql.InitMysql(Conf.Mysql)
	// 缓存初始化
	redis.InitRedis(Conf.Redis)
	// 服务注册
	r, err := etcd.NewEtcdRegistry([]string{Conf.Etcd.Addr})
	if err != nil {
		logger.Errorw("etcd connect failed", "err", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", Conf.Server.Addr)
	if err != nil {
		logger.Fatalw("resolve tcp addr failed", "err", err)
	}
	svr := userservice.NewServer(
		service.NewUserServiceImpl(),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: Conf.Server.Name}),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
		server.WithRegistry(r),
	)

	if err = svr.Run(); err != nil {
		logger.Fatalw("user app run failed", "err", err)
	}
}

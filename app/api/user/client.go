package user

import (
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"go-cloud-storage/app/api"
	"go-cloud-storage/pb/kitex_gen/api/v1/userservice"
	"go-cloud-storage/pkg/logger"
)

func InitUserClient() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		logger.Errorw("etcd resolver init failed", "err", err)
	}
	api.UserClient = userservice.MustNewClient("user", client.WithResolver(r))
}

package service

import (
	"context"
	"errors"
	"go-cloud-storage/app/user/global"
	"go-cloud-storage/app/user/internal/cache"
	"go-cloud-storage/app/user/internal/model"
	"go-cloud-storage/app/user/internal/store"
	"go-cloud-storage/common/errno"
	v1 "go-cloud-storage/pb/kitex_gen/api/v1"
	"go-cloud-storage/pkg/logger"
	"gorm.io/gorm"
	"time"
)

type UserServiceImpl struct {
	userCache *cache.UserCache
	userStore *store.UserStore
}

var _ v1.UserService = (*UserServiceImpl)(nil)

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		userCache: cache.NewUserCache(),
		userStore: store.NewUserStore(),
	}
}

// Signup 实现用户注册功能。
// 根据请求中的用户名和密码创建新用户，并返回注册结果。
func (u UserServiceImpl) Signup(ctx context.Context, req *v1.SignupRequest) (res *v1.SignupResponse, err error) {
	res = &v1.SignupResponse{
		Code:    int32(errno.OK.HTTP),
		Message: errno.OK.Message,
	}
	// 1、验证用户名是否已存在
	exist, err := u.userStore.ExistUsername(ctx, req.UserName)
	if err != nil {
		logger.Errorw("validate username error", "err", err)
		res.Code = int32(errno.InternalServerError.HTTP)
		res.Message = errno.InternalServerError.Message
		return res, nil
	}
	if exist {
		res.Code = int32(errno.ErrUserAlreadyExist.HTTP)
		res.Message = errno.ErrUserAlreadyExist.Message
		return res, nil
	}
	// 2、创建用户
	userID, err := global.SnowFlake.NextId()
	if err != nil {
		logger.Errorw("create userID error", "err", err)
		res.Code = int32(errno.InternalServerError.HTTP)
		res.Message = errno.InternalServerError.Message
		return res, nil
	}
	passwordMD5 := model.Md5Crypt(req.Password, userID)
	user := &model.User{
		Username: req.UserName,
		Password: passwordMD5,
	}
	user.ID = uint(userID)

	// 3、存入数据库
	if err := u.userStore.Create(ctx, user); err != nil {
		logger.Errorw("create user error", "err", err)
		res.Code = int32(errno.InternalServerError.HTTP)
		res.Message = errno.InternalServerError.Message
		return res, nil
	}
	res.UserId = int64(user.ID)
	// 4、存入缓存
	if err = cache.NewUserCache().HSetUserInfo(
		ctx,
		model.CreateUserKey(user.ID),
		model.CreateMapUserInfo(user),
	); err != nil {
		logger.Errorw("create user redis error", "err", err)
		res.Code = int32(errno.InternalServerError.HTTP)
		res.Message = errno.InternalServerError.Message
		return res, nil
	}
	// 5.生成token
	token, err := global.Jwt.Award(int(user.ID))
	if err != nil {
		logger.Errorw("create token error", "err", err)
		res.Code = int32(errno.InternalServerError.HTTP)
		res.Message = errno.InternalServerError.Message
		return res, nil
	}
	// 6.缓存token
	if err = cache.NewUserCache().SetToken(ctx, model.CreateUserKey(user.ID), token, 7*24*time.Hour); err != nil {
		logger.Errorw("create token redis error", "err", err)
		res.Code = int32(errno.InternalServerError.HTTP)
		res.Message = errno.InternalServerError.Message
		return res, err
	}
	res.Token = token
	return res, nil
}

// Login 实现用户登录逻辑。
// ctx: 上下文对象，用于控制请求的超时和传递请求相关的值。
// req: 包含用户名和密码的登录请求。
// 返回登录响应和可能的错误。
func (u UserServiceImpl) Login(ctx context.Context, req *v1.LoginRequest) (res *v1.LoginResponse, err error) {
	res = &v1.LoginResponse{
		Code:    int32(errno.OK.HTTP),
		Message: errno.OK.Message,
	}
	// 1、设置登录超时阈值
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	// 2、获取用户信息
	user, err := u.userStore.GetUserByUsername(ctx, req.UserName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		res.Code = int32(errno.ErrUserNotExist.HTTP)
		res.Message = errno.ErrUserNotExist.Message
		return res, nil
	}
	if err != nil {
		logger.Errorw("get user error", "err", err)
		res.Code = int32(errno.InternalServerError.HTTP)
		res.Message = errno.InternalServerError.Message
		return res, nil
	}
	// 3、用户密码判断
	if model.Md5Crypt(req.Password, user.ID) != user.Password {
		res.Code = int32(errno.ErrPasswordIncorrect.HTTP)
		res.Message = errno.ErrPasswordIncorrect.Message
		return res, nil
	}
	// 4、生成token
	token, err := global.Jwt.Award(int(user.ID))
	if err != nil {
		logger.Errorw("create token error", "err", err)
		res.Code = int32(errno.InternalServerError.HTTP)
		res.Message = errno.InternalServerError.Message
		return res, nil
	}
	// 5.缓存token
	if err = u.userCache.SetToken(ctx, model.CreateUserKey(user.ID), token, 7*24*time.Hour); err != nil {
		logger.Errorw("create token redis error", "err", err)
		res.Code = int32(errno.InternalServerError.HTTP)
		res.Message = errno.InternalServerError.Message
		return res, err
	}
	res.Token = token
	res.UserId = int64(user.ID)
	return res, nil
}

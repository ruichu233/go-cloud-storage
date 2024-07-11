package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go-cloud-storage/app/api"
	"go-cloud-storage/common/errno"
	"go-cloud-storage/common/response"
	v1 "go-cloud-storage/pb/kitex_gen/api/v1"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Signup(ctx context.Context, c *app.RequestContext) {
	var req v1.SignupRequest
	if err := c.Bind(&req); err != nil {
		response.WriteResponse(c, errno.ErrBind, nil)
	}

	r, _ := api.UserClient.Signup(ctx, &req)
	response.WriteResponse(c, nil, r)
}

func (h *Handler) Login(ctx context.Context, c *app.RequestContext) {

	var req v1.LoginRequest
	if err := c.Bind(&req); err != nil {
		response.WriteResponse(c, errno.ErrBind, nil)
	}

	r, _ := api.UserClient.Login(ctx, &req)
	response.WriteResponse(c, nil, r)
}

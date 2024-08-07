// Code generated by Kitex v0.10.1. DO NOT EDIT.

package userservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	v1 "go-cloud-storage/pb/kitex_gen/api/v1"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for app method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Signup": kitex.NewMethodInfo(
		signupHandler,
		newSignupArgs,
		newSignupResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newLoginArgs,
		newLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	userServiceServiceInfo                = NewServiceInfo()
	userServiceServiceInfoForClient       = NewServiceInfoForClient()
	userServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*v1.UserService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.10.1",
		Extra:           extra,
	}
	return svcInfo
}

func signupHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(v1.SignupRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(v1.UserService).Signup(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SignupArgs:
		success, err := handler.(v1.UserService).Signup(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SignupResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSignupArgs() interface{} {
	return &SignupArgs{}
}

func newSignupResult() interface{} {
	return &SignupResult{}
}

type SignupArgs struct {
	Req *v1.SignupRequest
}

func (p *SignupArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(v1.SignupRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SignupArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SignupArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SignupArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SignupArgs) Unmarshal(in []byte) error {
	msg := new(v1.SignupRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SignupArgs_Req_DEFAULT *v1.SignupRequest

func (p *SignupArgs) GetReq() *v1.SignupRequest {
	if !p.IsSetReq() {
		return SignupArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SignupArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SignupArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SignupResult struct {
	Success *v1.SignupResponse
}

var SignupResult_Success_DEFAULT *v1.SignupResponse

func (p *SignupResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(v1.SignupResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SignupResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SignupResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SignupResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SignupResult) Unmarshal(in []byte) error {
	msg := new(v1.SignupResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SignupResult) GetSuccess() *v1.SignupResponse {
	if !p.IsSetSuccess() {
		return SignupResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SignupResult) SetSuccess(x interface{}) {
	p.Success = x.(*v1.SignupResponse)
}

func (p *SignupResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SignupResult) GetResult() interface{} {
	return p.Success
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(v1.LoginRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(v1.UserService).Login(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *LoginArgs:
		success, err := handler.(v1.UserService).Login(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*LoginResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newLoginArgs() interface{} {
	return &LoginArgs{}
}

func newLoginResult() interface{} {
	return &LoginResult{}
}

type LoginArgs struct {
	Req *v1.LoginRequest
}

func (p *LoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(v1.LoginRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *LoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *LoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *LoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *LoginArgs) Unmarshal(in []byte) error {
	msg := new(v1.LoginRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var LoginArgs_Req_DEFAULT *v1.LoginRequest

func (p *LoginArgs) GetReq() *v1.LoginRequest {
	if !p.IsSetReq() {
		return LoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *LoginArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *LoginArgs) GetFirstArgument() interface{} {
	return p.Req
}

type LoginResult struct {
	Success *v1.LoginResponse
}

var LoginResult_Success_DEFAULT *v1.LoginResponse

func (p *LoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(v1.LoginResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *LoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *LoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *LoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *LoginResult) Unmarshal(in []byte) error {
	msg := new(v1.LoginResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *LoginResult) GetSuccess() *v1.LoginResponse {
	if !p.IsSetSuccess() {
		return LoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *LoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*v1.LoginResponse)
}

func (p *LoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *LoginResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Signup(ctx context.Context, Req *v1.SignupRequest) (r *v1.SignupResponse, err error) {
	var _args SignupArgs
	_args.Req = Req
	var _result SignupResult
	if err = p.c.Call(ctx, "Signup", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, Req *v1.LoginRequest) (r *v1.LoginResponse, err error) {
	var _args LoginArgs
	_args.Req = Req
	var _result LoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

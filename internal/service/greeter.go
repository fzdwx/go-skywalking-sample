package service

import (
	"context"
	"helloworld/pkg/logx"

	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	log logx.Logger
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, l logx.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: l}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	logx.Zap.Info("say hello " + in.Name)
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

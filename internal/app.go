package internal

import (
	"context"
	micro_old_server "github.com/pr0head/micro-old-server/pkg"
	microHttp "go.unistack.org/micro-network-transport-http/v3"
	etcd "go.unistack.org/micro-register-etcd/v3"
	"go.unistack.org/micro/v3"
	"go.unistack.org/micro/v3/register"
	"go.unistack.org/micro/v3/server"
)

type Application struct {
	service micro.Service
	svc     *Service
}

type Service struct {
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Run() {
	app.svc = &Service{}

	tr := microHttp.NewTransport()
	r := etcd.NewRegister(
		register.Addrs("127.0.0.1:2379"),
	)

	serverOptions := []server.Option{
		server.Register(r),
		server.Name(micro_old_server.ServiceName),
		server.Transport(tr),
	}
	microServer := server.NewServer(serverOptions...)

	options := []micro.Option{
		micro.Name(micro_old_server.ServiceName),
		micro.Server(microServer),
	}

	app.service = micro.NewService(options...)
	app.service.Init()

	if err := micro_old_server.RegisterMicroOldServerServer(app.service.Server(), app.svc); err != nil {
		panic("Can`t register service in micro")
	}

	if err := app.service.Run(); err != nil {
		panic("Can`t run service")
	}
}

func (s *Service) Test(_ context.Context, req *micro_old_server.TestRequest, res *micro_old_server.TestResponse) error {
	res.Output = req.Input

	return nil
}

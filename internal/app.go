package internal

import (
	"context"
	"github.com/micro/go-micro"
	micro_old_server "github.com/pr0head/micro-old-server/pkg"
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

	options := []micro.Option{
		micro.Name(micro_old_server.ServiceName),
	}

	app.service = micro.NewService(options...)
	app.service.Init()

	if err := micro_old_server.RegisterMicroOldServerHandler(app.service.Server(), app.svc); err != nil {
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

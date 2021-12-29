package container

import (
	"fmt"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/config"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller"
	controller_process_tracing "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/controller/product"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/infrastructure/database"
	grpcserver "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/infrastructure/grpc_server"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/infrastructure/jaeger"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/infrastructure/logrus"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/repository/postgres"
	service_process_tracing "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/service/product"
	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/utils"
	"go.uber.org/dig"
)

type Container struct {
	container *dig.Container
}

func BindTracingRepository(pgRep *postgres.Repository) utils.Repository {
	return pgRep
}

func BindTracingService(service *service_process_tracing.ProcessTracingService) service_process_tracing.Service {
	return service
}

func (c *Container) Configure() error {
	servicesConstructors := []interface{}{
		config.NewConfiguration,
		BindTracingRepository,
		BindTracingService,
		grpcserver.NewServer,
		database.NewServerBase,
		jaeger.NewJaeger,
		logrus.NewLog,
		controller.NewHealthZController,
		controller.NewPingPongController,
		controller_process_tracing.NewController,
		service_process_tracing.NewService,
		postgres.NewRepository,
		utils.NewUtils,
		utils.NewCustomValidator,
	}

	for _, service := range servicesConstructors {
		if err := c.container.Provide(service); err != nil {
			return err
		}
	}
	//appConfig := config.NewConfiguration()
	//jaeger.NewJaeger(appConfig)
	return nil
}

func (c *Container) Start() error {
	fmt.Println("Start Container")

	if err := c.container.Invoke(func(s *grpcserver.Server) {
		s.Start()
	}); err != nil {
		fmt.Printf("%s", err)

		return err
	}

	return nil
}

//MigrateDB ...
func (c *Container) MigrateDB() error {
	fmt.Println("Start Container DB")

	if err := c.container.Invoke(func(d *database.DB) {
		d.MigrateDB()
	}); err != nil {
		return err
	}

	return nil
}

func NewContainer() (*Container, error) {
	d := dig.New()

	container := &Container{
		container: d,
	}

	if err := container.Configure(); err != nil {
		return nil, err
	}

	return container, nil
}

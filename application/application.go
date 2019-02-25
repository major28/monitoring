package application

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/major28/monitoring/proto"
	"github.com/major28/monitoring/services"
	"github.com/major28/monitoring/utils/config"
	"github.com/major28/monitoring/utils/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type App struct {
	Resources IResources
	Logger    logger.ILogger
}

func NewApp() *App {
	return &App{}
}

func (s *App) Initialize(configPath string, env string) {
	config.LoadConfig(configPath, env)
	s.initLogger()
	s.initResources()
}

func (s *App) Run(ctx context.Context) {

	s.Logger.Warning("starting monitoring...")
	monitoringService := services.NewMonitoringService(s.Resources.GetDbManager(), s.Logger)
	err := monitoringService.Run(ctx)
	if err != nil {
		s.Logger.Critical("can't run monitoring servers %#v", err)
		return
	}

	s.Logger.Warning("starting gRPC server...")
	port := config.GetEnvConfig("port.api")
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		s.Logger.Critical("failed to listen %#v", err)
		return
	}
	s.Logger.Warning("start listening for monitoring at port %s", ":"+port)

	usersService := services.NewUsersService(s.Resources.GetDbManager(), s.Logger)
	server := NewGRPCServer(s.Logger, monitoringService, usersService)

	rpcServer := grpc.NewServer(grpc.UnaryInterceptor(server.AuthInterceptor))
	monitoring.RegisterMonitoringServer(rpcServer, server)
	reflection.Register(rpcServer)
	err = rpcServer.Serve(listener)
	if err != nil {
		s.Logger.Critical("failed to serve %#v", err)
		return
	}
}

func (s *App) initResources() {
	connectUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		config.GetEnvConfig("mysql.master.user"),
		config.GetEnvConfig("mysql.master.password"),
		config.GetEnvConfig("mysql.master.host"),
		config.GetEnvConfig("mysql.master.database"),
	)
	db, err := gorm.Open("mysql", connectUrl)
	if err != nil {
		panic(err)
	}

	s.Resources = &ResourceCollection{
		DbManager: db,
	}
}

func (s *App) initLogger() {
	s.Logger = logger.NewLogger(config.GetEnvConfig("log-level"))
}

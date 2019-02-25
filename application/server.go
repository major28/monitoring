package application

import (
	"context"

	"github.com/major28/monitoring/proto"
	"github.com/major28/monitoring/services"
	"github.com/major28/monitoring/utils/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

type server struct {
	logger            logger.ILogger
	monitoringService services.MonitoringIService
	usersService      services.UsersIService
}

func NewGRPCServer(logger logger.ILogger, monitoringService services.MonitoringIService, usersService services.UsersIService) *server {
	return &server{
		logger:            logger,
		monitoringService: monitoringService,
		usersService:      usersService,
	}
}

func (s *server) AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "access denied")
	}

	username, ok := md["username"]
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "access denied")
	}

	password, ok := md["password"]
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "access denied")
	}

	isAuth, err := s.usersService.CheckAuth(ctx, username[0], password[0])
	if err != nil {
		s.logger.Critical("CheckAuth error %#v", err)
		return nil, status.Errorf(codes.Unauthenticated, "access denied")
	}
	if !isAuth {
		return nil, status.Errorf(codes.Unauthenticated, "access denied")
	}

	ctx = context.WithValue(ctx, services.ContextKeyUsername, username)

	return handler(ctx, req)
}

func (s *server) GetServiceInfoByName(ctx context.Context, req *monitoring.GetServiceInfoByNameRequest) (*monitoring.GetServiceInfoByNameResponse, error) {

	serviceData, err := s.monitoringService.GetServiceByName(ctx, req.Name)
	if err != nil {
		s.logger.Error("Can't get service data (request: %#v, error: %#v)", req, err)
		return nil, status.Errorf(codes.Unknown, "Can't get information for service %s", req.Name)
	}

	err = s.usersService.AddRequestInHistory(ctx, "call method GetServiceInfoByName (service name: "+req.Name+")")
	if err != nil {
		s.logger.Error("can't save request in history (request: %#v, error: %#v)", req, err)
		return nil, status.Errorf(codes.Unknown, "unknown error")
	}

	responseTimeMs := *serviceData.ResponseTime / 1000000
	return &monitoring.GetServiceInfoByNameResponse{
		IsAvailable:      *serviceData.IsAvailable,
		LastResponseTime: responseTimeMs,
	}, nil
}

func (s *server) GetServiceWithMinResponseTime(ctx context.Context, req *monitoring.EmptyRequest) (*monitoring.ServiceInfoResponse, error) {

	serviceData, err := s.monitoringService.GetServiceWithMinOrMaxResponseTime(ctx, false)
	if err != nil {
		s.logger.Error("Can't get service data (request: %#v, error: %#v)", req, err)
		return nil, status.Errorf(codes.Unknown, "Can't get information")
	}

	err = s.usersService.AddRequestInHistory(ctx, "call method GetServiceWithMinResponseTime")
	if err != nil {
		s.logger.Error("can't save request in history (request: %#v, error: %#v)", req, err)
		return nil, status.Errorf(codes.Unknown, "unknown error")
	}

	return &monitoring.ServiceInfoResponse{Name: serviceData.Name}, nil
}

func (s *server) GetServiceWithMaxResponseTime(ctx context.Context, req *monitoring.EmptyRequest) (*monitoring.ServiceInfoResponse, error) {

	serviceData, err := s.monitoringService.GetServiceWithMinOrMaxResponseTime(ctx, true)
	if err != nil {
		s.logger.Error("Can't get service data (request: %#v, error: %#v)", req, err)
		return nil, status.Errorf(codes.Unknown, "Can't get information")
	}

	err = s.usersService.AddRequestInHistory(ctx, "call method GetServiceWithMaxResponseTime")
	if err != nil {
		s.logger.Error("can't save request in history (request: %#v, error: %#v)", req, err)
		return nil, status.Errorf(codes.Unknown, "unknown error")
	}

	return &monitoring.ServiceInfoResponse{Name: serviceData.Name}, nil
}

func (s *server) GetHistoryRequests(ctx context.Context, req *monitoring.EmptyRequest) (*monitoring.GetHistoryRequestsResponse, error) {

	historyRequests, err := s.usersService.GetRequestsFromHistory(ctx)
	if err != nil {
		s.logger.Error("Can't get history request (request: %#v, error: %#v)", req, err)
		return nil, status.Errorf(codes.Unknown, "Can't get information")
	}

	requests := make([]*monitoring.GetHistoryRequest, len(historyRequests))
	for i, request := range historyRequests {
		requests[i] = &monitoring.GetHistoryRequest{
			Request: request.Request,
			Date:    request.CreatedAt.Format(time.RFC3339),
		}
	}

	resp := &monitoring.GetHistoryRequestsResponse{
		GetHistoryRequests: requests,
	}

	return resp, nil
}

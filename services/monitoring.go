package services

import (
	"context"
	"crypto/tls"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/major28/monitoring/utils/config"
	"github.com/major28/monitoring/utils/logger"
	"net/http"
	"time"
)

type MonitoringIService interface {
	Run(ctx context.Context) error
	GetServiceByName(ctx context.Context, name string) (*service, error)
	GetServiceWithMinOrMaxResponseTime(ctx context.Context, isMax bool) (*service, error)
}

type monitoringService struct {
	db     *gorm.DB
	logger logger.ILogger
}

type service struct {
	gorm.Model
	Name         string
	Host         string
	IsAvailable  *bool
	ResponseTime *int64
}

func NewMonitoringService(db *gorm.DB, logger logger.ILogger) MonitoringIService {
	return &monitoringService{
		db:     db,
		logger: logger,
	}
}

func (s *monitoringService) Run(ctx context.Context) error {

	listServices, err := s.getAllServices()
	if err != nil {
		return err
	}

	interval := config.GetEnvConfigInt("checker.interval")
	timeout := config.GetEnvConfigInt("checker.timeout")

	intervalStartingGoroutine := interval / len(listServices)
	if intervalStartingGoroutine < 30 || intervalStartingGoroutine > 1000 {
		intervalStartingGoroutine = 0
	}

	for _, srv := range listServices {
		go s.checker(ctx, srv, interval, timeout)

		if intervalStartingGoroutine > 0 {
			time.Sleep(time.Duration(intervalStartingGoroutine) * time.Millisecond)
		}
	}

	return nil
}

func (s *monitoringService) GetServiceByName(ctx context.Context, name string) (*service, error) {

	resp := &service{}
	err := s.db.Where("name=?", name).Where("is_available IS NOT NULL").Where("response_time IS NOT NULL").First(resp).Error

	return resp, err
}

func (s *monitoringService) GetServiceWithMinOrMaxResponseTime(ctx context.Context, isMax bool) (*service, error) {

	resp := &service{}
	query := s.db.Where("is_available=?", true)
	if isMax {
		query = query.Order("response_time DESC")
	} else {
		query = query.Order("response_time")
	}

	err := query.Limit(1).Find(resp).Error

	return resp, err
}

func (s *monitoringService) checker(ctx context.Context, srv service, checkingInterval int, timeout int) {
	ticker := time.NewTicker(time.Millisecond * time.Duration(checkingInterval))
	for {
		<-ticker.C

		start := time.Now()
		err := s.sendRequest(ctx, srv, timeout)
		if err != nil {
			s.logger.Debug("%s %s \n", srv.Name, err.Error())
			responseTime := time.Since(start).Nanoseconds()
			isAvailable := false
			srv.ResponseTime = &responseTime
			srv.IsAvailable = &isAvailable
			err = s.db.Save(&srv).Error
			if err != nil {
				s.logger.Error("Can't update info for service (server: %#v ,error: %s)", srv, err.Error())
			}
			continue
		}

		responseTime := time.Since(start).Nanoseconds()
		isAvailable := true
		srv.ResponseTime = &responseTime
		srv.IsAvailable = &isAvailable
		err = s.db.Save(&srv).Error
		if err != nil {
			s.logger.Error("Can't update info for service (server: %#v ,error: %s)", srv, err.Error())
			continue
		}
	}
}

func (s *monitoringService) getAllServices() ([]service, error) {
	var results []service
	err := s.db.Find(&results).Error

	return results, err
}

func (s *monitoringService) sendRequest(ctx context.Context, srv service, timeout int) (err error) {

	req, err := http.NewRequest(http.MethodGet, srv.Host, nil)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Millisecond)
	defer cancel()
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer func() {
		if resp.Body == nil {
			return
		}
		if e := resp.Body.Close(); err == nil {
			err = e
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("service %s returned http code %d", srv.Name, resp.StatusCode)
	}

	return nil
}

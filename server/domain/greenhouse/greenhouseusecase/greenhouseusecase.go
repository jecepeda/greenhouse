package greenhouseusecase

import (
	"context"

	"github.com/jecepeda/greenhouse/server/domain/device"
	"github.com/jecepeda/greenhouse/server/domain/greenhouse"
	"github.com/jecepeda/greenhouse/server/gerror"
	"github.com/jecepeda/greenhouse/server/gsql"
)

type Service struct {
	txPool gsql.TransactionPool
	db     gsql.Common
	grepo  greenhouse.Repository
	drepo  device.Repository
}

func NewService(db gsql.Common, txPool gsql.TransactionPool, grepo greenhouse.Repository, drepo device.Repository) *Service {
	return &Service{
		txPool: txPool,
		db:     db,
		grepo:  grepo,
		drepo:  drepo,
	}
}

func (s *Service) SaveMonitoringData(ctx context.Context, data greenhouse.MonitoringData) (greenhouse.MonitoringData, error) {
	errMsg := "save monitoring data"

	_, err := s.drepo.FindByID(ctx, data.DeviceID, s.db)
	if err != nil {
		return greenhouse.MonitoringData{}, gerror.Wrap(err, errMsg)
	}

	tx, err := s.txPool.Start(ctx)
	if err != nil {
		return greenhouse.MonitoringData{}, gerror.Wrap(err, errMsg)
	}
	var success bool
	defer func() {
		gsql.RollbackIfFail(tx, success)
	}()

	data, err = s.grepo.StoreMonitoringData(ctx, tx.GetTx(), data)
	if err != nil {
		return greenhouse.MonitoringData{}, gerror.Wrap(err, errMsg)
	}

	err = gsql.Commit(tx)
	if err != nil {
		return greenhouse.MonitoringData{}, gerror.Wrap(err, errMsg)
	}
	success = true

	return data, nil
}

func (s *Service) FindByID(ctx context.Context, id uint64) (greenhouse.MonitoringData, error) {
	errMsg := "find by id"

	data, err := s.grepo.FindByID(ctx, s.db, id)
	if err != nil {
		return greenhouse.MonitoringData{}, gerror.Wrap(err, errMsg)
	}

	return data, nil
}

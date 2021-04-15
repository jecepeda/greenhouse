package deviceusecase

import (
	"context"

	"github.com/jecepeda/greenhouse/server/crypt"
	"github.com/jecepeda/greenhouse/server/domain/device"
	"github.com/jecepeda/greenhouse/server/gsql"
	"github.com/pkg/errors"
)

type Service struct {
	encrypter crypt.Encrypter
	txPoll    gsql.TransactionPool
	db        gsql.Common
	drepo     device.Repository
}

func NewService(encrypter crypt.Encrypter, db gsql.Common, txPool gsql.TransactionPool, drepo device.Repository) *Service {
	return &Service{
		encrypter: encrypter,
		db:        db,
		drepo:     drepo,
		txPoll:    txPool,
	}
}

func (s *Service) SaveDevice(ctx context.Context, name, password string) (device.Device, error) {
	errMsg := "save device"

	hashedPassword, err := s.encrypter.EncryptPassword(password)
	if err != nil {
		return device.Device{}, errors.Wrap(err, errMsg)
	}

	tx, err := s.txPoll.Start(ctx)
	if err != nil {
		return device.Device{}, errors.Wrap(err, errMsg)
	}
	var success bool
	defer func() {
		gsql.RollbackIfFail(tx, success)
	}()

	created, err := s.drepo.StoreDevice(ctx, tx.GetTx(), name, hashedPassword)
	if err != nil {
		return device.Device{}, errors.Wrap(err, errMsg)
	}
	if err = gsql.Commit(tx); err != nil {
		return device.Device{}, errors.Wrap(err, errMsg)
	}
	success = true

	return created, nil
}

func (s *Service) FindByID(ctx context.Context, deviceID uint64) (device.Device, error) {
	return s.drepo.FindByID(ctx, deviceID, s.db)
}

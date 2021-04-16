package handler

import (
	"github.com/jecepeda/greenhouse/server/crypt"
	"github.com/jecepeda/greenhouse/server/domain/device"
	"github.com/jecepeda/greenhouse/server/domain/greenhouse"
	"github.com/jecepeda/greenhouse/server/gsql"
)

type DependencyContainer interface {
	SetTransactionPool(gsql.TransactionPool)
	GetEncrypter() crypt.Encrypter
	SetEncrypter(encrypter crypt.Encrypter)
	GetDeviceService() device.Service
	GetGreenhouseService() greenhouse.Service
	Init()
	Serve(port string)
}

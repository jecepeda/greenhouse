// Package handler contains the necessary tools
// to run the server
package handler

import (
	"github.com/jecepeda/greenhouse/server/crypt"
	"github.com/jecepeda/greenhouse/server/domain/device"
	"github.com/jecepeda/greenhouse/server/domain/greenhouse"
	"github.com/jecepeda/greenhouse/server/gsql"
)

// DependencyContainer is the main interface
// which is used by handlers in order to access all the services
type DependencyContainer interface {
	SetTransactionPool(gsql.TransactionPool)
	GetEncrypter() crypt.Encrypter
	SetEncrypter(encrypter crypt.Encrypter)
	GetDeviceService() device.Service
	GetGreenhouseService() greenhouse.Service
	Init()
	Serve(port string)
}

package dcontainer

import (
	"github.com/jecepeda/greenhouse/server/crypt"
	"github.com/jecepeda/greenhouse/server/domain/device"
	"github.com/jecepeda/greenhouse/server/domain/device/devicerepo"
	"github.com/jecepeda/greenhouse/server/domain/device/deviceusecase"
	"github.com/jecepeda/greenhouse/server/gsql"
)

// repositories contains the repositories to make the server work
type repositories struct {
	deviceRepo device.Repository
}

// initRepositories inititualices the repositories to make the server work
func initRepositories() repositories {
	return repositories{
		deviceRepo: &devicerepo.Repository{},
	}
}

// DepContainer contains the necessary tools to join
// all services, repositories, third-party services, etc.
// and execute
type DepContainer struct {
	transactionPool gsql.TransactionPool
	db              gsql.Common

	// providers
	encrypter crypt.Encrypter

	// services
	deviceService device.Service
}

func NewDependencyContainer(db gsql.Common) *DepContainer {
	return &DepContainer{
		db: db,
	}
}

func (dc *DepContainer) SetTransactionPool(tp gsql.TransactionPool) {
	dc.transactionPool = tp
}

func (dc *DepContainer) SetEncrypter(encrypter crypt.Encrypter) {
	dc.encrypter = encrypter
}

func (dc *DepContainer) Init() {
	repos := initRepositories()

	dc.deviceService = deviceusecase.NewService(dc.encrypter, dc.db, dc.transactionPool, repos.deviceRepo)
}

package dcontainer

import (
	"github.com/gorilla/mux"
	"github.com/jecepeda/greenhouse/server/auth"
	"github.com/jecepeda/greenhouse/server/crypt"
	"github.com/jecepeda/greenhouse/server/domain/device"
	"github.com/jecepeda/greenhouse/server/domain/device/devicehandler"
	"github.com/jecepeda/greenhouse/server/domain/device/devicerepo"
	"github.com/jecepeda/greenhouse/server/domain/device/deviceusecase"
	"github.com/jecepeda/greenhouse/server/domain/greenhouse"
	"github.com/jecepeda/greenhouse/server/domain/greenhouse/greenhousehandler"
	"github.com/jecepeda/greenhouse/server/domain/greenhouse/greenhouserepo"
	"github.com/jecepeda/greenhouse/server/domain/greenhouse/greenhouseusecase"
	"github.com/jecepeda/greenhouse/server/gsql"
	"github.com/jecepeda/greenhouse/server/handler/router"
	"github.com/urfave/negroni"
)

// repositories contains the repositories to make the server work
type repositories struct {
	deviceRepo     device.Repository
	greenhouseRepo greenhouse.Repository
}

// initRepositories inititualices the repositories to make the server work
func initRepositories() repositories {
	return repositories{
		deviceRepo:     &devicerepo.Repository{},
		greenhouseRepo: &greenhouserepo.Repository{},
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
	deviceService     device.Service
	greenhouseService greenhouse.Service
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

func (dc *DepContainer) GetEncrypter() crypt.Encrypter {
	return dc.encrypter
}

func (dc *DepContainer) GetDeviceService() device.Service {
	return dc.deviceService
}

func (dc *DepContainer) GetGreenhouseService() greenhouse.Service {
	return dc.greenhouseService
}

func (dc *DepContainer) Init() {
	repos := initRepositories()

	dc.deviceService = deviceusecase.NewService(dc.encrypter, dc.db, dc.transactionPool, repos.deviceRepo)
	dc.greenhouseService = greenhouseusecase.NewService(dc.db, dc.transactionPool, repos.greenhouseRepo, repos.deviceRepo)
}

func (dc *DepContainer) Serve(port string) {
	routers := []router.Router{
		devicehandler.NewRouter(dc),
	}
	authRouters := []router.Router{
		greenhousehandler.NewAuthRouter(dc),
	}

	r := mux.NewRouter()
	for _, route := range routers {
		router.AddToMux(r, route)
	}
	ar := mux.NewRouter()
	for _, route := range authRouters {
		router.AddToMux(ar, route)
	}

	an := negroni.New(negroni.HandlerFunc(auth.AuthMiddleware), negroni.Wrap(ar))
	r.PathPrefix("/v1").Handler(an)
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())
	n.UseHandler(r)

	n.Run(":" + port)
}

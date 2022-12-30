package main

import (
	"fmt"
	"github.com/moolah-server-go/infrastructure/services"
	"github.com/moolah-server-go/values"

	_ "github.com/proullon/ramsql/driver"
)

type Application struct {
	accounts services.Accounts
	router   services.Router
}

func NewApplication(config values.Config) (Application, error) {
	accounts, err := services.NewAccounts(config)
	if err != nil {
		return Application{}, err
	}
	return Application{accounts: accounts, router: services.NewRouter()}, nil
}

func NullApplication(opts Application) Application {
	var accounts services.Accounts
	if opts.accounts != nil {
		accounts = opts.accounts
	} else {
		accounts = services.NullAccounts()
	}
	var router services.Router
	if opts.router != nil {
		router = opts.router
	} else {
		router = services.NullRouter()
	}
	return Application{
		accounts: accounts,
		router:   router,
	}
}

func (a *Application) RegisterHandlers() {
	a.router.Get("/api/accounts/", a.ListAccounts)
}

func (a *Application) ListAccounts() (any, error) {
	return a.accounts.List("user")
}

func (a *Application) Start() error {
	return a.router.Start(8080)
}

func main() {
	config := values.Config{DriverName: "ramsql", DataSourceName: "moolah"}
	application, err := NewApplication(config)
	if err != nil {
		fmt.Println("Failed to init application: ", err)
		return
	}
	application.RegisterHandlers()
	err = application.Start()
	if err != nil {
		fmt.Println("Exiting due to error: ", err)
	}
}

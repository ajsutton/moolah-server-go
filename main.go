package main

import (
	"fmt"
	"github.com/moolah-server-go/infrastructure/services/db"
	"github.com/moolah-server-go/infrastructure/services/web"
	"github.com/moolah-server-go/values"

	_ "github.com/proullon/ramsql/driver"
)

type Application struct {
	accounts db.Accounts
	router   web.Router
}

func NewApplication(config values.Config) (Application, error) {
	accounts, err := db.NewAccounts(config)
	if err != nil {
		return Application{}, err
	}
	return Application{accounts: accounts, router: web.NewRouter()}, nil
}

func NullApplication() Application {
	return Application{
		accounts: db.NullAccounts(),
		router:   web.NullRouter(),
	}
}

func (a *Application) RegisterHandlers() {
	a.router.Get("/api/accounts/", a.ListAccounts)
}

func (a *Application) ListAccounts(request web.Request) (any, error) {
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

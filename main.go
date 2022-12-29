package main

import (
	"fmt"
	"github.com/moolah-server-go/infrastructure/services"
)

type Application struct {
	accounts services.Accounts
	router   services.Router
}

func NewApplication() Application {
	return Application{accounts: services.NewAccounts(), router: services.NewRouter()}
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
	return a.accounts.List(), nil
}

func (a *Application) Start() error {
	return a.router.Start(8080)
}

func main() {
	application := NewApplication()
	application.RegisterHandlers()
	err := application.Start()
	if err != nil {
		fmt.Println("Exiting due to error: ", err)
	}
}

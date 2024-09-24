package app

import (
	"go-learning-project/config"
	"go-learning-project/web"
	"go-learning-project/web/utils"
	"sync"
)

type Application struct {
	wg sync.WaitGroup
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Init() {
	config.LoadConfig()
	utils.InitValidator()
}

func (app *Application) Run() {
	web.StartServer(&app.wg)

}

func (app *Application) Wait() {
	app.wg.Wait()
}

func (app *Application) Cleanup() {
	// close db
}

package main

import "go-learning-project/app"

func main() {

	application := app.NewApplication()
	application.Init()
	application.Run()
	application.Wait()
	application.Cleanup()
}

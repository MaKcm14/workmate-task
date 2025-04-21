package main

import "github.com/MaKcm14/workmate-task/internal/app"

func main() {
	service := app.NewService()
	service.Run()
}

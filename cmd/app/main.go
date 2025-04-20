package main

import "github.com/MaKcm14/workmate-task/internal/app"

func main() {
	storage := app.NewService()
	storage.Run()
}

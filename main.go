package main

import (
	config "backend/config"
	models "backend/models"
	routes "backend/routes"
	translate "backend/services/translate"
	"fmt"
	"runtime"
)

func Init() {
	config.InitConfig()
	models.ConnectDb()
	translate.InitDictionary()
	routes.SetupRouter()
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
func main() {
	Init()
	printMemUsage()
	routes.Router.Run(":8080")
	//defer models.DB.Close()
}

package main

import (
	"cmp"
	"fmt"
	"os"
)

const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

func main() {

	// fmt.Printf("%sHello This is Red%s\n", Red, Reset)
	// fmt.Printf("%sHello This is Green%s\n", Green, Reset)
	// fmt.Println(Bold + Yellow + "Hello this is Yellow" + Reset)

	// price := 250.45555
	// 250.46
	// strFormatedPrice := fmt.Sprintf("%.2f", price)
	// fmt.Println(strFormatedPrice)

	// slog.SetLogLoggerLevel(slog.LevelDebug)
	// logger := slog.New(
	// 	slog.NewJSONHandler(os.Stdout, nil),
	// )
	// logger.Debug("hello this is debug log")
	// logger.Info("hello this is info log")
	// logger.Warn("hi this is warning log")
	// log.Println("ok")

	port := cmp.Or(getFromEnv(), getPortFromFlag(), "7070")
	fmt.Println("application start on port:", port)
}

func getPortFromFlag() string {
	return "8080"
}

func getFromEnv() string {
	return os.Getenv("PORT")
}

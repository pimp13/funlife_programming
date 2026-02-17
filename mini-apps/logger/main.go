package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	var logger *slog.Logger
	const loggerFile = "app.log"
	handleWriterLog := "FILE"

	file, err := os.OpenFile(
		loggerFile,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		handleWriterLog = "STDOUT"
		log.Fatal(err)
	}
	defer file.Close()

	if handleWriterLog == "FILE" {
		logger = slog.New(
			slog.NewTextHandler(file, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)
	} else {
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)
	}

	logger.Info("Hi this is my custom logger save log in to file!")
}

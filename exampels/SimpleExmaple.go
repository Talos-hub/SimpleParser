package main

import (
	"Learn/pkg/parser"
	"log/slog"
	"os"
)

func SetUpLogger() *slog.Logger {
	file, err := os.OpenFile("bot.slog", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}

	return slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}

func main() {
	logger := SetUpLogger()
	p := parser.NewParser(logger, "Info/Data.json", `USD.*?([0-9,]+\.\d{4})`)
	p.StartData("https://www.cbr.ru/eng/currency_base/daily/")

}

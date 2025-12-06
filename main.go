package main

import (
	"os"
	"log/slog"

	"sourcery.dny.nu/pana"
	"github.com/alexflint/go-arg"
	"github.com/cato-001/twink/tui"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	var args struct{
		Auth *struct{
		} `arg:"subcommand:auth"`

		Tui *tui.Args `arg:"subcommand:tui"`
	}

	parser, err := arg.NewParser(arg.Config{}, &args)
	if err != nil {
		logger.Error("could not create cli argument parser", slog.Any("error", err))
		return
	}

	parser.MustParse(os.Args[1:])

	processor := pana.NewProcessor(logger)
	_ = processor

	if args.Tui != nil {
		err := tui.Start(*args.Tui)
		if err != nil {
			logger.Error("app exited unexpectedly", slog.Any("error", err))
		}
		return
	}
}

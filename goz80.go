package main

import (
	"fmt"
	"github.com/inkyblackness/imgui-go/v2"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"

	"github.com/XANi/goz80/platform"
	"github.com/XANi/goz80/renderer"
	"github.com/XANi/goz80/ui"
)

var version string
var log *zap.SugaredLogger
var debug = true
var exit = make(chan bool, 1)

func init() {
	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	// naive systemd detection. Drop timestamp if running under it
	if os.Getenv("INVOCATION_ID") != "" || os.Getenv("JOURNAL_STREAM") != "" {
		consoleEncoderConfig.TimeKey = ""
	}
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)
	consoleStderr := zapcore.Lock(os.Stderr)
	_ = consoleStderr

	// if needed point differnt priority log to different place
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, os.Stderr, lowPriority),
		zapcore.NewCore(consoleEncoder, os.Stderr, highPriority),
	)
	logger := zap.New(core)
	if debug {
		logger = logger.WithOptions(
			zap.Development(),
			zap.AddCaller(),
			zap.AddStacktrace(highPriority),
		)
	} else {
		logger = logger.WithOptions(
			zap.AddCaller(),
		)
	}
	log = logger.Sugar()

}

func main() {
	app := cli.NewApp()
	app.Name = "Artemix"
	app.Description = "MIX VM emulator"
	app.Version = version
	app.HideHelp = true
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "help, h", Usage: "show help"},
		cli.StringFlag{
			Name:   "url",
			Value:  "http://127.0.0.1",
			Usage:  "It's an url",
			EnvVar: "_URL",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.Bool("help") {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}

		log.Infof("Starting %s version: %s", app.Name, version)

		context := imgui.CreateContext(nil)
		defer context.Destroy()
		io := imgui.CurrentIO()

		platform, err := platform.NewSDL(io, platform.SDLClientAPIOpenGL3)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(-1)
		}
		defer platform.Dispose()

		renderer, err := renderer.NewOpenGL3(io)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(-1)
		}
		defer renderer.Dispose()

		ui.Run(platform, renderer)

		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:    "rem",
			Aliases: []string{"a"},
			Usage:   "example cmd",
			Action: func(c *cli.Context) error {
				log.Warn("running example cmd")
				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "example cmd",
			Action: func(c *cli.Context) error {
				log.Warn("running example cmd")
				return nil
			},
		},
	}
	// to sort do that
	//sort.Sort(cli.FlagsByName(app.Flags))
	//sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}

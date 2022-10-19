package app

import (
	"context"
	"fmt"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"net/http"
	"os"
)

type App struct {
	name       string
	host       string
	port       uint64
	log        string
	logger     *logrus.Logger
	httpServer *http.Server
}

func (app *App) buildTable(cf config.Http, pid int) *tablewriter.Table {
	addr := fmt.Sprintf("%s:%d", cf.Host, cf.Port)
	data := [][]string{
		{"1", "Application", ""},
		{"2", "Address", addr},
		{"3", "Port", cast.ToString(cf.Port)},
		{"4", "Pid", fmt.Sprintf("%d", pid)},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Info", "Desc"})
	table.AppendBulk(data)
	table.SetAlignment(tablewriter.ALIGN_LEFT) // Set Alignment
	return table
}

func (app *App) Start(ctx context.Context, cf config.Http) error {
	table := app.buildTable(cf, os.Getpid())
	table.Render()
	go func() {
		for {
			select {
			case <-ctx.Done():
				_ = app.httpServer.Shutdown(ctx)
			}
		}
	}()
	if err := app.httpServer.ListenAndServe(); err != nil {
		return errors.Wrap(err, "http server start error")
	}
	return nil

}

type Option func(app *App) error

func WithName(name string) Option {
	return func(app *App) error {
		app.name = name
		return nil
	}
}

func WithHost(host string) Option {
	return func(app *App) error {
		app.host = host
		return nil
	}
}

func WithPort(port uint64) Option {
	return func(app *App) error {
		app.port = port
		return nil
	}
}
func withLogName(log string) Option {
	return func(app *App) error {
		app.log = log
		return nil
	}
}

func WithLogger(logger *logrus.Logger) Option {
	return func(app *App) error {
		app.logger = logger
		return nil
	}
}

func WithHttpServer(httpServer *http.Server) Option {
	return func(app *App) error {
		app.httpServer = httpServer
		return nil
	}
}

func NewApp(opts ...Option) (*App, error) {
	instance := &App{}
	for _, opt := range opts {
		err := opt(instance)
		if err != nil {
			return nil, err
		}
	}
	return instance, nil
}

package grpc

import (
	"github.com/sirupsen/logrus"
)

type Application struct {
	name    string
	host    string
	port    uint64
	logName string
	logger  *logrus.Logger
}

type Option func(app *Application) error

func WithName(name string) Option {
	return func(app *Application) error {
		app.name = name
		return nil
	}
}

func WithHost(host string) Option {
	return func(app *Application) error {
		app.host = host
		return nil
	}
}

func WithPort(port uint64) Option {
	return func(app *Application) error {
		app.port = port
		return nil
	}
}

func WithLogName(logName string) Option {
	return func(app *Application) error {
		app.logName = logName
		return nil
	}
}

func WithLogger(logger *logrus.Logger) Option {
	return func(app *Application) error {
		app.logger = logger
		return nil
	}
}

func NewApplication(opts ...Option)(*Application,error)  {
	app:=&Application{}
	for _,opt:=range opts{
		err := opt(app)
		if err != nil {
			return nil, err
		}
	}
	return app,nil
}
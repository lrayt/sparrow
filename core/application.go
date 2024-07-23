package core

import (
	"fmt"
	"github.com/lrayt/sparrow/core/abstract"
	"github.com/lrayt/sparrow/core/runtime"
)

type Application struct {
	Env            *runtime.Env
	LicenseChecker abstract.LicenseChecker // 证书校验器
	LoggerProvider abstract.LoggerProvider
	ConfigProvider abstract.ConfigProvider
	Starters       []abstract.Starter
	Handlers       []abstract.Handler
}

func (app Application) Print() {
	fmt.Printf("AppName: %s\n", app.Env.AppName)
	fmt.Printf("RunEnv: %s\n", app.Env.RunEnv)
	fmt.Printf("Version: %s\n", app.Env.BuildVersion)
	fmt.Printf("WorkDir: %s\n", app.Env.WorkDir)
	fmt.Printf("VerifyLicense: %v\n", app.Env.VerifyLicense)
}

type Option func(app *Application)

func WithLicenseChecker(checker abstract.LicenseChecker) Option {
	return func(app *Application) {
		app.LicenseChecker = checker
	}
}

func WithWorkerDir(dir string) Option {
	return func(app *Application) {
		app.Env.WorkDir = dir
	}
}

func WithLogger(provider abstract.LoggerProvider) Option {
	return func(app *Application) {
		app.LoggerProvider = provider
	}
}

func WithConfigurator(provider abstract.ConfigProvider) Option {
	return func(app *Application) {
		app.ConfigProvider = provider
	}
}

func WithStarter(starters ...abstract.Starter) Option {
	return func(app *Application) {
		app.Starters = starters
	}
}

func WithHandler(handlers ...abstract.Handler) Option {
	return func(app *Application) {
		app.Handlers = handlers
	}
}

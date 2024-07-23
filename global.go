package sparrow

import (
	"github.com/lrayt/sparrow/core"
	"github.com/lrayt/sparrow/core/abstract"
	"github.com/lrayt/sparrow/core/kit"
	"github.com/lrayt/sparrow/core/runtime"
	"log"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

var (
	AppName       = "sparrow-app"
	Version       = "0.0.1"
	VerifyLicense = "false"
	app           = new(core.Application)
)

func InitApp(options ...core.Option) error {
	app.Env = runtime.NewEnv(AppName, Version, VerifyLicense)
	for _, option := range options {
		option(app)
	}
	// default workdir
	if len(app.Env.WorkDir) <= 0 {
		if err := app.Env.SetDefaultWorkDir(); err != nil {
			return err
		}
	}
	// print
	app.Print()
	// license verify
	if app.Env.VerifyLicense {
		if app.LicenseChecker == nil {
			core.WithLicenseChecker(kit.NewLicenseChecker())(app)
		}
		if err := app.LicenseChecker.Verify(); err != nil {
			return err
		}
	}
	// default configurator
	if app.ConfigProvider == nil {
		if provider, err := kit.NewYamlConfigProvider(app.Env); err != nil {
			return err
		} else {
			core.WithConfigurator(provider)(app)
		}
	}
	// default logger
	if app.LoggerProvider == nil {
		if provider, err := kit.NewLocalFileLogProvider(app.Env); err != nil {
			return err
		} else {
			core.WithLogger(provider)(app)
		}
	}
	return nil
}

func SetupApp() {
	var (
		errChan    = make(chan error, 1)
		signalChan = make(chan os.Signal, 1)
	)
	for _, starter := range app.Starters {
		if err := starter.Init(); err != nil {
			log.Fatalf("启动失败，err:%s\n", err.Error())
		} else {
			log.Printf("%s初始化成功\n", reflect.TypeOf(starter).String())
		}
	}
	for _, provider := range app.Handlers {
		if provider == nil {
			continue
		}
		go func(fn abstract.Handler) {
			if err := fn.Run(); err != nil {
				errChan <- err
			}
		}(provider)
	}

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-errChan:
		log.Fatalf("服务启动异常，err:%v", err)
	case <-signalChan:
		// shutdown handler
		for _, handler := range app.Handlers {
			if err := handler.Shutdown(); err != nil {
				log.Printf("shutdown handler:%s\n", err.Error())
			}
		}
		// close starter
		for _, starter := range app.Starters {
			if err := starter.Close(); err != nil {
				log.Printf("%s close err: %s\n", reflect.TypeOf(starter).String(), err.Error())
			} else {
				log.Printf("%s closed\n", reflect.TypeOf(starter).String())
			}
		}
	}
}

// GConfigs 全局配置
func GConfigs() abstract.ConfigProvider {
	return app.ConfigProvider
}

// GLoggerProvider 全局日志
func GLoggerProvider() abstract.LoggerProvider {
	return app.LoggerProvider
}

// GRunEnv 运行环境
func GRunEnv() runtime.RunEnv {
	return app.Env.RunEnv
}

func GBuildVersion() string {
	return app.Env.BuildVersion
}

func IsProdEnv() bool {
	return app.Env.RunEnv == runtime.RunProdEnv
}

func IsTestEnv() bool {
	return app.Env.RunEnv == runtime.RunTestEnv
}

func IsLocalEnv() bool {
	return app.Env.RunEnv == runtime.RunLocalEnv
}

func GWorkDir() {

}

func GResourceDir() {

}

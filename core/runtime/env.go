package runtime

import (
	"os"
)

type RunEnv string

const (
	RunLocalEnv = "local"
	RunTestEnv  = "test"
	RunProdEnv  = "prod"
)

func NewRunEnv(env string) RunEnv {
	if env == "prod" {
		return RunProdEnv
	} else if env == "test" {
		return RunTestEnv
	} else {
		return RunLocalEnv
	}
}

// Env 全局环境变量
type Env struct {
	AppName       string
	RunEnv        RunEnv
	WorkDir       string
	BuildVersion  string
	VerifyLicense bool
}

// SetDefaultWorkDir 设置默认地址
func (e *Env) SetDefaultWorkDir() error {
	if dir, err := os.Getwd(); err != nil {
		return err
	} else {
		e.WorkDir = dir
		return nil
	}
}

func NewEnv(appName, version, verifyLicense string) *Env {
	return &Env{
		AppName:       appName,
		RunEnv:        NewRunEnv(os.Getenv(appName)),
		VerifyLicense: verifyLicense == "true",
		BuildVersion:  version,
	}
}

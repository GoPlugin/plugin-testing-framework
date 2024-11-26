package main

import (
	"github.com/goplugin/plugin-testing-framework/lib/k8s/environment"
	"github.com/goplugin/plugin-testing-framework/lib/k8s/pkg/helm/plugin"
	"github.com/goplugin/plugin-testing-framework/lib/k8s/pkg/helm/ethereum"
)

func main() {
	e := environment.New(nil).
		AddHelm(ethereum.New(nil)).
		AddHelm(plugin.New(0, nil))
	if err := e.Run(); err != nil {
		panic(err)
	}
	if err := e.DumpLogs("logs/mytest"); err != nil {
		panic(err)
	}
}

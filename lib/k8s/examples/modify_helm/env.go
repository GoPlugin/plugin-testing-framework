package main

import (
	"fmt"

	"github.com/goplugin/plugin-testing-framework/lib/k8s/environment"
	"github.com/goplugin/plugin-testing-framework/lib/k8s/pkg/helm/plugin"
	"github.com/goplugin/plugin-testing-framework/lib/k8s/pkg/helm/ethereum"
	"github.com/goplugin/plugin-testing-framework/lib/k8s/pkg/helm/mockserver"
	mockservercfg "github.com/goplugin/plugin-testing-framework/lib/k8s/pkg/helm/mockserver-cfg"
)

func main() {
	e := environment.New(&environment.Config{
		NamespacePrefix: "modified-env",
		Labels:          []string{fmt.Sprintf("envType=Modified")},
	}).
		AddHelm(mockservercfg.New(nil)).
		AddHelm(mockserver.New(nil)).
		AddHelm(ethereum.New(nil)).
		AddHelm(plugin.New(0, map[string]any{
			"replicas": 1,
		}))
	err := e.Run()
	if err != nil {
		panic(err)
	}
	e.Cfg.KeepConnection = true
	e.Cfg.RemoveOnInterrupt = true
	e, err = e.
		ReplaceHelm("plugin-0", plugin.New(0, map[string]any{
			"replicas": 2,
		}))
	if err != nil {
		panic(err)
	}
	err = e.Run()
	if err != nil {
		panic(err)
	}
}

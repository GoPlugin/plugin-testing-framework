package main

import (
	"github.com/goplugin/plugin-testing-framework/lib/k8s/environment"
	"github.com/goplugin/plugin-testing-framework/lib/k8s/pkg/helm/plugin"
	"github.com/goplugin/plugin-testing-framework/lib/k8s/pkg/helm/ethereum"
	"github.com/goplugin/plugin-testing-framework/lib/k8s/pkg/helm/mockserver"
	mockservercfg "github.com/goplugin/plugin-testing-framework/lib/k8s/pkg/helm/mockserver-cfg"
)

func main() {
	err := environment.New(&environment.Config{
		NamespacePrefix:   "ztest",
		KeepConnection:    true,
		RemoveOnInterrupt: true,
	}).
		AddHelm(mockservercfg.New(nil)).
		AddHelm(mockserver.New(nil)).
		AddHelm(ethereum.New(nil)).
		AddHelm(plugin.New(0, map[string]interface{}{
			"replicas": 1,
		})).
		Run()
	if err != nil {
		panic(err)
	}
}

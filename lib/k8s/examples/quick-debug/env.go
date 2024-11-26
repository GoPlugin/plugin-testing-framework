package main

import (
	"github.com/goplugin/plugin-testing-framework/lib/k8s/environment"
	"github.com/goplugin/plugin-testing-framework/lib/k8s/pkg/helm/plugin"
	"github.com/goplugin/plugin-testing-framework/lib/k8s/pkg/helm/ethereum"
)

func main() {
	err := environment.New(&environment.Config{
		Labels:            []string{"type=construction-in-progress"},
		NamespacePrefix:   "new-environment",
		KeepConnection:    true,
		RemoveOnInterrupt: true,
	}).
		AddHelm(ethereum.New(nil)).
		AddHelm(plugin.New(0, nil)).
		Run()
	if err != nil {
		panic(err)
	}
}

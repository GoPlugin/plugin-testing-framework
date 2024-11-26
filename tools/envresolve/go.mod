module github.com/goplugin/plugin-testing-framework/tools/envresolve

go 1.22.5

require (
	//github.com/goplugin/plugin-testing-framework/lib v1.99.4-0.20240903123107-cd7909d3e9fb
	github.com/goplugin/plugin-testing-framework/lib v0.0.1 //plugin latest update
	github.com/spf13/cobra v1.8.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

retract [v1.999.0-test-release, v1.999.999-test-release]

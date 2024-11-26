package concurrent_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/goplugin/plugin-testing-framework/lib/k8s/environment"
	"github.com/goplugin/plugin-testing-framework/lib/k8s/pkg/helm/plugin"
)

func TestConcurrentEnvs(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		t.Parallel()
		e := environment.New(nil).
			AddHelm(plugin.New(0, nil))
		defer e.Shutdown()
		err := e.Run()
		require.NoError(t, err)
	})
	t.Run("test 2", func(t *testing.T) {
		t.Parallel()
		e := environment.New(nil).
			AddHelm(plugin.New(0, nil))
		defer e.Shutdown()
		err := e.Run()
		require.NoError(t, err)
		e, err = e.
			ReplaceHelm("plugin-0", plugin.New(0, map[string]any{
				"replicas": 2,
			}))
		require.NoError(t, err)
		err = e.Run()
		require.NoError(t, err)
	})
}

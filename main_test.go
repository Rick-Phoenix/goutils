package u_test

import (
	"testing"

	"github.com/Rick-Phoenix/goutils/scaffolder"
	"github.com/labstack/gommon/log"
)

func TestMain(t *testing.T) {
	tmplData := map[string]any{
		"PackageName": "svelteproj",
		"IsWails":     true,
	}
	err := scaffolder.ScaffoldSvelte("svelte", tmplData)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

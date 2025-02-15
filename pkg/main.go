package main

import (
	"os"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/tracing"

	"github.com/cloudrhinoltd/infinity-plus-datasource/pkg/pluginhost"
)

const pluginID = "infinity-plus-datasource"

func main() {
	dsOptions := datasource.ManageOpts{
		TracingOpts: tracing.Opts{},
	}
	if err := datasource.Manage(pluginID, pluginhost.NewDataSourceInstance, dsOptions); err != nil {
		backend.Logger.Error("error starting infinity plugin", "error", err.Error())
		os.Exit(1)
	}
}

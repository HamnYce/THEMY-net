package apiv1routes

import (
	"net/http"
	handlers "themynet/api/v1/handlers"
	"themynet/internal/debug"
)

func SetupRoutes() {
	// attach createHosts handler
	{
		debug.DebugPrintf("attaching createHost Handler\n")
		http.HandleFunc("/CreateHosts", handlers.CreateHostsHandler)
		debug.DebugPrintf("attached createHost Handler\n")
	}

	// attach RetrieveHosts handler
	{
		debug.DebugPrintf("attaching RetrieveHosts Handler\n")
		http.HandleFunc("/RetrieveHosts", handlers.RetrieveHostsHandler)
		debug.DebugPrintf("attached RetrieveHosts Handler\n")
	}

	// attach UpdateHosts handler
	{
		debug.DebugPrintf("attaching UpdateHost Handler\n")
		http.HandleFunc("/UpdateHosts", handlers.UpdateHostsHandler)
		debug.DebugPrintf("attached UpdateHost Handler\n")
	}

	// attach DeleteHosts handler
	{
		debug.DebugPrintf("attaching DeleteHost Handler\n")
		http.HandleFunc("/DeleteHosts", handlers.DeleteHostsHandler)
		debug.DebugPrintf("attached DeleteHost Handler\n")
	}
}

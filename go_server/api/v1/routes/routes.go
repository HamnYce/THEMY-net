package apiv1routes

import (
	"net/http"
	utils "themynet"
	handlers "themynet/api/v1/handlers"
	middleware "themynet/api/v1/middleware"
)

func SetupRoutes() {
	{
		utils.DebugPrintf("attaching createHost Handler\n")
		http.Handle("/CreateHosts", middleware.WrapperMiddleWare(http.HandlerFunc(handlers.CreateHostsHandler)))
		utils.DebugPrintf("attached createHost Handler\n")
	}

	{
		utils.DebugPrintf("attaching RetrieveHosts Handler\n")
		http.Handle("/RetrieveHosts", middleware.WrapperMiddleWare(http.HandlerFunc(handlers.RetrieveHostsHandler)))
		utils.DebugPrintf("attached RetrieveHosts Handler\n")
	}

	{
		utils.DebugPrintf("attaching UpdateHost Handler\n")
		http.Handle("/UpdateHosts", middleware.WrapperMiddleWare(http.HandlerFunc(handlers.UpdateHostsHandler)))
		utils.DebugPrintf("attached UpdateHost Handler\n")
	}

	{
		utils.DebugPrintf("attaching DeleteHost Handler\n")
		http.Handle("/DeleteHosts", middleware.WrapperMiddleWare(http.HandlerFunc(handlers.DeleteHostsHandler)))
		utils.DebugPrintf("attached DeleteHost Handler\n")
	}
}

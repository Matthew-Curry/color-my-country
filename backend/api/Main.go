package main

/********************************************************************************
* Description: The controller handles important user functions, and uses the    *
* dao package to actually connect to the database. Server is also intialized    *
*********************************************************************************/
import (
	serverinit "Go-directory/Server_Init"
)

func main() {
	// start http server to serve API requests
	serverinit.InitializeServer()
}

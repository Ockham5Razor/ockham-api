package main

import "ockham-api/cmd"

// @title Ockham API
// @version 1.0
// @description All APIs of Ockham Project
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @BasePath /api
func main() {
	cmd.Execute()
}

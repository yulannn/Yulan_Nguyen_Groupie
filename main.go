package main

import (
	"royaleapi/routeur"
	"royaleapi/templates"
)

func main() {
	templates.InitTemplate()
	routeur.Initserv()
}

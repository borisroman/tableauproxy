package atlassian

import (
	"sync"

	"github.com/borisroman/tableauproxy/pkg/azure"
	"github.com/borisroman/tableauproxy/pkg/tableau"
)

type Controller struct {
	LocalBaseUrl       string
	AzureSQLController *azure.Controller
	TableauController  *tableau.Controller
	mux                sync.Mutex
}

func GetController(domain string, azureSQLController *azure.Controller, tableauController *tableau.Controller) *Controller {
	return &Controller{
		LocalBaseUrl:       domain,
		AzureSQLController: azureSQLController,
		TableauController:  tableauController,
	}
}

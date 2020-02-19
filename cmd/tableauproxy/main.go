package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/borisroman/tableauproxy/pkg/atlassian"
	"github.com/borisroman/tableauproxy/pkg/azure"
	"github.com/borisroman/tableauproxy/pkg/tableau"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Controller struct {
	AtlassianAPI *atlassian.Controller
	AzureSQLAPI  *azure.Controller
}

func (c *Config) initController() *Controller {
	dbClient, err := azure.NewClient(c.AzureSQLServer, c.AzureSQLPort, c.AzureSQLUser, c.AzureSQLPassword, c.AzureSQLDatabase)
	if err != nil {
		log.Fatal(err)
	}

	tableauController := tableau.GetController(c.Domain)

	return &Controller{
		AtlassianAPI: atlassian.GetController(c.Domain, dbClient, tableauController),
		AzureSQLAPI:  dbClient,
	}
}

func (c *Controller) initHandlers() {
	rtr := mux.NewRouter()

	// Serve the App Descriptor - https://developer.atlassian.com/cloud/confluence/app-descriptor/
	rtr.HandleFunc("/atlassian-connect.json", c.AtlassianAPI.HandleAtlassianConnect)

	// Handle Install and Uninstall callbacks - https://developer.atlassian.com/cloud/confluence/app-descriptor/#lifecycle
	rtr.HandleFunc("/plugin-installed-callback", c.AtlassianAPI.HandleLifeCycleInstalled)
	rtr.HandleFunc("/plugin-uninstalled-callback", c.AtlassianAPI.HandleLifeCycleUninstalled)

	// Handle PersonalAccessToken CRUD
	rtr.HandleFunc("/personal-access-token", c.AtlassianAPI.HandlePersonalAccessToken)

	// Handle Tableau information requests
	rtr.HandleFunc("/tableau/sites", c.AtlassianAPI.HandleTableauSites)
	rtr.HandleFunc("/tableau/views", c.AtlassianAPI.HandleTableauViews)

	// Handle Table image requests
	rtr.HandleFunc("/macro-image.png", c.AtlassianAPI.HandleViewImagePng)    // The rendered image
	rtr.HandleFunc("/ui/macro/static-view", c.AtlassianAPI.HandleStaticView) // The static page

	// Show UI pages
	rtr.PathPrefix("/ui/admin/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./ui/build/index.html")
	})

	rtr.PathPrefix("/ui/macro/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./ui/build/index.html")
	})

	// Serve static UI files
	fs := http.FileServer(http.Dir("./ui/build"))
	rtr.PathPrefix("/ui/").Handler(http.StripPrefix("/ui/", fs))

	// Start routing!
	http.Handle("/", rtr)
}

func (c *Controller) startServer(port int) {
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%v", port),
		handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)),
	)
}

func (c *Config) initTableau() {
	tableau.GetController(c.Domain)
}

func main() {
	cfg := GetConfig()

	controller := cfg.initController()
	controller.initHandlers()
	controller.startServer(cfg.Port)
}

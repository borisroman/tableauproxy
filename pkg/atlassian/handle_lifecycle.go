package atlassian

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/borisroman/tableauproxy/pkg/models"
)

func (c *Controller) HandleAtlassianConnect(w http.ResponseWriter, r *http.Request) {
	fp := filepath.Join("templates", "atlassian-connect.json")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, "Unable to read template files - "+err.Error(), 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "atlassian-connect.json", c)
	if err != nil {
		http.Error(w, "Unable to execute template files - "+err.Error(), 500)
		return
	}
}

func (c *Controller) HandleLifeCycleInstalled(w http.ResponseWriter, r *http.Request) {
	payload, err := c.DecodeLifecyclePayload(r)
	if err != nil {
		http.Error(w, "Unable to decode lifecycle payload - "+err.Error(), 500)
		return
	}

	appInstance, err := c.AzureSQLController.ReadAppInstance(payload)
	if err != nil {
		http.Error(w, "Unable to persist app instance install - "+err.Error(), 500)
		return
	}
	if appInstance != nil {
		_, err = c.AzureSQLController.UpdateAppInstance(payload)
		if err != nil {
			http.Error(w, "Unable to persist app instance install - "+err.Error(), 500)
			return
		}
	} else {
		_, err = c.AzureSQLController.CreateAppInstance(payload)
		if err != nil {
			http.Error(w, "Unable to persist app instance install - "+err.Error(), 500)
			return
		}
	}

	w.WriteHeader(200)
}

func (c *Controller) HandleLifeCycleUninstalled(w http.ResponseWriter, r *http.Request) {
	// TODO Check if we need to handle JWT auth here

	payload, err := c.DecodeLifecyclePayload(r)
	if err != nil {
		http.Error(w, "Unable to decode lifecycle payload - "+err.Error(), 500)
		return
	}

	rows, err := c.AzureSQLController.DeleteAppInstance(payload)
	if err != nil || rows != 1 {
		http.Error(w, "Unable to persist app instance uninstall, please contact support", 500)
		return
	}

	w.WriteHeader(200)
}

func (c *Controller) DecodeLifecyclePayload(r *http.Request) (*models.LifecyclePayload, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var payload models.LifecyclePayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

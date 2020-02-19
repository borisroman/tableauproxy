package atlassian

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/borisroman/tableauproxy/pkg/models"
)

func (c *Controller) HandlePersonalAccessToken(w http.ResponseWriter, r *http.Request) {
	addDefaultPersonalAccessTokenHeaders(w, r)

	clientKey, err := c.Auth(r)
	if err != nil {
		http.Error(w, "error authenticating the request - "+err.Error(), 403)
		return
	}

	switch r.Method {
	// Create
	case http.MethodPost:
		c.HandleCreatePersonalAccessToken(w, r, clientKey)
		return

	// Read
	case http.MethodGet:
		c.HandleReadPersonalAccessToken(w, r, clientKey)
		return

	// Update
	case http.MethodPatch:
		c.HandleUpdatePersonalAccessToken(w, r, clientKey)
		return

	// Delete
	case http.MethodDelete:
		c.HandleDeletePersonalAccessToken(w, r, clientKey)
		return

	case http.MethodOptions:
		c.HandleOptionsPersonalAccessToken(w, r)
		return

	default:
		http.Error(w, "Only POST, GET, PATCH, DELETE supported", 500)
		return
	}
}

func (c *Controller) HandleCreatePersonalAccessToken(w http.ResponseWriter, r *http.Request, clientKey string) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read body - "+err.Error(), 500)
		return
	}

	var payload models.PersonalAccessToken
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, "unable to unmarshal body - "+err.Error(), 500)
		return
	}

	payload.ClientKey = clientKey

	// Check if it already exists, if so error
	personalAccessToken, err := c.AzureSQLController.ReadPersonalAccessTokenByName(payload.Name, clientKey)
	if err != nil {
		http.Error(w, "unable to check if token already exists - "+err.Error(), 500)
		return
	}
	if personalAccessToken != nil {
		http.Error(w, "token already exists use http method PATCH to update", 500)
		return
	}

	_, err = c.AzureSQLController.CreatePersonalAccessToken(&payload)
	if err != nil {
		http.Error(w, "unable to persist personal access token- "+err.Error(), 500)
		return
	}

	personalAccessToken, err = c.AzureSQLController.ReadPersonalAccessTokenByName(payload.Name, clientKey)
	if err != nil {
		http.Error(w, "unable to check if token already exists - "+err.Error(), 500)
		return
	}

	// Obfuscate response
	personalAccessToken.Secret = ""
	personalAccessToken.ClientKey = ""

	responseBytes, err := json.Marshal(personalAccessToken)
	if err != nil {
		http.Error(w, "unable to marshal response - "+err.Error(), 500)
		return
	}

	w.WriteHeader(200)
	_, _ = w.Write(responseBytes)
}

func (c *Controller) HandleReadPersonalAccessToken(w http.ResponseWriter, r *http.Request, clientKey string) {
	personalAccessTokens, err := c.AzureSQLController.ReadPersonalAccessToken(clientKey)
	if err != nil {
		http.Error(w, "unable to retrieve personal access tokens"+err.Error(), 500)
		return
	}

	// We never return the actual personalAccessTokenSecret in the API response
	for _, personalAccessToken := range personalAccessTokens {
		personalAccessToken.Secret = ""
		personalAccessToken.ClientKey = ""
	}

	response := &models.PersonalAccessTokenResponse{
		PersonalAccessTokens: personalAccessTokens,
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "unable to marshal response - "+err.Error(), 500)
		return
	}

	w.WriteHeader(200)
	_, _ = w.Write(responseBytes)
}

func (c *Controller) HandleUpdatePersonalAccessToken(w http.ResponseWriter, r *http.Request, clientKey string) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read body - "+err.Error(), 500)
		return
	}

	var payload models.PersonalAccessToken
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, "unable to unmarshal body - "+err.Error(), 500)
		return
	}

	payload.ClientKey = clientKey

	if payload.UUID == "" {
		http.Error(w, "uuid required PATCH call", 500)
		return
	}

	// Check if it already exists, if not error
	personalAccessToken, err := c.AzureSQLController.ReadPersonalAccessTokenByUUID(payload.UUID, clientKey)
	if err != nil {
		http.Error(w, "unable to check if token already exists - "+err.Error(), 500)
		return
	}
	if personalAccessToken == nil {
		http.Error(w, "token doesn't exists yet use http method POST to create", 500)
		return
	}

	_, err = c.AzureSQLController.UpdateAPersonalAccessToken(&payload)
	if err != nil {
		http.Error(w, "unable to persist access token - "+err.Error(), 500)
		return
	}

	personalAccessToken, err = c.AzureSQLController.ReadPersonalAccessTokenByName(payload.Name, clientKey)
	if err != nil {
		http.Error(w, "unable to check if token already exists - "+err.Error(), 500)
		return
	}

	// Obfuscate response
	personalAccessToken.Secret = ""
	personalAccessToken.ClientKey = ""

	responseBytes, err := json.Marshal(personalAccessToken)
	if err != nil {
		http.Error(w, "unable to marshal response - "+err.Error(), 500)
		return
	}

	w.WriteHeader(200)
	_, _ = w.Write(responseBytes)
}

func (c *Controller) HandleDeletePersonalAccessToken(w http.ResponseWriter, r *http.Request, clientKey string) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read body - "+err.Error(), 500)
		return
	}

	var payload models.PersonalAccessToken
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, "unable to unmarshal body - "+err.Error(), 500)
		return
	}

	payload.ClientKey = clientKey

	// Check if it already exists, if not error
	personalAccessToken, err := c.AzureSQLController.ReadPersonalAccessTokenByUUID(payload.UUID, clientKey)
	if err != nil {
		http.Error(w, "unable to check if token already exists - "+err.Error(), 500)
		return
	}
	if personalAccessToken == nil {
		http.Error(w, "token doesn't exists, can't delete", 500)
		return
	}

	_, err = c.AzureSQLController.DeletePersonalAccessToken(&payload)
	if err != nil {
		http.Error(w, "unable to persist delete access token - "+err.Error(), 500)
		return
	}

	w.WriteHeader(200)
}

func (c *Controller) HandleOptionsPersonalAccessToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

// Helper functions
func addDefaultPersonalAccessTokenHeaders(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type")
}

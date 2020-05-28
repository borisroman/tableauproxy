package atlassian

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/borisroman/tableauproxy/pkg/models"
)

func (c *Controller) HandleTableauSites(w http.ResponseWriter, r *http.Request) {
	addDefaultTableauInformationHeaders(w, r)

	clientKey, err := c.Auth(r)
	if err != nil {
		http.Error(w, "error authenticating the request - "+err.Error(), 403)
		return
	}

	switch r.Method {
	// Read
	case http.MethodGet:
		c.HandleTableauSitesGet(w, r, clientKey)
		return

	case http.MethodOptions:
		c.HandleOptionsTableauInformation(w, r)
		return

	default:
		http.Error(w, "Only GET and OPTIONS methods supported", 500)
		return
	}
}

func (c *Controller) HandleTableauSitesGet(w http.ResponseWriter, r *http.Request, clientKey string) {
	personalAccessTokenUUID, err := parsePersonalAccessToken(r)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	token, err := c.AzureSQLController.ReadPersonalAccessTokenByUUID(personalAccessTokenUUID, clientKey)
	if err != nil {
		http.Error(w, "unable to retrieve access token - "+err.Error(), 500)
		return
	}

	c.mux.Lock()
	sites, err := c.TableauController.GetSites(token)
	if err != nil {
		http.Error(w, "unable to retrieve sites from tableau - "+err.Error(), 500)
		c.mux.Unlock()
		return
	}
	c.mux.Unlock()

	var responseSites []*models.Site
	for _, site := range sites {
		responseSites = append(responseSites, &models.Site{
			Id:         site.Id,
			Name:       site.Name,
			ContentUrl: site.ContentUrl,
		})
	}

	response := &models.SitesResponse{
		Sites: responseSites,
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "unable to marshal response - "+err.Error(), 500)
		return
	}

	w.WriteHeader(200)
	_, _ = w.Write(responseBytes)
}

func (c *Controller) HandleTableauViews(w http.ResponseWriter, r *http.Request) {
	addDefaultTableauInformationHeaders(w, r)

	clientKey, err := c.Auth(r)
	if err != nil {
		http.Error(w, "error authenticating the request - "+err.Error(), 403)
		return
	}

	switch r.Method {
	// Read
	case http.MethodGet:
		c.HandleTableauViewsGet(w, r, clientKey)
		return

	case http.MethodOptions:
		c.HandleOptionsTableauInformation(w, r)
		return

	default:
		http.Error(w, "Only GET and OPTIONS methods supported", 500)
		return
	}
}

func (c *Controller) HandleTableauViewsGet(w http.ResponseWriter, r *http.Request, clientKey string) {
	personalAccessTokenUUID, err := parsePersonalAccessToken(r)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	siteId, err := parseSiteId(r)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	token, err := c.AzureSQLController.ReadPersonalAccessTokenByUUID(personalAccessTokenUUID, clientKey)
	if err != nil {
		http.Error(w, "unable to retrieve access token - "+err.Error(), 500)
		return
	}

	c.mux.Lock()
	views, err := c.TableauController.GetViews(token, siteId)
	if err != nil {
		http.Error(w, "unable to retrieve views from tableau - "+err.Error(), 500)
		c.mux.Unlock()
		return
	}
	c.mux.Unlock()

	var responseViews []*models.View
	for _, view := range views {
		responseViews = append(responseViews, &models.View{
			Id:         view.Id,
			Name:       view.Name,
			ContentUrl: view.ContentUrl,
		})
	}

	response := &models.ViewsResponse{
		Views: responseViews,
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "unable to marshal response - "+err.Error(), 500)
		return
	}

	w.WriteHeader(200)
	_, _ = w.Write(responseBytes)
}

func (c *Controller) HandleOptionsTableauInformation(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

// Helper functions
func addDefaultTableauInformationHeaders(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type")
}

func parsePersonalAccessToken(r *http.Request) (string, error) {
	personalAccessTokenUUID := r.URL.Query().Get("personalAccessTokenUUID")
	if personalAccessTokenUUID == "" {
		return "", errors.New("missing personalAccessTokenUUID parameter in URL")
	}

	return personalAccessTokenUUID, nil
}

func parseSiteId(r *http.Request) (string, error) {
	siteId := r.URL.Query().Get("siteId")
	if siteId == "" {
		return "", errors.New("missing siteId parameter in URL")
	}

	return siteId, nil
}

package atlassian

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/borisroman/tableauproxy/pkg/models"
)

const staticImageURL = "/macro-image.png"

func (c *Controller) HandleViewImagePng(w http.ResponseWriter, r *http.Request) {
	clientKey, err := c.Auth(r)
	if err != nil {
		http.Error(w, "error authenticating the request - "+err.Error(), 403)
		return
	}

	imageRequest, err := c.parseMacroUrl(r)
	if err != nil {
		http.Error(w, "Missing URL parameter - "+err.Error(), 500)
		return
	}

	token, err := c.AzureSQLController.ReadPersonalAccessTokenByUUID(imageRequest.PersonalAccessToken.UUID, clientKey)
	if err != nil {
		http.Error(w, "unable to retrieve access token - "+err.Error(), 500)
		return
	}
	if token == nil {
		http.Error(w, "personal access token not found for uuid - "+imageRequest.PersonalAccessToken.UUID, 500)
		return
	}

	imageRequest.PersonalAccessToken = token

	c.mux.Lock()
	image, err := c.TableauController.GetImage(imageRequest)
	if err != nil {
		http.Error(w, "Unable to retrieve image form Tableau - "+err.Error(), 500)
		c.mux.Unlock()
		return
	}
	c.mux.Unlock()

	_, err = w.Write(image)
	if err != nil {
		http.Error(w, "Unable to complete http request - "+err.Error(), 500)
		return
	}
}

func (c *Controller) HandleStaticView(w http.ResponseWriter, r *http.Request) {
	clientKey, err := c.Auth(r)
	if err != nil {
		http.Error(w, "error authenticating the request - "+err.Error(), 403)
		return
	}

	imageRequest, err := c.parseMacroUrl(r)
	if err != nil {
		http.Error(w, "Missing URL parameter - "+err.Error(), 500)
		return
	}

	token, err := c.AzureSQLController.ReadPersonalAccessTokenByUUID(imageRequest.PersonalAccessToken.UUID, clientKey)
	if err != nil {
		http.Error(w, "unable to retrieve access token - "+err.Error(), 500)
		return
	}
	if token == nil {
		http.Error(w, "personal access token not found for uuid - "+imageRequest.PersonalAccessToken.UUID, 500)
		return
	}

	imageRequest.PersonalAccessToken = token

	jwtToken, err := c.getJWTToken(r)
	if err != nil {
		http.Error(w, "jwt token not found", 500)
		return
	}

	response := fmt.Sprintf("<html><img src=\"%s%s?jwt=%s&siteId=%s&viewId=%s&personalAccessTokenUUID=%s\" style=\"%s\"/><html>", c.LocalBaseUrl, staticImageURL, jwtToken, imageRequest.SiteId, imageRequest.ViewId, imageRequest.PersonalAccessToken.UUID, imageRequest.ImageStyle)

	_, err = w.Write([]byte(response))
	if err != nil {
		http.Error(w, "Unable to complete http request - "+err.Error(), 500)
		return
	}
}

// Helper functions
func (c *Controller) parseMacroUrl(r *http.Request) (*models.ImageRequest, error) {
	siteId := r.URL.Query().Get("siteId")
	if siteId == "" {
		return nil, errors.New("missing siteId parameter in URL")
	}

	viewId := r.URL.Query().Get("viewId")
	if viewId == "" {
		return nil, errors.New("missing viewId parameter in URL")
	}

	personalAccessTokenUUID := r.URL.Query().Get("personalAccessTokenUUID")
	if personalAccessTokenUUID == "" {
		return nil, errors.New("missing personalAccessTokenUUID parameter in URL")
	}

	imageStyle := r.URL.Query().Get("imageStyle")

	return &models.ImageRequest{
		SiteId:     siteId,
		ViewId:     viewId,
		ImageStyle: imageStyle,
		PersonalAccessToken: &models.PersonalAccessToken{
			UUID: personalAccessTokenUUID,
		},
	}, nil
}

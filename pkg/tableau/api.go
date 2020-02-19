package tableau

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/borisroman/tableauproxy/pkg/models"
)

const (
	ApiPath    = "/api/3.6"
	SignInPath = ApiPath + "/auth/signin"
	SitePath   = ApiPath + "/sites"
)

type Controller struct {
	LocalBaseUrl string
}

func GetController(domain string) *Controller {
	return &Controller{
		LocalBaseUrl: domain,
	}
}

// Tableau Image
func (c *Controller) GetImage(imageRequest *models.ImageRequest) ([]byte, error) {
	siteContentUrl, err := c.getSiteContentUrl(imageRequest.PersonalAccessToken, imageRequest.SiteId)
	if err != nil {
		return nil, errors.New("Unable to get content url to login to Tableau - " + err.Error())
	}

	signInResponse, err := c.SignIn(imageRequest.PersonalAccessToken, &siteContentUrl)
	if err != nil {
		return nil, errors.New("Unable to login to Tableau - " + err.Error())
	}

	image, err := c.Image(imageRequest, signInResponse)
	if err != nil {
		return nil, errors.New("Unable to get image for view in Tableau - " + err.Error())
	}

	return image, nil
}

func (c *Controller) Image(imageRequest *models.ImageRequest, signInResponse *SignInResponse) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, imageRequest.PersonalAccessToken.BaseUrl+SitePath+"/"+signInResponse.Credentials.Site.Id+"/views/"+imageRequest.ViewId+"/image?resolution=high", nil)
	if err != nil {
		return nil, err
	}

	// req.Header.Set("Accept", "image/png")
	req.Header.Set("X-Tableau-Auth", signInResponse.Credentials.Token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got status code %v when trying to get image from Tableau", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

// Tableau Information
func (c *Controller) GetSites(personalAccessToken *models.PersonalAccessToken) ([]*SiteResponse, error) {
	signInResponse, err := c.SignIn(personalAccessToken, nil)
	if err != nil {
		return nil, errors.New("Unable to login to Tableau - " + err.Error())
	}

	getSitesResponse := &GetSitesResponse{}
	err, statusCode := c.doJSONRequest(personalAccessToken.BaseUrl+SitePath, http.MethodGet, signInResponse.Credentials.Token, nil, getSitesResponse)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, fmt.Errorf("got status code %v when trying to get sites from Tableau", statusCode)
	}

	return getSitesResponse.Sites.Site, nil
}

func (c *Controller) GetViews(personalAccessToken *models.PersonalAccessToken, siteId string) ([]*ViewResponse, error) {
	siteContentUrl, err := c.getSiteContentUrl(personalAccessToken, siteId)
	if err != nil {
		return nil, errors.New("Unable to get content url to login to Tableau - " + err.Error())
	}

	signInResponse, err := c.SignIn(personalAccessToken, &siteContentUrl)
	if err != nil {
		return nil, errors.New("Unable to login to Tableau - " + err.Error())
	}

	getViewsResponse := &GetViewsResponse{}
	url := personalAccessToken.BaseUrl + SitePath + "/" + siteId + "/views"
	err, statusCode := c.doJSONRequest(url, http.MethodGet, signInResponse.Credentials.Token, nil, getViewsResponse)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, fmt.Errorf("got status code %v when trying to get sites from Tableau", statusCode)
	}

	return getViewsResponse.Views.View, nil
}

// Tableau Sign In
func (c *Controller) SignIn(personalAccessToken *models.PersonalAccessToken, siteContentUrl *string) (*SignInResponse, error) {
	signin := &SignInRequest{
		Credentials: &CredentialsRequest{
			PersonalAccessTokenName:   personalAccessToken.Name,
			PersonalAccessTokenSecret: personalAccessToken.Secret,
			Site: &SiteRequest{
				ContentUrl: siteContentUrl,
			},
		},
	}

	signinJson, err := json.Marshal(signin)
	if err != nil {
		return nil, err
	}

	signInResponse := &SignInResponse{}
	err, statusCode := c.doJSONRequest(personalAccessToken.BaseUrl+SignInPath, http.MethodPost, "", bytes.NewBuffer(signinJson), signInResponse)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, fmt.Errorf("got status code %v when trying to login to Tableau", statusCode)
	}

	return signInResponse, nil
}

// Helper functions
func (c *Controller) doJSONRequest(url string, method string, authToken string, body io.Reader, response interface{}) (error, int) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err, 0
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if authToken != "" {
		req.Header.Set("X-Tableau-Auth", authToken)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err, 0
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, 0
	}

	err = json.Unmarshal(respBody, response)
	if err != nil {
		return err, 0
	}

	return nil, resp.StatusCode
}

func (c *Controller) getSiteContentUrl(personalAccessToken *models.PersonalAccessToken, siteId string) (string, error) {
	sites, err := c.GetSites(personalAccessToken)
	if err != nil {
		return "", err
	}

	for _, site := range sites {
		if site.Id == siteId {
			return site.ContentUrl, nil
		}
	}

	return "", errors.New("unable to find contentUrl for site")
}

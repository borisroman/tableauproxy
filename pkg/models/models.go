package models

type LifecyclePayload struct {
	Key                      string `json:"key"`
	ClientKey                string `json:"clientKey"`
	AccountID                string `json:"accountId"`
	SharedSecret             string `json:"sharedSecret"`
	BaseUrl                  string `json:"baseUrl"`
	DisplayURL               string `json:"displayUrl"`
	ProductType              string `json:"productType"`
	Description              string `json:"description"`
	ServiceEntitlementNumber string `json:"serviceEntitlementNumber"`
	OauthClientId            string `json:"oauthClientId"`
}

type PersonalAccessTokenResponse struct {
	PersonalAccessTokens []*PersonalAccessToken `json:"personalAccessTokens"`
}

type PersonalAccessToken struct {
	UUID      string `json:"uuid"`
	BaseUrl   string `json:"baseUrl"`
	ClientKey string `json:"clientKey,omitempty"`
	Name      string `json:"name"`
	Secret    string `json:"secret,omitempty"`
}

type SitesResponse struct {
	Sites []*Site `json:"sites"`
}

type Site struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	ContentUrl string `json:"contentUrl"`
}

type ViewsResponse struct {
	Views []*View `json:"views"`
}

type View struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	ContentUrl string `json:"contentUrl"`
}

type ImageRequest struct {
	SiteId              string
	ViewId              string
	PersonalAccessToken *PersonalAccessToken
	ImageStyle          string
}

type MacroTemplate struct {
	ImageBase64 string
}

type MacroStaticTemplate struct {
	ImageUrl string
}

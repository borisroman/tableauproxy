package tableau

/**
 * Requests
 */
type SignInRequest struct {
	Credentials *CredentialsRequest `json:"credentials"`
}

type CredentialsRequest struct {
	PersonalAccessTokenName   string       `json:"personalAccessTokenName"`
	PersonalAccessTokenSecret string       `json:"personalAccessTokenSecret"`
	Site                      *SiteRequest `json:"site"`
}

type SiteRequest struct {
	ContentUrl *string `json:"contentUrl"`
}

/**
 * Responses
 */
type SignInResponse struct {
	Credentials *CredentialsResponse `json:"credentials"`
}

type CredentialsResponse struct {
	Site  *SiteResponse `json:"site"`
	User  *UserResponse `json:"user"`
	Token string        `json:"token"`
}

type SiteResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	ContentUrl string `json:"contentUrl"`
}

type UserResponse struct {
	Id string `json:"id"`
}

type GetSitesResponse struct {
	Pagination *PaginationResponse `json:"pagination"`
	Sites      *GetSiteResponse    `json:"sites"`
}

type GetSiteResponse struct {
	Site []*SiteResponse `json:"site"`
}

type GetViewsResponse struct {
	Pagination *PaginationResponse `json:"pagination"`
	Views      *ViewsResponse      `json:"views"`
}

type GetViewByPathResponse struct {
	Pagination *PaginationResponse `json:"pagination"`
	Views      *ViewsResponse      `json:"views"`
}

type PaginationResponse struct {
	PageNumber     string `json:"pageNumber"`
	PageSize       string `json:"pageSize"`
	TotalAvailable string `json:"totalAvailable"`
}

type ViewsResponse struct {
	View []*ViewResponse `json:"view"`
}

type ViewResponse struct {
	Workbook    *WorkbookResponse `json:"workbook"`
	Owner       *OwnerResponse    `json:"owner"`
	Project     *ProjectResponse  `json:"project"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	ContentUrl  string            `json:"contentUrl"`
	CreatedAt   string            `json:"createdAt"`
	UpdatedAt   string            `json:"updatedAt"`
	ViewUrlName string            `json:"viewUrlName"`
}

type WorkbookResponse struct {
	Id string `json:"id"`
}

type OwnerResponse struct {
	Id string `json:"id"`
}

type ProjectResponse struct {
	Id string `json:"id"`
}

package atlassian

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/cristalhq/jwt"
)

func (c *Controller) Auth(r *http.Request) (string, error) {
	jwtToken, err := c.getJWTToken(r)
	if err != nil {
		return "", err
	}

	// 1. Decode the token
	token, err := jwt.ParseString(jwtToken)
	if err != nil {
		return "", err
	}

	// 2. Unmarshal the standardClaims to get the issuer
	var standardClaims jwt.StandardClaims
	err = json.Unmarshal(token.RawClaims(), &standardClaims)
	if err != nil {
		return "", err
	}

	// 3. The issuer is the clientKey of the Confluence plugin
	appInstance, err := c.AzureSQLController.ReadAppInstanceByClientKey(standardClaims.Issuer)
	if err != nil {
		return "", err
	}

	// 4. Only HS256 is supported
	if token.Header().Algorithm != jwt.HS256 {
		return "", errors.New("jwt algorithm is not supported, only HS256 is supported")
	}

	// 5. Create a signer using the sharedSecret retrieved from the database
	signer, err := jwt.NewHS256([]byte(appInstance.SharedSecret))
	if err != nil {
		return "", err
	}

	// 6. Verify the request using the signer
	err = signer.Verify(token.Payload(), token.Signature())
	if err != nil {
		return "", err
	}

	// 7. Check if the token is expired
	if standardClaims.IsExpired(time.Now()) {
		return "", errors.New("token is expired")
	}

	return standardClaims.Issuer, nil
}

func (c *Controller) getJWTToken(r *http.Request) (string, error) {
	// Get JWT token from Query parameter
	jwtTokenFromQuery := r.URL.Query().Get("jwt")

	// Get JWT token from Header
	jwtTokenFromHeaderParts := strings.Split(r.Header.Get("Authorization"), " ")
	jwtTokenFromHeader := ""
	if len(jwtTokenFromHeaderParts) == 2 {
		jwtTokenFromHeader = jwtTokenFromHeaderParts[1]
	}

	// Check if one of them is here
	if jwtTokenFromQuery == "" && jwtTokenFromHeader == "" {
		return "", errors.New("missing jwt parameter in URL and/or header")
	}

	if jwtTokenFromQuery != "" {
		return jwtTokenFromQuery, nil
	} else {
		return jwtTokenFromHeader, nil
	}
}

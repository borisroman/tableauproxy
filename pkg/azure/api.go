package azure

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/borisroman/tableauproxy/pkg/models"
	mssqldb "github.com/denisenkom/go-mssqldb"
)

type Controller struct {
	Server   string
	Port     int
	User     string
	Password string
	Database string

	ConnectionString string

	DB *sql.DB
}

func NewClient(server string, port int, user string, password string, database string) (*Controller, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("Connected to database %s", server)

	return &Controller{
		Server:           server,
		Port:             port,
		User:             user,
		Password:         password,
		Database:         database,
		ConnectionString: connString,
		DB:               db,
	}, nil
}

func (c *Controller) CreateAppInstance(payload *models.LifecyclePayload) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := c.DB.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := "INSERT INTO TableauProxy.AppInstances (AppKey, ClientKey, AccountID, SharedSecret, BaseUrl, DisplayURL, ProductType, Description, ServiceEntitlementNumber, OauthClientId) VALUES (@AppKey, @ClientKey, @AccountID, @SharedSecret, @BaseUrl, @DisplayURL, @ProductType, @Description, @ServiceEntitlementNumber, @OauthClientId); select convert(bigint, SCOPE_IDENTITY());"

	stmt, err := c.DB.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("AppKey", payload.Key),
		sql.Named("ClientKey", payload.ClientKey),
		sql.Named("AccountID", payload.AccountID),
		sql.Named("SharedSecret", payload.SharedSecret),
		sql.Named("BaseUrl", payload.BaseUrl),
		sql.Named("DisplayURL", payload.DisplayURL),
		sql.Named("ProductType", payload.ProductType),
		sql.Named("Description", payload.Description),
		sql.Named("ServiceEntitlementNumber", payload.ServiceEntitlementNumber),
		sql.Named("OauthClientId", payload.OauthClientId))
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}

	return newID, nil
}

func (c *Controller) ReadAppInstance(payload *models.LifecyclePayload) (*models.LifecyclePayload, error) {
	ctx := context.Background()

	err := c.DB.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	tsql := fmt.Sprintf("SELECT * FROM TableauProxy.AppInstances WHERE ClientKey = @ClientKey;")

	rows, err := c.DB.QueryContext(ctx, tsql, sql.Named("ClientKey", payload.ClientKey))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var id int
		var key, clientKey, accountID, sharedSecret, baseUrl, displayURL, productType, description, serviceEntitlementNumber, oauthClientId string

		err := rows.Scan(&id, &key, &clientKey, &accountID, &sharedSecret, &baseUrl, &displayURL, &productType, &description, &serviceEntitlementNumber, &oauthClientId)
		if err != nil {
			return nil, err
		}

		return &models.LifecyclePayload{
			Key:                      key,
			ClientKey:                clientKey,
			AccountID:                accountID,
			SharedSecret:             sharedSecret,
			BaseUrl:                  baseUrl,
			DisplayURL:               displayURL,
			ProductType:              productType,
			Description:              description,
			ServiceEntitlementNumber: serviceEntitlementNumber,
			OauthClientId:            oauthClientId,
		}, nil
	}

	return nil, nil
}

func (c *Controller) ReadAppInstanceByClientKey(clientKey string) (*models.LifecyclePayload, error) {
	ctx := context.Background()

	err := c.DB.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	tsql := fmt.Sprintf("SELECT * FROM TableauProxy.AppInstances WHERE ClientKey = @ClientKey;")

	rows, err := c.DB.QueryContext(ctx, tsql, sql.Named("ClientKey", clientKey))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var id int
		var key, clientKey, accountID, sharedSecret, baseUrl, displayURL, productType, description, serviceEntitlementNumber, oauthClientId string

		err := rows.Scan(&id, &key, &clientKey, &accountID, &sharedSecret, &baseUrl, &displayURL, &productType, &description, &serviceEntitlementNumber, &oauthClientId)
		if err != nil {
			return nil, err
		}

		return &models.LifecyclePayload{
			Key:                      key,
			ClientKey:                clientKey,
			AccountID:                accountID,
			SharedSecret:             sharedSecret,
			BaseUrl:                  baseUrl,
			DisplayURL:               displayURL,
			ProductType:              productType,
			Description:              description,
			ServiceEntitlementNumber: serviceEntitlementNumber,
			OauthClientId:            oauthClientId,
		}, nil
	}

	return nil, nil
}

func (c *Controller) UpdateAppInstance(payload *models.LifecyclePayload) (int64, error) {
	ctx := context.Background()

	err := c.DB.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("UPDATE TableauProxy.AppInstances SET AppKey = @AppKey, AccountID = @AccountID, SharedSecret = @SharedSecret, BaseUrl = @BaseUrl, DisplayURL = @DisplayURL, ProductType = @ProductType, Description = @Description, ServiceEntitlementNumber = @ServiceEntitlementNumber, OauthClientId = @OauthClientId WHERE ClientKey = @ClientKey")

	result, err := c.DB.ExecContext(
		ctx,
		tsql,
		sql.Named("AppKey", payload.Key),
		sql.Named("ClientKey", payload.ClientKey),
		sql.Named("AccountID", payload.AccountID),
		sql.Named("SharedSecret", payload.SharedSecret),
		sql.Named("BaseUrl", payload.BaseUrl),
		sql.Named("DisplayURL", payload.DisplayURL),
		sql.Named("ProductType", payload.ProductType),
		sql.Named("Description", payload.Description),
		sql.Named("ServiceEntitlementNumber", payload.ServiceEntitlementNumber),
		sql.Named("OauthClientId", payload.OauthClientId))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

func (c *Controller) DeleteAppInstance(payload *models.LifecyclePayload) (int64, error) {
	ctx := context.Background()

	err := c.DB.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("DELETE FROM TableauProxy.AppInstances WHERE ClientKey = @ClientKey;")

	result, err := c.DB.ExecContext(ctx, tsql, sql.Named("ClientKey", payload.ClientKey))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

func (c *Controller) CreatePersonalAccessToken(personalAccessToken *models.PersonalAccessToken) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := c.DB.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := "INSERT INTO TableauProxy.PersonalAccessTokens (UUID, BaseUrl, ClientKey, Name, Secret) VALUES (NEWID(), @BaseUrl, @ClientKey, @Name, @Secret); select convert(bigint, SCOPE_IDENTITY());"

	stmt, err := c.DB.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("BaseUrl", personalAccessToken.BaseUrl),
		sql.Named("ClientKey", personalAccessToken.ClientKey),
		sql.Named("Name", personalAccessToken.Name),
		sql.Named("Secret", personalAccessToken.Secret))
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}

	return newID, nil
}

func (c *Controller) ReadPersonalAccessToken(clientKey string) ([]*models.PersonalAccessToken, error) {
	var personalAccessTokens []*models.PersonalAccessToken

	ctx := context.Background()

	err := c.DB.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	tsql := fmt.Sprintf("SELECT * FROM TableauProxy.PersonalAccessTokens WHERE ClientKey = @ClientKey;")

	rows, err := c.DB.QueryContext(ctx, tsql, sql.Named("ClientKey", clientKey))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var baseUrl, clientKey, personalAccessTokenName, personalAccessTokenSecret string
		var uuid mssqldb.UniqueIdentifier

		err := rows.Scan(&id, &uuid, &baseUrl, &clientKey, &personalAccessTokenName, &personalAccessTokenSecret)
		if err != nil {
			return nil, err
		}

		personalAccessTokens = append(personalAccessTokens, &models.PersonalAccessToken{
			UUID:      uuid.String(),
			BaseUrl:   baseUrl,
			ClientKey: clientKey,
			Name:      personalAccessTokenName,
			Secret:    personalAccessTokenSecret,
		})
	}

	return personalAccessTokens, nil
}

func (c *Controller) ReadPersonalAccessTokenByUUID(uuid string, clientKey string) (*models.PersonalAccessToken, error) {
	ctx := context.Background()

	err := c.DB.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	tsql := fmt.Sprintf("SELECT * FROM TableauProxy.PersonalAccessTokens WHERE ClientKey = @ClientKey and UUID = @UUID;")

	rows, err := c.DB.QueryContext(ctx, tsql,
		sql.Named("ClientKey", clientKey),
		sql.Named("UUID", uuid))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var id int
		var baseUrl, clientKey, personalAccessTokenName, personalAccessTokenSecret string
		var uuid mssqldb.UniqueIdentifier

		err := rows.Scan(&id, &uuid, &baseUrl, &clientKey, &personalAccessTokenName, &personalAccessTokenSecret)
		if err != nil {
			return nil, err
		}

		return &models.PersonalAccessToken{
			UUID:      uuid.String(),
			BaseUrl:   baseUrl,
			ClientKey: clientKey,
			Name:      personalAccessTokenName,
			Secret:    personalAccessTokenSecret,
		}, nil
	}

	return nil, nil
}

func (c *Controller) ReadPersonalAccessTokenByName(name string, clientKey string) (*models.PersonalAccessToken, error) {
	ctx := context.Background()

	err := c.DB.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	tsql := fmt.Sprintf("SELECT * FROM TableauProxy.PersonalAccessTokens WHERE ClientKey = @ClientKey and name = @Name;")

	rows, err := c.DB.QueryContext(ctx, tsql,
		sql.Named("ClientKey", clientKey),
		sql.Named("Name", name))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var id int
		var baseUrl, clientKey, personalAccessTokenName, personalAccessTokenSecret string
		var uuid mssqldb.UniqueIdentifier

		err := rows.Scan(&id, &uuid, &baseUrl, &clientKey, &personalAccessTokenName, &personalAccessTokenSecret)
		if err != nil {
			return nil, err
		}

		return &models.PersonalAccessToken{
			UUID:      uuid.String(),
			BaseUrl:   baseUrl,
			ClientKey: clientKey,
			Name:      personalAccessTokenName,
			Secret:    personalAccessTokenSecret,
		}, nil
	}

	return nil, nil
}

func (c *Controller) UpdateAPersonalAccessToken(personalAccessToken *models.PersonalAccessToken) (int64, error) {
	ctx := context.Background()

	err := c.DB.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("UPDATE TableauProxy.PersonalAccessTokens SET Secret = @Secret, Name = @Name, BaseUrl = @BaseUrl WHERE ClientKey = @ClientKey AND UUID = @UUID;")

	result, err := c.DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ClientKey", personalAccessToken.ClientKey),
		sql.Named("UUID", personalAccessToken.UUID),
		sql.Named("BaseUrl", personalAccessToken.BaseUrl),
		sql.Named("Name", personalAccessToken.Name),
		sql.Named("Secret", personalAccessToken.Secret))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

func (c *Controller) DeletePersonalAccessToken(personalAccessToken *models.PersonalAccessToken) (int64, error) {
	ctx := context.Background()

	err := c.DB.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("DELETE FROM TableauProxy.PersonalAccessTokens WHERE ClientKey = @ClientKey AND UUID = @UUID;")

	result, err := c.DB.ExecContext(ctx, tsql,
		sql.Named("ClientKey", personalAccessToken.ClientKey),
		sql.Named("UUID", personalAccessToken.UUID))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

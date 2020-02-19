CREATE SCHEMA TableauProxy;
GO

CREATE TABLE TableauProxy.AppInstances (
    Id                          INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
    AppKey                      NVARCHAR(128),
    ClientKey                   NVARCHAR(128) NOT NULL UNIQUE,
    AccountID                   NVARCHAR(128),
    SharedSecret                NVARCHAR(128) NOT NULL,
    BaseUrl                     NVARCHAR(128) NOT NULL,
    DisplayURL                  NVARCHAR(128),
    ProductType                 NVARCHAR(128),
    Description                 NVARCHAR(128),
    ServiceEntitlementNumber    NVARCHAR(128),
    OauthClientId               NVARCHAR(128)
);
GO

CREATE TABLE TableauProxy.PersonalAccessTokens (
    Id          INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
    UUID        UNIQUEIDENTIFIER NOT NULL,
    BaseUrl     NVARCHAR(128) NOT NULL,
    ClientKey   NVARCHAR(128) NOT NULL,
    Name        NVARCHAR(128) NOT NULL,
    Secret      NVARCHAR(128) NOT NULL
);
GO

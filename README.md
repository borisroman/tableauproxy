# tableauproxy

A proxy for Tableau that uses JWT authentication to gather images from Views.


### Database migrations
Database migrations are handled by [Flyway](https://flywaydb.org/). Make sure you have the CLI client installed before continuing.
Change your working directory to `flyway` and fill in the correct credentials in the `flyway.conf` file. Next run `flyway migrate`.
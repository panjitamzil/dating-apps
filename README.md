# dating-apps
Explore the heart of our Dating App with this Golang-based REST API repository

## How to run 
1. Set up your database configuration (username and password) on `config.toml`
```
[Database]
  Host                  = 127.0.0.1
  Port                  = 5432
  User                  = 
  Password              = 
  Name                  = dating_apps
  MaxIdleConnection     = 10
	MaxOpenConnection     = 100
	MaxLifetimeConnection = 60
	MaxIdletimeConnection = 60
```

2. Run `go mod tidy` to update the dependencies

3. Create table using this query
```
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXIST users (
	id uuid DEFAULT uuid_generate_v4 (),
	email varchar(100) NOT NULL,
	password varchar(255) NOT NULL,
	fullname varchar(100),
	dob date,
	occupation varchar(100),
  subscription varchar(20),
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id)
)
```

4. Run the apps using `go run main.go`
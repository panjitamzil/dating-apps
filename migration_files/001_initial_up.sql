-- 001_initial_up.sql

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

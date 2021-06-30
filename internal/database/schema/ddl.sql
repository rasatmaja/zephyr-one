-- Install postgress extension for generate UUID
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create table auth
DROP TABLE IF EXISTS auth;
CREATE TABLE auth (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR ( 50 ) UNIQUE NOT NULL,
    passphrase VARCHAR ( 50 ) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create table account_info CREATE
DROP TABLE IF EXISTS account_info;
CREATE TABLE account_info (
    id uuid PRIMARY KEY,
    name VARCHAR ( 100 ) UNIQUE NOT NULL,
    photo_profile VARCHAR ( 200 ) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_id FOREIGN KEY (id) REFERENCES auth(id) ON DELETE CASCADE
);

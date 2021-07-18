-- AUTHOR: Rasio Ganang Atmaja <rasioatmaja29@gmail.com>

-- DROP Tables
DROP TABLE IF EXISTS contacts;
DROP TABLE IF EXISTS account_info;
DROP TABLE IF EXISTS auth;
DROP TABLE IF EXISTS contact_type;

-- Create table auth
CREATE TABLE auth (
    id uuid PRIMARY KEY NOT NULL,
    username VARCHAR ( 50 ) UNIQUE NOT NULL,
    passphrase VARCHAR ( 200 ) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create table account_info
CREATE TABLE account_info (
    id uuid PRIMARY KEY,
    name VARCHAR ( 100 ) NOT NULL,
    photo_profile VARCHAR ( 200 ) NOT NULL DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_id FOREIGN KEY (id) REFERENCES auth(id) ON DELETE CASCADE
);

-- Create table contact_type
CREATE TABLE contact_type (
    id SERIAL PRIMARY KEY,
    slug VARCHAR(10) UNIQUE NOT NULL,
    name VARCHAR(10) UNIQUE NOT NULL
);

-- Create seeder contact_type
INSERT INTO contact_type(slug, name) VALUES ('email', 'E-MAIL'), ('phone', 'PHONE');

-- Create table contacts
CREATE TABLE contacts (
    id SERIAL PRIMARY KEY,
    auth_id uuid,
    contact_type_id INT,
    contact VARCHAR(200) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_contact_type FOREIGN KEY(contact_type_id) REFERENCES contact_type(id),
    CONSTRAINT fk_auth FOREIGN KEY(auth_id) REFERENCES auth(id) ON DELETE CASCADE
);
USE bello_auth;

CREATE TABLE user (
    id BINARY(36) PRIMARY KEY,
    email NVARCHAR(50) NOT NULL,
    password NVARCHAR(100) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT,
    updated_at DATETIME NOT NULL
)
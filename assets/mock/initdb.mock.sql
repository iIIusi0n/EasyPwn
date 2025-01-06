CREATE FUNCTION BIN_TO_UUID(bytes BLOB) RETURNS TEXT AS
BEGIN
    RETURN lower(hex(substr(bytes,1,4))) || '-' || 
           lower(hex(substr(bytes,5,2))) || '-' || 
           lower(hex(substr(bytes,7,2))) || '-' ||
           lower(hex(substr(bytes,9,2))) || '-' ||
           lower(hex(substr(bytes,11,6)));
END;

CREATE FUNCTION UUID_TO_BIN(uuid TEXT) RETURNS BLOB AS
BEGIN
    RETURN hex(replace(uuid, '-', ''));
END;

CREATE TABLE IF NOT EXISTS user (
    id BINARY(16) PRIMARY KEY DEFAULT (UUID_TO_BIN(
        CONCAT(
            LOWER(HEX(RANDOM_BYTES(4))), '-',
            LOWER(HEX(RANDOM_BYTES(2))), '-4',
            SUBSTR(LOWER(HEX(RANDOM_BYTES(2))),2), '-',
            SUBSTR('89ab', FLOOR(1 + (RAND() * 4)), 1),
            SUBSTR(LOWER(HEX(RANDOM_BYTES(2))),2), '-',
            LOWER(HEX(RANDOM_BYTES(6)))
        )
    )),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_license_type (
    id BINARY(16) PRIMARY KEY DEFAULT (UUID_TO_BIN(
        CONCAT(
            LOWER(HEX(RANDOM_BYTES(4))), '-',
            LOWER(HEX(RANDOM_BYTES(2))), '-4',
            SUBSTR(LOWER(HEX(RANDOM_BYTES(2))),2), '-',
            SUBSTR('89ab', FLOOR(1 + (RAND() * 4)), 1),
            SUBSTR(LOWER(HEX(RANDOM_BYTES(2))),2), '-',
            LOWER(HEX(RANDOM_BYTES(6)))
        )
    )),
    name VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_license (
    id BINARY(16) PRIMARY KEY DEFAULT (UUID_TO_BIN(
        CONCAT(
            LOWER(HEX(RANDOM_BYTES(4))), '-',
            LOWER(HEX(RANDOM_BYTES(2))), '-4',
            SUBSTR(LOWER(HEX(RANDOM_BYTES(2))),2), '-',
            SUBSTR('89ab', FLOOR(1 + (RAND() * 4)), 1),
            SUBSTR(LOWER(HEX(RANDOM_BYTES(2))),2), '-',
            LOWER(HEX(RANDOM_BYTES(6)))
        )
    )),
    user_id BINARY(16) NOT NULL UNIQUE REFERENCES user(id),
    license_type_id BINARY(16) NOT NULL REFERENCES user_license_type(id),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO user_license_type (name) VALUES ('free');
INSERT INTO user_license_type (name) VALUES ('paid');

CREATE TABLE IF NOT EXISTS project_os (
    id BINARY(16) PRIMARY KEY DEFAULT (UUID_TO_BIN(
        CONCAT(
            LOWER(HEX(RANDOM_BYTES(4))), '-',
            LOWER(HEX(RANDOM_BYTES(2))), '-4',
            SUBSTR(LOWER(HEX(RANDOM_BYTES(2))),2), '-',
            SUBSTR('89ab', FLOOR(1 + (RAND() * 4)), 1),
            SUBSTR(LOWER(HEX(RANDOM_BYTES(2))),2), '-',
            LOWER(HEX(RANDOM_BYTES(6)))
        )
    )),
    name VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO project_os (name) VALUES ('ubuntu-2410');

CREATE TABLE IF NOT EXISTS project_plugin (
    id BINARY(16) PRIMARY KEY DEFAULT (UUID_TO_BIN(
        CONCAT(
            LOWER(HEX(RANDOM_BYTES(4))), '-',
            LOWER(HEX(RANDOM_BYTES(2))), '-4',
            SUBSTR(LOWER(HEX(RANDOM_BYTES(2))),2), '-',
            SUBSTR('89ab', FLOOR(1 + (RAND() * 4)), 1),
            SUBSTR(LOWER(HEX(RANDOM_BYTES(2))),2), '-',
            LOWER(HEX(RANDOM_BYTES(6)))
        )
    )),
    name VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO project_plugin (name) VALUES ('gef');
INSERT INTO project_plugin (name) VALUES ('pwndbg');

CREATE TABLE IF NOT EXISTS project (
    id BINARY(16) PRIMARY KEY DEFAULT (UUID_TO_BIN(
        CONCAT(
            LOWER(HEX(RANDOM_BYTES(4))), '-',
            LOWER(HEX(RANDOM_BYTES(2))), '-4',
            SUBSTR(LOWER(HEX(RANDOM_BYTES(2))),2), '-',
            SUBSTR('89ab', FLOOR(1 + (RAND() * 4)), 1),
            SUBSTR(LOWER(HEX(RANDOM_BYTES(2))),2), '-',
            LOWER(HEX(RANDOM_BYTES(6)))
        )
    )),
    name VARCHAR(255) NOT NULL,
    user_id BINARY(16) NOT NULL REFERENCES user(id),
    file_path VARCHAR(255) NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    os_id BINARY(16) NOT NULL REFERENCES project_os(id),
    plugin_id BINARY(16) NOT NULL REFERENCES project_plugin(id),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS instance (
    id BINARY(16) PRIMARY KEY DEFAULT (UUID_TO_BIN(
        CONCAT(
            LOWER(HEX(RANDOM_BYTES(4))), '-',
            LOWER(HEX(RANDOM_BYTES(2))), '-4',
            SUBSTR(LOWER(HEX(RANDOM_BYTES(2))),2), '-',
            SUBSTR('89ab', FLOOR(1 + (RAND() * 4)), 1),
            SUBSTR(LOWER(HEX(RANDOM_BYTES(2))),2), '-',
            LOWER(HEX(RANDOM_BYTES(6)))
        )
    )),
    project_id BINARY(16) NOT NULL REFERENCES project(id),
    container_id VARCHAR(64) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS instance_log (
    id BINARY(16) PRIMARY KEY DEFAULT (UUID_TO_BIN(
        CONCAT(
            LOWER(HEX(RANDOM_BYTES(4))), '-',
            LOWER(HEX(RANDOM_BYTES(2))), '-4',
            SUBSTR(LOWER(HEX(RANDOM_BYTES(2))),2), '-',
            SUBSTR('89ab', FLOOR(1 + (RAND() * 4)), 1),
            SUBSTR(LOWER(HEX(RANDOM_BYTES(2))),2), '-',
            LOWER(HEX(RANDOM_BYTES(6)))
        )
    )),
    instance_id BINARY(16) NOT NULL REFERENCES instance(id),
    log TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS email_confirmation (
    id BINARY(16) PRIMARY KEY DEFAULT (UUID_TO_BIN(
        CONCAT(
            LOWER(HEX(RANDOM_BYTES(4))), '-',
            LOWER(HEX(RANDOM_BYTES(2))), '-4',
            SUBSTR(LOWER(HEX(RANDOM_BYTES(2))),2), '-',
            SUBSTR('89ab', FLOOR(1 + (RAND() * 4)), 1),
            SUBSTR(LOWER(HEX(RANDOM_BYTES(2))),2), '-',
            LOWER(HEX(RANDOM_BYTES(6)))
        )
    )),
    email VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
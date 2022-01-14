CREATE EXTENSION postgis;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id         UUID PRIMARY KEY                     DEFAULT uuid_generate_v4(),
    full_name  VARCHAR(255)                NOT NULL,
    username   VARCHAR(255)                NOT NULL,
    email      VARCHAR(100)                NOT NULL,
    password   VARCHAR(2048)               NOT NULL,
    is_admin   BOOL                        NOT NULL DEFAULT false,
    is_active  BOOL                        NOT NULL DEFAULT false,
    settings   JSONB                       NOT NULL DEFAULT '{}'::jsonb,

    -- 'auth' is default, other authentication providers
    -- for example can be one of google, twitter etc.
    provider   VARCHAR(56)                 NULL     DEFAULT 'auth',
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP),
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP)
);

CREATE UNIQUE INDEX ix_email ON users (lower(email));
CREATE UNIQUE INDEX ix_username ON users (lower(username));
CREATE INDEX ix_full_name ON users (lower(full_name));


CREATE TABLE messages
(
    id           UUID PRIMARY KEY                     DEFAULT uuid_generate_v4(),
    owner_id     UUID                        NOT NULL,
    to_user_id   UUID                        NOT NULL,
    message      VARCHAR(1024)               NULL,
    raw_location VARCHAR(32)                 NULL,
    location     GEOMETRY(POINT, 4326)       NULL,
    created_at   TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP),
    expires_at   TIMESTAMP WITHOUT TIME ZONE NULL,

    CONSTRAINT fk_messages_owner
        FOREIGN KEY (owner_id)
            REFERENCES users (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_messages_to_user
        FOREIGN KEY (to_user_id)
            REFERENCES users (id)
            ON DELETE CASCADE
);

CREATE INDEX ix_owner ON messages (owner_id);
CREATE INDEX ix_to_user ON messages (to_user_id);
CREATE INDEX ix_geo_location ON messages USING gist (location);

CREATE TABLE files
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner_id   UUID NOT NULL,
    message_id UUID NOT NULL,

    CONSTRAINT fk_files_owner
        FOREIGN KEY (owner_id)
            REFERENCES users (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_files_message
        FOREIGN KEY (message_id)
            REFERENCES messages (id)
            ON DELETE CASCADE
);

CREATE INDEX ix_file_owner ON files (owner_id);
CREATE INDEX ix_related_message ON files (message_id);

INSERT INTO users
VALUES (uuid_generate_v4(),
        'Full Admin',
        'admin',
        'admin@example.com',
        'password',
        true,
        true,
        '{}',
        'auth',
        timezone('utc'::text, CURRENT_TIMESTAMP),
        timezone('utc'::text, CURRENT_TIMESTAMP));

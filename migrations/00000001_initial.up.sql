CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id         UUID PRIMARY KEY                     DEFAULT uuid_generate_v4(),
    full_name  VARCHAR(255)                NULL,
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

CREATE UNIQUE INDEX ix_users_email ON users (lower(email));
CREATE INDEX ix_users_full_name ON users (lower(full_name));

CREATE TABLE cards
(
    id             UUID PRIMARY KEY                     DEFAULT uuid_generate_v4(),
    user_id        UUID                        NOT NULL,
    encrypted_data TEXT                        NOT NULL,
    encrypted_key  VARCHAR(2048)               NOT NULL,
    expires_in     BIGINT                      NULL     DEFAULT 0,
    created_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP),

    CONSTRAINT fk_cards_user
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE CASCADE
);

CREATE INDEX ix_cards_user_id ON cards (user_id);

CREATE TABLE events
(
    id         UUID PRIMARY KEY                     DEFAULT uuid_generate_v4(),
    action     VARCHAR(20)                 NOT NULL, -- create:user, update:user etc.
    context    VARCHAR(20)                 NULL,     -- handler:users service:webhooks etc.
    source     VARCHAR(40)                 NULL,     -- user id, record id etc.
    data       JSONB                       NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP),
    expires_at TIMESTAMP WITHOUT TIME ZONE NULL
);

CREATE INDEX ix_events_action ON events (action);
CREATE INDEX ix_events_context ON events (context);
CREATE INDEX ix_events_source ON events (source);

INSERT INTO users
VALUES (uuid_generate_v4(),
        'Full Admin',
        'admin@example.com',
        'password',
        true,
        true,
        '{}',
        'auth',
        timezone('utc'::text, CURRENT_TIMESTAMP),
        timezone('utc'::text, CURRENT_TIMESTAMP));

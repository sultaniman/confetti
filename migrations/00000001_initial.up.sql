CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id           UUID PRIMARY KEY       DEFAULT uuid_generate_v4(),
    full_name    VARCHAR(255) NULL,
    email        VARCHAR(100)  NOT NULL,
    password     VARCHAR(2048) NOT NULL,
    is_admin     BOOL          NOT NULL DEFAULT false,
    is_active    BOOL          NOT NULL DEFAULT false,
    is_confirmed BOOL          NOT NULL DEFAULT false,
    settings     JSONB NULL DEFAULT '{}'::jsonb,

    -- 'auth' is default, other authentication providers
    -- for example can be one of google, twitter etc.
    provider     VARCHAR(56) NULL DEFAULT 'auth',
    created_at   TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP),
    updated_at   TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP)
);

CREATE UNIQUE INDEX ix_users_email ON users (lower(email));
CREATE INDEX ix_users_full_name ON users (lower(full_name));

CREATE TABLE user_confirmations
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id    UUID        NOT NULL,
    code       VARCHAR(40) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP),

    CONSTRAINT fk_user_confirmations_user
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE CASCADE
);

CREATE INDEX ix_user_confirmations_code ON user_confirmations (code);

CREATE TABLE password_resets
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id    UUID        NOT NULL,
    code       VARCHAR(40) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP),

    CONSTRAINT fk_password_resets_user
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE CASCADE
);

CREATE INDEX ix_password_resets_code ON password_resets (code);

CREATE TABLE cards
(
    id             UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id        UUID          NOT NULL,
    title          VARCHAR(255)  NOT NULL,
    encrypted_data TEXT          NOT NULL,
    encrypted_key  VARCHAR(2048) NOT NULL,
    key_id         VARCHAR(20) NULL,
    created_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP),
    updated_at     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP),

    CONSTRAINT fk_cards_user
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE CASCADE
);

CREATE INDEX ix_cards_user_id ON cards (user_id);

CREATE TABLE events
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    action     VARCHAR(20) NOT NULL, -- create:user, update:user etc.
    context    VARCHAR(20) NULL,     -- handler:users service:webhooks etc.
    source     VARCHAR(40) NULL,     -- user id, record id etc.
    data       JSONB NULL DEFAULT '{}'::jsonb,
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
        true,
        '{}',
        'auth',
        timezone('utc'::text, CURRENT_TIMESTAMP),
        timezone('utc'::text, CURRENT_TIMESTAMP));

CREATE TABLE IF NOT EXISTS users
(
    id            BIGINT PRIMARY KEY,
    name          TEXT                                   NOT NULL CONSTRAINT users_name_check CHECK (name <> ''::TEXT),
    email         TEXT                                   NOT NULL CONSTRAINT users_email_check CHECK (email <> ''::TEXT),
    password      TEXT                                   NOT NULL CONSTRAINT users_password_check CHECK (password <> ''::TEXT),
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    UNIQUE (email),
    UNIQUE (password)
    );
CREATE INDEX IF NOT EXISTS users_email_index ON users(email);
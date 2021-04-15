-- +migrate Up

CREATE TABLE devices(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);


-- +migrate Down

DROP TABLE `devices`;
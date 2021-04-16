-- +migrate Up

CREATE TABLE monitoring_data(
    id BIGSERIAL PRIMARY KEY,
    device_id BIGINT NOT NULL REFERENCES devices(id),
    temperature DOUBLE PRECISION NOT NULL,
    humidity DOUBLE PRECISION NOT NULL,
    heater_enabled BOOLEAN NOT NULL default false,
    humidifier_enabled  BOOLEAN NOT NULL default false,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +migrate Down

DROP TABLE monitoring_data;
BEGIN;

CREATE TYPE user_role AS ENUM (
    'admin',
    'employee',
    'project-manager'
    );

CREATE TYPE user_type AS ENUM (
    'full-time',
    'intern',
    'freelancer'
    );

CREATE TYPE asset_type AS ENUM (
    'laptop',
    'mobile',
    'mouse',
    'keyboard'
    );

CREATE TYPE asset_status AS ENUM (
    'available',
    'assigned',
    'needs_repair',
    'under_repair',
    'damaged'
    );

CREATE TYPE owner_type AS ENUM (
    'remotestate',
    'client'
    );

CREATE TYPE connection_type AS ENUM (
    'wired',
    'wireless',
    'bluetooth'
    );

CREATE TABLE IF NOT EXISTS users
(
    user_id     UUID PRIMARY KEY   DEFAULT gen_random_uuid(),
    name        TEXT      NOT NULL,
    email       TEXT      NOT NULL,
    phone_no    TEXT,
    role        user_role          DEFAULT 'employee',
    type        user_type NOT NULL DEFAULT 'full-time',
    password    TEXT      NOT NULL,
    created_at  TIMESTAMP          DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP,
    archived_at TIMESTAMPTZ
);

CREATE UNIQUE INDEX IF NOT EXISTS active_user ON users (TRIM(LOWER(email))) WHERE archived_at IS NULL;

CREATE TABLE IF NOT EXISTS assets
(
    asset_id       UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    serial_number  TEXT UNIQUE NOT NULL,
    brand          TEXT        NOT NULL,
    model          TEXT        NOT NULL,
    asset_type     asset_type  NOT NULL,
    status         asset_status     DEFAULT 'available',
    owner_type     owner_type       DEFAULT 'remotestate',
    warranty_start TIMESTAMP   NOT NULL,
    warranty_end   TIMESTAMP   NOT NULL,
    created_at     TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP,
    archived_at    TIMESTAMPTZ
);

create table laptop
(
    laptop_id        uuid primary key default gen_random_uuid(),
    asset_id         uuid unique references assets (asset_id) NOT NULL,
    ram              text,
    storage          text,
    operating_system text,
    processor        text,
    charger          boolean
);

CREATE TABLE mouse
(
    mouse_id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_id        UUID UNIQUE REFERENCES assets (asset_id) NOT NULL,
    connection_type connection_type                          NOT NULL,
    dpi             INTEGER
);

CREATE TABLE keyboard
(
    keyboard_id     UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_id        UUID UNIQUE REFERENCES assets (asset_id) NOT NULL,
    connection_type connection_type                          NOT NULL,
    layout          TEXT
);

CREATE TABLE mobile
(
    mobile_id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_id         UUID UNIQUE REFERENCES assets (asset_id) NOT NULL,
    operating_system TEXT,
    ram              INTEGER,
    storage          INTEGER,
    imei             TEXT UNIQUE,
    charger          boolean
);

CREATE TABLE asset_assignments
(
    assignment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_id      UUID NOT NULL REFERENCES assets (asset_id),
    assigned_to   UUID NOT NULL REFERENCES users (user_id),
    assigned_on   TIMESTAMPTZ      DEFAULT CURRENT_TIMESTAMP,
    returned_at   TIMESTAMPTZ,
    created_at    TIMESTAMPTZ      DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMPTZ,
    archived_at   TIMESTAMPTZ
);

CREATE TABLE asset_repairs
(
    repair_id           UUID PRIMARY KEY     DEFAULT gen_random_uuid(),
    asset_id            UUID        NOT NULL REFERENCES assets (asset_id),
    sent_for_repair_on  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    repair_completed_on TIMESTAMPTZ,
    created_at          TIMESTAMPTZ          DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ,
    archived_at         TIMESTAMPTZ
);

COMMIT;
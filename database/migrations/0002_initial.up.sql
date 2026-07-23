
ALTER TABLE users RENAME COLUMN user_id TO id;
ALTER TABLE assets RENAME COLUMN asset_id TO id;
ALTER TABLE laptop RENAME COLUMN laptop_id TO id;
ALTER TABLE mouse RENAME COLUMN mouse_id TO id;
ALTER TABLE keyboard RENAME COLUMN keyboard_id TO id;
ALTER TABLE mobile RENAME COLUMN mobile_id TO id;
ALTER TABLE asset_assignments RENAME COLUMN assignment_id TO id;
ALTER TABLE asset_repairs RENAME COLUMN repair_id TO id;

ALTER TABLE laptop
    ADD COLUMN created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN updated_at TIMESTAMPTZ,
    ADD COLUMN archived_at TIMESTAMPTZ;

ALTER TABLE mouse
    ADD COLUMN created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN updated_at TIMESTAMPTZ,
    ADD COLUMN archived_at TIMESTAMPTZ;

ALTER TABLE mobile
    ADD COLUMN created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN updated_at TIMESTAMPTZ,
    ADD COLUMN archived_at TIMESTAMPTZ;

ALTER TABLE keyboard
    ADD COLUMN created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN updated_at TIMESTAMPTZ,
    ADD COLUMN archived_at TIMESTAMPTZ;

ALTER TABLE users
    ALTER COLUMN email SET NOT NULL,
    ALTER COLUMN role SET NOT NULL,
    ALTER COLUMN created_at SET NOT NULL,
    ALTER COLUMN created_at TYPE TIMESTAMPTZ,
    ALTER COLUMN updated_at TYPE TIMESTAMPTZ;

ALTER TABLE assets
    ALTER COLUMN status SET NOT NULL,
    ALTER COLUMN owner_type SET NOT NULL,
    ALTER COLUMN created_at SET NOT NULL,
    ALTER COLUMN warranty_start TYPE TIMESTAMPTZ,
    ALTER COLUMN warranty_end TYPE TIMESTAMPTZ,
    ALTER COLUMN created_at TYPE TIMESTAMPTZ,
    ALTER COLUMN updated_at TYPE TIMESTAMPTZ,
    ADD CONSTRAINT assets_warranty_dates_check
        CHECK (warranty_end > warranty_start);

ALTER TABLE asset_assignments
    ALTER COLUMN assigned_on SET NOT NULL,
    ALTER COLUMN created_at SET NOT NULL;

ALTER TABLE asset_repairs
    ALTER COLUMN sent_for_repair_on SET NOT NULL,
    ALTER COLUMN created_at SET NOT NULL;

ALTER TABLE assets
    DROP CONSTRAINT assets_serial_number_key;

CREATE UNIQUE INDEX assets_serial_number_active_idx  ON assets (serial_number) WHERE archived_at IS NULL;
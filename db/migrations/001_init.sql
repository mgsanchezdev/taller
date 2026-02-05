-- Companies and workshops (id as VARCHAR for Go string UUIDs)
CREATE TABLE IF NOT EXISTS companies (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text,
    name VARCHAR(255) NOT NULL,
    tax_id VARCHAR(64),
    address TEXT,
    phone VARCHAR(64),
    email VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS workshops (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text,
    company_id VARCHAR(36) NOT NULL REFERENCES companies(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    address TEXT,
    phone VARCHAR(64),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Users
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text,
    workshop_id VARCHAR(36) NOT NULL REFERENCES workshops(id) ON DELETE CASCADE,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    role VARCHAR(32) NOT NULL DEFAULT 'tech',
    active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(workshop_id, email)
);

CREATE INDEX IF NOT EXISTS idx_users_workshop ON users(workshop_id);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(workshop_id, LOWER(email));

-- Vehicles
CREATE TABLE IF NOT EXISTS vehicles (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text,
    workshop_id VARCHAR(36) NOT NULL REFERENCES workshops(id) ON DELETE CASCADE,
    plate VARCHAR(32) NOT NULL,
    brand VARCHAR(128),
    model VARCHAR(128),
    year INT,
    vin VARCHAR(64),
    color VARCHAR(64),
    owner_name VARCHAR(255),
    owner_phone VARCHAR(64),
    owner_email VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_vehicles_workshop_plate ON vehicles(workshop_id, LOWER(plate));
CREATE INDEX IF NOT EXISTS idx_vehicles_workshop ON vehicles(workshop_id);

-- Workshop entries
CREATE TABLE IF NOT EXISTS workshop_entries (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text,
    workshop_id VARCHAR(36) NOT NULL REFERENCES workshops(id) ON DELETE CASCADE,
    vehicle_id VARCHAR(36) NOT NULL REFERENCES vehicles(id) ON DELETE CASCADE,
    entry_number VARCHAR(64) NOT NULL,
    entry_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    mileage INT DEFAULT 0,
    fuel_level VARCHAR(32),
    notes TEXT,
    status VARCHAR(32) NOT NULL DEFAULT 'received',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_entries_workshop ON workshop_entries(workshop_id);
CREATE INDEX IF NOT EXISTS idx_entries_vehicle ON workshop_entries(vehicle_id);
CREATE UNIQUE INDEX IF NOT EXISTS idx_entries_number ON workshop_entries(workshop_id, entry_number);

-- Inspection templates and items
CREATE TABLE IF NOT EXISTS inspection_templates (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text,
    workshop_id VARCHAR(36) NOT NULL REFERENCES workshops(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS inspection_items (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text,
    template_id VARCHAR(36) NOT NULL REFERENCES inspection_templates(id) ON DELETE CASCADE,
    label VARCHAR(255) NOT NULL,
    category VARCHAR(128),
    "order" INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Inspections and item results
CREATE TABLE IF NOT EXISTS inspections (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text,
    entry_id VARCHAR(36) NOT NULL REFERENCES workshop_entries(id) ON DELETE CASCADE,
    vehicle_id VARCHAR(36) NOT NULL REFERENCES vehicles(id) ON DELETE CASCADE,
    template_id VARCHAR(36) REFERENCES inspection_templates(id) ON DELETE SET NULL,
    performed_by VARCHAR(255),
    performed_at TIMESTAMPTZ,
    status VARCHAR(32) NOT NULL DEFAULT 'draft',
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS inspection_item_results (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text,
    inspection_id VARCHAR(36) NOT NULL REFERENCES inspections(id) ON DELETE CASCADE,
    item_id VARCHAR(36) NOT NULL REFERENCES inspection_items(id) ON DELETE CASCADE,
    status VARCHAR(32) NOT NULL,
    notes TEXT,
    photo_url TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_inspections_entry ON inspections(entry_id);
CREATE INDEX IF NOT EXISTS idx_inspections_vehicle ON inspections(vehicle_id);

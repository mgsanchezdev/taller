-- Faults
CREATE TABLE IF NOT EXISTS faults (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text,
    entry_id VARCHAR(36) NOT NULL REFERENCES workshop_entries(id) ON DELETE CASCADE,
    inspection_id VARCHAR(36) REFERENCES inspections(id) ON DELETE SET NULL,
    description TEXT NOT NULL,
    severity VARCHAR(32) NOT NULL DEFAULT 'medium',
    status VARCHAR(32) NOT NULL DEFAULT 'open',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Repairs
CREATE TABLE IF NOT EXISTS repairs (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text,
    entry_id VARCHAR(36) NOT NULL REFERENCES workshop_entries(id) ON DELETE CASCADE,
    vehicle_id VARCHAR(36) NOT NULL REFERENCES vehicles(id) ON DELETE CASCADE,
    fault_id VARCHAR(36) REFERENCES faults(id) ON DELETE SET NULL,
    description TEXT NOT NULL,
    labor_hours DECIMAL(10,2) DEFAULT 0,
    parts_cost DECIMAL(12,2) DEFAULT 0,
    status VARCHAR(32) NOT NULL DEFAULT 'pending',
    completed_at TIMESTAMPTZ,
    technician_id VARCHAR(36) REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Media (photos, videos, PDFs)
CREATE TABLE IF NOT EXISTS media (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text,
    entity_type VARCHAR(64) NOT NULL,
    entity_id VARCHAR(36) NOT NULL,
    type VARCHAR(32) NOT NULL,
    url TEXT NOT NULL,
    "key" TEXT NOT NULL,
    filename VARCHAR(255),
    size BIGINT DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Reports
CREATE TABLE IF NOT EXISTS reports (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text,
    workshop_id VARCHAR(36) NOT NULL REFERENCES workshops(id) ON DELETE CASCADE,
    entry_id VARCHAR(36) NOT NULL REFERENCES workshop_entries(id) ON DELETE CASCADE,
    vehicle_id VARCHAR(36) NOT NULL REFERENCES vehicles(id) ON DELETE CASCADE,
    type VARCHAR(32) NOT NULL,
    pdf_url TEXT NOT NULL,
    generated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Notifications
CREATE TABLE IF NOT EXISTS notifications (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text,
    entry_id VARCHAR(36) REFERENCES workshop_entries(id) ON DELETE SET NULL,
    vehicle_id VARCHAR(36) REFERENCES vehicles(id) ON DELETE SET NULL,
    channel VARCHAR(32) NOT NULL,
    "to" VARCHAR(255) NOT NULL,
    subject VARCHAR(255),
    body TEXT NOT NULL,
    status VARCHAR(32) NOT NULL DEFAULT 'pending',
    sent_at TIMESTAMPTZ,
    error TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_repairs_entry ON repairs(entry_id);
CREATE INDEX IF NOT EXISTS idx_repairs_vehicle ON repairs(vehicle_id);
CREATE INDEX IF NOT EXISTS idx_media_entity ON media(entity_type, entity_id);
CREATE INDEX IF NOT EXISTS idx_reports_entry ON reports(entry_id);
CREATE INDEX IF NOT EXISTS idx_reports_vehicle ON reports(vehicle_id);

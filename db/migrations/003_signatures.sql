-- Inspection signatures
CREATE TABLE IF NOT EXISTS inspection_signatures (
    id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text,
    inspection_id VARCHAR(36) NOT NULL REFERENCES inspections(id) ON DELETE CASCADE,
    signer_role VARCHAR(64) NOT NULL,
    signer_name VARCHAR(255) NOT NULL,
    signature_url TEXT NOT NULL,
    signed_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    ip VARCHAR(64),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_signatures_inspection ON inspection_signatures(inspection_id);

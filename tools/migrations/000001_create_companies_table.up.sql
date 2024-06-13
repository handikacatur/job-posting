CREATE TABLE companies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL
);

ALTER TABLE companies ADD COLUMN tsv TSVECTOR
    GENERATED ALWAYS AS (to_tsvector('english', name)) STORED;

CREATE INDEX idx_companies_tsv ON companies USING gin(tsv);

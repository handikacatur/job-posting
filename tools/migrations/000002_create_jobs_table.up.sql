CREATE TABLE jobs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    company_id UUID NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_company
      FOREIGN KEY(company_id)
      REFERENCES companies(id)
);

ALTER TABLE jobs ADD COLUMN tsv TSVECTOR
    GENERATED ALWAYS AS (to_tsvector('english', title || '' || description)) STORED;

CREATE INDEX idx_tsv ON jobs USING gin(tsv);
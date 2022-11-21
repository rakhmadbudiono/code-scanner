BEGIN;
CREATE TYPE enum_status AS ENUM (
    'QUEUED',
    'IN PROGRESS',
    'SUCCESS',
    'FAILURE'
);
CREATE TABLE IF NOT EXISTS results(
    id uuid NOT NULL PRIMARY KEY,
    repository_id uuid NOT NULL,
    status enum_status NOT NULL,
    findings JSONB,
    queued_at TIMESTAMP,
    scanning_at TIMESTAMP,
    finished_at TIMESTAMP,
    FOREIGN KEY (repository_id) REFERENCES repositories(id)
);
COMMIT;
-- Create the "servers" table then insert some data
CREATE TABLE IF NOT EXISTS servers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

INSERT INTO servers (name, type, status) VALUES ('Server 1', 'small', 'starting');
INSERT INTO servers (name, type, status) VALUES ('Server 2', 'medium', 'running');
INSERT INTO servers (name, type, status) VALUES ('Server 3', 'large', 'stopping');
INSERT INTO servers (name, type, status) VALUES ('Server 4', 'small', 'stopped');
INSERT INTO servers (name, type, status) VALUES ('Server 5', 'medium', 'starting');

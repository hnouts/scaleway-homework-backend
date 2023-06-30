CREATE USER myuser WITH PASSWORD 'example';
GRANT ALL PRIVILEGES ON DATABASE mydatabase TO myuser;

-- Create the "servers" table
CREATE TABLE IF NOT EXISTS servers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL
);
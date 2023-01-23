CREATE DATABASE website;

CREATE user admin WITH encrypted PASSWORD 'admin';

GRANT ALL PRIVILEGES ON DATABASE website TO "admin";

ALTER ROLE admin
SET
    client_encoding TO 'utf8';

ALTER ROLE admin
SET
    default_transaction_isolation TO 'read committed';

ALTER ROLE admin
SET
    timezone TO 'UTC';
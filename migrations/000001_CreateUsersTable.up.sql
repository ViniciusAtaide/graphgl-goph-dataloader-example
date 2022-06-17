CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE Users(
    id uuid DEFAULT uuid_generate_v4 (),
    name VARCHAR NOT NULL,
    emoji VARCHAR NOT NULL,
    PRIMARY KEY (id)
)

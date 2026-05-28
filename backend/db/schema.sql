CREATE TABLE IF NOT EXISTS scps (
    id                TEXT PRIMARY KEY,
    name              TEXT NOT NULL,
    containment_class TEXT NOT NULL,
    description       TEXT NOT NULL,
    lore              TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS locations (
    id          TEXT PRIMARY KEY,
    name        TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS rooms (
    id          TEXT PRIMARY KEY,
    location_id TEXT NOT NULL REFERENCES locations(id),
    name        TEXT NOT NULL,
    image_path  TEXT NOT NULL
);

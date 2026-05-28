-- name: ListSCPs :many
SELECT * FROM scps ORDER BY id;

-- name: GetSCP :one
SELECT * FROM scps WHERE id = ?;

-- name: ListLocations :many
SELECT * FROM locations ORDER BY id;

-- name: ListRoomsForLocation :many
SELECT * FROM rooms WHERE location_id = ? ORDER BY id;

-- name: UpsertSCP :exec
INSERT OR IGNORE INTO scps (id, name, containment_class, description, lore)
VALUES (?, ?, ?, ?, ?);

-- name: UpsertLocation :exec
INSERT OR IGNORE INTO locations (id, name, description)
VALUES (?, ?, ?);

-- name: UpsertRoom :exec
INSERT OR IGNORE INTO rooms (id, location_id, name, image_path)
VALUES (?, ?, ?, ?);

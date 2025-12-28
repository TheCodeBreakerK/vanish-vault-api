-- name: GetUserByProvider :one
SELECT * FROM users
WHERE provider = $1 AND provider_id = $2 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (email, provider, provider_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateRoom :one
INSERT INTO vault_rooms (owner_id, name, access_code, expires_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: AddMemberToRoom :one
INSERT INTO room_members (room_id, user_id, role)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListMyRooms :many
SELECT r.* FROM vault_rooms r
JOIN room_members m ON r.id = m.room_id
WHERE m.user_id = $1 AND r.is_active = true;

-- name: DeleteRoom :exec
DELETE FROM vault_rooms
WHERE id = $1 AND owner_id = $2;

-- name: CreateSecret :one
INSERT INTO secret_items (room_id, creator_id, encrypted_content, nonce)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListSecretsByRoom :many
SELECT id, creator_id, created_at, is_burned 
FROM secret_items
WHERE room_id = $1 AND is_burned = false;

-- name: GetSecretForView :one
SELECT * FROM secret_items
WHERE id = $1 AND room_id = $2 AND is_burned = false 
LIMIT 1;

-- name: BurnSecret :exec
UPDATE secret_items
SET is_burned = true, burned_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: GetMemberRole :one
SELECT role FROM room_members
WHERE room_id = $1 AND user_id = $2;
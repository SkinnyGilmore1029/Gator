-- name: CreateFeed :one
INSERT INTO feeds (name, url, user_id) VALUES ($1, $2, $3) RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds WHERE id = $1;

-- name: UpdateFeed :one
UPDATE feeds SET name = $1, url = $2 WHERE id = $3 RETURNING *;

-- name: DeleteFeed :one
DELETE FROM feeds WHERE id = $1 RETURNING *;
-- name: CreateFeed :one
INSERT INTO feeds (name, url, user_id) VALUES ($1, $2, $3) RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds WHERE id = $1;

-- name: UpdateFeed :one
UPDATE feeds SET name = $1, url = $2 WHERE id = $3 RETURNING *;

-- name: DeleteFeed :one
DELETE FROM feeds WHERE id = $1 RETURNING *;

-- name: ListFeeds :many
SELECT feeds.name AS feed_name, feeds.url, users.name AS user_name FROM feeds INNER JOIN users ON feeds.user_id = users.id;

-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;

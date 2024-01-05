-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE id = $1 LIMIT 1;

-- name: GetFeeds :many
SELECT * FROM feeds;


-- name: GetNextFeedsToFetch :many
-- This should fetch a batch of feeds by null or by oldest date of last_fetched_at
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = NOW(), updated_at = NOW()
WHERE id = $1
RETURNING *;

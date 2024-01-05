-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, description, published_at, url, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;


-- name: GetPostsOfFeedUserFollows :many
SELECT posts.*
FROM feed_follows
JOIN posts ON feed_follows.feed_id = posts.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC
LIMIT 20;

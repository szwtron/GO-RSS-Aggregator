-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6) 
RETURNING *;

-- name: SelectAllFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds ORDER BY last_fetch_at ASC NULLS FIRST LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds SET last_fetch_at = NOW(), updated_at = NOW() WHERE id = $1 RETURNING *;
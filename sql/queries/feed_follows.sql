-- name: CreateFeedFollow :one
WITH inserted_feed AS (
	INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
	VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	) RETURNING *
) SELECT inserted_feed.*, users.name AS user_name, feeds.name AS feed_name
FROM inserted_feed
JOIN users on inserted_feed.user_id = users.id
JOIN feeds on inserted_feed.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, users.name AS user_name, feeds.name AS feed_name
FROM feed_follows
JOIN users on inserted_feed.user_id = users.id
JOIN feeds on inserted_feed.feed_id = feeds.id
WHERE feed_follows.user_id = $1;

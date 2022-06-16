-- name: CreateFollower :one
INSERT INTO user_followers (
  follower_id,
  followee_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetFollower :one
SELECT * FROM user_followers
WHERE followee_id = $1 LIMIT 1;

-- name: GetFollowee :one
SELECT * FROM user_followers
WHERE follower_id = $1 LIMIT 1;

-- name: ListFollowers :many
SELECT * FROM user_followers
ORDER BY followee_id
LIMIT $1
OFFSET $2;

-- name: ListFollowees :many
SELECT * FROM user_followers
ORDER BY follower_id
LIMIT $1
OFFSET $2;

-- name: DeleteFollower :exec
DELETE FROM user_followers
WHERE follower_id = $1;
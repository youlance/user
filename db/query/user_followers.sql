-- name: CreateFollower :one
INSERT INTO user_followers (
  follower_id,
  followee_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: ListFollowers :many
SELECT * FROM user_followers
WHERE followee_id = $1
ORDER BY follower_id
LIMIT $2
OFFSET $3;

-- name: ListFollowees :many
SELECT * FROM user_followers
WHERE follower_id = $1
ORDER BY followee_id
LIMIT $2
OFFSET $3;

-- name: GetFolloweesCount :one
SELECT count(*) FROM user_followers
WHERE follower_id = $1;

-- name: GetFollowersCount :one
SELECT count(*) FROM user_followers
WHERE followee_id = $1;


-- name: DeleteFollower :exec
DELETE FROM user_followers
WHERE follower_id = $1 AND followee_id = $2;
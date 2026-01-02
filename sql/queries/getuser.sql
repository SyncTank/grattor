-- name: GetUser :one
SELECT users (id)
VALUES (
    $1,
)
RETURNING *;

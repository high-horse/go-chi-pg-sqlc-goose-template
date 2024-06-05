-- name: CreateUser :one
INSERT INTO users
(
    username,
    email,
    password_hash,
    created_at
) VALUES (
    $1, $2, $3, $4
) RETURNING * ;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM users WHERE username = $1 LIMIT 1;
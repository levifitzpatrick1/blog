-- name: GetPostBySlug :one
SELECT
    *
FROM
    posts
WHERE
    slug = ?
LIMIT
    1;

-- name: ListPosts :many
SELECT
    id,
    title,
    slug,
    published,
    SUBSTR (content, 0, 200) as excerpt
FROM
    posts
ORDER BY
    published DESC;

-- name: ListPostsOldestFirst :many
SELECT
    id,
    title,
    slug,
    published,
    SUBSTR (content, 0, 200) as excerpt
FROM
    posts
ORDER BY
    published ASC;

-- name: SearchPosts :many
SELECT
    id,
    title,
    slug,
    published,
    SUBSTR (content, 0, 200) as excerpt
FROM
    posts
WHERE
    title LIKE ?
    OR content LIKE ?
ORDER BY
    published DESC;

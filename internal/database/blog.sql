-- name: GetBlogBySlug :one
SELECT
    *
FROM
    blog
WHERE
    slug = ?
LIMIT
    1;

-- name: Listblog :many
SELECT
    *
FROM
    blog
ORDER BY
    published DESC;

-- name: ListblogOldestFirst :many
SELECT
    *
FROM
    blog
ORDER BY
    published ASC;

-- name: Searchblog :many
SELECT
    *
FROM
    blog
WHERE
    title LIKE ?
    OR content LIKE ?
ORDER BY
    published DESC;

-- name: GetLatestPost :one
SELECT
    title,
    slug
FROM
    blog
WHERE
    published IS NOT NULL
ORDER BY
    updated,
    published DESC
LIMIT
    1;

-- name: GetTagsBlog :many
SELECT
    tags.*
FROM
    tags
    JOIN tags_blog ON tags.id = tags_blog.tag_id
WHERE
    tags_blog.blog_id = ?

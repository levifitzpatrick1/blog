
-- name: GetFRCBySlug :one
SELECT
    *
FROM
    frc
WHERE
    slug = ?
LIMIT
    1;

-- name: GetFRCs :many
SELECT
    *
FROM
    frc
ORDER BY
    updated,
    created DESC;

-- name: GetTagsFrc :many
SELECT
    tags.*
FROM
    tags
    JOIN tags_frc ON tags.id = tags_frc.tag_id
WHERE
    tags_frc.frc_id = ?

-- name: CreatePost :one
INSERT INTO posts(id,created_at,updated_at,title,url,description,published_at,feed_id)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;

-- name: GetPosts :many
SELECT feeds.url AS "feed_url",sub_two.id,sub_two.title,sub_two.description,sub_two.url,sub_two.published_at FROM
    feeds INNER JOIN
    (SELECT posts.id,posts.title,posts.description,posts.url,posts.published_at,posts.feed_id FROM
        posts INNER JOIN 
        (SELECT * FROM feed_follows
        WHERE feed_follows.user_id = $1) sub_one
        ON posts.feed_id = sub_one.feed_id) sub_two
    ON feeds.id = sub_two.feed_id
ORDER BY sub_two.published_at DESC
LIMIT $2;

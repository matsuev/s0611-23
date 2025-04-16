
-- name: GetProfileByID :one
SELECT * FROM "public"."profile"
WHERE "id"=$1
;
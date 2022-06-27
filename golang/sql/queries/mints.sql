-- name: InsertMintWithImageURL :exec
INSERT INTO mints (block_num, collection_address, token_id, token_url, image_url, token_name)
  VALUES (@block_num::int, @collection_address::varchar, @token_id::varchar, @token_url::varchar, @image_url::varchar,@token_name::varchar)
  ON CONFLICT DO NOTHING;

-- name: InsertMintWithImageData :exec
INSERT INTO mints (block_num, collection_address, token_id, token_name, image_data)
  VALUES (@block_num::int, @collection_address::varchar, @token_id::varchar, @token_name::varchar, @image_data::varchar)
  ON CONFLICT DO NOTHING;

-- name: GetFlawedMintsForBlock :many
SELECT * FROM mints WHERE block_num = $1 AND image_url IS NULL;

-- name: GetMintsForBlock :many
SELECT * FROM mints WHERE block_num = $1;

-- name: GetHighestBlock :one
SELECT MAX(block_num) FROM mints;
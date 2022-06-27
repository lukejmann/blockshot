CREATE SCHEMA mints;

CREATE TABLE mints (
  block_num int NOT NULL,
  collection_address varchar(42) NOT NULL,
  token_id varchar(200) NOT NULL,
  token_url varchar(2000),
  token_name varchar(200),
  image_url varchar(1000),
  image_data varchar(30000),
  PRIMARY KEY (collection_address, token_id)
);


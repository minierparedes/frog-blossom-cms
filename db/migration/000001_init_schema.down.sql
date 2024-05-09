-- Drop foreign key constraint before dropping table
ALTER TABLE website DROP CONSTRAINT website_owner_id_fkey;

-- Drop tables that do not have dependencies
DROP TABLE IF EXISTS meta;
DROP TABLE IF EXISTS page_components;
DROP TABLE IF EXISTS pages;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS websites;

-- Drop table after dropping the foreign key constraint
DROP TABLE IF EXISTS users;


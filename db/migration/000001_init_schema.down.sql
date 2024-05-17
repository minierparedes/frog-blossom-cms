-- Migration Down: Drop tables that do not have dependencies
DROP TABLE IF EXISTS meta;
DROP TABLE IF EXISTS pages;
DROP TABLE IF EXISTS posts;

DROP TABLE IF EXISTS users;
-- Migration Down: Drop table after dropping the foreign key constraint


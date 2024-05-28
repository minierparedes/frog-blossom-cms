-- Migration Down: Rollback the changes to is_deleted column
DO $$
BEGIN
    IF EXISTS(SELECT 1 FROM information_schema.tables WHERE table_name = 'users') THEN
        ALTER TABLE users
            DROP COLUMN IF EXISTS is_deleted;
    END IF;
END $$;

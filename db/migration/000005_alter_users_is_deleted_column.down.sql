-- Migration Down: Rollback the changes to is_deleted column
DO $$
BEGIN
    IF EXISTS(SELECT 1 FROM information_schema.columns WHERE table_name = 'users' AND column_name = 'is_deleted') THEN
        ALTER TABLE users
            DROP COLUMN IF EXISTS is_deleted;
    END IF;
END $$;

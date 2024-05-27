DO $$
BEGIN
    IF EXISTS(SELECT 1 FROM information_schema.tables WHERE table_name = 'users') THEN
        ALTER TABLE users
            DROP COLUMN IF EXISTS is_deleted;

        ALTER TABLE users
            ADD COLUMN IF NOT EXISTS is_deleted BOOLEAN DEFAULT FALSE;
    END IF;
END $$;

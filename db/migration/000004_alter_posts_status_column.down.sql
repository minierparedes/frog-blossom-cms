-- Migration Down: Rollback the changes to status column
DO $$
BEGIN
    IF EXISTS(SELECT 1 FROM information_schema.tables WHERE table_name = 'posts') THEN
        ALTER TABLE posts
            DROP COLUMN IF EXISTS status;
        
        ALTER TABLE posts
            ADD COLUMN IF NOT EXISTS status varchar(255);
    END IF;
END $$;

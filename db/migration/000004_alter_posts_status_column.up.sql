DO $$
BEGIN
    IF EXISTS(SELECT 1 FROM information_schema.tables WHERE table_name = 'posts') THEN
        ALTER TABLE posts
            DROP COLUMN IF EXISTS status;

        ALTER TABLE posts
            ADD COLUMN IF NOT EXISTS status access NOT NULL;
    END IF;
END $$;

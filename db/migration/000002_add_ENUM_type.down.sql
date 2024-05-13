-- Migration Down: Rollback creation of enum types

-- Drop enum type "access"
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'access') THEN
        DROP TYPE access;
    END IF;
END $$;


-- Drop enum type "level"
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'level') THEN
        DROP TYPE level;
    END IF;
END $$;


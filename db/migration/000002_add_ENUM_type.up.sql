DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'access') THEN
        CREATE TYPE access AS ENUM ('admin', 'user');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'level') THEN
        CREATE TYPE level AS ENUM ('draft', 'pending', 'private', 'publish');
    END IF;
END $$;

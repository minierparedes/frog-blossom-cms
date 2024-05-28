DO $$
BEGIN
    BEGIN
        CREATE TYPE access AS ENUM ('admin', 'user');
    EXCEPTION WHEN duplicate_object THEN
        NULL;
    END;

    BEGIN
        CREATE TYPE level AS ENUM ('draft', 'pending', 'private', 'publish');
    EXCEPTION WHEN duplicate_object THEN
        NULL;
    END;
END $$;

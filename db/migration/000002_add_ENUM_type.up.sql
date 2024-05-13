DO $$
BEGIN
    BEGIN
        -- Attempt to create the enum type
        CREATE TYPE access AS ENUM ('admin', 'user');
    EXCEPTION WHEN duplicate_object THEN
        -- If the type already exists, do nothing
        NULL;
    END;

    BEGIN
        -- Attempt to create the enum type
        CREATE TYPE level AS ENUM ('draft', 'pending', 'private', 'publish');
    EXCEPTION WHEN duplicate_object THEN
        -- If the type already exists, do nothing
        NULL;
    END;
END $$;

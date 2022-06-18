CREATE TABLE transfer (id UUID DEFAULT public.uuid_generate_v4() NOT NULL PRIMARY KEY,
                       account_origin_id uuid NOT NULL,
                       account_destination_id uuid NOT NULL,
                       amount INTEGER NOT NULL, created_at TIMESTAMPTZ NOT NULL DEFAULT NOW());

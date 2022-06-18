CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE account (id UUID DEFAULT public.uuid_generate_v4() NOT NULL PRIMARY KEY,
                                    name VARCHAR(100) NOT NULL,
                                    cpf VARCHAR(100) UNIQUE NOT NULL,
                                    secret VARCHAR(255) NOT NULL,
                                    balance INTEGER NOT NULL,
                                    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW());

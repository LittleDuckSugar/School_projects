-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
    users_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    email varchar NOT NULL,
    "password" varchar NOT NULL,
    username varchar NOT NULL,
    tel varchar NOT NULL,
    age int4 NOT NULL,
    "location" varchar NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (users_id)
);
-- public."admin" definition

-- Drop table

-- DROP TABLE public."admin";

CREATE TABLE public."admin" (
    admin_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    username varchar NOT NULL,
    email varchar NOT NULL,
    "password" varchar NOT NULL,
    CONSTRAINT admin_pkey PRIMARY KEY (admin_id)
);
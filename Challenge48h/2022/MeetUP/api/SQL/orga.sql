-- public.orga definition

-- Drop table

-- DROP TABLE public.orga;

CREATE TABLE public.orga (
    orga_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    email varchar NOT NULL,
    "password" varchar NOT NULL,
    username varchar NOT NULL,
    tel varchar NOT NULL,
    note int4 NOT NULL,
    CONSTRAINT orga_pkey PRIMARY KEY (orga_id)
);
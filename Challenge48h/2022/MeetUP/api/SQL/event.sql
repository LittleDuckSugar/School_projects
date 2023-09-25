-- public."event" definition

-- Drop table

-- DROP TABLE public."event";

CREATE TABLE public."event" (
    event_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    nom varchar NOT NULL,
    description varchar NOT NULL,
    note int4 NOT NULL,
    "owner" uuid NOT NULL,
    "type" bool NOT NULL,
    category_id uuid NOT NULL,
    users _uuid NOT NULL,
    tags _varchar NOT NULL,
    "date" timestamp NOT NULL,
    duration int8 NOT NULL,
    CONSTRAINT event_pkey PRIMARY KEY (event_id)
);
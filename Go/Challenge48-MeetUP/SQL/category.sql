-- public.category definition

-- Drop table

-- DROP TABLE public.category;

CREATE TABLE public.category (
    category_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    "name" varchar NOT NULL,
    CONSTRAINT category_pkey PRIMARY KEY (category_id)
);
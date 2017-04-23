CREATE EXTENSION IF NOT EXISTS pgcrypto
    SCHEMA public
    VERSION "1.3";

CREATE TABLE IF NOT EXISTS public.characters
(
    uuid uuid NOT NULL DEFAULT gen_random_uuid(),
    name character varying(15) COLLATE pg_catalog."default" NOT NULL,
    class character varying(15) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT characters_pkey PRIMARY KEY (uuid)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.characters
    OWNER to postgres;

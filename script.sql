CREATE TABLE IF NOT EXISTS public.article
(
    content text COLLATE pg_catalog."default",
    title text COLLATE pg_catalog."default",
    "desc" text COLLATE pg_catalog."default"
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.article
    OWNER to postgres;
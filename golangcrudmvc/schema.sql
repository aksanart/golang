CREATE DATABASE golangcrud
-- 
CREATE TABLE public.buku
(
    harga text,
    id_buku SERIAL,
    judul text,
    path text,
    PRIMARY KEY (id_buku)
);
-- 

CREATE TABLE public.user_rest
(
    id_user integer NOT NULL DEFAULT nextval('user_rest_id_user_seq'::regclass),
    username text COLLATE pg_catalog."default" NOT NULL,
    key text COLLATE pg_catalog."default" NOT NULL,
    hash text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT user_rest_pkey PRIMARY KEY (id_user)
)
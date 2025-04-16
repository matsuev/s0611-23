CREATE TABLE IF NOT EXISTS public.profile
(
    id bigint NOT NULL,
    given_name character varying(50) NOT NULL DEFAULT '',
    family_name character varying(50) NOT NULL DEFAULT '',
    email character varying(100) NOT NULL DEFAULT '',
    PRIMARY KEY (id),
    FOREIGN KEY (id)
        REFERENCES public.account (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);
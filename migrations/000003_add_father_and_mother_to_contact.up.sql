ALTER TABLE IF EXISTS public.contact
   ADD COLUMN IF NOT EXISTS father_name character varying(100) NOT NULL DEFAULT '';


ALTER TABLE IF EXISTS public.contact
   ADD COLUMN IF NOT EXISTS mother_name character varying(100) NOT NULL DEFAULT '';
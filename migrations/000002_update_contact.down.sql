-- Table: public.contact

ALTER TABLE IF EXISTS public.contact
   DROP COLUMN IF EXISTS full_name;

ALTER TABLE IF EXISTS public.contact
   DROP COLUMN IF EXISTS family_name;

ALTER TABLE IF EXISTS public.contact
   DROP COLUMN IF EXISTS given_name;
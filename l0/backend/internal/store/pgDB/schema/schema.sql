-- DROP TABLE public.orders;

CREATE TABLE public.entity (
	id varchar(255) NOT NULL,
	orders jsonb NOT NULL,
	CONSTRAINT entity_pkey PRIMARY KEY (id)
);
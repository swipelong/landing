--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.9
-- Dumped by pg_dump version 9.6.9

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: emails_entered; Type: TABLE; Schema: public; Owner: jamiedover
--

CREATE TABLE public.emails_entered (
    id integer NOT NULL,
    email_address character varying(255) NOT NULL,
    creation_timestamp timestamp without time zone DEFAULT now() NOT NULL,
    ip_address character varying(64) NOT NULL,
    status_sent boolean,
    CONSTRAINT emails_entered_email_address_check CHECK (((email_address)::text <> ''::text))
);


ALTER TABLE public.emails_entered OWNER TO jamiedover;

--
-- Name: emails_entered_id_seq; Type: SEQUENCE; Schema: public; Owner: jamiedover
--

CREATE SEQUENCE public.emails_entered_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.emails_entered_id_seq OWNER TO jamiedover;

--
-- Name: emails_entered_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jamiedover
--

ALTER SEQUENCE public.emails_entered_id_seq OWNED BY public.emails_entered.id;


--
-- Name: emails_entered id; Type: DEFAULT; Schema: public; Owner: jamiedover
--

ALTER TABLE ONLY public.emails_entered ALTER COLUMN id SET DEFAULT nextval('public.emails_entered_id_seq'::regclass);


--
-- Data for Name: emails_entered; Type: TABLE DATA; Schema: public; Owner: jamiedover
--

COPY public.emails_entered (id, email_address, creation_timestamp, ip_address, status_sent) FROM stdin;
\.


--
-- Name: emails_entered_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jamiedover
--

SELECT pg_catalog.setval('public.emails_entered_id_seq', 194, true);


--
-- Name: emails_entered emails_entered_email_address_key; Type: CONSTRAINT; Schema: public; Owner: jamiedover
--

ALTER TABLE ONLY public.emails_entered
    ADD CONSTRAINT emails_entered_email_address_key UNIQUE (email_address);


--
-- Name: emails_entered emails_entered_pkey; Type: CONSTRAINT; Schema: public; Owner: jamiedover
--

ALTER TABLE ONLY public.emails_entered
    ADD CONSTRAINT emails_entered_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--


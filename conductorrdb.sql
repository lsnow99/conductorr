-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.2-beta2
-- PostgreSQL version: 12.0
-- Project Site: pgmodeler.io
-- Model Author: ---

-- object: conductorr | type: ROLE --
-- DROP ROLE IF EXISTS conductorr;
CREATE ROLE conductorr WITH 
	CREATEDB
	LOGIN
	ENCRYPTED PASSWORD '5AxJRGPATCqsh';
-- ddl-end --


-- Database creation must be done outside a multicommand file.
-- These commands were put in this file only as a convenience.
-- -- object: conductorrdb | type: DATABASE --
-- -- DROP DATABASE IF EXISTS conductorrdb;
-- CREATE DATABASE conductorrdb;
-- -- ddl-end --
-- 

-- object: public.media_type | type: TYPE --
-- DROP TYPE IF EXISTS public.media_type CASCADE;
CREATE TYPE public.media_type AS
 ENUM ('TV','MOVIE');
-- ddl-end --
ALTER TYPE public.media_type OWNER TO postgres;
-- ddl-end --

-- object: public.status | type: TYPE --
-- DROP TYPE IF EXISTS public.status CASCADE;
CREATE TYPE public.status AS
 ENUM ('GRABBED','FILEBOT_RUNNING','PLEX_SCANNING','COMPLETED');
-- ddl-end --
ALTER TYPE public.status OWNER TO postgres;
-- ddl-end --

-- object: public."Jobs" | type: TABLE --
-- DROP TABLE IF EXISTS public."Jobs" CASCADE;
CREATE TABLE public."Jobs" (
	grabber_internal_id integer NOT NULL,
	job_id serial NOT NULL,
	torrent_linker_id text NOT NULL,
	nzb_linker_id text NOT NULL,
	time_grabbed timestamptz DEFAULT now(),
	title text NOT NULL,
	imdb_id text NOT NULL,
	release_title text NOT NULL,
	content_type text NOT NULL,
	download_client text,
	download_directory text,
	status text NOT NULL,
	filebot_logs text,
	scanner_logs text,
	grabbed_quality text NOT NULL,
	grabbed_size text NOT NULL,
	time_filebot_started timestamptz,
	time_filebot_done timestamptz,
	time_scan_started timestamptz,
	time_scan_done timestamptz,
	CONSTRAINT "Jobs_pk" PRIMARY KEY (job_id)

);
-- ddl-end --
ALTER TABLE public."Jobs" OWNER TO conductorr;
-- ddl-end --

-- object: public."SystemConfiguration" | type: TABLE --
-- DROP TABLE IF EXISTS public."SystemConfiguration" CASCADE;
CREATE TABLE public."SystemConfiguration" (
	system_configuration_id bool NOT NULL DEFAULT true,
	username varchar(32),
	password bytea,
	filebot_logs_enabled bool DEFAULT true,
	plex_scanner_logs_enabled bool DEFAULT true,
	CONSTRAINT system_config_pk PRIMARY KEY (system_configuration_id),
	CONSTRAINT configuration_one_row CHECK (system_configuration_id)

);
-- ddl-end --
ALTER TABLE public."SystemConfiguration" OWNER TO conductorr;
-- ddl-end --



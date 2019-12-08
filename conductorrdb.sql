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

-- object: public.jobs | type: TABLE --
-- DROP TABLE IF EXISTS public.jobs CASCADE;
CREATE TABLE public.jobs (
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
	CONSTRAINT jobs_pk PRIMARY KEY (job_id)

);
-- ddl-end --
ALTER TABLE public.jobs OWNER TO conductorr;
-- ddl-end --

-- object: public.system_configuration | type: TABLE --
-- DROP TABLE IF EXISTS public.system_configuration CASCADE;
CREATE TABLE public.system_configuration (
	system_configuration_id bool NOT NULL DEFAULT true,
	username varchar(32),
	password bytea,
	filebot_logs_enabled bool DEFAULT true,
	plex_scanner_logs_enabled bool DEFAULT true,
	CONSTRAINT system_config_pk PRIMARY KEY (system_configuration_id),
	CONSTRAINT system_config_one_row CHECK (system_configuration_id)

);
-- ddl-end --
ALTER TABLE public.system_configuration OWNER TO conductorr;
-- ddl-end --

-- object: public.sonarr_configuration | type: TABLE --
-- DROP TABLE IF EXISTS public.sonarr_configuration CASCADE;
CREATE TABLE public.sonarr_configuration (
	sonarr_configuration_id bool NOT NULL DEFAULT true,
	sonarr_url text,
	sonarr_api_key varchar(128),
	sonarr_category text,
	CONSTRAINT sonarr_config_pk PRIMARY KEY (sonarr_configuration_id),
	CONSTRAINT sonarr_config_one_row CHECK (sonarr_configuration_id)

);
-- ddl-end --
ALTER TABLE public.sonarr_configuration OWNER TO conductorr;
-- ddl-end --

-- object: public.radarr_configuration | type: TABLE --
-- DROP TABLE IF EXISTS public.radarr_configuration CASCADE;
CREATE TABLE public.radarr_configuration (
	radarr_configuration_id bool NOT NULL DEFAULT true,
	radarr_url text,
	radarr_api_key varchar(128),
	radarr_category text,
	CONSTRAINT radarr_config_pk PRIMARY KEY (radarr_configuration_id),
	CONSTRAINT radarr_config_one_row CHECK (radarr_configuration_id)

);
-- ddl-end --
ALTER TABLE public.radarr_configuration OWNER TO conductorr;
-- ddl-end --

-- object: public.filbot_configuration | type: TABLE --
-- DROP TABLE IF EXISTS public.filbot_configuration CASCADE;
CREATE TABLE public.filbot_configuration (
	filebot_configuration_id bool NOT NULL DEFAULT true,
	fb_output_dir text,
	fb_subtitles_locale varchar(16),
	fb_action varchar(16),
	fb_exclude_list text,
	fb_amc_log text,
	fb_kodi text,
	fb_plex text,
	fb_emby text,
	fb_pushover text,
	fb_discord text,
	fb_gmail text,
	fb_mail text,
	fb_mailto text,
	fb_report_error bool DEFAULT true,
	fb_store_report bool DEFAULT true,
	fb_extract_folder text,
	fb_exec text,
	fb_ignore text,
	fb_artwork bool DEFAULT true,
	fb_skip_extract bool DEFAULT true,
	fb_delete_after_extract bool DEFAULT true,
	fb_clean bool DEFAULT true,
	fb_unsorted bool DEFAULT true,
	fb_home_dir text,
	fb_extras bool DEFAULT true,
	fb_deployment_name text,
	fb_namespace text,
	CONSTRAINT filebot_config_pk PRIMARY KEY (filebot_configuration_id),
	CONSTRAINT filebot_config_one_row CHECK (filebot_configuration_id)

);
-- ddl-end --
ALTER TABLE public.filbot_configuration OWNER TO conductorr;
-- ddl-end --

-- object: public.plex_configuration | type: TABLE --
-- DROP TABLE IF EXISTS public.plex_configuration CASCADE;
CREATE TABLE public.plex_configuration (
	plex_configuration_id bool NOT NULL DEFAULT true,
	plex_namespace text,
	plex_deployment_name text,
	plex_auth_token text,
	plex_base_url text,
	CONSTRAINT plex_config_pk PRIMARY KEY (plex_configuration_id),
	CONSTRAINT plex_config_one_row CHECK (plex_configuration_id)

);
-- ddl-end --
ALTER TABLE public.plex_configuration OWNER TO conductorr;
-- ddl-end --



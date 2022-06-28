-- create in postgres_server
-- load extension first time after install
-- created should in target database, by user postgres
CREATE EXTENSION mysql_fdw;

-- create server object
CREATE SERVER shopeexb_mysql_server
    FOREIGN DATA WRAPPER mysql_fdw
    OPTIONS (host '10.20.11.117', port '30320');

-- grant other users
GRANT USAGE ON FOREIGN SERVER shopeexb_mysql_server to analysis;


SET SESSION AUTHORIZATION analysis;
-- created by user sensors_dev
-- create user mapping
DROP user mapping for  analysis server shopeexb_mysql_server;
CREATE USER MAPPING FOR analysis -- db.user
    SERVER shopeexb_mysql_server
    OPTIONS (username '', password '');

CREATE SCHEMA newtond_shopee_xb;

-- Import table definitions from a remote schema foreign_films on server film_server, creating the foreign tables in local schema films
IMPORT FOREIGN SCHEMA newtond_shopee_xb FROM SERVER shopeexb_mysql_server INTO newtond_shopee_xb;
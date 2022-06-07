#!/bin/bash
# Original script https://github.com/MartinKaburu/docker-postgresql-multiple-databases

set -e
set -u

function create_user_and_database() {
	local database=$(echo $1)
	echo "  Creating user and database '$database'"
	psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
	    CREATE DATABASE $database;
	    GRANT ALL PRIVILEGES ON DATABASE $database TO $POSTGRES_USER;
EOSQL
}

function create_procedures() {
	local database=$(echo $1)
	echo "  Creating procedures '$database'"
	psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" -d $database <<-EOSQL
	    CREATE OR REPLACE PROCEDURE truncate_tables(_schema character varying) AS \$\$
      declare
          r record;
      begin
        FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = current_schema() and tablename not in ('goose_db_version')) LOOP
          EXECUTE 'TRUNCATE TABLE ' || quote_ident(r.tablename) || ' RESTART IDENTITY CASCADE';
        END LOOP;
      end;
      \$\$
      LANGUAGE plpgsql;
EOSQL
}

if [ -n "$POSTGRES_MULTIPLE_DATABASES" ]; then
	echo "Multiple database creation requested: $POSTGRES_MULTIPLE_DATABASES"
	for db in $(echo $POSTGRES_MULTIPLE_DATABASES | tr ',' '\n'); do
		create_user_and_database $db
		create_procedures $db
	done
	echo "Multiple databases created"
fi
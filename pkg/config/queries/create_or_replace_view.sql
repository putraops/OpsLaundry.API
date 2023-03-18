CREATE OR REPLACE FUNCTION public.ufn_create_or_replace_view(
	_schema_name text,
	_table_name text)
    RETURNS boolean
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
DECLARE
	_sql TEXT;
	_column RECORD;
BEGIN

    _sql := ' DROP VIEW IF EXISTS ' || _schema_name || '.vw_' || _table_name;
	EXECUTE _sql;
		
	_sql := ' CREATE OR REPLACE VIEW ' || _schema_name || '.vw_' || _table_name || ' AS' || E'\n' || '';		
	_sql := _sql || ' SELECT' || E'\n' || '';
	FOR _column IN 
		SELECT *
		FROM information_schema.columns 
		WHERE table_schema = _schema_name
			AND table_name = _table_name
	LOOP
		_sql := _sql || 'r.' || _column.column_name || ', ' || E'\n' || '';
	END LOOP;
	_sql := _sql || 'o1.name AS organization_name, ' || E'\n' || '';

	_sql := _sql || 'CASE WHEN u1.last_name IS NULL OR u1.last_name = '''' THEN u1.first_name ELSE concat(u1.first_name, '' '', u1.last_name) END AS record_created, ' || E'\n' || '';
	_sql := _sql || 'CASE WHEN u2.last_name IS NULL OR u2.last_name = '''' THEN u2.first_name ELSE concat(u2.first_name, '' '', u2.last_name) END AS record_updated, ' || E'\n' || '';
	_sql := _sql || 'CASE WHEN u3.last_name IS NULL OR u3.last_name = '''' THEN u3.first_name ELSE concat(u3.first_name, '' '', u3.last_name) END AS record_submitted, ' || E'\n' || '';
	_sql := _sql || 'CASE WHEN u4.last_name IS NULL OR u4.last_name = '''' THEN u4.first_name ELSE concat(u4.first_name, '' '', u4.last_name) END AS record_approved ' || E'\n' || '';
	-- _sql := _sql || 't1.team_name AS record_owning_team';
	_sql := _sql || 'FROM ' || _schema_name || '.' || _table_name || ' r' || E'\n' || '';
	_sql := _sql || 'LEFT JOIN ' || _schema_name || '.organization o1 ON o1.id = r.organization_id' || E'\n' || '';
	_sql := _sql || 'LEFT JOIN ' || _schema_name || '.application_user u1 ON u1.id = r.created_by' || E'\n' || '';
	_sql := _sql || 'LEFT JOIN ' || _schema_name || '.application_user u2 ON u2.id = r.updated_by' || E'\n' || '';
	_sql := _sql || 'LEFT JOIN ' || _schema_name || '.application_user u3 ON u3.id = r.submitted_by' || E'\n' || '';
	_sql := _sql || 'LEFT JOIN ' || _schema_name || '.application_user u4 ON u4.id = r.approved_by' || E'\n' || '';

	-- _sql := _sql || ' LEFT JOIN ' || _schema_name || '.team t1 ON t1.id = r.owner_id';
	EXECUTE _sql;
	
	_sql := ' ALTER TABLE ' || _schema_name || '.vw_' || _table_name;
	_sql := _sql || ' OWNER TO postgres';
	EXECUTE _sql;

	RETURN true;
END;
$BODY$;

ALTER FUNCTION public.ufn_create_or_replace_view(text, text)
    OWNER TO postgres;

GRANT EXECUTE ON FUNCTION public.ufn_create_or_replace_view(text, text) TO PUBLIC;

GRANT EXECUTE ON FUNCTION public.ufn_create_or_replace_view(text, text) TO postgres;
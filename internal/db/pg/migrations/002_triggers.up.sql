CREATE OR REPLACE FUNCTION validate_attributes()
RETURNS TRIGGER AS $$
DECLARE
	required_table_name TEXT;
	place_type RECORD;
	place_ids_count INT;
BEGIN
	SELECT attributes_table_name
	INTO required_table_name
	FROM place_types
	WHERE id = (SELECT type_id FROM places WHERE id = NEW.place_id);

	IF required_table_name IS NULL THEN
		RAISE EXCEPTION 'required table name not found for place id %', NEW.place_id;
	END IF;

	IF TG_TABLE_NAME <> required_table_name THEN
		RAISE EXCEPTION 'attributes of place with ID % must be in % table', NEW.place_id, required_table_name;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION add_attributes_validator()
RETURNS TRIGGER AS $$
DECLARE
	query TEXT;
BEGIN
	query := format(
		'CREATE OR REPLACE TRIGGER validate_%I_trigger
		 BEFORE INSERT OR UPDATE ON %I
		 FOR EACH ROW EXECUTE FUNCTION validate_attributes();',
		NEW.attributes_table_name,
		NEW.attributes_table_name
	);

	EXECUTE query;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER add_attributes_validator_trigger
BEFORE INSERT OR UPDATE ON place_types
FOR EACH ROW EXECUTE FUNCTION add_attributes_validator();

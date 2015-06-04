--
-- This is example schema for sqlite3
--
-- TODO: Schema should be generated some other way
--       or database table should be initialize
--       some other way.
--

DROP TABLE IF EXISTS acl_rules;
CREATE TABLE acl_rules (
	id INTEGER PRIMARY KEY,
	actor TEXT,
	target TEXT,
	scope TEXT
);

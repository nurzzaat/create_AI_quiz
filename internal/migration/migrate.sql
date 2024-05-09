create table users(
	id serial primary key,
	email text default '',
	password text default '',
	firstname text default '',
	lastname text default '',
	createdat text default '',
	roleid integer default 0
);
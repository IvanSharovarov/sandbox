create table users
(
	id serial not null,
	name varchar,
	email varchar not null,
	password varchar not null,
	created_at timestamp,
	updated_at timestamp
);

create unique index users_email_uindex
	on users (email);

create unique index users_id_uindex
	on users (id);

alter table users
	add constraint users_pk
		primary key (id);


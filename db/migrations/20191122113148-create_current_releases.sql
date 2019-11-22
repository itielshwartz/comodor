-- +migrate Up
-- +migrate Up
create table comodor.releases
(
    row_id         serial not null
        constraint releases_pk
            primary key,
    name           text   not null,
    namespace      text   not null,
    cluster        text   not null,
    status         text   not null,
    created_at     date   not null,
    revision       integer,
    schema_version integer default '-1'::integer
);

alter table comodor.releases
    owner to postgres;

create unique index releases_row_id_uindex
    on comodor.releases (row_id);

create unique index releases_cluster_name_namespace_uindex
    on comodor.releases (cluster, name, namespace);

create table comodor.services
(
	row_id serial not null
		constraint services_pk
			primary key,
	name text not null,
	type text not null,
	cluster_ip text not null,
	external_ip text not null,
	ports text not null,
	created_at date not null
);

create table comodor.pods
(
	row_id serial not null
		constraint pods_pk
			primary key,
	name text not null,
	ready integer,
	total integer,
	status text not null,
	restarts integer not null,
	created_at date not null
);

create table comodor.statefulsets
(
	row_id serial not null
		constraint statefulsets_pk
			primary key,
	name text not null,
	ready integer,
	total integer,
	created_at date not null
);

-- +migrate Down

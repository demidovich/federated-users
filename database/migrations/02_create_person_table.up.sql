create table person (
    uuid uuid primary key,
    federation_uuid uuid,
    attrs jsonb
);

create index person_federation_uuid_idx on person (federation_uuid);
create index person_attrs_idx on person using gin (attrs jsonb_path_value_ops);

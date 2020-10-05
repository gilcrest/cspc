create table lookup.county_lkup
(
    county_id uuid not null
        constraint county_lkup_pk
            primary key,
    state_prov_id uuid not null,
    county_cd varchar not null,
    county_name varchar not null,
    latitude_average varchar,
    longitude_average varchar,
    create_username varchar not null,
    create_timestamp timestamp with time zone default CURRENT_TIMESTAMP not null,
    update_username varchar not null,
    update_timestamp timestamp with time zone default CURRENT_TIMESTAMP not null,
    constraint county_lkup_fk
        foreign key (state_prov_id) references lookup.state_prov_lkup (state_prov_id)
);

comment on table lookup.county_lkup is 'County lookup table stores counties for a country/state. A county is a geographical region of a country used for administrative or other purposes, in certain modern nations.';
comment on column lookup.county_lkup.county_id is 'Arbitrary unique ID for primary key.';
comment on column lookup.county_lkup.state_prov_id is 'The state or province ID for the region of the country. This is an FK to lookup.state_prov_lkup';
comment on column lookup.county_lkup.county_cd is 'A unique code for the county. For the US, this are Federal Information Processing System (FIPS) Codes for States and Counties. FIPS codes are numbers which uniquely identify geographic areas.';
comment on column lookup.county_lkup.latitude_average is 'Average latitude of the county territory';
comment on column lookup.county_lkup.longitude_average is 'Average longitude of the state or province territory';

alter table lookup.county_lkup owner to postgres;

create unique index county_lkup_ui2
    on lookup.county_lkup (county_cd);


create table lookup.state_prov_lkup
(
    state_prov_id uuid not null
        constraint state_prov_lkup_pk
            primary key,
    country_alpha_2_cd varchar(2) not null
        constraint state_prov_lkup_fk
            references lookup.country_lkup (country_alpha_2_cd),
    state_prov_cd varchar not null,
    state_name varchar not null,
    state_fips_cd integer,
    latitude_average varchar,
    longitude_average varchar,
    create_username varchar not null,
    create_timestamp timestamp with time zone default CURRENT_TIMESTAMP not null,
    update_username varchar not null,
    update_timestamp timestamp with time zone default CURRENT_TIMESTAMP not null,
    constraint state_prov_lkup_pk_2
        unique (country_alpha_2_cd, state_prov_cd)
);

comment on table lookup.state_prov_lkup is 'State and Province Code by Country Lookup';
comment on column lookup.state_prov_lkup.state_prov_cd is 'The state or province code for the region of the country';
comment on column lookup.state_prov_lkup.state_name is 'State or Province name';
comment on column lookup.state_prov_lkup.latitude_average is 'Average latitude of the state or province territory';
comment on column lookup.state_prov_lkup.longitude_average is 'Average longitude of the state or province territory';
comment on column lookup.state_prov_lkup.state_fips_cd is 'US Only. FIPS codes are numbers which uniquely identify geographic areas.  The number of digits in FIPS codes vary depending on the level of geography.  State-level FIPS codes have two digits, county-level FIPS codes have five digits of which the first two are the FIPS code of the state to which the county belongs.';

alter table lookup.state_prov_lkup owner to postgres;

create unique index state_prov_lkup_ui1
    on lookup.state_prov_lkup (state_fips_cd);


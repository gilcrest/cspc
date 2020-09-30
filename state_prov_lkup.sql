create table lookup.state_prov_cd_lkup
(
    state_prov_id uuid not null
        constraint state_prov_cd_lkup_pk
            primary key,
    country_alpha_2_cd varchar(2) not null
        constraint country_cd_lkup_fk
            references lookup.country_cd_lkup (country_alpha_2_cd),
    state_prov_cd varchar not null,
    state_name varchar not null,
    latitude_average varchar,
    longitude_average varchar,
    create_username varchar not null,
    create_timestamp timestamp with time zone default CURRENT_TIMESTAMP not null,
    update_username varchar not null,
    update_timestamp timestamp with time zone default CURRENT_TIMESTAMP not null,
    constraint state_prov_cd_lkup_pk_2
        unique (country_alpha_2_cd, state_prov_cd)
);

comment on table lookup.state_prov_cd_lkup is 'State and Province Code by Country Lookup';
comment on column lookup.state_prov_cd_lkup.state_prov_cd is 'The state or province code for the region of the country';
comment on column lookup.state_prov_cd_lkup.state_name is 'State or Province name';
comment on column lookup.state_prov_cd_lkup.latitude_average is 'Average latitude of the state or province territory';
comment on column lookup.state_prov_cd_lkup.longitude_average is 'Average longitude of the state or province territory';

alter table lookup.state_prov_cd_lkup owner to postgres;


create schema lookup;

comment on schema lookup is 'Reference tables that can apply to multiple different applications, such as a country code lookup.';

alter schema lookup owner to postgres;

create table lookup.country_lkup
(
    country_id uuid not null
        constraint country_lkup_pk
            primary key,
    country_alpha_2_cd varchar(2) not null,
    country_alpha_3_cd varchar(3) not null,
    country_un_m49_cd bigint not null,
    country_name varchar(100),
    latitude_average varchar(20),
    longitude_average varchar(20),
    create_username varchar not null,
    create_timestamp timestamp with time zone default CURRENT_TIMESTAMP not null,
    update_username varchar not null,
    update_timestamp timestamp with time zone default CURRENT_TIMESTAMP not null
);

comment on column lookup.country_lkup.country_id is 'Unique ID for the country';
comment on column lookup.country_lkup.country_alpha_2_cd is 'ISO 3166-1 Alpha 2 Code';
comment on column lookup.country_lkup.country_alpha_3_cd is 'ISO 3166-1 Alpha 3 Code';
comment on column lookup.country_lkup.country_un_m49_cd is 'Standard Country or Area Codes for Statistical Use (Series M, No. 49) is a standard for area codes used by the United Nations for statistical purposes, developed and maintained by the United Nations Statistics Division.';
comment on column lookup.country_lkup.latitude_average is 'Average latitude of the country territory';
comment on column lookup.country_lkup.longitude_average is 'Average longitude of the country territory';

alter table lookup.country_lkup owner to postgres;

create unique index country_lkup_country_alpha_2_cd_uindex
    on lookup.country_lkup (country_alpha_2_cd);

create unique index country_lkup_country_alpha_3_cd_uindex
    on lookup.country_lkup (country_alpha_3_cd);

create unique index country_lkup_country_un_m49_cd_uindex
    on lookup.country_lkup (country_un_m49_cd);

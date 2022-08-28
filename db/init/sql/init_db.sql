create table users
(
    id         varchar(36)                                        not null
        constraint users_pkc
            primary key,
    user_name  varchar(50)                                        not null,
    password   varchar(1000)                                      not null,
    email      varchar(255)                                       not null,
    created_at timestamp with time zone default CURRENT_TIMESTAMP not null,
    updated_at timestamp with time zone default CURRENT_TIMESTAMP not null,
    is_del     smallint                 default 0                 not null,
    version    bigint                   default 1                 not null
);

alter table users
    owner to common;

create index users_ix1
    on users (is_del);

create unique index users_ix2
    on users (user_name);



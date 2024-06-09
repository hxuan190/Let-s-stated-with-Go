create table user
(
    id         char(36)     not null
        primary key,
    first_name varchar(255) not null,
    last_name  varchar(255) not null,
    email      varchar(254) not null,
    pass_word  char(60)     not null,
    salt       char(29)     not null,
    role       int          not null,
    created_at timestamp    not null,
    updated_at timestamp    not null,
    constraint email
        unique (email)
);


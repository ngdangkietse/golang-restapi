create table tbl_user
(
    id         int primary key auto_increment,
    name       varchar(255) character set utf8 not null,
    age        int,
    address    text,
    email      varchar(100)                    not null,
    password   varchar(100)                    not null,
    role_id    int                             not null,
    created_at timestamp                       not null default current_timestamp,
    updated_at timestamp                       not null on update current_timestamp
)
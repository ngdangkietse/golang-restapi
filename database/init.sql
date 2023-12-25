create table tbl_user
(
    id         int primary key auto_increment,
    name       varchar(255) character set utf8 not null,
    age        int,
    address    text,
    created_at timestamp                       not null default current_timestamp,
    updated_at timestamp                       not null on update current_timestamp
)
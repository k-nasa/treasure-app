-- +goose Up
create table tags(
  id int(10) UNSIGNED auto_increment,
  name text not null,
  created_at timestamp not null default NOW(),
  updated_at timestamp not null default NOW() ON UPDATE CURRENT_TIMESTAMP,

  primary key(id)
) ENGINE=InnoDB default charset=utf8mb4;

insert into tag(name) values("hoge")
insert into tag(name) values("nya-n")

-- +goose Down
drop table tags;

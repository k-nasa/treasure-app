-- +goose Up
create table comments(
  id int(10) UNSIGNED auto_increment,
  user_id int(10) UNSIGNED not null,
  article_id int(10) UNSIGNED not null,
  body text not null,
  created_at timestamp not null default NOW(),
  updated_at timestamp not null default NOW() ON UPDATE CURRENT_TIMESTAMP,

  primary key(id),
  foreign key comment_fk_user (user_id) references user(id),
  foreign key comment_fk_article (article_id) references article(id)
) ENGINE=InnoDB default charset=utf8mb4;

-- +goose Down
drop table comments;


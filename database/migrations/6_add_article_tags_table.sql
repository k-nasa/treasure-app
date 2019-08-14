-- +goose Up
create table article_tag(
  article_id int(10) UNSIGNED not null,
  tag_id int(10) UNSIGNED not null,
  created_at timestamp not null default CURRENT_TIMESTAMP,

  primary key(article_id, tag_id),
  foreign key article_tag_fk_article (article_id) references article(id),
  foreign key article_tag_fk_tag (tag_id) references tags(id)
) ENGINE=InnoDB default charset=utf8mb4;


-- +goose Down
drop table article_tag;

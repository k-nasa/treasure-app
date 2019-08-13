-- +goose Up
alter table article add user_id int(10) UNSIGNED;
alter table article add constraint article_fk_user foreign key (user_id) references user(id);

-- +goose Down
alter table article drop foreign key article_fk_user;
alter table article drop column user_id;


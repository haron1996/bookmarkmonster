-- +goose Up
DROP TABLE IF EXISTS bookmark_tag;

CREATE TABLE bookmark_tag (
    bookmark_id text not null references bookmark(id) on delete cascade,
    tag_id text not null references tag(id) on delete cascade,
    constraint bookmark_tag_pk primary key(bookmark_id, tag_id),
    constraint bookmark_tag_unique unique(bookmark_id, tag_id)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS bookmark_tag;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

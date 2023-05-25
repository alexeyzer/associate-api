-- +goose Up
-- +goose StatementBegin
CREATE TABLE stimus_word(
    id SERIAL,
    name text not null
);

CREATE UNIQUE INDEX stimus_word_uindx on stimus_word(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE stimus_word;
-- +goose StatementEnd

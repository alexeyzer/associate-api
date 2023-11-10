-- +goose Up
-- +goose StatementBegin
CREATE TABLE associate_word(
    id SERIAL,
    name text not null
);

CREATE UNIQUE INDEX  associate_word_uindx on associate_word(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE associate_word;
-- +goose StatementEnd

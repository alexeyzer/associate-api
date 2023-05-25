-- +goose Up
-- +goose StatementBegin
CREATE TABLE experiment_result(
    id SERIAL,
    experiment_id bigint not null,
    user_id bigint not null,
    is_anonymous bool not null,
    stimus_word_id bigint not null,
    assotiation_word_id bigint,
    time_spend bigint
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE experiment_result;
-- +goose StatementEnd

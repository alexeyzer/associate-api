-- +goose Up
-- +goose StatementBegin
CREATE TABLE experiment_result_calculated(
    experiment_id bigint not null,
    stimus_word_id bigint not null,
    assotiation_word_id bigint not null,
    PRIMARY KEY(experiment_id, stimus_word_id, assotiation_word_id),
    amount bigint not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE experiment_result_calculated;
-- +goose StatementEnd

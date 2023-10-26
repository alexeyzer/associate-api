-- +goose Up
-- +goose StatementBegin
CREATE TYPE experiment_status AS ENUM ('CREATED','IN_PROGRESS', 'CANCELED', 'DONE');

CREATE TABLE experiment(
    id SERIAL,
    name text not null,
    description text,
    creator_id bigint not null,
    conducducted_amount bigint,
    required_amount bigint,
    status experiment_status not null,
    experiment_stimuses jsonb not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE experiment;

DROP TYPE experiment_status;
-- +goose StatementEnd

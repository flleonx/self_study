INSERT INTO orders (
    aggregate_id,
    event_type,
    event_data,
    version
) VALUES (
    :aggregate_id,
    :event_type,
    :event_data,
    :version
)
RETURNING *;

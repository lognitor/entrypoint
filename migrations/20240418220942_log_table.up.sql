create table if not exists log_request
(
    id          UUID default generateUUIDv4(),
    token       String,
    type        ENUM('text/plain' = 1, 'text/json' = 2, 'binary/download' = 3, 'binary/view' = 4),
    level       ENUM('DEBUG' = 0, 'INFO' = 1, 'WARN' = 2, 'ERROR' = 3, 'FATAL' = 4, 'PANIC' = 5),
    prefix      String,
    ip          Nullable(IPv4),
    agent       Nullable(String),
    message     String,
    trace       Array(Map(String, String)),
    source      Nullable(String),
    created_at  DATETIME
)
    engine = MergeTree()
        PRIMARY KEY id
        PARTITION BY toYYYYMM(created_at)
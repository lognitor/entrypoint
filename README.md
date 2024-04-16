## Entrypoint for write log data to kafka and transport to kafka worker

### Environment variables

| ENV                     | Demo                            | Default | Description                      |
|-------------------------|---------------------------------|---------|----------------------------------|
| GRPC_PORT               | 50051                           | 50051   | Port for GRPC server             |
| GRPC_ALLOW_IP           | 0.0.0.0                         | 0.0.0.0 | Allow ip address                 |
| HTTP_PORT               | 8080                            | 8080    | Port for http server             |
| HTTP_ALLOW_IP           | 0.0.0.0                         | 0.0.0.0 | Allow ip address                 |
| KAFKA_BROKERS           | 100.1.0.2:19092,100.1.0.3:19092 | -       | Brokers address "," split        |
| KAFKA_TOPIC             | log.request                     | -       | Kafka topic for writting         |
| CLICKHOUSE_HOSTS        | localhost:9092,localhost:9091   | -       | Hosts list for clickhouse        |
| CLICKHOUSE_USER         | dbuser                          | -       | Username for clickhouse          |
| CLICKHOUSE_PASSWORD     | dbpassoword                     | -       | Database password for clickhouse |
| CLICKHOUSE_DATABASE     | dbname                          | -       | Database name for clickhouse     |

### Worker service

This service worker for writing data from kafka to clickhouse

```shell
docker run -d --name log-writer \
    -e KAFKA_BROKERS=<your_brokers_address>
    -e KAFKA_TOPIC=<your_topic_name>
    -e CLICKHOUSE_HOSTS=<>
    -e CLICKHOUSE_USER=<>
    -e CLICKHOUSE_PASSWORD=<>
    -e CLICKHOUSE_DATABASE=<>
    ghcr.io/lognitor/entrypoint:latest \
    /app/worker
```

### Entrypoint service

This service for writing data to kafka topic from HTTP or GRPC

```shell
docker run -d --name log-writer \
    -e KAFKA_BROKERS=<your_brokers_address>
    -e KAFKA_TOPIC=<your_topic_name>
    -p 50051:50051 \
    -p 8080:8080 \
    ghcr.io/lognitor/entrypoint:latest
```


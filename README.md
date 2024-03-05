## Entrypoint for write log data to kafka

### Environment variables
| ENV           | Demo                            | Default | Description               |
|---------------|---------------------------------|---------|---------------------------|
| GRPC_PORT     | 50051                           | 50051   | Port for GRPC server      |
| GRPC_ALLOW_IP | 0.0.0.0                         | 0.0.0.0 | Allow ip address          |
| HTTP_PORT     | 8080                            | 8080    | Port for http server      |
| HTTP_ALLOW_IP | 0.0.0.0                         | 0.0.0.0 | Allow ip address          |
| KAFKA_BROKERS | 100.1.0.2:19092,100.1.0.3:19092 | -       | Brokers address "," split |
| KAFKA_TOPIC   | in.log                          | -       | Kafka topic for writting  |


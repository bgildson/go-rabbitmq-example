# Go RabbitMQ Example

A small example of how to use RabbitMQ with GoLang to produce and consume messages with AMQP lib

<p align="center">
  <img src="./example.gif" width="1024" alt="example" />
</p>

## How to run the example

- Ensure that you have GoLang properly installed, like described [here](https://golang.org/doc/install).

- Ensure that you have Docker properly installed, like described [here](https://docs.docker.com/get-docker/).

```sh
# clone this repository
git clone https://github.com/bgildson/go-rabbitmq-example.git

# step into the repository
cd go-rabbitmq-example

# execute the follow command and wait the service start
docker run --rm -p 5672:5672 -e RABBITMQ_DEFAULT_USER=guest -e RABBITMQ_DEFAULT_PASS=guest -v ${PWD}/rabbit.definitions.json:/rabbit.definitions.json:ro --name rabbitmq rabbitmq:3.8.4-management

# in another terminal, execute the follow command to restore the rabbitmq schema used in the example
docker exec rabbitmq rabbitmqadmin import /rabbit.definitions.json

# clone the .env.example with default environment variables
cp .env.example .env

# install the golang dependencies
go mod download

# now you have all the base configurations

# in the terminal, execute the follow command to start a consumer in catalogs_que queue
go run consumer/main.go -queue=catalogs_que

# in another terminal, execute the follow command to start a producer in products_ex exchange
go run producer/main.go -exchange=products_ex

# all the sended commands in the producer terminal should be printed in consumer terminal
```

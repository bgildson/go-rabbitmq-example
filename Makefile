consumer-%:
	go run consumer/main.go --queue=$(*)_que

producer:
	go run producer/main.go --exchange=products_ex

import:
	docker-compose exec rabbitmq rabbitmqadmin import /rabbit.definitions.json

.PHONY: consumer producer import

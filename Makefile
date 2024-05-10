go-service:
	cd ./go-service && go run main.go

nodejs-consumer-1:
	cd ./nodejs-consumer-1 && npm run start

nodejs-consumer-2:
	cd ./nodejs-consumer-2 && npm run start

.PHONY: go-service nodejs-consumer-1 nodejs-consumer-2
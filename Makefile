# build the application
build:
	go mod tidy
	go build -o shipment-calculator main.go

# start the application
start:
	./shipment-calculator -p 8085 -l ./priv/sizes.json start

docker-up:
	docker-compose -p re-packs up -d

docker-down:
	docker-compose -p re-packs down
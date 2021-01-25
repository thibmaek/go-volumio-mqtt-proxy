include Makefile.vars

LDFLAGS="-X 'main.broker=$(MQTT_BROKER)' -X 'main.user=$(MQTT_USER)' -X 'main.password=$(MQTT_PASSWORD)'"

run:
	go run -ldflags=$(LDFLAGS) .

build:
	go build -ldflags=$(LDFLAGS) .

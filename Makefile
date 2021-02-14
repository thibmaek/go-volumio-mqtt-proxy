include Makefile.vars

LDFLAGS="-X 'main.broker=$(MQTT_BROKER)' -X 'main.user=$(MQTT_USER)' -X 'main.password=$(MQTT_PASSWORD)'"

.DEFAULT_GOAL := build

clean:
	rm -rf bin

run:
	go run -ldflags=$(LDFLAGS) .

build_x64:
	GOOS=darwin go build -ldflags=$(LDFLAGS) -o bin/go-volumio-mqtt_mac .
	GOOS=linux go build -ldflags=$(LDFLAGS) -o bin/go-volumio-mqtt_linux .

build_arm:
	GOOS=linux GOARCH=arm GOARM=6 go build -ldflags=$(LDFLAGS) -o bin/go-volumio-mqtt_arm .

build: build_x64 build_arm

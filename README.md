# Go Volumio Proxy

## Install

```shell
$ go mod download
```

## Usage

Add a `Makefile.vars` file with the required MQTT credentials:

```make
MQTT_BROKER=localhost
MQTT_USER=user
MQTT_PASSWORD=password
```

Then run/build the application with make

```shell
$ make run
```

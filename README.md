go osc capture
===
A simple CLI to record/replay OSC(Open Sound Control).

## Download
Download [the newest binary](https://github.com/asus4/go-osc-capture/releases)



## Usage
```
# Record to capture.csv
osccap record -port 7000 osc_capture.csv

# Listen from multicast addreass and record to the csv file
osccap record -multicast 225.6.7.8 -port 7000 osc_capture.csv
```

```
# Play the capture
osccap play -addr 127.0.0.1 -port 8000 osc_capture.csv
```

## Dependences
```
go get github.com/urfave/cli
go get github.com/hypebeast/go-osc/osc
```


## Build
```
./build.sh
```
go-osc-capture
===
A simple CLI to record/replay osc (open sound control) packets.


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
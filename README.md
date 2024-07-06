# pcap2go - Create go test-data from pcap files

This is (very) small utility program that convert pcap files to go
`[][]byte` declarations. It's intended for unit-tests to test packet
parsing with real packets.

A normal use-case is to append a `[][]byte` variable with test-data to
your unit test.

```
# (capture data for instance with "tcpdump -ni eth0 -w my-capture.pcap udp")
pcap2go -variable udpPackets my-capture.pcap >> packet_test.go
```

Test with:
```
go build -o pcap2go pcap2go.go
./pcap2go udp.pcap
```

## HW offload

To capture fragments it is often necessary to turn off HW offload
(even on virtio devices):

```
ethtool -K eth1 gro off gso off tso off
```

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	variable := flag.String("variable", "pcap_packets", "Name of the variable")
	eth := flag.Bool("eth", false, "Include the eth header")
	cap := flag.Int("cap", 0, "If >0 cap packet length")
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println(`
pcap2go [options] <file>

Read a pcap-file and emit packets declared as go []byte
slices. Intended for including captured packets in unit tests.
`)
		flag.PrintDefaults()
		return
	}
	if err := readFile(flag.Arg(0), *variable, *eth, *cap); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readFile(file, variable string, includeEth bool, cap int) error {
	handle, err := pcap.OpenOffline(file)
	if err != nil {
		return err
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Printf("var %s = [][]byte{\n", variable)
	for packet := range packetSource.Packets() {
		pdata := packet.Data()
		if ! includeEth {
			pdata = pdata[14:]
		}
		if cap > 0 && cap < len(pdata) {
			pdata = pdata[:cap]
		}
		fmt.Printf("\t[]byte{\n")
		printBytes(pdata)
		fmt.Printf("\t},\n")
	}
	fmt.Printf("}\n")
	return nil
}
func printBytes(b []byte) {
	for i := 0; i < len(b); i++ {
		if (i % 16) == 0 {
			fmt.Printf("\t\t")
		}
		if (i % 16) == 15 {
			fmt.Printf("0x%02x,\n", b[i])
		} else {
			fmt.Printf("0x%02x, ", b[i])
		}
	}
	if (len(b) % 16) != 0 {
		fmt.Printf("\n")
	}
}

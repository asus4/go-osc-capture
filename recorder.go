package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/hypebeast/go-osc/osc"
)

func osc2csv(time float32, msg *osc.Message) []string {
	cells := []string{
		fmt.Sprintf("%.6f", time),
		msg.Address,
	}

	for _, arg := range msg.Arguments {
		switch arg.(type) {
		case float32, float64:
			cells = append(cells, fmt.Sprintf("%.6f", arg))
		default:
			cells = append(cells, fmt.Sprint(arg))
		}

	}
	return cells
}

func readRoutine(conn *net.UDPConn, writer *csv.Writer, isClose chan bool) error {
	startTime := time.Now()

	server := &osc.Server{}
	for {
		// Check is closed
		select {
		case _close := <-isClose:
			if _close {
				fmt.Println("server is closed")
				return nil
			}
		default:
		}

		// Receive packet
		packet, err := server.ReceivePacket(conn)
		if err != nil {
			fmt.Println("Server error: " + err.Error())
			os.Exit(1)
		}
		if packet == nil {
			continue
		}

		elapsed := float32(time.Now().Sub(startTime)) / float32(time.Second)

		switch packet.(type) {
		default:
			fmt.Println("Unknow packet type")
			continue

		case *osc.Message:
			writer.Write(osc2csv(elapsed, packet.(*osc.Message)))
			fmt.Println(packet.(*osc.Message))

		case *osc.Bundle:
			bundle := packet.(*osc.Bundle)
			for _, message := range bundle.Messages {
				writer.Write(osc2csv(elapsed, message))
				fmt.Println(message)
			}
		}
	}
}

func recorder(port uint, path string, multicastAddr string) error {

	// make File
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	err = file.Truncate(0)
	if err != nil {
		return err
	}
	writer := csv.NewWriter(file)

	// open server
	var conn *net.UDPConn
	multicastIP := net.ParseIP(multicastAddr)
	if multicastIP.IsMulticast() {
		addr := &net.UDPAddr{
			IP:   multicastIP,
			Port: int(port),
		}
		conn, err = net.ListenMulticastUDP("udp", nil, addr)
	} else {
		addr := &net.UDPAddr{
			IP:   net.IPv4zero,
			Port: int(port),
		}
		conn, err = net.ListenUDP("udp", addr)
	}
	if err != nil {
		return err
	}
	defer conn.Close()
	// }

	// start go routine
	fmt.Println("Press \"q\" to exit")
	serverIsClose := make(chan bool)
	go readRoutine(conn, writer, serverIsClose)

	// read exit status
	reader := bufio.NewReader(os.Stdin)
	for {
		c, err := reader.ReadByte()
		if err != nil || c == 'q' {
			serverIsClose <- true
			return err
		}
	}

}

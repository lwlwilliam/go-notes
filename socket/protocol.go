// 对接收的各层协议消息进行解剖
package main

import (
	"fmt"
	"strings"
)

func main() {
	header("respEthernet")

	respEthernet := map[string][]int{
		"destination": {0x00, 0x0c, 0x29, 0x3a, 0x66, 0x8c},
		"source":      {0xbc, 0x54, 0x36, 0xce, 0x03, 0x72},
		"type":        {0x0800},
	}

	for k, v := range respEthernet {
		for subk, subv := range v {
			switch k {
			case "destination", "source":
				if subk == 0 {
					s := []string{}
					for _, subvv := range v {
						s = append(s, fmt.Sprintf("%x", subvv))
					}
					mac := strings.Join(s, ":")
					fmt.Printf("%-20s:%s\n", k, mac)
				}
			case "type":
				fmt.Printf("%-20s:%#x\n", k, subv)
			}
		}
	}

	header("respIP")

	respIP := map[string][]int{
		"version":                     {0x4},
		"headerLength":                {0x5},
		"differentiatedServicesField": {0x00},
		"totalLength":                 {0x0264},
		"identification":              {0x0bb7},
		"flagsAndFragmentOffset":      {0x4000}, // TODO: wireshark 这里的 Flags 和 Fragment offset 是有重合的
		"timeToLive":                  {0x40},
		"protocol":                    {0x06},
		"headerChecksum":              {0xfcac},
		"source":                      {0xc0, 0xa8, 0x0a, 0x44},
		"destination":                 {0x0e, 0x1d, 0x57, 0x27},
	}

	for k, v := range respIP {
		for subk, subv := range v {
			switch k {
			case "version", "headerLength", "", "flagsAndFragmentOffset", "protocol":
				fmt.Printf("%-30s:%b\n", k, subv)
			case "differentiatedServicesField", "identification", "headerChecksum":
				fmt.Printf("%-30s:%#x\n", k, subv)
			case "totalLength", "timeToLive":
				fmt.Printf("%-30s:%d\n", k, subv)
			case "source", "destination":
				if subk == 0 {
					var strV []string
					for _, subvv := range v {
						strV = append(strV, fmt.Sprintf("%d", subvv))
					}

					ip := strings.Join(strV, ".")
					fmt.Printf("%-30s:%s\n", k, ip)
				}
			}
		}
	}

	header("respTCP")

	respTCP := map[string][]int{
		//{0x00, 0x50},
		"sourcePort": {0x0050}, // 80

		//{0xcc, 0xa7},
		"destinationPort": {0xcca7}, // 52391

		//{0x5f, 0xe7, 0x88, 0xd7},
		"sequenceNumber": {0x5fe788d7}, // 1609009367 TODO: wireshark 是 1 (relative sequence number)

		//{0xd2, 0xc4, 0x19, 0xf4},
		"acknowledgmentNumber": {0xd2c419f4}, // 3536067060 TODO: wireshark 是 561 (relative sequence number)

		//{0x80, 0x18},
		"flags": {0x8018}, // TODO: 这里不确定是多少位。看 wireshark 好像是 Header Length 和 Flags 有重合的地方

		//{0x00, 0xeb},
		"windowSize": {0x00eb}, // 235

		//{0xce, 0x7a},
		"checksum": {0xce7a}, // 0xce7a

		//{0x00, 0x00},
		"urgentPointer": {0x0000}, // 0x0000

		//{0x01, 0x01, 0x08, 0x0a, 0xf0, 0x4e, 0xd0, 0x56, 0x13, 0xee, 0xef, 0x96},
		// 00000001, 00000001, 00001000, 00001010, 11110000, 01001110, 11010000, 01010110, 00010011, 11101110, 11101111, 10010110
		"options": {0x01, 0x01, 0x08, 0x0a, 0xf0, 0x4e, 0xd0, 0x56, 0x13, 0xee, 0xef, 0x96},
	}

	for k, v := range respTCP {
		for _, subv := range v {
			switch k {
			case
				"sourcePort",
				"destinationPort",
				"sequenceNumber",
				"acknowledgmentNumber",
				"windowSize":
				fmt.Printf("%-20s:%v", k, subv)
			case "flags":
				fmt.Printf("%-20s:%16b", k, subv)
			case "checksum", "urgentPointer":
				fmt.Printf("%-20s:%#x", k, subv)
			case "options":
				fmt.Printf("%-20s:%08b\t", k, subv)
			}
			fmt.Printf("\n")
		}
	}

	header("respHTTP")

	respHTTP := [][]byte{
		// HTTP/1.1 304 Not Modified\r\n
		{0x48, 0x54, 0x54, 0x50, 0x2f, 0x31, 0x2e, 0x31, 0x20, 0x33, 0x30, 0x34, 0x20, 0x4e, 0x6f, 0x74, 0x20, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x0d, 0x0a},

		// Server: nginx/1.12.2\r\n
		{0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x20, 0x6e, 0x67, 0x69, 0x6e, 0x78, 0x2f, 0x31, 0x2e, 0x31, 0x32, 0x2e, 0x32, 0x0d, 0x0a},

		// Last-Modified: Tue, 06 Mar 2018 09:26:21 GMT\r\n
		{0x4c, 0x61, 0x73, 0x74, 0x2d, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x3a, 0x20, 0x54, 0x75, 0x65, 0x2c, 0x20, 0x30, 0x36, 0x20, 0x4d, 0x61, 0x72, 0x20, 0x32, 0x30, 0x31, 0x38, 0x20, 0x30, 0x39, 0x3a, 0x32, 0x36, 0x3a, 0x32, 0x31, 0x20, 0x47, 0x4d, 0x54, 0x0d, 0x0a},

		// Connection: keep-alive\r\n
		{0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x6b, 0x65, 0x65, 0x70, 0x2d, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x0d, 0x0a},

		// ETag: "5a9e5ebd-e74"\r\n
		{0x45, 0x54, 0x61, 0x67, 0x3a, 0x20, 0x22, 0x35, 0x61, 0x39, 0x65, 0x35, 0x65, 0x62, 0x64, 0x2d, 0x65, 0x37,
			0x34, 0x22, 0x0d, 0x0a},

		// \r\n
		{0x0d, 0x0a},
	}

	for _, v := range respHTTP {
		l := len(v)
		for k, subv := range v {
			if k >= l-2 {
				fmt.Print(strings.Replace(
					fmt.Sprintf("%q", subv), "'", "", -1))
				if k == l-1 {
					fmt.Printf("\n")
				}
			} else {
				fmt.Printf("%c", subv)
			}
		}
	}
}

func header(h string) {
	fmt.Printf("\n%s\t%s\n\n", h, "====================================================")
}

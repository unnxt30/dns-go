package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	"github.com/unnxt30/dns-go/cmd"
	"github.com/unnxt30/dns-go/models"
)

type DNSClient struct {
	conn       *net.UDPConn
	serverAddr *net.UDPAddr
	timeout    time.Duration
}

func NewDNSClient(serverAddr string) (*DNSClient, error) {
	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		return nil, fmt.Errorf("couldn't resolve UDP address: %v", err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, fmt.Errorf("couldn't establish UDP connection: %v", err)
	}

	return &DNSClient{
		conn:       conn,
		serverAddr: addr,
		timeout:    time.Second * 5,
	}, nil
}

func (c *DNSClient) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

func (c *DNSClient) Close() error {
	return c.conn.Close()
}

func (c *DNSClient) Query(domain string, queryType uint16) (*QueryResult, error) {
	message := models.DNSMessage{
		Header: models.DNSHeader{
			ID:      22,
			Flags:   models.HeaderFlags{RD: 0},
			QDCount: 1,
			ANCount: 0,
			NSCount: 0,
			ARCount: 0,
		},
		Question: models.DNSQuestion{
			QName:  domain,
			QType:  queryType,
			QClass: 1,
		},
	}

	encoded, err := cmd.EncodeMessage(message)
	if err != nil {
		return nil, fmt.Errorf("error encoding message: %v", err)
	}

	if err := c.conn.SetDeadline(time.Now().Add(c.timeout)); err != nil {
		return nil, fmt.Errorf("error setting timeout: %v", err)
	}

	resp, err := cmd.ExchangeMessage(c.conn, encoded)
	if err != nil {
		return nil, fmt.Errorf("error exchanging message: %v", err)
	}

	decoder := cmd.DNSDecoder{
		Encoded: resp,
		Offset:  0,
	}

	return decodeResponse(&decoder)
}

type QueryResult struct {
	Header    models.DNSHeader
	Question  models.DNSQuestion
	Answer    []models.ResponseStruct
	NSRecords []string
	IPRecords []string
}

func decodeResponse(decoder *cmd.DNSDecoder) (*QueryResult, error) {
	result := &QueryResult{}

	header, err := decoder.DecodeHeader()
	if err != nil {
		return nil, fmt.Errorf("error decoding header: %v", err)
	}
	result.Header = header

	question, err := decoder.DecodeQuestion()
	if err != nil {
		return nil, fmt.Errorf("error decoding question: %v", err)
	}
	result.Question = question
	if header.ANCount > 0 {
		ans, err := decoder.DecodeAnswers(int(header.ANCount))

		if err != nil {
			return nil, fmt.Errorf("error decoding NS records: %v", err)
		}
		result.Answer = ans
	}
	if header.NSCount > 0 {
		_, err = decoder.DecodeAnswers(int(header.NSCount))
		if err != nil {
			return nil, fmt.Errorf("error decoding NS records: %v", err)
		}
		result.NSRecords = decoder.NSRecords
	}

	if header.ARCount > 0 {
		_, err = decoder.DecodeAnswers(int(header.ARCount))
		if err != nil {
			return nil, fmt.Errorf("error decoding additional records: %v", err)
		}
		result.IPRecords = decoder.IPRecords
	}

	return result, nil
}

func resolve(domain string, queryType uint16, serverAddr string) (*QueryResult, error) {
	client, err := NewDNSClient(serverAddr)
	if err != nil {
		return nil, fmt.Errorf("couldn't create DNS client: %v", err)
	}
	defer client.Close()

	client.SetTimeout(time.Second * 10)

	result, err := client.Query(domain, queryType)
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}

	if result.Header.ANCount > 0 {
		answers := result.Answer
		for _, x := range answers {
			fmt.Printf("Resolved IP for %s: %s\n", domain, x.RData)
		}
		return result, nil
	}

	// If the response contains NS records, follow the referrals
	if result.Header.NSCount > 0 {
		for _, ns := range result.NSRecords {
			fmt.Printf("Following referral to authoritative name server: %s\n", ns)

			var nsIP string
			for _, ip := range result.IPRecords {
				if ip != "" {
					nsIP = ip
					break
				}
			}

			if nsIP == "" {
				fmt.Printf("Resolving IP for name server: %s\n", ns)
				nsResult, err := resolve(ns, 1, "198.41.0.4:53")
				if err != nil {
					fmt.Printf("Failed to resolve NS %s: %v\n", ns, err)
					continue
				}
				if len(nsResult.IPRecords) > 0 {
					nsIP = nsResult.IPRecords[0]
				}
			}

			if nsIP == "" {
				fmt.Printf("No IP found for name server: %s\n", ns)
				continue
			}

			fmt.Printf("Querying authoritative name server: %s (%s)\n", ns, nsIP)
			return resolve(domain, queryType, nsIP+":53")
		}
	}

	return nil, fmt.Errorf("no answer found for %s", domain)
}

func main() {
	domain := flag.String("domain", "", "The domain to resolve")
	flag.Parse()

	if *domain == "" {
		fmt.Println("Please provide a domain to resolve using the -domain flag.")
		return
	}

	root := "198.41.0.4:53"
	_, err := resolve(*domain, 1, root)
	if err != nil {
		fmt.Printf("Error resolving %s: %v\n", *domain, err)
		return
	}

}

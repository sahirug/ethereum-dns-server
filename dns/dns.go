package dns

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/sahirug/ethereum-dns/contracts"
	"github.com/sahirug/ethereum-dns/errorhandler"
	"github.com/sahirug/ethereum-dns/util"
	"golang.org/x/net/dns/dnsmessage"
	"log"
	"net"
	"strings"
	"time"
)

const (
	udpPort int = 53
	packetLen int = 512
)

type DDNServer interface {
	Listen()
	Query(Packet)
}

type DDNService struct {
	conn *net.UDPConn
}

type Packet struct {
	addr net.UDPAddr
	message dnsmessage.Message
}

func (s *DDNService) Listen(e *eddns.Eddns, auth *bind.TransactOpts) {
	var err error
	s.conn, err = net.ListenUDP("udp", &net.UDPAddr{Port:udpPort})
	if err != nil {
		errorhandler.HandleErr(err, 11)
	}
	defer s.conn.Close()

	log.Println("Waiting for DNS requests...")
	for true {
		buf := make([]byte, packetLen)
		_, addr, err := s.conn.ReadFromUDP(buf)

		if err != nil {
			log.Println(err)
			time.Sleep(time.Second)
			continue
		}

		var m dnsmessage.Message

		err = m.Unpack(buf)

		if err != nil {
			errorhandler.HandleErr(err, 13)
		}


		if len(m.Questions) == 0 {
			log.Print("No DNS requests found. Retrying....")
			continue
		}

		s.Query(Packet{*addr, m}, e, auth)
	}
}

func (s *DDNService) Query(packet Packet, contract *eddns.Eddns, auth *bind.TransactOpts) {
	if packet.message.Questions[0].Type == dnsmessage.TypeA || packet.message.Questions[0].Type == dnsmessage.TypeAAAA {
		log.Printf("Message type: %s", packet.message.Questions[0].Type.String())
		if packet.message.Header.Response {
			// do stuff
		}
		log.Print("Querying blockchain")
		q := packet.message.Questions[0].Name.String()
		fullQ := strings.Split(q, ".")
		domain, tld1 := fullQ[0], fullQ[1]
		log.Printf("Domain is %s tld is %s", domain, tld1)

		var domainName [32]byte
		var tld [12]byte
		var rType [32]byte

		copy(domainName[:], []byte(domain))
		copy(tld[:], []byte(tld1))
		if packet.message.Questions[0].Type == dnsmessage.TypeA {
			copy(rType[:], []byte("A"))
		} else {
			copy(rType[:], []byte("AAAA"))
		}

		opts := bind.CallOpts{
			Pending: true,
			From: auth.From,
		}

		ip, err := contract.GetIp(&opts, domainName, tld, rType)

		if err != nil {
			log.Printf("Error when trying to resolve %s", fullQ)
			log.Printf("Error is %v", err)
		} else {
			stringIps := util.ConvertByteArrToString(ip)
			resources, err := convertToDNSResource(q, stringIps, packet.message.Questions[0].Type)
			if err != nil {
				log.Println("error when trying to create resource record")
				log.Println(err)
			}
			packet.message.Answers = append(packet.message.Answers, resources...)
			sendPacket(s.conn, packet.message, packet.addr)
		}
	}
}

func sendPacket(conn *net.UDPConn, message dnsmessage.Message, addr net.UDPAddr) {
	packed, err := message.Pack()
	if err != nil {
		log.Println("error when trying to pack message")
		log.Println(err)
		return
	}
	t, err := conn.WriteToUDP(packed, &addr)
	if err != nil {
		log.Println("error when trying to write to udp")
		log.Println(err)
	}
	_ = t
}

func convertToDNSResource(domain string, data []string, rType dnsmessage.Type) ([]dnsmessage.Resource, error) {
	//log.Fatalf("Ips are %v", data)
	rName, err := dnsmessage.NewName(domain)
	none := []dnsmessage.Resource{}

	var resources []dnsmessage.Resource;

	if err != nil {
		return none, err
	}

	for i := 0; i < len(data); i++ {
		//var rType dnsmessage.Type
		var rBody dnsmessage.ResourceBody

		stringIp := data[i]

		//rType = dnsmessage.TypeA
		ip := net.ParseIP(stringIp)
		if ip == nil {
			return none, errors.New("invalid ip")
		}
		if rType == dnsmessage.TypeA {
			rBody = &dnsmessage.AResource{A: [4]byte{ip[12], ip[13], ip[14], ip[15]}}
		} else {
			var ipV6 [16]byte
			copy(ipV6[:], ip)
			rBody = &dnsmessage.AAAAResource{AAAA: ipV6}
		}
		resource := dnsmessage.Resource{
			Header: dnsmessage.ResourceHeader{
				Name:rName,
				Type:rType,
				Class:dnsmessage.ClassCSNET,
				TTL:500,
			},
			Body: rBody,
		}
		resources = append(resources, resource)
	}
	//log.Fatalf("resources are %v", resources)
	return resources, nil
}


func New() DDNService {
	return DDNService{}
}

func Start(e *eddns.Eddns, auth *bind.TransactOpts) *DDNService {
	s := New()
	s.Listen(e, auth)

	return &s
}
package dns

import (
	"github.com/sahirug/ethereum-dns/contracts"
	"github.com/sahirug/ethereum-dns/errorhandler"
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

func (s *DDNService) Listen(e *eddns.Eddns) {
	var err error
	s.conn, err = net.ListenUDP("udp", &net.UDPAddr{Port:udpPort})
	if err != nil {
		errorhandler.HandleErr(err, 11)
	}
	defer s.conn.Close()

	log.Println("Polling for DNS requests...")
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

		s.Query(Packet{*addr, m}, e)
	}
}

func (s *DDNService) Query(packet Packet, contract *eddns.Eddns) {
	if packet.message.Questions[0].Type != dnsmessage.TypePTR {
		if packet.message.Header.Response {
			// do stuff
		}
		log.Print("Querying blockchain")
		q := packet.message.Questions[0].Name.String()
		fullQ := strings.Split(q, ".")
		domain, tld1 := fullQ[0], fullQ[1]
		log.Printf("Domain is %s tld is %s", domain, tld1)

		var tld [12]byte

		copy(tld[:], []byte(tld1))

		ip, err := contract.GetIp(nil, []byte(domain), tld)

		if err != nil {
			log.Printf("Error when trying to resolve %s", fullQ)
		} else {
			log.Printf("IP: %s", string(ip[0][:]))
			//packet.message.Answers = append()

		}
	}
}


func New() DDNService {
	return DDNService{}
}

func convertToDNSResource() {

}

func Start(e *eddns.Eddns) *DDNService {
	s := New()
	s.Listen(e)

	return &s
}
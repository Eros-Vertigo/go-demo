package pcap

import (
	"bufio"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
)

func init() {
	log.Info("module pcap")
	setUp()
}

type httpStreamFactory struct{}

func (f *httpStreamFactory) New(a, b gopacket.Flow) tcpassembly.Stream {
	r := tcpreader.NewReaderStream()
	go printRequests(&r, a, b)
	return &r
}

func printRequests(r io.Reader, a, b gopacket.Flow) {
	buf := bufio.NewReader(r)
	for {
		if req, err := http.ReadRequest(buf); err == io.EOF {
			return
		} else if err != nil {
			log.Println("Error parsing HTTP requests:", err)
		} else {
			fmt.Println(a, b)
			fmt.Println("HTTP REQUEST:", req)
			fmt.Println("Body contains", tcpreader.DiscardBytesToEOF(req.Body), "bytes")
		}
	}
}

func setUp() {
	var handle *pcap.Handle
	//var dec gopacket.Decoder
	var err error
	if handle, err = pcap.OpenOffline("/Users/yuantong/Developer/go/src/srun4-antiproxy/common/config/min.pcap"); err != nil {
		log.Error("PCAP OpenOffline error:", err)
		return
	}
	// set up assembly
	//streamFactory := &httpStreamFactory{}
	//streamPool := tcpassembly.NewStreamPool(streamFactory)
	//assembly := tcpassembly.NewAssembler(streamPool)
	source := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range source.Packets() {
		//analyzeTCP(packet, assembly)
		//analyzeUDP(packet)
		analyzeTLS(packet)
	}
}

func analyzeTCP(packet gopacket.Packet, assembly *tcpassembly.Assembler) {
	if packet.TransportLayer() == nil {
		return
	}
	if packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
		return
	}
	tcp := packet.TransportLayer().(*layers.TCP)

	assembly.Assemble(packet.NetworkLayer().NetworkFlow(), tcp)
}

func analyzeUDP(packet gopacket.Packet) {
	if packet.TransportLayer() == nil {
		return
	}
	if packet.TransportLayer().LayerType() != layers.LayerTypeUDP {
		return
	}
	udp := packet.TransportLayer().(*layers.UDP)

	if udp.DstPort == 8000 {
		payload := packet.ApplicationLayer().Payload()
		flag := fmt.Sprintf("%x", payload[0:1])
		if flag == "02" {
			qq, err := strconv.ParseUint(fmt.Sprintf("%x", payload[7:11]), 16, 32)
			if err != nil {
				return
			}
			log.Infof("qq number is %d", qq)
		}
	}
}

func analyzeTLS(packet gopacket.Packet) {
	if packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
		return
	}
	if packet.ApplicationLayer() == nil {
		return
	}
	payload := packet.ApplicationLayer().Payload()
	p := gopacket.NewPacket(payload, layers.LayerTypeTLS, gopacket.DecodeOptions{
		SkipDecodeRecovery:       true,
		DecodeStreamsAsDatagrams: true,
	})
	l := p.Layer(layers.LayerTypeTLS)
	if l == nil {
		return
	}
	temp := l.(*layers.TLS)
	if len(temp.Handshake) != 1 {
		return
	}
	if temp.Handshake[0].Length < 70 {
		return
	}
	con := payload[uint16(len(payload))-temp.Handshake[0].Length:]
	ht, _ := strconv.ParseUint(fmt.Sprintf("%x", con[0:1]), 16, 32)
	if ht == 1 {
		fmt.Printf("src port [%s]", packet.TransportLayer().TransportFlow().Src().String())
		decoder(con)
	}
}

func decoder(con []byte) {
	current := uint64(38)
	sessionL := con[current : current+1]
	sessionLength, _ := strconv.ParseUint(fmt.Sprintf("%x", sessionL), 16, 32)
	current += sessionLength + 1
	cipherL := con[current : current+2]
	cipherLength, _ := strconv.ParseUint(fmt.Sprintf("%x", cipherL), 16, 32)
	current = current + cipherLength + 6
	notSNI := false
	for current < uint64(len(con)) {
		temp := con[current : current+2]
		if fmt.Sprintf("0x%.2X", temp) == "0x0000" {
			current += 7
			length := con[current : current+2]
			l, _ := strconv.ParseUint(fmt.Sprintf("%x", length), 16, 32)
			current += 2
			serverName := con[current : current+l]
			log.Infof("server_name is %s\n", serverName)
			return
		} else {
			if notSNI {
				return
			}
			current += 4
			notSNI = true
			continue
		}
	}
}

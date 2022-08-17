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
	if len(temp.Handshake) == 0 {
		return
	}
	fmt.Println(temp.Handshake)
	fmt.Println(p)
}

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
	if handle, err = pcap.OpenOffline("/Users/yuantong/Desktop/ios.pcap"); err != nil {
		log.Error("PCAP OpenOffline error:", err)
		return
	}
	// set up assembly
	streamFactory := &httpStreamFactory{}
	streamPool := tcpassembly.NewStreamPool(streamFactory)
	assembly := tcpassembly.NewAssembler(streamPool)

	source := gopacket.NewPacketSource(handle, gopacket.DecodersByLayerName[handle.LinkType().String()])
	for packet := range source.Packets() {
		if packet.TransportLayer() == nil {
			continue
		}
		if packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
			continue
		}
		tcp := packet.TransportLayer().(*layers.TCP)

		assembly.Assemble(packet.NetworkLayer().NetworkFlow(), tcp)
		//fmt.Println(packet)
	}
}

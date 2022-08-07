package uaparse

import (
	log "github.com/sirupsen/logrus"
	"github.com/ua-parser/uap-go/uaparser"
)

func init() {
	uagent := "Mozilla/5.0 (Linux; Android 12; V2185A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.134 Mobile Safari/537.36 EdgA/103.0.1264.77"

	parser := uaparser.NewFromSaved()
	//if err != nil {
	//	log.Fatal(err)
	//}

	client := parser.Parse(uagent)

	log.Infof("ua.Os = [%s]", client.Os.Family)
	log.Infof("ua.Os = [%s]", client.Os.Minor)
	log.Infof("ua.Os = [%s]", client.Os.Major)
	log.Infof("ua.Os = [%s]", client.Os.Patch)
	log.Infof("ua.Os = [%s]", client.Os.PatchMinor)
	log.Infof("ua.Ver = [%s]", client.Os.ToVersionString())
	log.Infof("ua.Device = [%s]", client.Device.Family)
}

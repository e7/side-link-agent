package main

import (
	"flag"
	"net"
	"strings"

	"github.com/hashicorp/memberlist"
    log "github.com/sirupsen/logrus"
)

const (
	ActionCreate = "create"
	ActionJoin   = "join"
)

var (
	flgAction     string
	flgJoinTarget string
)

func init() {
	flag.StringVar(&flgAction, "action", "join", "-action create or join")
	flag.StringVar(&flgJoinTarget, "target", "", "-target host	required if join")
	flag.Parse()
}

func GetOutBoundIP() (string, error) {
	conn, err := net.Dial("udp", "12.34.56.78:53")
	if err != nil {
		return "", err
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return strings.Split(localAddr.String(), ":")[0], nil
}

func main() {
    log.SetLevel(log.InfoLevel)

	mList, err := memberlist.Create(memberlist.DefaultLANConfig())
	if err != nil {
		log.Fatalf("Create memberlist failed, err: %s", err.Error())
	}

	if ActionCreate == flgAction {
        log.Infof("create new cluster, health: %d", mList.GetHealthScore())
	} else if ActionJoin == flgAction {
		if flgJoinTarget == "" {
			flag.Usage()
			return
		}

        if _, err = mList.Join([]string{flgJoinTarget}); nil != err {
        	log.Fatalf("Join memberlist failed, err: %s", err.Error())
		}
    } else {
        flag.Usage()
        return
    }

    
}

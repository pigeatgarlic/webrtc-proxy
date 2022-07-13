package main

import (
	proxy "github.com/pigeatgarlic/webrtc-proxy"
	"github.com/pigeatgarlic/webrtc-proxy/util/config"
	"github.com/pion/webrtc/v3"
)

func main() {
	grpc := config.GrpcConfig{
		Port:          8000,
		ServerAddress: "localhost",
	}
	rtc := config.WebRTCConfig{
		Ices: []webrtc.ICEServer{webrtc.ICEServer{
			URLs: []string{"stun:stun.l.google.com:19302"},
		}},
	}
	br := []*config.BroadcasterConfig{
		&config.BroadcasterConfig{
			Port: 5001,
			Protocol: "udp",
			BufferSize: 1028,

			Type: "video",
			Name: "rtp2",
			Codec: webrtc.MimeTypeH264,
		},
	}
	lis := []*config.ListenerConfig{
		&config.ListenerConfig{
			Port: 6000,
			Protocol: "udp",
			BufferSize: 1028,

			Type: "video",
			Name: "rtp2",
			Codec: webrtc.MimeTypeH264,
		},
	}

	_,err := proxy.InitWebRTCProxy(nil,&grpc,&rtc,br,lis);
	if err != nil {
		panic(err);
	}
	shut := make(chan bool)
	<- shut
}

package notify

import (
	"net"
	"net/http"
	"time"

	"github.com/appleboy/go-fcm"
	"github.com/sideshow/apns2"
)

var (
	// ApnsClient is apns client
	ApnsClient *apns2.Client
	// FCMClient is apns client
	FCMClient *fcm.Client
	// MaxConcurrentIOSPushes pool to limit the number of concurrent iOS pushes
	MaxConcurrentIOSPushes chan struct{}

	transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
		MaxIdleConns:        5,
		MaxIdleConnsPerHost: 5,
		MaxConnsPerHost:     20,
	}
	feedbackClient = &http.Client{
		Transport: transport,
	}
)

const (
	HIGH   = "high"
	NORMAL = "nornal"
)

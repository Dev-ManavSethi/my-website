package utils

import (
	"net/http"
	"strings"
)

func GetUserIP(r *http.Request) string{


	IPAddress := r.Header.Get("X-Real-Ip")

	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}

	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}


	Addresses := strings.Split(IPAddress, ":")
	IPAddress = Addresses[0]

	return IPAddress

}

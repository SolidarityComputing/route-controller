package configs

import (
	"log"
	"os"
	"strings"
	"time"
)

var (
	ServeAddress = ":8080"

	AuthAccount  = ""
	AuthPassword = ""

	EtcdUsername  = ""
	EtcdPassword  = ""
	EtcdTimeout   = 10 * time.Second
	EtcdEndPoints = []string{"http://rancher.kfcoding.com:2379"}

	TraefikPrefix = "/kfcoding/traefik/"
)

func InitEnvs() {
	if t := os.Getenv("ServeAddress"); t != "" {
		ServeAddress = t
	}
	if t := os.Getenv("AuthAccount"); t != "" {
		AuthAccount = t
	}
	if t := os.Getenv("AuthPassword"); t != "" {
		AuthPassword = t
	}
	if t := os.Getenv("EtcdEndPoints"); t != "" {
		EtcdEndPoints = strings.Split(t, ",")
	}
	if t := os.Getenv("EtcdUsername"); t != "" {
		EtcdUsername = t
	}
	if t := os.Getenv("EtcdPassword"); t != "" {
		EtcdPassword = t
	}
	if t := os.Getenv("TraefikPrefix"); t != "" {
		TraefikPrefix = t
	}

	log.Print("---> Init Environments")
	log.Print("ServeAddress:        ", ServeAddress)
	log.Print("AuthAccount:         ", AuthAccount)
	log.Print("AuthPassword:        ", AuthPassword)
	log.Print("EtcdEndPoints:       ", EtcdEndPoints)
	log.Print("EtcdUsername:        ", EtcdUsername)
	log.Print("EtcdPassword:        ", EtcdPassword)
	log.Print("TraefikPrefix:       ", TraefikPrefix)
}

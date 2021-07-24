package main

import (
	"fmt"
	"log"
	"os"

	service "github.com/cloudnativego/backing-fulfillment/service"
	"github.com/hudl/fargo"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3001"
	}

	eurekaUrl := os.Getenv("EUREKA_URL")
	if len(eurekaUrl) == 0 {
		log.Fatal("Missing ENV variable: EUREKA_URL")
	}

	c := fargo.NewConn(eurekaUrl)

	i := fargo.Instance{
		HostName:         "i-6543",
		Port:             3001,
		App:              "BACKING_FULFILLMENT",
		IPAddr:           "127.0.0.10",
		VipAddress:       "127.0.0.10",
		SecureVipAddress: "127.0.0.10",
		DataCenterInfo:   fargo.DataCenterInfo{Name: fargo.MyOwn},
		Status:           fargo.UP,
	}

	c.RegisterInstance(&i)
	f, _ := c.GetApps()

	for key, theApp := range f {
		fmt.Println("Registered App:", key, " First Host Name:", theApp.Instances[0].HostName)
	}

	// Ordinarily we'd use a CF environment here, but we don't need it for
	// the fake data we're returning.
	server := service.NewServer()
	server.Run(":" + port)
}

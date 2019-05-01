package main

import (
	"golang.org/x/net/context"
	"github.com/containernetworking/cni/libcni"
	// "github.com/labstack/echo/middleware"
	// "github.com/labstack/echo"
	// "github.com/madoka/app/handler"
	"fmt"
)

func main() {
	// e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// e.GET("/auth/test",handler.AuthTest())
	// e.POST("/container/network",handler.AddNetworkHandler())

	// fmt.Println("madoka")
	// e.Start(":8081")
	fmt.Println()

	netconf, err := libcni.LoadConfList("/etc/cni/net.d", "mynet")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(netconf)
	cninet := &libcni.CNIConfig{
		Path: []string{"/opt/cni/bin"},
	}
	fmt.Println(cninet)

	containerID := "de72f482b5c907afebc1983e249a6b94c34252050b0b8f7656d0d64662662ea9"
	netns := "/var/run/docker/netns/d314b983720c"
	ifName := "eth0"
	cniArgs := [][2]string{}
	var capabilityArgs map[string]interface{}

	fmt.Println(containerID)
	fmt.Println(netns)
	fmt.Println(ifName)
	fmt.Println(cniArgs)

	rt := &libcni.RuntimeConf{
		ContainerID:    containerID,
		NetNS:          netns,
		IfName:         ifName,
		Args:           cniArgs,
		CapabilityArgs: capabilityArgs,
	}

	// cninet.DelNetworkList(context.TODO(), netconf, rt)
	cninet.AddNetworkList(context.TODO(), netconf, rt)
}

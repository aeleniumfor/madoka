package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/madoka/app/handler"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/auth/test", handler.AuthTest())
	e.POST("/container/ip", handler.AddNetworkHandler())

	fmt.Println("madoka")
	e.Start(":8081")
	fmt.Println()

	// netconf, err := libcni.LoadConfList("/etc/cni/net.d", "mynet")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(netconf)
	// cninet := &libcni.CNIConfig{
	// 	Path: []string{"/opt/cni/bin"},
	// }
	// fmt.Println(cninet)

	// containerID := "e35d1b5c6f4779416d20b31175765542673caf85f22a5f6f4007418992220356"
	// netns := "/var/run/docker/netns/8ae6ce49f163"
	// ifName := "eth0"
	// cniArgs := [][2]string{}
	// var capabilityArgs map[string]interface{}

	// // fmt.Println(containerID)
	// // fmt.Println(netns)
	// // fmt.Println(ifName)
	// // fmt.Println(cniArgs)

	// rt := &libcni.RuntimeConf{
	// 	ContainerID:    containerID,
	// 	NetNS:          netns,
	// 	IfName:         ifName,
	// 	Args:           cniArgs,
	// 	CapabilityArgs: capabilityArgs,
	// }

	// // cninet.DelNetworkList(context.TODO(), netconf, rt)
	// cninet.AddNetworkList(context.TODO(), netconf, rt)
	// cninet.CheckNetworkList(context.TODO(), netconf, rt)
	// // 10.22.0.3/16
}

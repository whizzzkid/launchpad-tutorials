package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
)

func createNode (addressesString string) host.Host {
	node, error := libp2p.New(libp2p.ListenAddrStrings(addressesString))
	if error != nil {
		panic(error)
	}

	return node
}

func createSourceNode() host.Host {
	return createNode("/ip4/0.0.0.0/tcp/7777")
}

func createTargetNode() host.Host {
	return createNode("/ip4/0.0.0.0/tcp/8888")
}

func connectToTargetNode(sourceNode host.Host, targetNode host.Host) {
	error := sourceNode.Connect(context.Background(), *host.InfoFromHost(targetNode))
	if error != nil {
		panic(error)
	}
}

func countSourceNodePeers(sourceNode host.Host) int {
	return len(sourceNode.Network().Peers())
}

func printNodeID(host host.Host) {
	println(fmt.Sprintf("ID: %s", host.ID().String()))
}

func printNodeAddresses(host host.Host) {
	addressesString := make([]string, 0)
	for _, address := range host.Addrs() {
		addressesString = append(addressesString, address.String())
	}

	println(fmt.Sprintf("Multiaddresses: %s", strings.Join(addressesString, ", ")))
}

func main() {
	sourceNode := createSourceNode()
	println("-- SOURCE NODE INFORMATION --")
	printNodeID(sourceNode)
	printNodeAddresses(sourceNode)

	targetNode := createTargetNode()
	println("-- TARGET NODE INFORMATION --")
	printNodeID(targetNode)
	printNodeAddresses(targetNode)

	connectToTargetNode(sourceNode, targetNode)

	println(fmt.Sprintf("Source node peers: %d", countSourceNodePeers(sourceNode)))
}

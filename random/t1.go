package main

import "fmt"


type Network struct {
	Value string
}

type Config struct {
	Network Network
	Networks []Network
}

func main(){

	config := Config{
		Network: Network{
			Value: "t1",
		},
		Networks: []Network{
			{
				Value: "t2",
			},
			{
				Value: "t3",
			},
			{
				Value: "t4",
			},
			{
				Value: "t5",
			},
		},
	}


	var network string;
	var networks []string;


	for _, network := range config.Networks {
		networks = append(networks, network.Value)
	}

	fmt.Println("Appended Networks", networks)

	if len(networks) > 1 {
		network = networks[0]
		networks = networks[1:]
	}
	
	fmt.Println("Network", network)
	fmt.Println("Networks", networks)

}
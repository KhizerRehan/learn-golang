package main

import "fmt"

type CloudProvider string

const (
	CloudProviderAWS                 CloudProvider = "aws"
	CloudProviderAzure               CloudProvider = "azure"
	CloudProviderDigitalocean        CloudProvider = "digitalocean"
	CloudProviderGoogle              CloudProvider = "gce"
	CloudProviderEquinixMetal        CloudProvider = "equinixmetal"
	CloudProviderPacket              CloudProvider = "packet"
	CloudProviderHetzner             CloudProvider = "hetzner"
	CloudProviderKubeVirt            CloudProvider = "kubevirt"
	CloudProviderLinode              CloudProvider = "linode"
	CloudProviderNutanix             CloudProvider = "nutanix"
	CloudProviderOpenstack           CloudProvider = "openstack"
	CloudProviderVsphere             CloudProvider = "vsphere"
	CloudProviderVultr               CloudProvider = "vultr"
	CloudProviderVMwareCloudDirector CloudProvider = "vmware-cloud-director"
	CloudProviderFake                CloudProvider = "fake"
	CloudProviderEdge                CloudProvider = "edge"
	CloudProviderAlibaba             CloudProvider = "alibaba"
	CloudProviderAnexia              CloudProvider = "anexia"
	CloudProviderScaleway            CloudProvider = "scaleway"
	CloudProviderBaremetal           CloudProvider = "baremetal"
	CloudProviderExternal            CloudProvider = "external"
	CloudProviderOpenNebula          CloudProvider = "opennebula"
)

func main() {

	fmt.Println(CloudProviderAWS)
	fmt.Println(CloudProviderAzure)

}

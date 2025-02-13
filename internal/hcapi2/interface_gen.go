package hcapi2

//go:generate interfacer -for github.com/hetznercloud/hcloud-go/hcloud.CertificateClient -as hcapi2.CertificateClientBase -o zz_certificate_client_base.go
//go:generate interfacer -for github.com/hetznercloud/hcloud-go/hcloud.DatacenterClient -as hcapi2.DatacenterClientBase -o zz_datacenter_client_base.go
//go:generate interfacer -for github.com/hetznercloud/hcloud-go/hcloud.ImageClient -as hcapi2.ImageClientBase -o zz_image_client_base.go
//go:generate interfacer -for github.com/hetznercloud/hcloud-go/hcloud.ISOClient -as hcapi2.ISOClientBase -o zz_iso_client_base.go
//go:generate interfacer -for github.com/hetznercloud/hcloud-go/hcloud.FirewallClient -as hcapi2.FirewallClientBase -o zz_firewall_client_base.go
//go:generate interfacer -for github.com/hetznercloud/hcloud-go/hcloud.FloatingIPClient -as hcapi2.FloatingIPClientBase -o zz_floating_ip_client_base.go
//go:generate interfacer -for github.com/hetznercloud/hcloud-go/hcloud.LocationClient -as hcapi2.LocationClientBase -o zz_location_client_base.go
//go:generate interfacer -for github.com/hetznercloud/hcloud-go/hcloud.LoadBalancerClient -as hcapi2.LoadBalancerClientBase -o zz_loadbalancer_client_base.go
//go:generate interfacer -for github.com/hetznercloud/hcloud-go/hcloud.LoadBalancerTypeClient -as hcapi2.LoadBalancerTypeClientBase -o zz_loadbalancer_type_client_base.go
//go:generate interfacer -for github.com/hetznercloud/hcloud-go/hcloud.NetworkClient -as hcapi2.NetworkClientBase -o zz_network_client_base.go
//go:generate interfacer -for github.com/hetznercloud/hcloud-go/hcloud.ServerClient -as hcapi2.ServerClientBase -o zz_server_client_base.go
//go:generate interfacer -for github.com/hetznercloud/hcloud-go/hcloud.ServerTypeClient -as hcapi2.ServerTypeClientBase -o zz_server_type_client_base.go
//go:generate interfacer -for github.com/hetznercloud/hcloud-go/hcloud.SSHKeyClient -as hcapi2.SSHKeyClientBase -o zz_ssh_key_client_base.go
//go:generate interfacer -for github.com/hetznercloud/hcloud-go/hcloud.VolumeClient -as hcapi2.VolumeClientBase -o zz_volume_client_base.go
//go:generate interfacer -for github.com/hetznercloud/hcloud-go/hcloud.PlacementGroupClient -as hcapi2.PlacementGroupClientBase -o zz_placement_group_client_base.go
//go:generate interfacer -for github.com/hetznercloud/hcloud-go/hcloud.RDNSClient -as hcapi2.RDNSClientBase -o zz_rdns_client_base.go

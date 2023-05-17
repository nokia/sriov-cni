module github.com/nokia/sriov-cni/plugins

go 1.20

require (
	github.com/Masterminds/semver v1.5.0
	github.com/k8snetworkplumbingwg/sriov-cni v2.1.0+incompatible
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.27.6
)

require (
	github.com/containernetworking/cni v1.1.2 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/vishvananda/netlink v1.2.1-beta.2 // indirect
	github.com/vishvananda/netns v0.0.2 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/k8snetworkplumbingwg/sriov-cni v2.1.0+incompatible => ../

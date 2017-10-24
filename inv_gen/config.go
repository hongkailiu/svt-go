package inv_gen

import "github.com/hongkailiu/svt-go/log"

type Hosts struct {
	master_nodes   map[string]struct{}
	etcd_nodes   map[string] struct{}
	infra_nodes   map[string] struct{}
	compute_nodes   map[string] struct{}
	lb_nodes   map[string] struct{}
	glusterfs_nodes   map[string] struct{}
}


func (hosts Hosts) addNodes(role string, nodes []string) {
	switch role {
	case master_key:
		addAllNodes(&(hosts.master_nodes), nodes)
	case etcd_key:
		addAllNodes(&(hosts.etcd_nodes), nodes)
	case compute_key:
		addAllNodes(&(hosts.compute_nodes), nodes)
	case lb_key:
		addAllNodes(&(hosts.lb_nodes), nodes)
	case glusterfs_key:
		addAllNodes(&(hosts.glusterfs_nodes), nodes)
	}
}

func addAllNodes(mP *map[string]struct{}, nodes []string) {
	for _, node := range nodes {
		addNode(mP, node)
	}
}

func addNode(mP *map[string]struct{}, node string) {
	if node == "" {
		return
	}
	m := *mP
	if _, ok := m[node]; ok {

	} else {
		m[node] = struct{}{}
	}
}

func (hosts Hosts) genInv(varString string, path string) {
	log.Debug("===master")
	for k, v := range hosts.master_nodes {
		log.Debug("k:", k, "v:", v)
	}
	log.Debug("===etcd")
	for k, v := range hosts.etcd_nodes {
		log.Debug("k:", k, "v:", v)
	}
}
package inv_gen

import (
	"github.com/hongkailiu/svt-go/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

const (
	node_vars_key = "node_vars"
	master_key = "master"
	etcd_key = "etcd"
	infra_key = "infra"
	compute_key = "compute"
	lb_key = "lb"
	glusterfs_key = "glusterfs"
)

type Args struct {
	ConfigFile string
}

type InventoryGenerator interface {
	Run(args Args) error
}

type myInventoryGenerator struct {

}

func (ig myInventoryGenerator) Run(args Args) error {
	log.Info("Generating inventory file")
	source, err := ioutil.ReadFile(args.ConfigFile)
	if err != nil {
		return err
	}

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(source), &m)
	if err != nil {
		return err
	}
	hostM := Hosts{
		master_nodes: make(map[string]struct{}),
		etcd_nodes: make(map[string]struct{}),
		infra_nodes: make(map[string]struct{}),
		compute_nodes: make(map[string]struct{}),
		lb_nodes: make(map[string]struct{}),
		glusterfs_nodes: make(map[string]struct{}),
	}

	//hostVarM := map[string]map[string]string{}
	log.Info("aaa")

	for k, v := range m {
		log.Debug("k:", k, "v:", v)
		kString := k.(string)
		if strings.Contains(kString, master_key) {
			log.Info("bbb")
			hostM.addNodes(master_key, getNodes(v))
			log.Info("ccc")
		}
		log.Info("===")
		if strings.Contains(kString, infra_key) {
			log.Info("000")
			hostM.addNodes(infra_key, getNodes(v))
		}
		if strings.Contains(kString, etcd_key) {
			log.Info("111")
			hostM.addNodes(etcd_key, getNodes(v))
		}
		if strings.Contains(kString, lb_key) {
			log.Info("222")
			hostM.addNodes(lb_key, getNodes(v))
		}
		if strings.Contains(kString, compute_key) {
			log.Info("333")
			hostM.addNodes(compute_key, getNodes(v))
		}
	}

	hostM.genInv("aaa=bbb", "/tmp/2.file")
	return nil
}

func getNodes(v interface{}) []string {
	aInterface := v.([]interface{})
	aString := make([]string, len(aInterface))
	for _, v := range aInterface {
		aString = append(aString, v.(string))
	}
	return aString
}

func GetInventoryGenerator() InventoryGenerator {
	return myInventoryGenerator{}
}
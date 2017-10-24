package inv_gen

import (
	"github.com/hongkailiu/svt-go/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
	"bytes"
	"os"
	"bufio"
)

const (
	master_key = "master"
	etcd_key = "etcd"
	infra_key = "infra"
	compute_key = "compute"
	lb_key = "lb"
	glusterfs_key = "glusterfs"
	//
	new_line  = "\n"
	openshift_public_hostname = "openshift_public_hostname"
	//
	auto_gen = "\"${auto-gen}\""
	openshift_master_default_subdomain = "openshift_master_default_subdomain"
	openshift_cloudprovider_aws_access_key = "openshift_cloudprovider_aws_access_key"
	openshift_cloudprovider_aws_secret_key = "openshift_cloudprovider_aws_secret_key"
	openshift_hosted_registry_storage_s3_accesskey = "openshift_hosted_registry_storage_s3_accesskey"
	openshift_hosted_registry_storage_s3_secretkey = "openshift_hosted_registry_storage_s3_secretkey"
	//
	AWS_ACCESS_KEY_ID  = "AWS_ACCESS_KEY_ID"
	AWS_SECRET_ACCESS_KEY = "AWS_SECRET_ACCESS_KEY"
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
		if strings.Contains(kString, glusterfs_key) {
			log.Info("333")
			hostM.addNodes(glusterfs_key, getNodes(v))
		}
	}

	return hostM.genInv(args.ConfigFile, "/tmp/2.file")
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


func (hosts Hosts) genInv(configFile string, invFile string) error {
	var buffer bytes.Buffer
	buffer.WriteString("[OSEv3:children]" + new_line)
	buffer.WriteString("masters" + new_line)
	buffer.WriteString("nodes" + new_line)
	buffer.WriteString("etcd" + new_line)
	buffer.WriteString("lb" + new_line)
	buffer.WriteString("glusterfs" + new_line)


	buffer.WriteString(new_line)
	buffer.WriteString("[OSEv3:vars]" + new_line)
	file, err := os.Open(configFile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var isVar = false
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		log.Debug(line)
		if line == "###Do not modify this line:  begin###" {
			isVar = true
		}
		if line == "###Do not modify this line:  end###" {
			isVar = false
		}
		if isVar {
			if strings.HasPrefix(line, "#") || strings.Index(line, ":") < 0 {
				continue
			}
			k:= strings.Trim(line[0:strings.Index(line, ":")], " ")
			v:= strings.Trim(line[strings.Index(line, ":")+1:], " ")
			if v == auto_gen {
				switch k {
				case openshift_master_default_subdomain:
					log.Debug("===111hosts.getSubDomain", hosts.getSubDomain())
					v = hosts.getSubDomain()
				case openshift_cloudprovider_aws_access_key:
					osV := os.Getenv(AWS_ACCESS_KEY_ID)
					if osV != "" {
						v = osV
					}
				case openshift_hosted_registry_storage_s3_accesskey:
					osV := os.Getenv(AWS_ACCESS_KEY_ID)
					if osV != "" {
						v = osV
					}
				case openshift_cloudprovider_aws_secret_key:
					osV := os.Getenv(AWS_SECRET_ACCESS_KEY)
					if osV != "" {
						v = osV
					}
				case openshift_hosted_registry_storage_s3_secretkey:
					osV := os.Getenv(AWS_SECRET_ACCESS_KEY)
					if osV != "" {
						v = osV
					}
				}
			}
			if k!="" && v!="" {
				buffer.WriteString(k+"="+v + new_line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	log.Debug("===lb")
	buffer.WriteString(new_line)
	buffer.WriteString("[lb]" + new_line)
	for k, v := range hosts.lb_nodes {
		log.Debug("k:", k, "v:", v)
		buffer.WriteString(k + new_line)
	}

	log.Debug("===etcd")
	buffer.WriteString(new_line)
	buffer.WriteString("[etcd]" + new_line)
	for k, v := range hosts.etcd_nodes {
		log.Debug("k:", k, "v:", v)
		line := k + " " + openshift_public_hostname + "=" + k
		buffer.WriteString(line + new_line)
	}

	log.Debug("===master")
	buffer.WriteString(new_line)
	buffer.WriteString("[masters]" + new_line)
	for k, v := range hosts.master_nodes {
		log.Debug("k:", k, "v:", v)
		line := k + " " + openshift_public_hostname + "=" + k
		buffer.WriteString(line + new_line)
	}

	buffer.WriteString(new_line)
	buffer.WriteString("[nodes]" + new_line)
	for k, v := range hosts.master_nodes {
		log.Debug("k:", k, "v:", v)
		line := k + " " + openshift_public_hostname + "=" + k + " " + "openshift_node_labels=\"{'region': 'infra', 'zone': 'default'}\" openshift_scheduleable=false"
		buffer.WriteString(line + new_line)
	}
	log.Debug("===infra")
	for k, v := range hosts.infra_nodes {
		log.Debug("k:", k, "v:", v)
		line := k + " " + openshift_public_hostname + "=" + k + " " + "openshift_node_labels=\"{'region': 'infra', 'zone': 'default'}\""
		buffer.WriteString(line + new_line)
	}

	log.Debug("===compute")
	for k, v := range hosts.compute_nodes {
		log.Debug("k:", k, "v:", v)
		line := k + " " + openshift_public_hostname + "=" + k + " " + "openshift_node_labels=\"{'region': 'primary', 'zone': 'default'}\""
		buffer.WriteString(line + new_line)
	}
	log.Debug("===glusterfs")
	buffer.WriteString(new_line)
	buffer.WriteString("[glusterfs]" + new_line)
	for k, v := range hosts.glusterfs_nodes {
		log.Debug("k:", k, "v:", v)
		line := k + " " + openshift_public_hostname + "=" + k + " " + "openshift_node_labels=\"{'region': 'primary', 'zone': 'default'}\""
		buffer.WriteString(line + new_line)
	}

	return ioutil.WriteFile(invFile, buffer.Bytes(), 0644)
}


package cluster_loader

import (
	"github.com/hongkailiu/svt-go/task"
	"github.com/hongkailiu/svt-go/log"
	"github.com/hongkailiu/svt-go/oc"
	"fmt"
)

type Args struct {
	PoolSize uint
	ConfigFile string
}

type ClusterLoader interface {
	Run(args Args) error
}

type myClusterLoader struct {

}

func (cl myClusterLoader) Run(args Args) error {
	configP, err := task.LoadFromFile(args.ConfigFile)
	if err != nil {
		return err
	}
	//config := *configP
	log.Debug(*configP)
	oc.SetPoolSize(args.PoolSize)

	oc.WhoAmI()

	for i, project := range (*configP).Projects {
		oc.IsProjectExisting(fmt.Sprintf("%s%d", project.Basename, i))
	}


	return nil
}

func GetClusterLoader() ClusterLoader {
	return myClusterLoader{}
}
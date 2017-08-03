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
		projectName := fmt.Sprintf("%s%d", project.Basename, i)
		isExisting, _ := oc.IsProjectExisting(projectName)
		log.Debug(fmt.Sprintf("isExisting: %t", isExisting))
		if !isExisting {
			oc.CreateProject(projectName)
			oc.Label(oc.Namespace, projectName, "purpose", "test", "--overwrite")

		} else {
			log.Fatal(fmt.Sprintf("project %s already exitsed", projectName))
		}
	}


	return nil
}

func GetClusterLoader() ClusterLoader {
	return myClusterLoader{}
}
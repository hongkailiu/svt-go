package cluster_loader

import (
	"github.com/hongkailiu/svt-go/task"
	"github.com/hongkailiu/svt-go/log"
	"github.com/hongkailiu/svt-go/oc"
	"fmt"
	"errors"
	"sync"
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
	defer oc.ClosePool()

	oc.WhoAmI()

	if err := prepareProjects(configP); err != nil {
		return err
	}
	var wg sync.WaitGroup
	for _, project := range configP.Projects {
		tuning := project.Tuning
		tuningSet := findTuningSet(tuning, configP.TuningSets)
		wg.Add(1)
		oc.HandleProject(&project, &wg, tuningSet)
	}
	log.Info("wait until all projects are handled ...")
	wg.Wait()
	log.Info("all projects are handled")
	return nil
}

func findTuningSet(tuning string, tuningSets []task.TuningSet) *task.TuningSet {
	for _, tuningSet := range tuningSets {
		if tuningSet.Name == tuning {
			return &tuningSet
		}
	}
	return nil
}

func prepareProjects(config *task.Config) error {
	for _, project := range config.Projects {
		for i := 0; i < project.Number; i++ {
			projectName := fmt.Sprintf("%s%d", project.Basename, i)
			isExisting, _ := oc.IsProjectExisting(projectName)
			log.Debug(fmt.Sprintf("isExisting: %t", isExisting))
			if !isExisting {
				oc.NewProject(projectName)
				oc.Label(oc.Namespace, projectName, "purpose", "test", "--overwrite")

			} else {
				return errors.New(fmt.Sprintf("project %s already exitsed", projectName))
			}
		}
	}
	return nil
}

func GetClusterLoader() ClusterLoader {
	return myClusterLoader{}
}
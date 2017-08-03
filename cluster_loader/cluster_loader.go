package cluster_loader

type ClusterLoader interface {
	Run(args Args) error
}

type Args struct {
	PoolSize int
	ConfigFile string
}

type myClusterLoader struct {

}


func (cl myClusterLoader) Run(args Args) error {
	return nil
}

func GetClusterLoader() ClusterLoader {
	return myClusterLoader{}
}
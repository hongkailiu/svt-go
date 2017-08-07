package task

import (
	"testing"
	"github.com/hongkailiu/svt-go/log"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestRootHandler(t *testing.T) {

	assert := assert.New(t)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	configP, err := LoadFromFile(dir+"/../conf/pyconfigMasterVirtScalePause.yaml")
	assert.Nil(err)
	config := *configP
	assert.NotNil(config)

	assert.Len(config.Projects, 1)
	p := config.Projects[0]
	assert.Equal(3, p.Number)
	assert.Equal("default", p.Tuning)
	assert.Equal("c", p.Basename)

	assert.Len(p.Templates, 7)
	t0 := Template{Number:3, File:"content/build-config-template.json", Parameters:nil}
	t1 := Template{Number:6, File:"content/build-template.json", Parameters:nil}
	t2 := Template{Number:1, File:"content/image-stream-template.json", Parameters:nil}
	t3 := Template{Number:2, File:"content/deployment-config-1rep-pause-template.json",
		Parameters:map[string]string{"ENV_VALUE":"asodfn209e8j0eij0emc2oed2ed2ed2e2easodfn209e8j0eij0emc2oed2ed2ed2e2easodfn209e8j0eij0emc2oed2ed2ed2e2easodfn209e8j0eij0emc2oed2ed2ed2e2easodfn209e8j0eij0emc2oed2ed2ed2e2easodfn209e8j0eij0emc2oed2ed2ed2e2easodfn209e8j0eij0emc2oed2ed2ed2e2easodfn209e8j0eij12"}}
	t4 := Template{Number:1, File:"content/deployment-config-2rep-pause-template.json",
		Parameters:map[string]string{"ENV_VALUE":"asodfn209e8j0eij0emc2oed2ed2ed2e2easodfn209e8j0eij0emc2oed2ed2ed2e2easodfn209e8j0eij0emc2oed2ed2ed2e2easodfn209e8j0eij0emc2oed2ed2ed2e2easodfn209e8j0eij0emc2oed2ed2ed2e2easodfn209e8j0eij0emc2oed2ed2ed2e2easodfn209e8j0eij0emc2oed2ed2ed2e2easodfn209e8j0eij12"}}
	t5 := Template{Number:20, File:"content/ssh-secret-template.json"}
	t6 := Template{Number:3, File:"content/route-template.json"}
	assert.Contains(p.Templates, t0)
	assert.Contains(p.Templates, t1)
	assert.Contains(p.Templates, t2)
	assert.Contains(p.Templates, t3)
	assert.Contains(p.Templates, t4)
	assert.Contains(p.Templates, t5)
	assert.Contains(p.Templates, t6)

	assert.Len(config.TuningSets, 1)
	ts := TuningSet{Name:"default"}
	ts.PodsInTuningSet.Stepping.StepSize = 5
	ts.PodsInTuningSet.Stepping.Pause = "10 s"
	ts.PodsInTuningSet.RateLimit.Delay = "250 ms"
	assert.Contains(config.TuningSets, ts)
}
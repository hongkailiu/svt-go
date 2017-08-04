package task

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

func LoadFromBytes(source []byte) (*Config, error) {
	var config Config
	err := yaml.Unmarshal(source, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func LoadFromFile(file string) (*Config, error) {
	source, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return LoadFromBytes(source)
}

type Config struct {
	Projects   []Project `yaml:"projects"`
	Quotas     []Quota `yaml:"quotas"`
	TuningSets []TuningSet `yaml:"tuningsets"`
}

type Project struct {
	Number   int `yaml:"num"`
	Basename string `yaml:"basename"`
	Tuning   string `yaml:"tuning"`
	Quota    string `yaml:"quota"`
	Users    []User `yaml:"users"`
	Templates    []Template `yaml:"templates"`
}

type User struct {
	Number       int `yaml:"num"`
	Role         string `yaml:"role"`
	Basename     string `yaml:"basename"`
	Password     string `yaml:"password"`
	UserPassFile string `yaml:"userpassfile"`
}

type Template struct {
	Number   int `yaml:"num"`
	File     string `yaml:"file"`
	Parameters    map[string]string `yaml:"parameters"`
}

type Quota struct {
	Name string `yaml:"name"`
	File string `yaml:"file"`
}

type TuningSet struct {
	Name            string `yaml:"name"`
	PodsInTuningSet struct {
				RateLimit struct {
						  Delay string `yaml:"delay"`
					  } `yaml:"rate_limit"`
				Stepping  struct {
						  StepSize int `yaml:"stepsize"`
						  Pause    string `yaml:"pause"`
					  } `yaml:"stepping"`
			}   `yaml:"pods"`
}
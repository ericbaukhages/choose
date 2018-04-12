package choose

import (
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func TestAdd(t *testing.T) {
	config := makeTestConfig()
	name := "bob"

	err := config.Add(name, "~")
	if err != nil {
		t.Errorf("Adding failed, %v", err)
	}

	if config.Values[name] == "" {
		t.Errorf("\"%s\" was not saved in Values", name)
	}

	if !stringInSlice(name, config.keys) {
		t.Errorf("\"%s\" was not saved in keys", name)
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func makeTestConfig() Config {
	values := make(map[string]string)
	//TODO: catch error from Expand
	values["go"], _ = homedir.Expand("~/go/src/ebaukhages")

	keys := makeKeys(values)

	config := Config{
		Values: values,
		keys:   keys,
	}
	return config
}

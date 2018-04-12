package choose

import (
	"os"
	"reflect"
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

func TestSave(t *testing.T) {
	testConfigFile := "test.config"

	config := makeTestConfig()
	config.Location = testConfigFile

	err := config.Save()
	if err != nil {
		t.Errorf("Save failed: %v", err)
	}

	newConfig := Config{
		Location: testConfigFile,
	}
	newConfig.Parse()

	if !reflect.DeepEqual(config.keys, newConfig.keys) {
		t.Errorf("New config keys are not equal to old config keys")
	}

	if !reflect.DeepEqual(config.Values, newConfig.Values) {
		t.Errorf("New config values are not equal to old config values")
	}

	err = os.Remove(testConfigFile)
	if err != nil {
		t.Errorf("Unable to delete test config: %v", err)
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

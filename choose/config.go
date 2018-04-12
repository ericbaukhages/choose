package choose

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

// Config stores configuration data as well as its location
type Config struct {
	Location string
	Values   map[string]string
	keys     []string
}

// Parse sets up the config
func (c *Config) Parse() {
	c.Values = parseConfig(c.Location)
	c.keys = makeKeys(c.Values)
}

func parseConfig(logFile string) map[string]string {
	f, err := os.OpenFile(logFile, os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("open file error: %v", err)
	}

	defer f.Close()

	config := make(map[string]string)
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
		}

		values := strings.Split(line, " ")
		name, path := values[0], strings.TrimSpace(values[1])
		config[name] = path
	}

	return config
}

func makeKeys(config map[string]string) []string {
	keys := make([]string, len(config))

	i := 0
	for k := range config {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	return keys
}

// IsValid is a function
func (c *Config) IsValid(name string, path string) error {
	// check if path is valid directory
	stat, err := os.Stat(path)
	if !(err == nil && stat.IsDir()) {
		return fmt.Errorf("path is not valid directory")
	}

	// check if name already exists
	if c.Values[name] != "" {
		return fmt.Errorf("%s already exists", name)
	}

	return nil
}

// Add a new session configuration
func (c *Config) Add(name string, path string) error {
	path, err := homedir.Expand(path)
	if err != nil {
		return err
	}

	err = c.IsValid(name, path)
	if err != nil {
		return err
	}

	c.keys = append(c.keys, name)
	c.Values[name] = path

	return nil
}

// Save a new configuration object
func (c *Config) Save() error {
	// create the file if it doesn't already exist
	_, err := os.Stat(c.Location)
	if os.IsNotExist(err) {
		file, err := os.Create(c.Location)
		if err != nil {
			return err
		}
		file.Close()
	}

	// open the file
	file, err := os.OpenFile(c.Location, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// write config to file
	for key, value := range c.Values {
		_, err := file.WriteString(fmt.Sprintf("%s %s\n", key, value))
		if err != nil {
			return err
		}
	}

	// save file
	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}

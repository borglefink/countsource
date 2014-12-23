// Copyright 2014 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
 * config.go
 *
 */
package config

// Importing libraries
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

/*
 * Config
 */
type Config struct {
	Extensions         []string
	BinaryExtensions   []string
	ExcludeDirectories []string
	ExcludeFiles       []string
}

var configFilename = ""

func init() {
	// Set the config file name to [thisexecutablefilename].config
	configFilename = strings.Replace(filepath.Base(os.Args[0]), ".exe", ".config", 1)
}

func LoadConfig() Config {
	var sc Config

	// Read whole the file
	var jsonstring, err = ioutil.ReadFile(configFilename)
	if err != nil {
		createConfig()

		// Read the file again, now it should exist
		jsonstring, err = ioutil.ReadFile(configFilename)
	}

	err = json.Unmarshal(jsonstring, &sc)

	if err != nil {
		fmt.Printf("json.Unmarshal, %v\n", err)
		sc = Config{}
	}

	return sc
}

/*
 * createConfig - initializes a config file if one is not found
 */
func createConfig() {
	var sc Config

	// Extensions to count
	sc.Extensions = []string{
		".css",
		".go",
		".htm",
		".html",
		".js",
		".json",
		".less",
		".sass",
		".xml",
		".xsd",
	}

	// Binary extensions to count
	sc.BinaryExtensions = []string{
		".gif",
		".ico",
		".jpg",
		".png",
	}

	// Directories to exclude
	sc.ExcludeDirectories = []string{
		".svn",
		"bin",
		"obj",
		"Scripts\\jquery.ui",
		"_svn",
	}

	// Directories to exclude
	sc.ExcludeFiles = []string{
		configFilename,
		"jquery.min.js",
		"jquery.ui.js",
	}

	save(sc)
}

/*
 * Save Config to disk as JSON
 */
func save(sc Config) {
	var jsonstring, err = json.MarshalIndent(&sc, "", "  ")
	if err != nil {
		fmt.Printf("json.Marshal(sc), %s %v\n", string(jsonstring), err)
	}

	err = ioutil.WriteFile(configFilename, jsonstring, 0666)
	if err != nil {
		fmt.Printf("ioutil.WriteFile, %v\n", err)
	}
}

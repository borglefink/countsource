// Copyright 2014 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"utils"
)

// Config
type Config struct {
	Extensions         []string
	BinaryExtensions   []string
	ExcludeDirectories []string
	ExcludeFiles       []string
}

var pathSeparator = "/"
var configFilename = ""

// ------------------------------------------
// init
// ------------------------------------------
func init() {
	pathSeparator = utils.GetPathSeparator()

	var fullFilePath, _ = filepath.Abs(os.Args[0])

	// Set the config file name to [thisexecutablepath\thisexecutablefilename].config
	configFilename = fmt.Sprintf("%s%s%s", filepath.Dir(fullFilePath), pathSeparator, strings.Replace(filepath.Base(fullFilePath), ".exe", ".config", 1))
}

// ------------------------------------------
// LoadConfig
// ------------------------------------------
func LoadConfig() Config {
	var c Config

	// Read whole the file
	var jsonstring, err = ioutil.ReadFile(configFilename)
	if err != nil {
		createConfig()

		// Read the file again, now it should exist
		jsonstring, err = ioutil.ReadFile(configFilename)
	}

	// Strip comments from config file
	var re = regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")
	var newJsonstring = re.ReplaceAll(jsonstring, nil)

	// Create config to be returned
	err = json.Unmarshal(newJsonstring, &c)

	if err != nil {
		fmt.Printf("Could not read the config, %v\n", err)
		c = Config{}
	}

	return c
}

// ------------------------------------------
// createConfig
// ------------------------------------------
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
		fmt.Sprintf("Scripts%sjquery.ui", pathSeparator),
		"_svn",
	}

	// Directories to exclude
	sc.ExcludeFiles = []string{
		filepath.Base(configFilename),
		"jquery.min.js",
		"jquery.ui.js",
	}

	save(sc)
}

// ------------------------------------------
// save
// ------------------------------------------
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

[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.png?v=103)](https://opensource.org/licenses/mit-license.php)
[![works badge](https://cdn.rawgit.com/nikku/works-on-my-machine/v0.2.0/badge.svg)](https://github.com/nikku/works-on-my-machine)

## Description

*countsource* is a small command line utility for counting source code lines, written in Go (http://golang.org/). 
It can also count binaries (number of files and filesize).
There is a config file to configure what files to count (see config section below).
When counting source code lines, newline will be determined as type windows (CRLF), unix (LF) or older mac (CR) for each file.

The result will look along the lines of this:
```
Directory processed:
c:\projectdirectory
---------------------------------------------------------------
filetype        #files       #lines  line%          size  size%
---------------------------------------------------------------
.css                 9        3 512   42.3        92 168   12.0
.html                1          229    2.8         7 626    1.0
.js                 22        4 563   54.9       195 256   25.5
.jpg                 7                           260 274   33.9
.png               120                           211 318   27.6
---------------------------------------------------------------
Total:             159        8 304  100.0       766 642  100.0
```

## Usage

Give a directory as a parameter. If none is given, the current directory is used.
All sub-directories will be searched as well, and included in the result.

```
countsource [directory] [-c fullpathtoconfigfile] [-dir] [-file] [-inc] [-excl]
```

The optional parameters *-dir*, *-file* *-inc* and *-excl* are for analysis/debug, showing which directories or files are included or excluded. If *-dir* or *-file* is given, both included and excluded items are shown, unless *-inc* or *-excl* are added.

Use *countsource -?* to show usage.

## Config file

The configuration file can be specified with *-c "full config file name"*. 

If no config file is specified, it is read from the current directory. 
If not found in the current directory, the config is expected to be found in the same directory as the executable. 

If a config file does not exist, one is created in the current directory,
with default values similar to the following:

```JSON
/*
 * Config file for example project
 * FOr binary files, only file size is counted
 */
{
  "Extensions": [
    ".css",
    ".html",
    ".js",
    ".jpg",
    ".png"
  ],
  "ExcludeDirectories": [
    "bin",
    "obj",
    "Scripts\\jquery.ui"
  ],
  "ExcludeFiles": [
    "countsource.config",
    "jquery.min.js"
  ]
}
```

Note that directories can be specified in a bit more detail, to be sure to exclude elements you don't want counted.

It is possible to put comments in the config file. 
Note that comments are normally not allowed in json, so these comments are stripped from the config file before it is read.
Only Go-type comments are allowed, single line comments starting with //, or block comments enclosed by /* and */.

## Installation

Clone the repository into your GOPATH somewhere and do a **go build**.
Create a config file for a project you want to count source code for, and put the config file in the root of that directory.
If you have several projects using identical config files, use a single config file and refer to it with the *-c* parameter when counting.

## Background

I wanted to count the number of source code lines for all the source code in an ASP.NET MVC project to keep track of the size of it. So I just wrote this.

## License

A MIT license is used here - do what you want with this. Nice though if improvements and corrections could trickle back to me somehow. :-)

## Description

*countsource* is a small command line utility for counting source code lines, written in Go (http://golang.org/). 
It can also count binaries (number of files and filesize).
There is a config file to configure what to count (see config section below).
When counting source code lines, newline will be determined as type windows (CRLF), unix (LF) or older mac (CR) for each file.

The result will look along the lines of this:
```
Directory processed:
c:\mydirectory\exampleprojectdirectory
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
All sub-directories will be included in the result as well.

```
countsource [directory] [--dir] [--file] [--inc] [--excl]
```

The optional parameters is for analysis/debug, showing which directories or files are 
included or excluded. By default shows all files and directories.

Use *countsource --?* to show usage.

## Config file

The config file is expected to be found in the same directory as the executable.
If a config file does not exist, one is created, with simple default values 
along the lines of this:

```JSON
/*
 * Config file for exampleproject
 */
{
  "Extensions": [
    ".css",
    ".html",
    ".js"
  ],
  "BinaryExtensions": [
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
Note that comments are not allowed in json, so these comments are stripped from the config when interpreting it.
Only Go-type comments are allowed, like line type comments starting with //, or block comments enclosed by /* and */.

## Background

I wanted to count the number of source code lines for all the source code in an ASP.NET MVC project 
to keep track of the size of it. So I just wrote this.

## License

A MIT license is used here - do what you want with this. Nice though if improvements and corrections could trickle back to me somehow. :-)

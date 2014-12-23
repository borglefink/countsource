countsource
===========
A small command line utility written in Go (http://golang.org/), for counting source code lines. 
Also counts binaries, but only number of files and filesize.
Will for each file determine if newline is of type windows (CRLF), unix (LF) or older mac (CR), 
and then use the correct one when counting number of source code lines in file.

The result will look along the lines of this:
```
Directory processed:
c:\mydirectory\myproject
-------------------------------------------------
filetype        #files       #lines          size
-------------------------------------------------
.css                10        3 583        95 921
.html                1           54         3 722
.js                 23        4 628       197 086
.jpg                 7                    260 274
.png               307                    495 174
-------------------------------------------------
Total:             348        8 265     1 052 177
```

Usage
---------
Give a directory as a parameter. If none is given, the executable's directory is used. 
All sub-directories will be included in the result as well.

```
countsource [directory]
```

Config file
---------
If a config file does not exist, one is created, with simple default values. 
Along the lines of this:

```JSON
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

Background
----------
I wanted to count the number of source code lines for all the source code in an ASP.NET MVC project to keep track of the size of it. So I just wrote this.

License
----------
A MIT license is used here - do what you want with this. Nice though if improvements and corrections could trickle back to me somehow. :-)

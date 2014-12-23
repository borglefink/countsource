countsource
===========
A small Go utility for counting source code lines.

Usage
---------
Give a directory as a parameter. If none is given, the executable's directory is used.
```
countsource [directory]
```

Config file
---------
If a config file does not exist, one is created, with simple default values. Along the lines of this:

```json
{
  "Extensions": [
    ".css",
    ".html",
    ".js",
  ],
  "BinaryExtensions": [
    ".jpg",
    ".png"
  ],
  "ExcludeDirectories": [
    "bin",
    "obj",
    "Scripts\\jquery.ui",
  ],
  "ExcludeFiles": [
    "countsource.config",
    "jquery.min.js",
  ]
}
```
Note that directories can be specified further, to be sure to exclude elements you don't want counted.

Background
----------
I wanted to count the number of source code lines for all the source code in an ASP.NET MVC project to keep track of the size of it. So I just wrote this.

License
----------
A MIT license is used here - do what you want with this. Nice though if improvements and corrections could trickle back to me somehow. :-)

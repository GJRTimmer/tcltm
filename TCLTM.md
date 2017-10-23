# .tcltm Specification

A ```tcltm``` specification is defined in [Yaml](https://en.wikipedia.org/wiki/YAML).

Multiple packages can be defined within a single ```.tcltm``` file, the ```Package``` key
is an array of Tcl packages.

# Required Keys
The following keys are required.
- Name
- Version
- Tcl
- Files

# Binary Files
All files defined with the key ```Files``` which do not have the extension ```.tcl``` are treated as
binary files.

Additionally there is the possibility to auto-load libraries.
This can be achieved be defining the file as ```file:load```.

The ```.tm``` file will extract the embedded libary and directly issue the Tcl ```load``` command.


```yaml
---
Package:
  - 
    Name: sqlite
    Version: 3.7.14.1
    Summary: SQLite for Tcl
    Description: Package for handling SQLite
    License: ~
    Tcl: 8.6
    Dependencies: ~
    Files:
      - sqlite.tcl
      - tclsqlite3.dll:load
    InitScript: ~
```

# Example

```yaml
--- # Package SMTP
Package:
  - 
    Name: smtp
    Version: 1.4.5
    Summary: ~
    Description: ~
    License: ~
    Tcl: 8.3
    Dependencies:
      - mime 1.4.1
      - SASL 1.0
      - SALS::NTLM 1.0
    Files:
      - smtp.tcl
    InitScript: ~
```
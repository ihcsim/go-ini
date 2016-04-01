# go-ini
An INI config parser.

## Introduction
The INI file format is an informal standard for configuration files for some platforms or software. INI files are simple text files with a basic structure composed of sections, properties, and values. 

Although made popular by Windows, INI files can be used on any system thanks to their flexibility. They allow a program to store configuration data, which can then be easily parsed and changed. Two notable systems that use the INI format are Samba and Trac. Linux and Unix systems also use a similar file format for system configuration. In addition, platform-agnostic software may use this file format for configuration. It is human-readable and simple to parse, so it is a usable format for configuration files that do not require much greater complexity.

For more information on the INI file format, refer to this [Wikipedia page](https://en.wikipedia.org/wiki/INI_file).

## INI File Structure
The basic element contained in an INI file is the property. A property is a name/value pair, delimited by an equals sign =.

```
name=value
```

Section declarations start with [ and end with ]. All properties after the section declaration will be associated with that section.

```
[section1]
var1=val1

[section2]
var2=val2
```

All lines beginning with a semicolon ; or a number sign # are considered to be comments. Comment lines are ignored when parsing INI files.

Example of an INI file:

```
[section1]
; some comment on section1
var1 = val1
var2 = val2

[section2]
# another comment
var1 = val1
var2 = val2
```

## Getting started

To get the library:

```sh
go get github.com/ihcsim/go-ini
go test -v -cover github.com/ihcsim/go-ini
```

To use in code:

```go
import (
  ini "github.com/ihcsim/go-ini"
  "fmt"
)

func main(){
  // ...

  var iniMap map[string]map[string]string
  iniMap = ini.Parse(someInIString)
  fmt.Printf("%+v\n", iniMap)

  // ...
}
```

## License
This project is under Apache v2 License. See the [LICENSE file](LICENSE) for the full license text.

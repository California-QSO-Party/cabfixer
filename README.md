CABFIXER 0.2
========

Cabfixer detects columns in cabrillo files.
It outputs a copy of the file with pipe-separated
columns in QSO lines.

Installation
------------
### Windows

### Linux
To install on Linux first install golang toolchain:
```
$ sudo apt-get install golang
```
Enter the project's root directory:
```
$ cd <project's root directory>
```
Build the project:
```
$ go build
```

Usage
-----

```
$ ./cabfixer *.log
```

Cabfixer accepts globs in command line arguments.

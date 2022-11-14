Rank is a command line utility for calculating the rankings given a series of game results.

Building
========

To run the unit tests:

```shell
go test -v ./...
```

To build _rank_ run the following from a command line terminal:

```shell
go get
go build -o rank *.go
```

An executable named _rank_ will be produced that is compatible with the architecture that it was built on. 

Usage
=====

The use the compiled utility, run the following from a command line terminal:

#### POSIX OS:

```shell
./rank PATH_TO_INPUT # POSIX
```

##### Windows:

```shell
rank.exe PATH_TO_INPUT # Windows
```

The _PATH_TO_INPUT_ placeholder must be replaced with the path to an actual input file.
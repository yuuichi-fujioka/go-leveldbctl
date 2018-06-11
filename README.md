# leveldbctl

[![Build Status](https://travis-ci.org/yuuichi-fujioka/go-leveldbctl.svg?branch=master)](https://travis-ci.org/yuuichi-fujioka/go-leveldbctl)

LevelDB control command.

This command provides easy way to CRUD operation on LevelDB.

```
$ leveldbctl put foo bar
put foo: bar into ./.
$ leveldbctl get foo
bar
```

## Features

* Initialize LevelDB
* Put key/value into LevelDB
* Get value with key
* Delete key
* Dump all key/values in LevelDB
* Print all keys

## Install

`$ go get github.com/yuuichi-fujioka/go-leveldbctl/cmd/leveldbctl`

## Usage

```
$ leveldbctl -h
NAME:
   leveldbctl - A new cli application

USAGE:
   leveldbctl [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     init, i    Initialize a LevelDB
     walk, w    Walk in a LevelDB
     keys, k    Serach all keys in a LevelDB
     put, p     Put a value into a LevelDB
     get, g     Gut a value from a LevelDB
     delete, d  Delete a value from a LevelDB
     help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --dbdir value, -d value  LevelDB Directory (default: "./") [$LEVELDB_DIR]
   --help, -h               show help
   --version, -v            print the version
```


# leveldbctl

[![Build Status](https://travis-ci.org/yuuichi-fujioka/go-leveldbctl.svg?branch=master)](https://travis-ci.org/yuuichi-fujioka/go-leveldbctl)
[![Coverage Status](https://coveralls.io/repos/github/yuuichi-fujioka/go-leveldbctl/badge.svg?branch=master)](https://coveralls.io/github/yuuichi-fujioka/go-leveldbctl?branch=master)

LevelDB control command.

This command provides easy way to CRUD operation on LevelDB.

```sh
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

```sh
$ export GO111MODULE=on
$ go get github.com/yuuichi-fujioka/go-leveldbctl/cmd/leveldbctl
```

## Usage

```sh
NAME:
   leveldbctl - A new cli application

USAGE:
   leveldbctl [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     init, i    Initialize a LevelDB
     walk, w    Walk in a LevelDB
     keys, k    Search all keys in a LevelDB
     put, p     Put a value into a LevelDB
     get, g     Gut a value from a LevelDB
     delete, d  Delete a value from a LevelDB
     help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --dbdir value, -d value  LevelDB Directory (default: "./") [$LEVELDB_DIR]
   --hexkey, --xk           get / put hexadecimal keys
   --hexvalue, --xv         get / put hexadecimal values
   --help, -h               show help
   --version, -v            print the version
```

For hexadecimal keys and values:

```sh
$ export LEVELDB_DIR=${HOME}/.bitcoin/index
$ leveldbctl -xk g 62f2a1f90489f1f74e441f325ec6f532df8286847d7c7a14000000000000000000|xxd -p
89fe04a3db1d801d92188d350880fec55300008020edd4d15faba7c63dd7
c83961bf6783a691fb8f5f6887120000000000000000009f413c1df7e296
4af9babb54e46d4414eaad550b27b409e29ab80a832ac64ce9966ab95ddf
8e1417baf3db320a
```


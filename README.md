# golorich

## Overview

Golang log standard package extension Logger

## Description

- Add a prefix of each log level to the log.
- You can specify the log level to output. Logs less than the specified log level are not output.<br>
For example, error in production environment, debugging in development environment.

### e.g.

```golang
logger.Debugln() // [DEBUG]
logger.Infoln() // [INFO]
logger.Warnln() // [WARN]
logger.Errorln() // [ERROR]
logger.Fatalln() // [FATAL]
```

## Requirement

- [github.com/hashicorp/logutils](https://github.com/hashicorp/logutils)

## Install

```sh
go get github.com/dip-dev/golorich
```

## Usage

```golang
logger := golorich.New(os.Stdout, "prefix_", log.LstdFlags|log.Lshortfile, golorich.Debug)
logger.Debugf("logging: %s", "debug message")
// prefix_2018/09/04 04:12:11 main.go:52: [DEBUG] logging: debug message

logger.Errorln("logging:", "error", "message")
// prefix_2018/09/04 04:12:11 main.go:53: [ERROR] logging: error message
```

### Use environment variable for log level

```golang
logger := golorich.New(os.Stdout, "", log.LstdFlags|log.Lshortfile, golorich.GetLevelFromString(os.Getenv("LOG_LEVEL")))
```

## Feature

- Specifying log format
- Specify log output destination

## Licence

MIT License

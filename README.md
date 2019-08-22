# Star Trek

[![Build Status](https://travis-ci.org/hackercompany/StarTrek.svg?branch=master)](https://travis-ci.org/hackercompany/StarTrek) [![Go Report Card](https://goreportcard.com/badge/github.com/hackercompany/MayTheForceBeWithYou)](https://goreportcard.com/report/github.com/hackercompany/MayTheForceBeWithYou)

This application is to translate English names to their Klingon equivalent and also gives the species of the given character(if exists) taking data from http://stapi.co

Caching has been implemented into the system so as to preserver on api calls and a file storage is used. Once a succesful api call is made, it is stored in `data.txt` and future occurances of the same name are served from cache.

The external library stapi is implemented as an interface and injected into the cache mechanism dynamically. This makes the system future proof in terms of integration of other external libs.

The codex implements a Klingon model which is in form of a lookup table. The Klingon dictionary is stored as a json file in the codex directory.

# Installation
  - Download and install GO version 1.12 from [here](https://golang.org/dl/)
  - `make` should run the required test cases and build the binary

# Run
```
./StarTrek Uhura
0xF8E5 0xF8D6 0xF8E5 0xF8E1 0xF8D0
Human
```

# Verbose
Add flag `-v` to print verbose logging

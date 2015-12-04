#!/bin/sh
godoc2md $(pwd | gsed "s%$(go env GOPATH)/src/%%") |ghead -n -5 > README.md

#!/bin/sh
godoc2md $(pwd | sed "s%$(go env GOPATH)/src/%%") |head -n -5 > README.md

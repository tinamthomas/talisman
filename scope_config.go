package main

const 	scopeConfig = `
go:
 - makefile
 - go.mod
 - go.sum
 - Gopkg.toml
 - Gopkg.lock
 - glide.yaml
 - glide.lock
 - vendor/
java:
 - .gradle
node:
 - yarn.lock
 - package-lock.json
 - node_modules/
`
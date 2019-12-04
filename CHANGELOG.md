# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

Fix docker-compose file.

Attempt makefile - needs a script to shut down only the shippy containers. Currently it stops all running containers.

## v0.0.6 [2019-12-04]

Fixes from tutorial repo

## v0.0.5 [2019-12-04]

Much refactoring and jiggering about with repos.

## v0.0.4 [2019-12-02]

Both services now running and talking to each other in docker with go-micro enabled.

## v0.0.3 [2019-12-01] Part 1 completed

cli now talking to the service and getting a list of consignments added.

## v0.0.2 [2019-12-01] Rebooting code

Err redo from start.

## v0.0.1 [2019-11-30] Part 1 Completed

Initial project creation.

Initial gRPC server built.

cli service created. Note the tutorial says to run `go run main.go` but you need to run `go run cli.go`.

Added method to get a list of consignments.

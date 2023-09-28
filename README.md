# EMR Pricing API

## Introduction

The EMR Pricing API is a Go-based web service that provides pricing information for EMR and EC2 instances.

## Prerequisites

Before getting started, make sure you have the following prerequisites installed:

- [Go](https://golang.org/doc/install): The Go programming language.
- [Echo](https://github.com/labstack/echo): A high-performance, minimalist web framework for Go.

## How to run
$go run main.go

## Example of API requests
- $curl -X GET "http://localhost:8080/price?serviceCode=ElasticMapReduce&instanceType=m4.large"
- $curl -X GET "http://localhost:8080/price?serviceCode=AmazonEC2&instanceType=C6g.12xlarge"

## How to test
$go test ./...
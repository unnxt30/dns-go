# dns-go

A DNS resolver written in Go. This project is a recreational project or an attempt to get a basic understanding of how DNS works.

## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)

## Introduction

This project is a simple DNS resolver implemented in Go. It sends DNS queries to a specified DNS server and parses the responses. The project aims to provide a basic understanding of how DNS works and how to implement a DNS resolver from scratch.

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/unnxt30/dns-go.git
    cd dns-go
    ```

2. Install the required Go version (1.22.1 or later).

3. Build the project:

    ```sh
    go build -o dns-go main.go
    ```

## Usage

To use the DNS resolver, run the compiled binary with the `-domain` flag to specify the domain you want to resolve:

```sh
./dns-go -domain example.com
```

# go-vbox-api

`go-vbox-api` is a Go module designed to interact with the VirtualBox SOAP API by making calls to the `vboxwebsrv` service. This module provides an interface for managing VirtualBox VMs programmatically, allowing you to create, configure, and control VMs directly from your Go applications.

SOAP glue code generated using [gowsdl](https://github.com/hooklift/gowsdl) 

## Features

- Compatible with VirtualBox SOAP API based on VirtualBoxSDK-7.0.8-156879
## Features (desired)

- Simple and intuitive API for managing VirtualBox VMs built on top the SOAP glue code
- Support for common VM operations such as create, start, stop, pause, and resume
- Access to VM settings like memory, CPU, storage, and networking configurations

## Prerequisites

- VirtualBox (version 6.0 or later) installed on your system
- `vboxwebsrv` running and properly configured to allow connections from your Go application

## Installation

To install the `go-vbox-api` module, use the following command:

```sh
go install -u github.com/0n0sendai/go-vbox-api
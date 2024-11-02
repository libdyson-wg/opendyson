# opendyson

[![GitHub (Pre-)Release Date](https://img.shields.io/github/release-date-pre/libdyson-wg/opendyson)](https://github.com/libdyson-wg/opendyson/releases/)
[![Latest Release](https://badgen.net/github/release/libdyson-wg/opendyson)](https://github.com/libdyson-wg/opendyson/releases/)
[![Latest Commit](https://badgen.net/github/last-commit/libdyson-wg/opendyson/main)](https://github.com/libdyson-wg/opendyson/commit/HEAD)

A library for talking to Dyson devices in Go, with a simple CLI tool.

## Installation

### `go install` method

If you have Go installed and configured on your system, you can install it with `go install github.com/libdyson-wg/opendyson` 

### Download and run

You can download the latest portable executables from the [releases page](https://github.com/libdyson-wg/opendyson/releases), selecting the
file that matches your operating system and platform. After downloading the file, you may need to mark it as executable on your system.

After downloading the file, you can run it from a terminal or command prompt.

## Using the opendyson CLI

You can run `opendyson help` to get a list of commands.

The first time you run any command, or upon running `opendyson login`, you will be prompted to log into your MyDyson account. After a
successfully logging in, your login session will be saved to a config file for future use. The location of the file will be displayed,
should you wish to review or delete it.

### List Devices and Data

`opendyson devices` will fetch all the information available about your devices from the Dyson API, and will use Zeroconf to attempt to
find each device's IP Address on your local network. Upon completion, all registered devices, with all available information will be
displayed in yml format. 

The output of this includes sensitive information which should **not** be shared with anyone you do not trust. IoT Credentials will allow
anyone to control a device or read its sensors remotely from anywhere in the world.

### Listen to MQTT messages

`opendyson listen SERIALNUMBER` will subscribe to the device's status, error, and command message topics through MQTT. If the `--iot` flag
is included, opendyson will connect to the cloud-based IoT service instead of attempting to connect directly to the device.

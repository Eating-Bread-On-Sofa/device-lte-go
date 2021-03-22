// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/edgexfoundry/device-lte-go"
	"github.com/edgexfoundry/device-lte-go/driver"
	"github.com/edgexfoundry/device-sdk-go/pkg/startup"
)

const (
	serviceName string = "edgex-device-lte"
)

func main() {
	d := driver.NewProtocolDriver()
	startup.Bootstrap(serviceName, device_lte_go.Version, d)
}

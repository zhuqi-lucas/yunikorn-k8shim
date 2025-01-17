/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package main

import (
	"os"

	"k8s.io/kubernetes/cmd/kube-scheduler/app"

	"github.com/apache/yunikorn-k8shim/pkg/conf"
	"github.com/apache/yunikorn-k8shim/pkg/schedulerplugin"
	pluginconf "github.com/apache/yunikorn-k8shim/pkg/schedulerplugin/conf"
)

var (
	version string
	date    string
)

func main() {
	// override the default config handling when in plugin mode
	conf.SetSchedulerConfFactory(pluginconf.NewSchedulerConf)

	conf.BuildVersion = version
	conf.BuildDate = date
	conf.IsPluginVersion = true

	command := app.NewSchedulerCommand(
		app.WithPlugin(schedulerplugin.SchedulerPluginName, schedulerplugin.NewSchedulerPlugin))

	pluginconf.InitCliFlagSet(command)

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

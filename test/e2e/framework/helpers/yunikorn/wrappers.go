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

package yunikorn

import (
	"time"

	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"

	"github.com/apache/yunikorn-core/pkg/common/configs"

	"github.com/apache/yunikorn-k8shim/test/e2e/framework/configmanager"
	"github.com/apache/yunikorn-k8shim/test/e2e/framework/helpers/common"
	"github.com/apache/yunikorn-k8shim/test/e2e/framework/helpers/k8s"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var k = k8s.KubeCtl{}

func UpdateConfigMapWrapper(oldConfigMap *v1.ConfigMap, schedPolicy string, annotation string) {
	Ω(k.SetClient()).To(BeNil())
	By("Port-forward the scheduler pod")
	fwdErr := k.PortForwardYkSchedulerPod()
	Ω(fwdErr).NotTo(HaveOccurred())

	By("Enabling new scheduling config")

	// Save old configMap
	Ω(k.SetClient()).To(BeNil())
	var c, err = k.GetConfigMaps(configmanager.YuniKornTestConfig.YkNamespace,
		configmanager.DefaultYuniKornConfigMap)
	Ω(err).NotTo(HaveOccurred())
	Ω(c).NotTo(BeNil())
	c.DeepCopyInto(oldConfigMap)
	Ω(c).Should(BeEquivalentTo(oldConfigMap))

	// Create configMap with timestamp + schedPolicy
	sc := common.CreateBasicConfigMap()
	if schedPolicy != "" {
		err = common.SetSchedulingPolicy(sc, "default", "root", schedPolicy)
		Ω(err).NotTo(HaveOccurred())
	}
	ts, tsErr := common.SetQueueTimestamp(sc, "default", "root")
	Ω(tsErr).NotTo(HaveOccurred())
	configStr, yamlErr := common.ToYAML(sc)
	Ω(yamlErr).NotTo(HaveOccurred())
	c.Data[configmanager.DefaultPolicyGroup] = configStr
	var d, err3 = k.UpdateConfigMap(c, configmanager.YuniKornTestConfig.YkNamespace)
	Ω(err3).NotTo(HaveOccurred())
	Ω(d).NotTo(BeNil())

	// Updating scheduler pod annotation to trigger force refresh of configmaps
	Ω(k.UpdateYunikornSchedulerPodAnnotation(annotation)).NotTo(HaveOccurred())
	err = WaitForQueueTS("root", ts, 2*time.Minute)
	Ω(err).NotTo(HaveOccurred())
}

func RestoreConfigMapWrapper(oldConfigMap *v1.ConfigMap, annotation string) {
	Ω(k.SetClient()).To(BeNil())
	By("Restoring the old config maps")
	var c, err = k.GetConfigMaps(configmanager.YuniKornTestConfig.YkNamespace,
		configmanager.DefaultYuniKornConfigMap)
	Ω(err).NotTo(HaveOccurred())
	Ω(c).NotTo(BeNil())

	oldSC := new(configs.SchedulerConfig)
	err = yaml.Unmarshal([]byte(oldConfigMap.Data[configmanager.DefaultPolicyGroup]), oldSC)
	Ω(err).NotTo(HaveOccurred())
	ts, tsErr := common.SetQueueTimestamp(oldSC, "default", "root")
	Ω(tsErr).NotTo(HaveOccurred())
	c.Data[configmanager.DefaultPolicyGroup], err = common.ToYAML(oldSC)
	Ω(err).NotTo(HaveOccurred())

	var e, err3 = k.UpdateConfigMap(c, configmanager.YuniKornTestConfig.YkNamespace)
	Ω(err3).NotTo(HaveOccurred())
	Ω(e).NotTo(BeNil())
	// Updating scheduler pod annotation to trigger force refresh of configmaps
	Ω(k.RemoveYunikornSchedulerPodAnnotation(annotation)).NotTo(HaveOccurred())
	err = WaitForQueueTS("root", ts, 2*time.Minute)
	Ω(err).NotTo(HaveOccurred())
}

var Describe = ginkgo.Describe
var It = ginkgo.It
var By = ginkgo.By
var BeforeSuite = ginkgo.BeforeSuite
var AfterSuite = ginkgo.AfterSuite
var BeforeEach = ginkgo.BeforeEach
var AfterEach = ginkgo.AfterEach

var Equal = gomega.Equal
var Ω = gomega.Expect
var BeNil = gomega.BeNil
var HaveOccurred = gomega.HaveOccurred
var BeEquivalentTo = gomega.BeEquivalentTo

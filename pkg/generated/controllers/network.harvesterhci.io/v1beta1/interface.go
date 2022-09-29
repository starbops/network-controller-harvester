/*
Copyright 2019 Harvester Network Controller Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/harvester/harvester-network-controller/pkg/apis/network.harvesterhci.io/v1beta1"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v1beta1.AddToScheme)
}

type Interface interface {
	ClusterNetwork() ClusterNetworkController
	NodeNetwork() NodeNetworkController
	VlanConfig() VlanConfigController
	VlanStatus() VlanStatusController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (c *version) ClusterNetwork() ClusterNetworkController {
	return NewClusterNetworkController(schema.GroupVersionKind{Group: "network.harvesterhci.io", Version: "v1beta1", Kind: "ClusterNetwork"}, "clusternetworks", false, c.controllerFactory)
}
func (c *version) NodeNetwork() NodeNetworkController {
	return NewNodeNetworkController(schema.GroupVersionKind{Group: "network.harvesterhci.io", Version: "v1beta1", Kind: "NodeNetwork"}, "nodenetworks", false, c.controllerFactory)
}
func (c *version) VlanConfig() VlanConfigController {
	return NewVlanConfigController(schema.GroupVersionKind{Group: "network.harvesterhci.io", Version: "v1beta1", Kind: "VlanConfig"}, "vlanconfigs", false, c.controllerFactory)
}
func (c *version) VlanStatus() VlanStatusController {
	return NewVlanStatusController(schema.GroupVersionKind{Group: "network.harvesterhci.io", Version: "v1beta1", Kind: "VlanStatus"}, "vlanstatuses", false, c.controllerFactory)
}

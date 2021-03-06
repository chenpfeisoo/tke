/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
	platformv1 "tkestack.io/tke/api/platform/v1"
)

// FakeClusterAddons implements ClusterAddonInterface
type FakeClusterAddons struct {
	Fake *FakePlatformV1
}

var clusteraddonsResource = schema.GroupVersionResource{Group: "platform.tkestack.io", Version: "v1", Resource: "clusteraddons"}

var clusteraddonsKind = schema.GroupVersionKind{Group: "platform.tkestack.io", Version: "v1", Kind: "ClusterAddon"}

// Get takes name of the clusterAddon, and returns the corresponding clusterAddon object, and an error if there is any.
func (c *FakeClusterAddons) Get(ctx context.Context, name string, options v1.GetOptions) (result *platformv1.ClusterAddon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusteraddonsResource, name), &platformv1.ClusterAddon{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.ClusterAddon), err
}

// List takes label and field selectors, and returns the list of ClusterAddons that match those selectors.
func (c *FakeClusterAddons) List(ctx context.Context, opts v1.ListOptions) (result *platformv1.ClusterAddonList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusteraddonsResource, clusteraddonsKind, opts), &platformv1.ClusterAddonList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &platformv1.ClusterAddonList{ListMeta: obj.(*platformv1.ClusterAddonList).ListMeta}
	for _, item := range obj.(*platformv1.ClusterAddonList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

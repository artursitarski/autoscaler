// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package metatypes

import (
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func DeepCopy_metatypes_OwnerReference(in OwnerReference, out *OwnerReference, c *conversion.Cloner) error {
	out.APIVersion = in.APIVersion
	out.Kind = in.Kind
	out.UID = in.UID
	out.Name = in.Name
	if in.Controller != nil {
		value := *in.Controller
		out.Controller = &value
	}
	return nil
}
/*
Package kit for Kubernetes operators

Copyright 2016 The Rook Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Some of the code below came from https://github.com/coreos/etcd-operator
which also has the apache 2.0 license.
*/
package main

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SampleList is a list of samples from the TPR.
type SampleList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	Metadata metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of third party objects
	Items []Sample `json:"items"`
}

// There is known issue with TPR in client-go:
//   https://github.com/kubernetes/client-go/issues/8
// Workarounds:
// - We include `Metadata` field in object explicitly.
// - we have the code below to work around a known problem with third-party resources and ugorji.

type SampleListCopy SampleList
type SampleCopy Sample

func (p *Sample) UnmarshalJSON(data []byte) error {
	tmp := SampleCopy{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	tmp2 := Sample(tmp)
	*p = tmp2
	return nil
}

func (pl *SampleList) UnmarshalJSON(data []byte) error {
	tmp := SampleListCopy{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	tmp2 := SampleList(tmp)
	*pl = tmp2
	return nil
}

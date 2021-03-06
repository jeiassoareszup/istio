// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// DEPRECATED - These commands are deprecated and will be removed in future releases.

package cmd

import (
	"istio.io/istio/pilot/pkg/config/kube/crd/controller"
	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pkg/config/schema/collections"
)

var (
	// Create a model.ConfigStore (or sortedConfigStore)
	configStoreFactory = newConfigStore
)

func newConfigStore() (model.ConfigStore, error) {
	return controller.NewClient(kubeconfig, configContext, collections.Pilot,
		"", &model.DisabledLedger{}, "")
}

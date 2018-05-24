/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package apis

import (
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	"github.com/nervanasystems/carbon/applications/test-exp-service/pkg/apis/aggregator"
	aggregatorv1 "github.com/nervanasystems/carbon/applications/test-exp-service/pkg/apis/aggregator/v1"
)

// GetAllApiBuilders returns all known APIGroupBuilders
// so they can be registered with the apiserver
func GetAllApiBuilders() []*builders.APIGroupBuilder {
	return []*builders.APIGroupBuilder{
		GetAggregatorAPIBuilder(),
	}
}

var aggregatorApiGroup = builders.NewApiGroupBuilder(
	"aggregator.aipg.intel.com",
	"github.com/nervanasystems/carbon/applications/test-exp-service/pkg/apis/aggregator").
	WithUnVersionedApi(aggregator.ApiVersion).
	WithVersionedApis(
		aggregatorv1.ApiVersion,
	).
	WithRootScopedKinds()

func GetAggregatorAPIBuilder() *builders.APIGroupBuilder {
	return aggregatorApiGroup
}

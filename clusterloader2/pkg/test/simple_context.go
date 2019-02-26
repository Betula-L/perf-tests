/*
Copyright 2018 The Kubernetes Authors.

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

package test

import (
	"path/filepath"

	"k8s.io/perf-tests/clusterloader2/pkg/chaos"
	"k8s.io/perf-tests/clusterloader2/pkg/config"
	"k8s.io/perf-tests/clusterloader2/pkg/framework"
	"k8s.io/perf-tests/clusterloader2/pkg/measurement"
	"k8s.io/perf-tests/clusterloader2/pkg/state"
	"k8s.io/perf-tests/clusterloader2/pkg/tuningset"
)

type simpleContext struct {
	clusterLoaderConfig *config.ClusterLoaderConfig
	framework           *framework.Framework
	state               *state.State
	templateProvider    *config.TemplateProvider
	tuningSetFactory    tuningset.TuningSetFactory
	measurementManager  *measurement.MeasurementManager
	chaosMonkey         *chaos.Monkey
}

func createSimpleContext(c *config.ClusterLoaderConfig, f *framework.Framework, s *state.State) Context {
	templateProvider := config.NewTemplateProvider(filepath.Dir(c.TestConfigPath))
	return &simpleContext{
		clusterLoaderConfig: c,
		framework:           f,
		state:               s,
		templateProvider:    templateProvider,
		tuningSetFactory:    tuningset.NewTuningSetFactory(),
		measurementManager:  measurement.CreateMeasurementManager(f, &c.ClusterConfig, templateProvider),
		chaosMonkey:         chaos.NewMonkey(f.GetClientSets().GetClient(), c.ClusterConfig.Provider),
	}
}

// GetClusterConfig return cluster config.
func (sc *simpleContext) GetClusterLoaderConfig() *config.ClusterLoaderConfig {
	return sc.clusterLoaderConfig
}

// GetFramework returns framework.
func (sc *simpleContext) GetFramework() *framework.Framework {
	return sc.framework
}

// GetState returns current test state.
func (sc *simpleContext) GetState() *state.State {
	return sc.state
}

// GetTemplateProvider returns template provider.
func (sc *simpleContext) GetTemplateProvider() *config.TemplateProvider {
	return sc.templateProvider
}

// GetTickerFactory returns tuning set factory.
func (sc *simpleContext) GetTuningSetFactory() tuningset.TuningSetFactory {
	return sc.tuningSetFactory
}

// GetMeasurementManager returns measurement manager.
func (sc *simpleContext) GetMeasurementManager() *measurement.MeasurementManager {
	return sc.measurementManager
}

// GetChaosMonkey returns chaos monkey.
func (sc *simpleContext) GetChaosMonkey() *chaos.Monkey {
	return sc.chaosMonkey
}

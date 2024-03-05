/*
Copyright 2021 The cert-manager Authors.

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

package v1alpha1

import (
	time "time"

	"k8s.io/apimachinery/pkg/runtime"
	logsapi "k8s.io/component-base/logs/api/v1"
	"k8s.io/utils/ptr"

	"github.com/cert-manager/cert-manager/pkg/apis/config/cainjector/v1alpha1"
)

var (
	defaultLeaderElect                 = true
	defaultLeaderElectionNamespace     = "kube-system"
	defaultLeaderElectionLeaseDuration = 60 * time.Second
	defaultLeaderElectionRenewDeadline = 40 * time.Second
	defaultLeaderElectionRetryPeriod   = 15 * time.Second
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

func SetDefaults_CAInjectorConfiguration(obj *v1alpha1.CAInjectorConfiguration) {
	if obj.PprofAddress == "" {
		obj.PprofAddress = "localhost:6060"
	}

	logsapi.SetRecommendedLoggingConfiguration(&obj.Logging)
}

func SetDefaults_LeaderElectionConfig(obj *v1alpha1.LeaderElectionConfig) {
	if obj.Enabled == nil {
		obj.Enabled = &defaultLeaderElect
	}

	if obj.Namespace == "" {
		obj.Namespace = defaultLeaderElectionNamespace
	}

	// TODO: Does it make sense to have a duration of 0?
	if obj.LeaseDuration == time.Duration(0) {
		obj.LeaseDuration = defaultLeaderElectionLeaseDuration
	}

	if obj.RenewDeadline == time.Duration(0) {
		obj.RenewDeadline = defaultLeaderElectionRenewDeadline
	}

	if obj.RetryPeriod == time.Duration(0) {
		obj.RetryPeriod = defaultLeaderElectionRetryPeriod
	}
}

func SetDefaults_EnableDataSourceConfig(obj *v1alpha1.EnableDataSourceConfig) {
	if obj.Certificates == nil {
		obj.Certificates = ptr.To(true)
	}
}

func SetDefaults_EnableInjectableConfig(obj *v1alpha1.EnableInjectableConfig) {
	if obj.MutatingWebhookConfigurations == nil {
		obj.MutatingWebhookConfigurations = ptr.To(true)
	}
	if obj.ValidatingWebhookConfigurations == nil {
		obj.ValidatingWebhookConfigurations = ptr.To(true)
	}
	if obj.CustomResourceDefinitions == nil {
		obj.CustomResourceDefinitions = ptr.To(true)
	}
	if obj.APIServices == nil {
		obj.APIServices = ptr.To(true)
	}
}

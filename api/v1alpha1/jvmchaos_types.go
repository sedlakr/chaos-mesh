// Copyright 2020 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// JVMChaosSpec defines the desired state of JVMChaos
type JVMChaosSpec struct {
	// Mode defines the mode to run chaos action.
	// Supported mode: one / all / fixed / fixed-percent / random-max-percent
	Mode PodMode `json:"mode"`

	// Value is required when the mode is set to `FixedPodMode` / `FixedPercentPodMod` / `RandomMaxPercentPodMod`.
	// If `FixedPodMode`, provide an integer of pods to do chaos action.
	// If `FixedPercentPodMod`, provide a number from 0-100 to specify the max % of pods the server can do chaos action.
	// If `RandomMaxPercentPodMod`,  provide a number from 0-100 to specify the % of pods to do chaos action
	// +optional
	Value string `json:"value"`

	// Selector is used to select pods that are used to inject chaos action.
	Selector SelectorSpec `json:"selector"`

	// Duration represents the duration of the chaos action
	// +optional
	Duration *string `json:"duration,omitempty"`

	// Scheduler defines some schedule rules to control the running time of the chaos experiment about time.
	// +optional
	Scheduler *SchedulerSpec `json:"scheduler,omitempty"`

	// Action defines the specific jvm chaos action.
	// Supported action: delay, return, script, cfl, oom, ccf, tce, delay4servlet, tce4servlet
	// +kubebuilder:validation:Enum=delay;return;script;cfl;oom;ccf;tce;delay4servlet;tce4servlet
	Action JVMChaosAction `json:"action"`

	// JVMParameter represents the detail about jvm chaos action definition
	// +optional
	JVMParameter `json:",inline"`

	// Target defines the specific jvm chaos target.
	// Supported target: servlet;psql;jvm;jedis;http;dubbo;rocketmq;tars;mysql;druid
	// +kubebuilder:validation:Enum=servlet;psql;jvm;jedis;http;dubbo;rocketmq;tars;mysql;druid
	Target JVMChaosTarget `json:"target"`
}

// GetSelector is a getter for Selector (for implementing SelectSpec)
func (in *JVMChaosSpec) GetSelector() SelectorSpec {
	return in.Selector
}

// GetMode is a getter for Mode (for implementing SelectSpec)
func (in *JVMChaosSpec) GetMode() PodMode {
	return in.Mode
}

// GetValue is a getter for Value (for implementing SelectSpec)
func (in *JVMChaosSpec) GetValue() string {
	return in.Value
}

type JVMChaosTarget string

const (
	SERVLET JVMChaosTarget = "servlet"

	PSQL JVMChaosTarget = "psql"

	JVM JVMChaosTarget = "jvm"

	JEDIS JVMChaosTarget = "jedis"

	HTTP JVMChaosTarget = "http"

	DUBBO JVMChaosTarget = "dubbo"

	ROCKETMQ JVMChaosTarget = "rocketmq"

	MYSQL JVMChaosTarget = "mysql"

	DRUID JVMChaosTarget = "druid"

	TARS JVMChaosTarget = "tars"
)

// JVMChaosAction represents the chaos action about jvm
type JVMChaosAction string

const (
	// JVMDelayAction represents the JVM chaos action of invoke delay
	JVMDelayAction JVMChaosAction = "delay"

	// JVMReturnAction represents the JVM chaos action of return value
	JVMReturnAction JVMChaosAction = "return"

	// JVMReturnAction represents the JVM chaos action for complex failure scenarios.
	// Write Java or Groovy scripts, such as tampering with parameters, modifying return values,
	// throwing custom exceptions, and so on
	JVMScriptAction JVMChaosAction = "script"

	// JVMCpuFullloadAction represents the JVM chaos action of CPU is full
	JVMCpuFullloadAction JVMChaosAction = "cfl"

	// JVMOOMAction represents the JVM chaos action of OOM exception
	JVMOOMAction JVMChaosAction = "oom"

	// JVMCodeCacheFillingAction represents the JVM chaos action of code cache filling
	JVMCodeCacheFillingAction JVMChaosAction = "ccf"

	// JVMExceptionAction represents the JVM chaos action of throwing custom exceptions
	JVMExceptionAction JVMChaosAction = "tce"

	// JVMConnectionPoolFullAction represents the JVM chaos action of Connection Pool Full
	JVMConnectionPoolFullAction JVMChaosAction = "cpf"

	// JVMThrowDeclaredExceptionAction represents the JVM chaos action of throwing declared exception
	JVMThrowDeclaredExceptionAction JVMChaosAction = "tde"

	// JVMThreadPoolFullAction represents the JVM chaos action of thread pool full
	JVMThreadPoolFullAction JVMChaosAction = "tpf"
)

// JVMParameter represents the detail about jvm chaos action definition
type JVMParameter struct {

	// Flags represents the flags of action
	// +optional
	Flags map[string]string `json:"flags,omitempty"`

	// Matchers represents the matching rules for the target
	// +optional
	Matchers map[string]string `json:"matchers,omitempty"`
}

// JVMChaosStatus defines the observed state of JVMChaos
type JVMChaosStatus struct {
	ChaosStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +chaos-mesh:base

// JVMChaos is the Schema for the jvmchaos API
type JVMChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JVMChaosSpec   `json:"spec,omitempty"`
	Status JVMChaosStatus `json:"status,omitempty"`
}

func init() {
	SchemeBuilder.Register(&JVMChaos{}, &JVMChaosList{})
}
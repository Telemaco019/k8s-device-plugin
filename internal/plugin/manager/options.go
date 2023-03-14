/**
# Copyright (c) NVIDIA CORPORATION.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
**/

package manager

import (
	spec "github.com/NVIDIA/k8s-device-plugin/api/config/v1"
	"github.com/NVIDIA/k8s-device-plugin/internal/cdi"
	"gitlab.com/nvidia/cloud-native/go-nvlib/pkg/nvml"
)

type Option func(*manager)

// WithCDIHandler sets the CDI handler for the manager
func WithCDIHandler(handler cdi.Interface) Option {
	return func(m *manager) {
		m.cdiHandler = handler
	}
}

// WithNVML sets the NVML handler for the manager
func WithNVML(nvmllib nvml.Interface) Option {
	return func(m *manager) {
		m.nvmllib = nvmllib
	}
}

// WithConfig sets the config reference for the manager
func WithConfig(config *spec.Config) Option {
	return func(m *manager) {
		m.config = config
	}
}

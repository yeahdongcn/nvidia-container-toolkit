/*
# Copyright (c) 2021, NVIDIA CORPORATION.  All rights reserved.
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
*/

package filter

import "github.com/NVIDIA/nvidia-container-toolkit/pkg/discover"

type all struct {
	selectors []Selector
}

// All returns a selector that evaluates true if EACH of the specified selectors
// are selected.
func All(selectors ...Selector) Selector {
	s := all{
		selectors: selectors,
	}
	return &s
}

func (s all) Selected(device discover.Device) bool {
	for _, si := range s.selectors {
		if !si.Selected(device) {
			return false
		}
	}
	return true
}

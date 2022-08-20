// Copyright 2022 Ben Hale
// SPDX-License-Identifier: Apache-2.0

package ha

type Device struct {
	Identifiers     []string `json:"identifiers,omitempty"`
	Manufacturer    string   `json:"manufacturer,omitempty"`
	Model           string   `json:"model,omitempty"`
	Name            string   `json:"name,omitempty"`
	SuggestedArea   string   `json:"suggested_area,omitempty"`
	SoftwareVersion string   `json:"sw_version,omitempty"`
}

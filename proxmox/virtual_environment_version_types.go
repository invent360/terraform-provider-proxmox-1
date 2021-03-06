/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmox

// VirtualEnvironmentVersionResponseBody contains the body from a version response.
type VirtualEnvironmentVersionResponseBody struct {
	Data *VirtualEnvironmentVersionResponseData `json:"data,omitempty"`
}

// VirtualEnvironmentVersionResponseData contains the data from a version response.
type VirtualEnvironmentVersionResponseData struct {
	Keyboard     string `json:"keyboard"`
	Release      string `json:"release"`
	RepositoryID string `json:"repoid"`
	Version      string `json:"version"`
}

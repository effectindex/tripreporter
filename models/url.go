// SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package models

type URL struct {
	Unique
	Name string `json:"name"`
	URL  string `json:"url"`
}

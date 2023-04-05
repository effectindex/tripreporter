// SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

/* eslint-disable no-console */

export default function log(message, ...optionalParams) {
  if (process.env.NODE_ENV !== "production") {
    console.log(message, ...optionalParams)
  }
}

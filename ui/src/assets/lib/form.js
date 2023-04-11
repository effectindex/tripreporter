// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

export function getTextLength(text, max) {
  if (!text) {
    return ""
  }

  return `(${text.length.toLocaleString()} / ${max.toLocaleString()})`
}

// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

export function titleCase(str) {
  if (str === str.toUpperCase() || !isNaN(str.charAt(0))) {
    return str
  }

  str = str.toLowerCase().split(' ');
  for (let i = 0; i < str.length; i++) {
    str[i] = str[i].charAt(0).toUpperCase() + str[i].slice(1);
  }

  return str.join(' ');
}

export function getOrDefault(s1, s2) {
  if (!s1 || s1 === "") {
    return s2
  }

  return s1
}

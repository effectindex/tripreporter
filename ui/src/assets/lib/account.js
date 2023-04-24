// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import log from "@/assets/lib/logger";

export default class Account { // TODO: Rewrite in TS for #106
  constructor({ id, email, username }) {
    this.id = id;
    this.email = email;
    this.username = username;

    log("account.js", this)
    return this;
  }
}

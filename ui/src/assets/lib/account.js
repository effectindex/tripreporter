// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import log from "@/assets/lib/logger";

export default class Account { // TODO: Rewrite in TS for #106
  constructor({ id, email, username, display_name }) {
    this.id = id;
    this.email = email;
    this.username = username;
    this.display_name = display_name;
    this.default_name = display_name ? display_name : username;

    log("account.js", this)
    return this;
  }
}

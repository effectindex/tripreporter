// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import Timestamp from "@/assets/lib/timestamp";

export default class User { // TODO: Rewrite in TS for #106
  constructor({ obj, id, display_name, created }) {
    this.id = id;
    this.display_name = display_name;
    this.created = new Timestamp({ date: created, longFormat: true })
    obj && Object.assign(this, obj);
    return this;
  }
}

// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import Timestamp from "@/assets/lib/timestamp";
import Report from "@/assets/lib/report";
import log from "@/assets/lib/logger";

export default class User { // TODO: Rewrite in TS for #106
  constructor({ id, username, display_name, created, reports }) {
    this.id = id;
    this.username = username;
    this.display_name = display_name;
    this.reports = [];
    this.created = new Timestamp({ date: created, longFormat: true });

    for (const n in reports) {
      let report = new Report(reports[n]);
      this.reports.push(report)
    }

    log("user.js", this)
    return this;
  }
}

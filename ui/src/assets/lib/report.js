// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import ReportEvent from "@/assets/lib/report-event";
import log from "@/assets/lib/logger";

export default class Report { // TODO: Rewrite in TS for #106
  constructor({ id, user, account_id, creation_time, modified_time, report_date, title, setting, report_subject, report_events }) {
    this.id = id;
    this.user = user;
    this.account_id = account_id;
    this.creation_time = creation_time;
    this.modified_time = modified_time;
    this.report_date = report_date;
    this.title = title;
    this.setting = setting;
    this.medications = [];
    this.report_subject = report_subject;
    this.report_events = [];

    if (report_subject && report_subject["medications"]) {
      for (const n in report_subject["medications"]) {
        let event = new ReportEvent({ obj: report_subject["medications"][n] })
        this.medications.push(event)
      }
    }

    for (const n in report_events) {
      let event = new ReportEvent({ obj: report_events[n] })
      this.report_events.push(event)
    }

    log("report.js", this)
    return this;
  }
}

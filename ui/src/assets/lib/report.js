// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import ReportEvent from "@/assets/lib/report-event";
import log from "@/assets/lib/logger";
import DrugData from "@/assets/lib/drug-data";

export default class Report { // TODO: Rewrite in TS for #106
  constructor({ id, user, account_id, creation_time, modified_time, report_date, title, setting, report_subject, report_events, drugs }) {
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
    this.drugs = [];

    if (report_subject && report_subject["medications"]) {
      for (const n in report_subject["medications"]) {
        let event = new ReportEvent({ drug: report_subject["medications"][n] })
        event["type"] = 2
        this.medications.push(event)
      }
    }

    for (const n in report_events) {
      let event = new ReportEvent({ obj: report_events[n] })
      this.report_events.push(event)
    }

    for (const n in drugs) {
      let drug = new DrugData({ obj: drugs[n] })
      this.drugs.push(drug)
    }

    log("report.js", this)
    return this;
  }
}

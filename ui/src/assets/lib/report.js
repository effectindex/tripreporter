import ReportEvent from "@/assets/lib/report-event";
import log from "@/assets/lib/logger";

export default class Report { // TODO: Rewrite in TS for #106
  constructor({ id, account_id, creation_time, modified_time, report_date, title, setting, report_events }) {
    this.id = id;
    this.account_id = account_id;
    this.creation_time = creation_time;
    this.modified_time = modified_time;
    this.report_date = report_date;
    this.title = title;
    this.setting = setting;
    this.report_events = [];
    for (const n in report_events) {
      let event = new ReportEvent({ obj: report_events[n] })
      this.report_events.push(event)
    }

    log("report.js", this)
    return this;
  }
}

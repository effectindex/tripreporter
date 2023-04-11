// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import DrugData from "@/assets/lib/drug-data";

export default class ReportEvent { // TODO: Rewrite in TS for #106
  constructor({ obj, report_id, index, timestamp, section, content, drug }) {
    this.report_id = report_id;
    this.index = index;
    this.timestamp = timestamp;
    this.section = section;
    this.content = content;
    this.drug = new DrugData({ obj: drug });
    obj && Object.assign(this, obj);
    return this;
  }
}

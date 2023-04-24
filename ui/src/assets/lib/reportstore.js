// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import { ref } from "vue";
import { defineStore } from "pinia";
import log from "@/assets/lib/logger";
import Report from "@/assets/lib/report";
import Timestamp from "@/assets/lib/timestamp";
import User from "@/assets/lib/user";
import ReportSubject from "@/assets/lib/report-subject";

export const useReportStore = defineStore('report', {
  state: () => {
    return {
      m: new Map(),
    }
  },
  actions: {
    updateData(id, status, data) {
      let r = this.m.get(id)
      if (r === undefined) {
        r = {
          reportJson: new Report({}),
          reportUser: ref(null),
          reportDate: ref(null),
          reportSubject: ref(null),
          hideMessage: false,
          apiSuccess: false,
        }
      }

      r.apiSuccess = status === 200;

      if (r.apiSuccess) {
        log("Loading report data", typeof r.reportJson, typeof data)
        r.reportJson = new Report(data);
        r.reportUser = new User(r.reportJson.user);
        r.reportDate = new Timestamp({ date: r.reportJson.report_date, longFormat: true });
        r.reportSubject = new ReportSubject({ obj: r.reportJson.report_subject })
        r.hideMessage = true;
        this.m.set(id, r);
        log("Loaded report store", typeof r.reportJson)
      }
    },
    isLoaded(id) {
      const r = this.m.get(id);
      return r.apiSuccess && r.reportJson && r.reportJson !== {} && r.reportUser !== null // subject can be null?
    }
  },
})

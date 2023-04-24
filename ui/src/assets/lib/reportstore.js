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
      reportJson: new Report({}),
      reportUser: ref(null),
      reportDate: ref(null),
      reportSubject: ref(null),
      hideMessage: false,
      apiSuccess: false
    }
  },
  actions: {
    updateData(status, data) {
      this.apiSuccess = status === 200;

      if (this.apiSuccess) {
        log("Loading report data", typeof this.reportJson, typeof data)
        this.reportJson = new Report(data);
        this.reportUser = new User({
          id: this.reportJson.user.id,
          display_name: this.reportJson.user.display_name,
          created: this.reportJson.user.created
        });
        this.reportDate = new Timestamp({ date: this.reportJson.report_date, longFormat: true });
        this.reportSubject = new ReportSubject({ obj: this.reportJson.report_subject })
        this.hideMessage = true;
        log("Loaded report store", typeof this.reportJson)
      }
    },
    isLoaded() {
      return this.apiSuccess && this.reportJson && this.reportJson !== {} && this.reportUser !== null // subject can be null?
    }
  },
})

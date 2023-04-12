// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import { ref } from "vue";
import { defineStore } from "pinia";
import log from "@/assets/lib/logger";
import Report from "@/assets/lib/report";
import Timestamp from "@/assets/lib/timestamp";

export const useReportsStore = defineStore('reports', {
  state: () => {
    return {
      reportJson: new Report({}),
      reportDate: ref(null),
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
        this.reportDate = new Timestamp({date: this.reportJson.report_date, longFormat: true});
        this.hideMessage = true;
        log("Loaded reports store", typeof this.reportJson)
      }
    },
    isLoaded() {
      return this.apiSuccess && this.reportJson && this.reportJson !== {}
    }
  },
})

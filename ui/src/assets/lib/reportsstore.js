// SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import { defineStore } from "pinia";
import log from "@/assets/lib/logger";
import Report from "@/assets/lib/report";

export const useReportsStore = defineStore('reports', {
  state: () => {
    return {
      reportJson: new Report({}),
      hideMessage: false,
      apiSuccess: true
    }
  },
  actions: {
    updateData(status, data) {
      this.apiSuccess = status === 200;

      if (this.apiSuccess) {
        log("Loading report data", typeof this.reportJson, typeof data)
        this.reportJson = new Report(data);
        this.hideMessage = true;
        log("Loaded reports store", typeof this.reportJson)
      }
    },
    isLoaded() {
      return this.apiSuccess && this.reportJson && this.reportJson !== {}
    }
  },
})

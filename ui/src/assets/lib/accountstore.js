// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import { defineStore } from "pinia";
import log from "@/assets/lib/logger";
import { ref } from "vue";
import Account from "@/assets/lib/account";

export const useAccountStore = defineStore('account', {
  state: () => {
    return {
      accountJson: ref(null),
      hideMessage: false,
      apiSuccess: false
    }
  },
  actions: {
    updateData(status, data) {
      this.apiSuccess = status === 200;

      if (this.apiSuccess) {
        log("Loading account data", typeof this.accountJson, typeof data)
        this.accountJson = new Account(data);
        this.hideMessage = true;
        log("Loaded account store", typeof this.accountJson)
      }
    },
    isLoaded() {
      return this.apiSuccess && this.accountJson !== null
    }
  },
})

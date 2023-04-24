// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import { ref } from "vue";
import { defineStore } from 'pinia'
import { validateSession } from "@/assets/lib/session";
import log from "@/assets/lib/logger";

export const useSessionStore = defineStore('session', {
  state: () => {
    return {
      updatedPreviously: false,
      activeSession: false,
      lastUsername: "",
      createAccountForm: ref({}), // Used when creating an account
      createUserForm: ref({}), // Used when creating an account user
    }
  },
  actions: {
    updateSession(axios) {
      validateSession(axios).then((res) => {
        this.updatedPreviously = true;
        this.activeSession = res;
        log(`Loaded store session: ${this.activeSession ? "active session" : "no session"}`)
      })
    },
    invalidateSession() {
      this.updatedPreviously = false;
      this.activeSession = false;
    }
  },
})

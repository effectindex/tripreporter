// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import { defineStore } from "pinia";
import log from "@/assets/lib/logger";
import { ref } from "vue";
import Timestamp from "@/assets/lib/timestamp";
import User from "@/assets/lib/user";

export const useUserStore = defineStore('user', {
  state: () => {
    return {
      user: new User({}),
      createdDate: ref(null),
      hideMessage: false,
      apiSuccess: false
    }
  },
  actions: {
    updateData(status, data) {
      this.apiSuccess = status === 200;

      if (this.apiSuccess) {
        log("Loading user data", typeof this.user, typeof data)
        this.user = new User(data);
        this.createdDate = new Timestamp({ date: data.created, showTime: true, longFormat: true });
        this.hideMessage = true;
        log("Loaded user store", typeof this.user, this.user)
      }
    },
    isLoaded() {
      return this.apiSuccess && this.user !== null
    }
  },
})

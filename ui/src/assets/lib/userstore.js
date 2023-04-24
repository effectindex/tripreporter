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
      m: new Map(),
    }
  },
  actions: {
    updateData(id, status, data) {
      log("userstore", "here1")
      let u = this.m.get(id)
      if (u === undefined) {
       u = {
         user: new User({}),
         createdDate: ref(null),
         hideMessage: false,
         apiSuccess: false,
       }
      }
      log("userstore", "here2")

      u.apiSuccess = status === 200;
      log("userstore", "here3")

      if (u.apiSuccess) {
        log("Loading user data", typeof u.user, typeof data)
        u.user = new User(data);
        u.createdDate = new Timestamp({ date: data.created, showTime: true, longFormat: true });
        u.hideMessage = true;
        this.m.set(id, u);
        log("Loaded user store", typeof u.user, u.user)
      }
    },
    hideMessage(id) {
      return this.m !== undefined && this.m.get(id) && this.m.get(id).hideMessage === true;
    },
    isLoaded(id) {
      const u = this.m.get(id)
      return u.apiSuccess && u.user !== null
    }
  },
})

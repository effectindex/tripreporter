// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import { defineStore } from 'pinia'
import { ref } from "vue";

export const useCreateStore = defineStore('create', {
  state: () => {
    return {
      hasSubject: ref(false),
      useImperial: ref(false),
      subjectInfo: ref({
        age: Number,
        gender: ref(),
        heightCm: Number,
        weightKg: Number,
        heightFt: Number,
        heightIn: Number,
        weightLbs: Number
      })
    }
  }
})

// SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import { defineStore } from 'pinia'

export const useCreateStore = defineStore('create', {
  state: () => {
    return {
      submitClass: 'formkit-outer',
      page: 0
    }
  }
})

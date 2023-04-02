import { defineStore } from 'pinia'

export const useCreateStore = defineStore('create', {
  state: () => {
    return {
      submitClass: 'formkit-outer',
      page: 0
    }
  }
})

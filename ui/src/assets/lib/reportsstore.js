import { defineStore } from "pinia";

export const useReportsStore = defineStore('reports', {
    state: () => {
        return {
            id: "",
            data: undefined,
            apiSuccess: true,
            apiMessage: ""
        }
    }
})

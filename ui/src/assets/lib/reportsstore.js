import { defineStore } from "pinia";
import log from "@/assets/lib/logger";

export const useReportsStore = defineStore('reports', {
    state: () => {
        return {
            data: undefined,
            hideMessage: false,
            apiSuccess: true
        }
    },
    actions: {
        updateData(status, data) {
            this.apiSuccess = status === 200;

            if (this.apiSuccess) {
                this.data = data;
                this.hideMessage = true;
                log("Loaded reports store", this.data)
            }
        },
    },
})

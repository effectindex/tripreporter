import { defineStore } from 'pinia'
import { validateSession } from "@/assets/lib/session";
import log from "@/assets/lib/logger";

export const useSessionStore = defineStore('session', {
    state: () => {
        return {
            updatedPreviously: false,
            activeSession: false,
            lastUsername: ""
        }
    },
    actions: {
        updateSession(axios) {
            validateSession(axios).then((res) => {
                this.activeSession = res;
                this.updatedPreviously = true;
                log(`Loaded store session: ${this.activeSession ? "active session" : "no session"}`)
            })
        },
    },
})

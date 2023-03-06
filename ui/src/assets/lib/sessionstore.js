import { defineStore } from 'pinia'
import { validateSession } from "@/assets/lib/session";
import log from "@/assets/lib/logger";

export const useSessionStore = defineStore('session', {
    state: () => {
        return { activeSession: false }
    },
    actions: {
        updateSession(axios) {
            validateSession(axios).then((res) => {
                this.activeSession = res;
                log(`store session is ${this.activeSession}`)
                // emit("sessionUpdated", "hello from setup!") // TODO
            })
        },
    },
})

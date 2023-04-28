<template>
  <slot v-if="store.activeSession" />
  <div v-else-if="store.updatedPreviously" v-show="false">
    {{ router.push('/login') }}
  </div>
</template>

<script>
export default {
  name: "AuthWrapper"
}
</script>

<script setup>
import { inject } from "vue";
import { useSessionStore } from "@/assets/lib/sessionstore";

const router = inject('router')
const store = useSessionStore();
</script>

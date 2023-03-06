<template>
  <div class="create">
    <div v-if="!store.activeSession" class="no-session">
      <not-found/>
    </div>
    <div v-else>
      <h1>This page is temporarily disabled, come back later!</h1>
    </div>
  </div>
</template>

<script>
import LayoutDefault from "@/layouts/LayoutDefault.vue";

export default {
  name: "CreateView",
  created() {
    this.$emit('update:layout', LayoutDefault);
  }
}
</script>

<script setup>
import {useSessionStore} from '@/assets/lib/sessionstore'
import NotFound from "@/views/NotFound.vue";

const store = useSessionStore();
</script>

<style scoped>

</style>

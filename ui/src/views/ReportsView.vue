<template>
  <div class="reports">
    <h1>This is a reports page</h1>
  </div>
</template>

<script>
import LayoutDefault from "@/layouts/LayoutDefault.vue";

export default {
  name: "ReportsView",
  created() {
    this.$emit('update:layout', LayoutDefault);
  }
}
</script>

<style scoped>

</style>

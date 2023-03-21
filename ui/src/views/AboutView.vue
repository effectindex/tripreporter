<template>
  <div class="about">
    <h1 class="--tr-not-bold">This is an about page</h1>
  </div>
</template>

<script>
import LayoutDefault from "@/layouts/LayoutDefault.vue";

export default {
  name: "AboutView",
  created() {
    this.$emit('update:layout', LayoutDefault);
  }
}
</script>

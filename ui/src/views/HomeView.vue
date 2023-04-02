<template>
  <div class="Home">
    <h1 class="--tr-header-h1">Home</h1>
    <p>
      Welcome to the <span class="--tr-muted-text">subjective.report</span> home page.
    </p>
  </div>
</template>

<script>
import LayoutDefault from "@/layouts/LayoutDefault.vue";

export default {
  name: 'HomeView',
  created() {
    this.$emit('update:layout', LayoutDefault);
  }
}
</script>

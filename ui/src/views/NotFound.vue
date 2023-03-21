<template>
  <div>
    <h1 class="--tr-not-bold">Not Found</h1>
    <p>
      Oops, we couldn't find that page. Try going
      <router-link to="/">home</router-link>
    </p>
  </div>
</template>

<script>
import LayoutDefault from "@/layouts/LayoutDefault.vue";

export default {
  name: "NotFound",
  created() {
    this.$emit('update:layout', LayoutDefault);
  }
}
</script>

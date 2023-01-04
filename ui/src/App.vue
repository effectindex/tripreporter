<template>
  <layout-default>
    <router-view/>
  </layout-default>
</template>

<style>
html,
body {
  height: 100vh;
  width: 100vw;
  margin: 0;
}

#app {
  height: inherit;
  width: inherit;
}
</style>

<script setup>
import LayoutDefault from "@/layouts/LayoutDefault.vue";
</script>

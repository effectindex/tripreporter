<!--
SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <component :is="layout">
    <router-view v-model:layout="layout"/>
  </component>
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

<script>
export default {
  name: 'App',
  data() {
    return {
      layout: 'div',
    };
  },
};
</script>

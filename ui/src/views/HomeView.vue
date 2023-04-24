<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <div class="Home">
    <h1 class="--tr-header-h1">Home</h1>
    <p>
      Welcome to the <router-link to="/" class="--tr-link">subjective.report</router-link> home page.
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

<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <div class="about">
    <h1 class="--tr-header-h1">This is an account page</h1>
  </div>
</template>

<script>
import LayoutDefault from "@/layouts/LayoutDefault.vue";

export default {
  name: "AccountView",
  created() {
    this.$emit('update:layout', LayoutDefault);
  }
}
</script>

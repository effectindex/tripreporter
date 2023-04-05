<!--
SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <div>
    <h1 class="--tr-header-h1">Not Found</h1>
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

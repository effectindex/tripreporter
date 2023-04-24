<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <div class="Home">
    <h1 class="--tr-header-h1">Home</h1>
    <p>
      Welcome to the <subjective-report-link /> home page.
    </p>
  </div>
</template>

<script>
import LayoutDefault from "@/layouts/LayoutDefault.vue";
import SubjectiveReportLink from "@/components/SubjectiveReportLink.vue";

export default {
  name: 'HomeView',
  components: { SubjectiveReportLink },
  created() {
    this.$emit('update:layout', LayoutDefault);
  }
}
</script>

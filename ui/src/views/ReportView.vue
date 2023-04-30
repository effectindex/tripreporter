<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <AuthWrapper>
    <div class="reports">
      <Suspense>
        <report-box :id="$route.params.id"/>
      </Suspense>
    </div>
  </AuthWrapper>
</template>

<script>
import LayoutDefault from "@/layouts/LayoutDefault.vue";
import ReportBox from "@/components/ReportBox.vue";

export default {
  name: "ReportView",
  components: {
    ReportBox
  },
  created() {
    this.$emit('update:layout', LayoutDefault);
  }
}
</script>

<script setup>
import AuthWrapper from "@/components/AuthWrapper.vue";
</script>

<style scoped>

</style>

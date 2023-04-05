<!--
SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <div class="reports">
    <Suspense>
      <report-box :id="$route.query.id"/>
    </Suspense>
    <div v-if="!store.hideMessage" class="DefaultView__message" id="DefaultView__message">
      <div class="DefaultView__message_text" id="DefaultView__message_text"></div>
    </div>
  </div>
</template>

<script>
import LayoutDefault from "@/layouts/LayoutDefault.vue";
import ReportBox from "@/components/ReportBox.vue";

export default {
  name: "ReportsView",
  components: {
    ReportBox
  },
  created() {
    this.$emit('update:layout', LayoutDefault);
  }
}
</script>

<script setup>
import { useReportsStore } from "@/assets/lib/reportsstore";

const store = useReportsStore();
</script>

<style scoped>

</style>

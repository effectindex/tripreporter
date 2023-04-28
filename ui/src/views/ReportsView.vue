<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <AuthWrapper>
    <div class="reports">
      <Suspense>
        <report-box :id="$route.query.id"/>
      </Suspense>
      <div v-if="!store.hideMessage($route.query.id)" class="DefaultView__message" id="DefaultView__message">
        <div class="DefaultView__message_text" id="DefaultView__message_text"></div>
      </div>
    </div>
  </AuthWrapper>
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
import { useReportStore } from "@/assets/lib/reportstore";
import AuthWrapper from "@/components/AuthWrapper.vue";

const store = useReportStore();
</script>

<style scoped>

</style>

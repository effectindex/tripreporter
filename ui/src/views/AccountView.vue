<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <AuthWrapper>
    <div class="about">
      <Suspense>
        <account-box />
      </Suspense>
    </div>
  </AuthWrapper>
</template>

<script>
import LayoutDefault from "@/layouts/LayoutDefault.vue";
import AccountBox from "@/components/AccountBox.vue";

export default {
  name: "AccountView",
  components: {
    AccountBox
  },
  created() {
    this.$emit('update:layout', LayoutDefault);
  }
}
</script>

<script setup>
import AuthWrapper from "@/components/AuthWrapper.vue";
</script>

<style>
@import url(@/assets/css/forms.css);
</style>

<style scoped>
@import url(@/assets/css/message_util.css);
</style>

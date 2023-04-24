<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <div class="about">
    <Suspense>
      <account-box :id="$route.query.id"/>
    </Suspense>
    <div v-if="!store.hideMessage" class="DefaultView__message" id="DefaultView__message">
      <div class="DefaultView__message_text" id="DefaultView__message_text"></div>
    </div>
  </div>
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
import { useAccountStore } from "@/assets/lib/accountstore";

const store = useAccountStore();
</script>

<style>
@import url(@/assets/css/forms.css);
</style>

<style scoped>
@import url(@/assets/css/message_util.css);
</style>

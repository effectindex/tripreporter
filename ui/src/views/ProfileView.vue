<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <AuthWrapper>
    <div class="reports">
      <Suspense>
        <user-box :id="$route.params.id"/>
      </Suspense>
      <div v-if="!store.hideMessage($route.params.id)" class="DefaultView__message" id="DefaultView__message">
        <div class="DefaultView__message_text" id="DefaultView__message_text"></div>
      </div>
    </div>
  </AuthWrapper>
</template>

<script>
import LayoutDefault from "@/layouts/LayoutDefault.vue";
import UserBox from "@/components/UserBox.vue";

export default {
  name: "ProfileView",
  components: {
    UserBox
  },
  created() {
    this.$emit('update:layout', LayoutDefault);
  }
}
</script>

<script setup>
import { useUserStore } from "@/assets/lib/userstore";
import AuthWrapper from "@/components/AuthWrapper.vue";

const store = useUserStore();
</script>

<style scoped>

</style>

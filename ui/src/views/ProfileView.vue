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
import AuthWrapper from "@/components/AuthWrapper.vue";
</script>

<style scoped>

</style>

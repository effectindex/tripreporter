<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <h1 class="--tr-header-h1">You're logged in! 🎉</h1>
  <p>
    You can go to <router-link to="/create" class="--tr-link">create a report</router-link>, or <router-link to="/account" class="--tr-link">manage your account</router-link>.
  </p>
  <p>
    Click <router-link to="/logout" class="--tr-link-red">here</router-link> to logout.
  </p>
</template>

<script>
export default {
  name: "AlreadyLoggedIn"
}
</script>

<style scoped>

</style>

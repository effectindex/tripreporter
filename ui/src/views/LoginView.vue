<!--
SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <div class="login">
    <div v-if="!store.activeSession" class="no-session">
      <h1 class="--tr-header-h1">Login to your <span class="--tr-muted-text">subjective.report</span> account ✨</h1>

      <div class="DefaultView__message" id="DefaultView__message">
        <div class="DefaultView__message_text" id="DefaultView__message_text"></div>
      </div>
      <div class="DefaultView__form">
        <FormKit type="form" @submit="submitForm" #default="{ state: { errors } }" :actions="false">
          <FormKit
              type="text"
              name="username"
              id="username"
              label="Username"
              validation="required"
              placeholder="lyv76"
              :value="store.lastUsername"
          />

          <FormKit
              type="password"
              name="password"
              id="password"
              label="Password"
              validation="required"
              placeholder="----------"
          />

          <FormKit type="submit" label="Login" data-next="true" :disabled="errors && submitting"/>
        </FormKit>
      </div>
    </div>
    <div v-else>
      <h1 class="--tr-header-h1">You're logged in! 🎉</h1>
    </div>
  </div>
</template>

<script>
import LayoutDefault from "@/layouts/LayoutDefault.vue";

export default {
  name: "LoginView",
  created() {
    this.$emit('update:layout', LayoutDefault);
  }
}
</script>

<script setup>
import { inject, ref } from 'vue'
import { handleMessageError, setMessage } from '@/assets/lib/message_util';
import { useSessionStore } from '@/assets/lib/sessionstore'

const axios = inject('axios')
const store = useSessionStore();

const messageSuccess = "Successfully logged in!";
let success = ref(false);
let submitting = ref(false);

const submitForm = async (fields) => {
  store.lastUsername = fields.username;
  submitting.value = true;

  await axios.post('/account/login', fields).then(function (response) {
    success.value = response.status === 200;
    submitting.value = false;
    setMessage(response.data.msg, messageSuccess, success.value);
  }).catch(function (error) {
    success.value = error.response.status === 200;
    submitting.value = false;
    setMessage(error.response.data.msg, messageSuccess, success.value);
    handleMessageError(error);
  })

  if (success.value) {
    store.updateSession(axios); // TODO: Make login UX less confusing
  }
}
</script>

<style>
@import url(@/assets/css/forms.css);
</style>

<style scoped>
@import url(@/assets/css/message_util.css);
</style>

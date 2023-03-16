<template>
  <div class="signup">
    <h1>Create a <span class="--tr-muted-text">subjective.report</span> account 🚀</h1>

    <div class="DefaultView__message" id="DefaultView__message">
      <div class="DefaultView__message_text" id="DefaultView__message_text"></div>
    </div>
    <div class="DefaultView__form">
      <FormKit type="form" @submit="submitForm" submit-label="Signup">
        <!-- TODO: Make optional.
        -->
        <FormKit
            type="email"
            name="email"
            id="email"
            label="Email address"
            help="Used for password recovery. (Optional)"
            validation="required|email"
            placeholder="example@effectindex.com"
        />

        <!-- TODO: How do we handle errors here?
             TODO: The validation obviously isn't perfect, as it's quite annoying to replicate the ruleset with regex.
             TODO: We can use API calls as validation with FormKit, so we should use our backend in the future instead.
        -->
        <FormKit
            type="text"
            name="username"
            id="username"
            label="Username"
            help="Used to login. You can use letters, numbers and symbols."
            validation="required|length:3,32|matches:/^[a-z0-9_-]+$/"
            :validation-messages="{ matches: 'Must contain only lowercase letters, numbers, underscores or dashes.' }"
            placeholder="mark76"
        />

        <FormKit
            type="password"
            name="password"
            id="password"
            label="Password"
            help="Used to login. Must contain at least 2 symbols and 3 letters / numbers."
            validation="required|length:8,32"
            placeholder="----------"
        />
      </FormKit>
    </div>
  </div>
</template>

<script>
import LayoutDefault from "@/layouts/LayoutDefault.vue";

export default {
  name: "SignupView",
  created() {
    this.$emit('update:layout', LayoutDefault);
  }
}
</script>

<script setup>
import {inject} from 'vue'
import {handleMessageError, setMessage} from '@/assets/lib/message_util';

const router = inject('router')
const axios = inject('axios')
const messageSuccess = "Account successfully created!<br>You will be redirected to login in 3 seconds.";
let success = false;

const submitForm = async (fields) => {
  // don't do anything if the user presses the button again, for example, while waiting for a redirect
  if (success) {
    return
  }

  const location = `/login?username=${fields.username}`;

  axios.post('/account', fields).then(function (response) {
    success = response.status === 201;
    setMessage(response.data.msg, messageSuccess, success, router, location, 3000);
  }).catch(function (error) {
    success = error.response.status === 201;
    setMessage(error.response.data.msg, messageSuccess, success, router, location, 3000);
    handleMessageError(error)
  })
}
</script>

<style>
@import url(@/assets/css/forms.css);
</style>

<style scoped>
@import url(@/assets/css/message_util.css);
</style>

<template>
  <div class="login">
    <h1>Login to your TripReporter Account</h1>

    <div class="DefaultView__message" id="DefaultView__message">
      <div class="DefaultView__message_text" id="DefaultView__message_text"></div>
    </div>
    <div class="LoginView__form">
      <FormKit type="form" @submit="submitForm" submit-label="Login">
        <FormKit
            type="text"
            name="username"
            id="username"
            label="Username"
            validation="required"
            placeholder="mark76"
        />

        <FormKit
            type="password"
            name="password"
            id="password"
            label="Password"
            validation="required"
            placeholder="----------"
        />
      </FormKit>
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
import {inject} from 'vue'
import {handleMessageError, setMessage} from '@/assets/lib/message_util';

const axios = inject('axios')
const messageSuccess = "Successfully logged in!";

const submitForm = async (fields) => {
  // TODO: This is only DELETE for testing purposes (it's the only endpoint that verifies password hash).
  // TODO: Change ASAP.
  axios.delete('/account', {data: fields}).then(function (response) {
    setMessage(response.data.msg, messageSuccess, response.status === 200);
  }).catch(function (error) {
    setMessage(error.response.data.msg, messageSuccess, error.response.status === 200);
    handleMessageError(error)
  })
}
</script>

<style>
[data-type="submit"] .formkit-input {
  background: #3d9991;
}

[data-type="submit"] .formkit-input:hover {
  background: #3d9991;
  filter: brightness(75%);
}
</style>

<style scoped>
@import url(@/assets/css/message_util.css);

.LoginView__form {
  max-width: 25em;
  margin: auto;
  text-align: left;
}
</style>

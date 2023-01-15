<template>
  <div class="signup">
    <h1>Create your TripReporter Account</h1>

    <div class="DefaultView__message" id="DefaultView__message">
      <div class="DefaultView__message_text" id="DefaultView__message_text"></div>
    </div>
    <div class="SignupView__form">
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
export default {
  name: "SignupView"
}
</script>

<script setup>
import {inject} from 'vue'
import {setMessage} from '@/assets/lib/message_util';

const axios = inject('axios')
const messageSuccess = "Account successfully created!<br>You will be redirected to login in 3 seconds.";

const submitForm = async (fields) => {
  axios.post('/account', fields).then(function (response) {
    console.log(response)
    console.log(response.data)
    setMessage(response.data.msg, messageSuccess, response.status === 201);
  }).catch(function (error) {
    setMessage(error.response.data.msg, messageSuccess,error.response.status === 201);

    if (!error.response && error.request) {
      console.log(error.request);
    } else {
      console.log(`Error: ${error.message} - ${error}`);
    }
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

.SignupView__form {
  max-width: 25em;
  margin: auto;
  text-align: left;
}
</style>

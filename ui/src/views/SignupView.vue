<template>
  <div class="signup">
    <h1>Create your TripReporter Account</h1>

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
            validation="required|length:3,32|matches:/^[a-z0-9_-]+$/"
            placeholder="trmark76&!"
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
import { inject } from 'vue'
const axios = inject('axios')

const submitForm = async (fields) => {
  await new Promise((r) => setTimeout(r, 100))
  alert(JSON.stringify(fields))

  axios.get('/account').then(function (response) {
    console.log(response.data);
    console.log(response.status);
    console.log(response.statusText);
    console.log(response.headers);
    console.log(response.config);
  })
}
</script>

<style>
[data-type="submit"] .formkit-input {
  background: #3d9991;
}
</style>

<style scoped>
.SignupView__form {
  max-width: 25em;
  margin: auto;
  text-align: left;
}
</style>

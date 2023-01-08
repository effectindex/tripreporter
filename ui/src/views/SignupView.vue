<template>
  <div class="signup">
    <h1>Create your TripReporter Account</h1>

    <div class="SignupView__message" id="SignupView__message">
      <div class="SignupView__message_text" id="SignupView__message_text"></div>
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
import {inject} from 'vue'

const axios = inject('axios')

const submitForm = async (fields) => {
  axios.post('/account', fields).then(function (response) {
    console.log(response)
    console.log(response.data)
    setMessage(response.data.msg, response.status);
  }).catch(function (error) {
    setMessage(error.response.data.msg, error.response.status);

    if (!error.response && error.request) {
      console.log(error.request);
    } else {
      console.log(`Error: ${error.message} - ${error}`);
    }
  })
}

function setMessage(message, status) {
  const elemText = document.getElementById("SignupView__message_text");
  elemText.textContent = message;

  const elem = document.getElementById("SignupView__message")
  elem.style.display = 'block';

  if (status === 201) {
    elem.style.background = '#3d9991'
    elemText.innerHTML = "Account successfully created!<br>You will be redirected to login in 3 seconds.";
    window.setTimeout(function () {
      window.location.href = "/login";
    }, 3000);
  } else {
    elem.style.background = '#a83232'
  }
}

</script>

<style>
[data-type="submit"] .formkit-input {
  background: #3d9991;
}
</style>

<style scoped>
.SignupView__message {
  max-width: 25em;
  margin: auto;
  color: #ffffff;
  background: #a83232;
  border-radius: 2em;
  display: none;
}

.SignupView__message_text {
  display: flex;
  justify-content: center;
  align-content: center;
  flex-direction: column;
  height: 4em;
  line-height: 1.3em;
  margin-bottom: 15px;
}

.SignupView__form {
  max-width: 25em;
  margin: auto;
  text-align: left;
}
</style>

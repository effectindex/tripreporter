<template>
  <div class="login">
    <h1>Login to your TripReporter Account</h1>

    <div class="LoginView__message" id="LoginView__message">
      <div class="LoginView__message_text" id="LoginView__message_text"></div>
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
            placeholder="trmark76&!"
        />
      </FormKit>
    </div>
  </div>
</template>

<script>
export default {
  name: "LoginView"
}
</script>

<script setup>
import {inject} from 'vue'

const axios = inject('axios')

const submitForm = async (fields) => {
  // TODO: This is only DELETE for testing purposes (it's the only endpoint that verifies password hash).
  // TODO: Change ASAP.
  axios.delete('/account', {data: fields}).then(function (response) {
    console.log(fields)
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
  const elemText = document.getElementById("LoginView__message_text");
  elemText.textContent = message;

  const elem = document.getElementById("LoginView__message")
  elem.style.display = 'block';

  if (status === 200) {
    elem.style.background = '#3d9991'
    elemText.innerHTML = "Successfully logged in!";
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
.LoginView__message {
  max-width: 25em;
  margin: auto;
  color: #ffffff;
  background: #a83232;
  border-radius: 2em;
  display: none;
}

.LoginView__message_text {
  display: flex;
  justify-content: center;
  align-content: center;
  flex-direction: column;
  height: 4em;
  line-height: 1.3em;
  margin-bottom: 15px;
}

.LoginView__form {
  max-width: 25em;
  margin: auto;
  text-align: left;
}
</style>

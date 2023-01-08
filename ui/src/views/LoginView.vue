<template>
  <div class="login">
    <h1>Login to your TripReporter Account</h1>

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
const submitForm = async (fields) => {
  await new Promise((r) => setTimeout(r, 1000))
  alert(JSON.stringify(fields))
}
</script>

<style>
[data-type="submit"] .formkit-input {
  background: #3d9991;
}
</style>

<style scoped>
.LoginView__form {
  max-width: 25em;
  margin: auto;
  text-align: left;
}
</style>

<template>
  <div class="login">
    <h1>Login to your TripReporter Account</h1>

    <div class="LoginView__form">
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
    </div>
  </div>
</template>

<script>
export default {
  name: "LoginView"
}
</script>

<style scoped>
.LoginView__form {
  max-width: 25em;
  margin: auto;
}
</style>

<template>
  <div class="signup">
    <h1 class="--tr-not-bold">Create a <span class="--tr-muted-text">subjective.report</span> account 🚀</h1>

    <div class="DefaultView__message" id="DefaultView__message">
      <div class="DefaultView__message_text" id="DefaultView__message_text"></div>
    </div>

    <!--    meow-->
    <div class="DefaultView__form">
      <!--      <pre wrap>{{ store.createAccountForm }}</pre>-->
    </div>
    <div class="DefaultView__form">
      <FormKit type="form" @submit="submitForm" v-model="store.createAccountForm">

        <!--        <FormKit type="multi-step" tab-style="tab"></FormKit>-->
        <!--          <FormKit type="step"></FormKit>-->
        <!-- TODO: Make email optional. -->
        <!-- TODO: Implement user signup (#77) -->
        <div>
          <FormKit
              type="email"
              name="email"
              id="email"
              label="Email address"
              help="Used for password recovery."
              validation="required|email|(500)validateAccount"
              :validation-rules="{ validateAccount }"
              placeholder="lyv@effectindex.com"
          />

          <FormKit
              type="text"
              name="username"
              id="username"
              label="Username"
              help="Used to login. You can use letters, numbers and symbols."
              validation="required|length:3,32|(500)validateAccount"
              :validation-rules="{ validateAccount }"
              placeholder="lyv76"
          />

          <FormKit
              type="password"
              name="password"
              id="password"
              label="Password"
              help="Used to login. Must contain at least 2 symbols."
              validation="required|length:8,32|(100)validateAccount"
              :validation-rules="{ validateAccount }"
              placeholder="----------"
          />
        </div>
<!--        TODO: Implement user signup (#77) -->
<!--        <div v-show="!showFirstPage()">-->
<!--          <FormKit type="group" name="new_user" group="new_user">-->
<!--            <FormKit-->
<!--                type="text"-->
<!--                name="display_name"-->
<!--                id="display_name"-->
<!--                label="Display Name"-->
<!--                placeholder="Lyvergic Acid"-->
<!--                help="Shown to other users when viewing your profile."-->
<!--            />-->
<!--          </FormKit>-->
<!--        </div>-->
        <!--        <FormKit type="button" @click="" :disabled="!pageUser" label="Back"/>-->
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
  // TODO: Implement user signup (#77)
  // methods: {
  //   getUserSubmitLabel(value) {
  //     const labelEmpty = "Skip"
  //     const labelSubmit = "Submit"
  //
  //     // TODO: Add value.date_of_birth, value.height and value.weight once we have encryption implemented.
  //     if (value && value.display_name) {
  //       return labelSubmit
  //     }
  //
  //     return labelEmpty
  //   }
  // }
}
</script>

<script setup>
import {inject, ref} from 'vue'
import {handleMessageError, setMessage} from '@/assets/lib/message_util';
import {useSessionStore} from "@/assets/lib/sessionstore";

const router = inject('router')
const axios = inject('axios')
const store = useSessionStore();

const messageSuccess = "Account successfully created!<br>You will be redirected to login in 3 seconds.";
// let lastResponse = ref("");
// let pageUser = ref(false);
let success = ref(false);

const submitForm = async (fields) => {
  // don't do anything if the user presses the button again, for example, while waiting for a redirect
  if (success.value) {
    return
  }

  store.lastUsername = fields.username;

  // if (store.createAccountForm && store.createAccountForm.display_name) {
  //   fields.new_user = {"display_name": store.createAccountForm.display_name}
  // }

  // TODO: Implement user signup (#77)
  axios.post('/account', fields).then(function (response) {
    success.value = response.status === 201;
    // lastResponse.value = response.data.msg;
    // pageUser.value = lastResponse.value.startsWith("user: ")
    setMessage(response.data.msg, messageSuccess, success.value, router, '/login', 3000);
  }).catch(function (error) {
    success.value = error.response.status === 201;
    // lastResponse.value = error.response.data.msg;
    // pageUser.value = lastResponse.value.startsWith("user: ")
    setMessage(error.response.data.msg, messageSuccess, success.value, router, '/login', 3000);
    handleMessageError(error);
  })
}

const validateAccount = async (node) => {
  function getMsg(msg) {
    if (msg.startsWith("Email or username")) {
      return {[node.name]: `${node.name.charAt(0).toUpperCase()}${node.name.slice(1)} ${msg.slice(18)}`}
    }

    return {[node.name]: msg}
  }

  axios.post('/account/validate', {[node.name]: node.value}).then(function (response) {
    if (response.status !== 200) {
      node.setErrors([], getMsg(response.data.msg))
    }
    return response.status === 200
  }).catch(function (error) {
    if (error.response.status !== 200) {
      node.setErrors([], getMsg(error.response.data.msg))
    }
    return error.response.status === 200
  })
}
</script>

<style>
@import url(@/assets/css/forms.css);
</style>

<style scoped>
@import url(@/assets/css/message_util.css);
</style>

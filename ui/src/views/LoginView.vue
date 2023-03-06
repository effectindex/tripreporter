<template>
  <div class="login">
    <div v-if="!store.activeSession" class="no-session">
      <h1>Login to your TripReporter Account</h1>

      <div class="DefaultView__message" id="DefaultView__message">
        <div class="DefaultView__message_text" id="DefaultView__message_text"></div>
      </div>
      <div class="DefaultView__form">
        <FormKit type="form" @submit="submitForm" submit-label="Login">
          <FormKit
              type="text"
              name="username"
              id="username"
              label="Username"
              validation="required"
              placeholder="mark76"
              :value="queryUsername()"
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
    <div v-else>
      <h1>You're logged in! 🎉</h1>
    </div>
  </div>
</template>

<script>
import LayoutDefault from "@/layouts/LayoutDefault.vue";

export default {
  name: "LoginView",
  created() {
    this.$emit('update:layout', LayoutDefault);
  },
  methods: {
    queryUsername() {
      return this.$route.query.username ? this.$route.query.username : "";
    }
  }
}
</script>

<script setup>
import {inject} from 'vue'
import {handleMessageError, setMessage} from '@/assets/lib/message_util';
import {useSessionStore} from '@/assets/lib/sessionstore'

const axios = inject('axios')
const store = useSessionStore();

const messageSuccess = "Successfully logged in!";
let success = false;

const submitForm = async (fields) => {
  // don't do anything if the user presses the button again, for example, while waiting for a redirect
  if (success) {
    return
  }

  await axios.post('/account/login', fields).then(function (response) {
    success = response.status === 200;
    setMessage(response.data.msg, messageSuccess, success);
  }).catch(function (error) {
    success = error.response.status === 200;
    setMessage(error.response.data.msg, messageSuccess, success);
    handleMessageError(error)
  })

  if (success) {
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

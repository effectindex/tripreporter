<template>
  <div class="signup">
    <h1 class="--tr-not-bold">Create a <span class="--tr-muted-text">subjective.report</span> account 🚀</h1>

    <div class="DefaultView__message" id="DefaultView__message">
      <div class="DefaultView__message_text" id="DefaultView__message_text"></div>
    </div>

    <!--    meow-->
    <div class="DefaultView__form">
      <pre wrap>{{ store.createAccountForm }}</pre>
      <pre wrap>{{ store.createUserForm }}</pre>
    </div>
    <div class="DefaultView__form">
      <FormKit type="form" @submit="submitForm" :actions="false">

        <!--        <FormKit type="multi-step" tab-style="tab"></FormKit>-->
        <!--          <FormKit type="step"></FormKit>-->
        <!-- TODO: Make email optional. -->
        <!-- TODO: Implement user signup (#77) -->
        <FormKit type="multi-step" name="account_form" tab-style="progress" :hide-progress-labels="true"
                 :allow-incomplete="false">
          <FormKit type="step" name="account_info" v-model="store.createAccountForm">
            <FormKit
                type="email"
                name="email"
                id="email"
                label="Email address"
                help="Used for password recovery."
                preserve-errors="true"
                validation="required|email|validateAccount"
                validation-visibility="dirty"
                :validation-rules="{ validateAccount }"
                placeholder="lyv@effectindex.com"
            />

            <FormKit
                type="text"
                name="username"
                id="username"
                label="Username"
                help="Used to login. You can use letters, numbers and symbols."
                preserve-errors="true"
                validation="required|length:3,32|validateAccount"
                validation-visibility="dirty"
                :validation-rules="{ validateAccount }"
                placeholder="lyv76"
            />

            <FormKit
                type="password"
                name="password"
                id="password"
                label="Password"
                help="Used to login. Must contain at least 2 symbols."
                preserve-errors="true"
                validation="required|length:8,32|validateAccount"
                validation-visibility="dirty"
                :validation-rules="{ validateAccount }"
                placeholder="----------"
            />
          </FormKit>
          <FormKit type="step" name="user_info" v-model="store.createUserForm">
            <FormKit
                type="text"
                name="display_name"
                id="display_name"
                label="Display Name"
                placeholder="Lyvergic Acid"
                help="Shown to other users when viewing your profile."
            />

            <!--suppress VueUnrecognizedSlot -->
            <template #stepNext>
              <FormKit type="submit" data-next="true"/>
            </template>
          </FormKit>
        </FormKit>
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
import {inject, ref} from "vue"
import {handleMessageError, setMessage} from "@/assets/lib/message_util";
import log from "@/assets/lib/logger";
import {useSessionStore} from "@/assets/lib/sessionstore";

const router = inject('router')
const axios = inject('axios')
const store = useSessionStore();

// const accountFormID = "create-account-form";
const messageSuccess = "Account successfully created!<br>You will be redirected to login in 3 seconds.";
// let lastResponse = ref("");
// let pageUser = ref(false);
let success = ref(false);

// eslint-disable-next-line no-unused-vars
const submitForm = async (fields, handlers) => {
  log("submitForm", fields)
  // don't do anything if the user presses the button again, for example, while waiting for a redirect
  // if (success.value) {
  //   return
  // }

  let lastPage = true;

  handlers.children[0].walk(child => {
    if (child.name === "account_info") {
      child.context.handlers.incrementStep(1, child.context)()
      lastPage = false
      log("account_info", lastPage, child.context)
    }

    if (child.name === "user_info") {
      lastPage = child.context.isLastStep
      log("user_info", lastPage, child.context)
    }
  })

  log("submitForm: if", lastPage)
  if (lastPage) {
    store.lastUsername = fields.username;

    log("submitForm: got store", store.createAccountForm)
    if (store.createAccountForm) {
      store.createAccountForm.new_user = store.createUserForm
    }
    log("submitForm: got full store", store.createAccountForm)

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
}

let validationCache = ref({})
let nodeListeners = ref({})

// This is used to validate our account fields with the API, as well as caching the validation to limit API requests.
const validateAccount = async (node) => {
  // getCached will look for a [node.name][node.value] and see if we already have a cached error for this
  // specific node's input value.
  function getCached() {
    if (validationCache.value[node.name] && validationCache.value[node.name][`${node.value}`]) {
      return validationCache.value[node.name][`${node.value}`]
    }

    return undefined
  }

  // Insert into the cache when we get an error from the API.
  function setCached(msg) {
    if (validationCache.value[node.name]) {
      validationCache.value[node.name][`${node.value}`] = msg
    } else {
      validationCache.value[node.name] = {[`${node.value}`]: msg}
    }
  }

  // This will create a struct to use for node.setErrors, in addition to setting the cache.
  function getMsg(msg, setCache) {
    if (setCache) {
      setCached(msg)
    }

    if (msg.startsWith("Email or username")) {
      return {[node.name]: `${node.name.charAt(0).toUpperCase()}${node.name.slice(1)} ${msg.slice(18)}`}
    }

    return {[node.name]: msg}
  }

  // If we haven't created a listener for this node yet, make one
  if (!nodeListeners.value[node.name]) {
    nodeListeners.value[node.name] = node.on('commit', ({payload}) => {
      // We do this to ensure previous API errors are not displayed when the input is being validated by 'required'
      // or another validation rule that is not ours.
      node.clearErrors()
    })
  }

  // First we want to check if we already have a cached error for the [node.name][node.input] value, to avoid making
  // an unnecessary API request.
  const cachedMsg = getCached()
  if (cachedMsg !== undefined) {
    return new Promise((resolve) => {
      node.setErrors([], getMsg(cachedMsg, false))
      resolve(cachedMsg)
    })
  }

  // Now that we know we don't have a cached error, check if the input is valid using the API.
  axios.post('/account/validate', {[node.name]: node.value}).then(function (response) {
    if (response.status !== 200) {
      node.setErrors([], getMsg(response.data.msg, true))
    } else {
      node.clearErrors()
    }
    return response.status === 200
  }).catch(function (error) {
    if (error.response.status !== 200) {
      node.setErrors([], getMsg(error.response.data.msg, true))
    } else {
      node.clearErrors()
    }
    return error.response.status === 200
  })
}
</script>

<style>
@import url(@/assets/css/forms.css);
@import url(@/assets/css/forms-multi-step.css);
</style>

<style scoped>
@import url(@/assets/css/message_util.css);
</style>

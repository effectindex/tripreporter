<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <div class="signup">
    <h1 class="--tr-header-h1">Create a <span class="--tr-muted-text">subjective.report</span> account 🚀</h1>

    <div class="DefaultView__message" id="DefaultView__message">
      <div class="DefaultView__message_text" id="DefaultView__message_text"></div>
    </div>

    <div class="DefaultView__form">
      <FormKit type="form" @submit="submitForm" #default="{ state: { errors } }" :actions="false">
        <!-- TODO: Make email optional. -->
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
                help="(optional) Shown to other users when viewing your profile."
            />

            <!--suppress VueUnrecognizedSlot -->
            <template #stepNext>
              <FormKit type="submit" data-next="true" :disabled="errors && submitting"/>
            </template>
          </FormKit>
        </FormKit>
      </FormKit>
    </div>
  </div>
</template>

<script>
import LayoutDefault from "@/layouts/LayoutDefault.vue";
import { getNode } from "@formkit/core";

export default {
  name: "SignupView",
  created() {
    this.$emit('update:layout', LayoutDefault);
  },
  mounted() {
    // Ensure username is lowercase when entering it
    getNode('username').hook.commit((payload, next) => {
      return next(payload.toLowerCase());
    });
  }
}
</script>

<script setup>
import { inject, ref } from "vue"
import { handleMessageError, setMessage } from "@/assets/lib/message_util";
import log from "@/assets/lib/logger";
import { useSessionStore } from "@/assets/lib/sessionstore";

const router = inject('router')
const axios = inject('axios')
const store = useSessionStore();

const messageSuccess = "Account successfully created!<br>You will be redirected to login in 3 seconds.";
let success = ref(false);
let submitting = ref(false);

const submitForm = async (fields, handlers) => {
  log("submitForm", fields)

  let lastPage = true;
  let makeUserInputActive = false;

  handlers.children[0].walk(child => {
    if (child.name === "account_info") {
      // Used to disable submit button if on last page
      // Increment to the next page if the user pressed enter on the first page.
      // TODO: Workaround for https://github.com/formkit/formkit/issues/641
      if (child.context.isActiveStep) {
        lastPage = false
        makeUserInputActive = true
        child.context.handlers.next()
      }
    }

    // Focus the next text input if the user presses enter, and we've switched to the last page
    if (child.name === "user_info" && makeUserInputActive) {
      setTimeout(function () {
        const el = document.getElementById(child.context.node.children[0].props.id)
        el.focus()
        el.click()
      }, 20)
    }
  })

  submitting.value = true;

  if (lastPage) {
    if (store.createAccountForm) {
      store.lastUsername = store.createAccountForm["username"];
      store.createAccountForm.new_user = store.createUserForm
    }

    log("submitForm: got store", store.createAccountForm)

    axios.post('/account', store.createAccountForm).then(function (response) {
      success.value = response.status === 201;
      submitting.value = false;
      validationCache.value = {};
      setMessage(response.data.msg, messageSuccess, success.value, router, '/login', 3000);
    }).catch(function (error) {
      success.value = error.response.status === 201;
      submitting.value = false;
      validationCache.value = {};
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
      validationCache.value[node.name] = { [`${node.value}`]: msg }
    }
  }

  // This will create a struct to use for node.setErrors, in addition to setting the cache.
  function getMsg(msg, setCache) {
    if (setCache) {
      setCached(msg)
    }

    if (msg.startsWith("Email or username")) {
      return { [node.name]: `${node.name.charAt(0).toUpperCase()}${node.name.slice(1)} ${msg.slice(18)}` }
    }

    return { [node.name]: msg }
  }

  // If we haven't created a listener for this node yet, make one
  if (!nodeListeners.value[node.name]) {
    nodeListeners.value[node.name] = node.on('commit', ({ payload }) => {
      // We do this to ensure previous API errors are not displayed when the input is being validated by 'required'
      // or another validation rule that is not ours.
      node.clearErrors()
    })
  }

  // Next we want to check if we already have a cached error for the [node.name][node.input] value, to avoid making
  // an unnecessary API request.
  const cachedMsg = getCached()
  if (cachedMsg !== undefined) {
    return new Promise((resolve) => {
      node.setErrors([], getMsg(cachedMsg, false))
      resolve(cachedMsg)
    })
  }

  // Pause validation if we're currently submitting the form
  if (submitting.value) {
    return new Promise((resolve) => {
      log("skipping validation", validationCache.value, node.store.submitted)
      resolve(true)
    })
  }

  // Now that we know we don't have a cached error, check if the input is valid using the API.
  axios.post('/account/validate', { [node.name]: node.value }).then(function (response) {
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

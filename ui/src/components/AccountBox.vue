<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <h1 v-if="!getStore().showDeleteForm" class="--tr-header-h1">Manage Account</h1>
  <h1 v-else class="--tr-header-h1">Delete Account</h1>

  <div class="DefaultView__message" id="DefaultView__message">
    <div class="DefaultView__message_text" id="DefaultView__message_text"></div>
  </div>

  <div class="LayoutAccount__main" v-if="isLoaded()">
    <div v-show="!getStore().showDeleteForm">
      <div v-for="(account, index) in [getStore().accountJson]" :key="index">
        <div class="LayoutAccount__account">
          <div class="LayoutAccount__info">
            <div class="LayoutAccount__info_entry">
              <!-- TODO: Add edit buttons -->
              <header-row-box
                  class="LayoutAccount__info_entry_box"
                  style="margin-left: 0;"
                  header="Account"
                  icon="user"
                  :columns="['Username', 'Email']"
                  :rows="{
                    'Username': getStore().accountJson.username,
                    'Email': getStore().accountJson.email,
                  }"
              />

              <div class="LayoutAccount_buttons">
                <FormKit
                    type="button"
                    class="LayoutAccount__buttons_button"
                    label="Logout"
                    @click="logOut"
                    :disabled="getSession().activeSession === false"
                />
                <FormKit
                    type="button"
                    style="background-color: var(--tr-error)"
                    class="LayoutAccount__buttons_button"
                    label="Delete Account"
                    @click="showDeleteForm"
                    :disabled="getSession().activeSession === false"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div v-show="getStore().showDeleteForm">

      <div class="DefaultView__form">
        <FormKit type="form" @submit="submitForm" #default="{ state: { valid } }" :actions="false">
          <FormKit
              type="email"
              name="email"
              id="email"
              label="Email address"
              help="The email associated with your account."
              validation="required|email"
              placeholder="lyv@effectindex.com"
          />

          <FormKit
              type="password"
              name="password"
              id="password"
              label="Password"
              help="The password associated with your account."
              validation="required|length:8,32"
              placeholder="----------"
          />

          <div class="LayoutAccount_buttons">
            <FormKit
                type="button"
                class="LayoutAccount__buttons_button"
                label="Back"
                @click="hideDeleteForm"
                :disabled="getSession().activeSession === false"
            />
            <FormKit
                type="submit"
                style="background-color: var(--tr-error)"
                class="LayoutAccount__buttons_button"
                label="Delete Account"
                data-next="true"
                :disabled="!valid || getSession().activeSession === false"
            />
          </div>
        </FormKit>
      </div>
    </div>
  </div>
</template>

<script>
import HeaderRowBox from "@/components/HeaderRowBox.vue";
import { useAccountStore } from "@/assets/lib/accountstore";
import { useSessionStore } from "@/assets/lib/sessionstore";

const store = useAccountStore();
const sessionStore = useSessionStore();

export default {
  name: "AccountBox",
  components: { HeaderRowBox },
  methods: {
    getStore() {
      return store
    },
    getSession() {
      return sessionStore
    },
    isLoaded() {
      return store.isLoaded()
    }
  }
}
</script>

<script async setup>
import { inject, ref } from "vue";
import { handleMessageError, setMessage } from "@/assets/lib/message_util";
import log from "@/assets/lib/logger";

const router = inject('router')
const axios = inject('axios')

const messageSuccess = "Account successfully deleted!<br>You will be redirected to the home page in 3 seconds.";
const messageLoggedOut = "Logged out successfully!<br>You will be redirected to login in 3 seconds.";
let success = ref(false);
let loggedOut = ref(false);
let submitting = ref(false);
let ranSetup = false;

const submitForm = async (fields) => {
  log("submitForm", fields)

  submitting.value = true;

  axios.delete('/account', { data: fields }).then(function (response) {
    success.value = response.status === 200;
    submitting.value = false;
    if (success.value === true) {
      store.showDeleteForm = false;
      sessionStore.invalidateSession();
    }
    setMessage(response.data.msg, messageSuccess, success.value, router, '/', 3000);
  }).catch(function (error) {
    success.value = error.response.status === 200;
    submitting.value = false;
    if (success.value === true) {
      store.showDeleteForm = false;
      sessionStore.invalidateSession();
    }
    setMessage(error.response.data.msg, messageSuccess, success.value, router, '/', 3000);
    handleMessageError(error);
  })
}

const showDeleteForm = (e) => {
  store.showDeleteForm = true
}

const hideDeleteForm = (e) => {
  store.showDeleteForm = false
}

const logOut = (e) => {
  submitting.value = true

  axios.delete('/session').then(function (response) {
    loggedOut.value = response.status === 200;
    submitting.value = false;
    if (loggedOut.value === true) {
      store.showDeleteForm = false;
      sessionStore.invalidateSession();
    }
    setMessage(response.data.msg, messageLoggedOut, loggedOut.value, router, '/login', 3000);
  }).catch(function (error) {
    loggedOut.value = error.response.status === 200;
    submitting.value = false;
    if (loggedOut.value === true) {
      store.showDeleteForm = false;
      sessionStore.invalidateSession();
    }
    setMessage(error.response.data.msg, messageLoggedOut, loggedOut.value, router, '/login', 3000);
    handleMessageError(error);
  })
}

if (ranSetup !== true) {
  ranSetup = true

  await axios.get('/account').then(function (response) {
    store.updateData(response.status, response.data)
    setMessage(response.data.msg, "", store.apiSuccess);
  }).catch(function (error) {
    log("error", error)
    store.updateData(error.response.status, error.response.data)
    setMessage(error.response.data.msg, "", store.apiSuccess);
    handleMessageError(error);
  })
}
</script>
<style scoped>
.LayoutAccount__main {
    text-align: left;
}

.LayoutAccount__main h1 {
    text-align: center;
}

.LayoutAccount__account {
    max-width: 25em;
    margin: auto;
}

.LayoutAccount__info_entry_box {
    margin-bottom: 1em;
}

.LayoutAccount_buttons {
    display: flex;
    flex-wrap: wrap;
    flex-direction: row;
    justify-content: left;
}

.LayoutAccount__buttons_button {
    flex-grow: 1;
}
</style>
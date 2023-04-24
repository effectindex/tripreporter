<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <div class="LayoutAccount__main" v-if="isLoaded()">
    <div v-for="(account, index) in [getStore().accountJson]" :key="index">
      <h1 class="--tr-header-h1">Manage Account</h1>

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
                  class="LayoutAccount__buttons_button"
                  type="button"
                  label="Logout"
                  @click="randomColor"
              />
              <FormKit
                  class="LayoutAccount__buttons_button"
                  type="button"
                  label="Delete Account"
                  @click="randomColor"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import HeaderRowBox from "@/components/HeaderRowBox.vue";
import { useAccountStore } from "@/assets/lib/accountstore";
import { inject } from "vue";
import { handleMessageError, setMessage } from "@/assets/lib/message_util";
import log from "@/assets/lib/logger";

const store = useAccountStore();
let ranSetup = false;

export default {
  name: "AccountBox",
  components: { HeaderRowBox },
  methods: {
    getStore() {
      return store
    },
    isLoaded() {
      return store.isLoaded()
    }
  },
  async setup() {
    if (ranSetup === true) {
      return
    }
    ranSetup = true

    const axios = inject('axios')
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
}
</script>

<script setup>
const randomColor = (e) => {
  const hex = Math
      .floor(Math.random() * 16777215)
      .toString(16)
  e.target.setAttribute(
      'style',
      'background-color: #' + hex
  )
}
</script>
<style scoped>
@import url(@/assets/css/message_util.css);

.LayoutAccount__main {
    text-align: left;
}

.LayoutAccount__main h1 {
    text-align: center;
}

.LayoutAccount__account {
    max-width: 90%;
    margin: auto;
}

/* override LayoutAccount__account for desktop browsers */
@media only screen and (min-width: 680px) {
    .LayoutAccount__account {
        max-width: 75vw;
    }
}

.LayoutAccount__info_entry_box {
    margin-bottom: 1em;
}

.LayoutAccount_buttons {
    display: flex;
    flex-wrap: wrap;
    flex-direction: row;
    align-items: baseline;
    justify-content: left;
    margin-bottom: 1em;
}

.LayoutAccount__buttons_button {
    flex-grow: 1;
    margin-bottom: 1em;
}

.LayoutAccount__buttons_button:last-child {
    margin-bottom: 0;
}
</style>
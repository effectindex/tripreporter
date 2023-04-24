<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <div class="LayoutUser__main" v-if="getStore().isLoaded()">
    <div v-for="(user, index) in [getStore().user]" :key="index">
      <h1 class="--tr-header-h1">{{ user.display_name }}</h1>

      <div class="LayoutUser__user">
        <div class="LayoutReport__report_summary">
          <div class="LayoutReport__report_summary_entry">
            <header-row-box
                style="margin-left: 0;"
                header="User"
                icon="user"
                :columns="['Name', 'Created']"
                :rows="{
                  'Name': getStore().user.display_name,
                  'Created': getStore().createdDate.get()
                }"
            />
          </div>
        </div>
        <pre wrap>{{ getStore().user }}</pre>

      </div>
    </div>
  </div>
</template>

<script>
import { useUserStore } from "@/assets/lib/userstore";
import { inject } from "vue";
import { handleMessageError, setMessage } from "@/assets/lib/message_util";
import log from "@/assets/lib/logger";
import HeaderRowBox from "@/components/HeaderRowBox.vue";

const store = useUserStore();
let ranSetup = false

export default {
  name: "UserBox",
  components: { HeaderRowBox },
  props: {
    id: String
  },
  methods: {
    getStore() {
      return store
    },
  },
  async setup(props) {
    if (ranSetup) {
      return
    }
    ranSetup = true

    const axios = inject('axios')
    await axios.get('/user/' + props.id).then(function (response) {
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

<style scoped>
@import url(@/assets/css/message_util.css);

.LayoutUser__main {
    text-align: left;
}

.LayoutUser__main h1 {
    text-align: center;
}

.LayoutUser__user {
    max-width: 75%;
    margin: auto;
}
</style>
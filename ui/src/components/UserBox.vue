<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <div class="LayoutUser__main" v-if="isLoaded()">
    <div v-for="(user, index) in [getStore().user]" :key="index">
      <h1 class="--tr-header-h1">{{ user.display_name }}</h1>

      <div class="LayoutUser__user">
        <div class="LayoutUser__report_summary">
          <div class="LayoutUser__report_summary_entry">
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
        <header-column-box
            v-if="user.reports && user.reports.length > 0"
            header="Reports"
            icon="report"
            :columns="['Title', 'Date', 'Substances']"
            :rows="getRows(user.reports)"
            :links="getLinks(user.reports)"
        />
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
import HeaderColumnBox from "@/components/HeaderColumnBox.vue";
import Timestamp from "@/assets/lib/timestamp";
import titleCase from "@/assets/lib/string_util";

const store = useUserStore();
let state = new Map();

export default {
  name: "UserBox",
  computed: {
    Timestamp() {
      return Timestamp
    }
  },
  components: { HeaderColumnBox, HeaderRowBox },
  props: {
    id: String
  },
  methods: {
    getStore() {
      return store ? store.m.get(this.id) : undefined
    },
    isLoaded() {
      return this.getStore() !== undefined && store.isLoaded(this.id);
    },
    getRows(reports) {
      let rows = []
      reports.forEach(r => {
        const substances = r.drugs.map(drug => titleCase(drug.name));
        rows.push({ 'Title': r.title, 'Date': new Timestamp({date: r.report_date, longFormat: true}).get(), 'Substances': substances.join(", ") })
      })
      return rows
    },
    getLinks(reports) {
      let links = []
      reports.forEach(r => {
        links.push({ 'Title': `/report/${r.id}` })
      })
      return links
    }
  },
  async setup(props) {
    let user = state.get(props.id)
    if (user === undefined) {
      user = {
        ranSetup: Boolean
      }
    }

    if (user.ranSetup === true) {
      return
    }
    user.ranSetup = true
    state.set(props.id, user)

    const axios = inject('axios')
    await axios.get('/user/' + props.id).then(function (response) {
      store.updateData(props.id, response.status, response.data)
      setMessage(response.data.msg, "", store.m.get(props.id) ? store.m.get(props.id).apiSuccess : false);
    }).catch(function (error) {
      log("error", error)
      store.updateData(props.id, error.response.status, error.response.data)
      setMessage(error.response.data.msg, "", store.m.get(props.id) ? store.m.get(props.id).apiSuccess : false);
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
    max-width: 90%;
    margin: auto auto 1em auto;
}

/* override LayoutUser__user for desktop browsers */
@media only screen and (min-width: 680px) {
    .LayoutUser__user {
        max-width: 75vw;
    }
}

.LayoutUser__report_summary {
    display: flex;
    flex-direction: row;
    align-items: baseline;
    justify-content: center;
}

.LayoutUser__report_summary_entry {
    flex-grow: 1;
    margin-right: 1em;
}

.LayoutUser__report_summary_entry:last-child {
    margin-right: 0;
}
</style>
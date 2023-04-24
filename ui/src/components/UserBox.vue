<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <div class="LayoutUser__main" v-if="getStore().isLoaded()">
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
let ranSetup = false

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
      return store
    },
    getRows(reports) {
      let rows = []
      reports.forEach(r => {
        const substances = r.drugs.map(drug => titleCase(drug.name));
        rows.push({ 'Title': r.title, 'Date': new Timestamp({date: r.report_date, longFormat: true}).get(), 'Substances': substances.join(", ") })
      })
      console.log("rows", rows)
      return rows
    },
    getLinks(reports) {
      let links = []
      reports.forEach(r => {
        links.push({ 'Title': `/reports?id=${r.id}` })
      })
      console.log("links", links)
      return links
    }
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


.LayoutUser__report_summary {
    display: flex;
    flex-direction: row;
    align-items: baseline;
    justify-content: center;
    margin-bottom: 1em;
}

.LayoutUser__report_summary_entry {
    flex-grow: 1;
    margin-right: 1em;
}

.LayoutUser__report_summary_entry:last-child {
    margin-right: 0;
}
</style>
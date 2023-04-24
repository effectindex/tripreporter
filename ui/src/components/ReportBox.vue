<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <div class="LayoutReport__main" v-if="isLoaded()">
    <div v-for="(report, index) in [getStore().reportJson]" :key="index">
      <h1 class="--tr-header-h1">{{ report.title }}</h1>

      <div class="LayoutReport__report">
        <div class="LayoutReport__report_summary">
          <!--          TODO: See #99 / #100 -->
          <div class="LayoutReport__report_summary_entry">
            <header-row-box
                style="margin-left: 0;"
                header="User"
                icon="user"
                :columns="['Name', 'Date', 'Gender', 'Height', 'Weight']"
                :rows="{
                  'Name': getStore().reportUser.display_name,
                  'Date': getStore().reportDate.get(),
                  'Gender': getStore().reportSubject.gender,
                  'Height': getStore().reportSubject.height(),
                  'Weight': getStore().reportSubject.weight(),
                }"
                :links="{
                  'Name': `/profile?id=${getStore().reportUser.id}`
                }"
            />
          </div>
          <div class="LayoutReport__report_summary_entry">
            <drug-summary-box header="Substances" :events="report.report_events"/>
          </div>
          <div class="LayoutReport__report_summary_entry">
            <drug-summary-box style="margin-right: 0;" header="Medication" :events="report.medications"/>
          </div>
        </div>

        <div class="LayoutReport__setting">
          <div :class="{'LayoutReportBox': true, 'LayoutReportBox_last': true}">
            <div v-if="getStore().reportDate.valid()" class="LayoutReport__setting_date">
              Experienced on <timestamp-text :timestamp="getStore().reportDate"/>
            </div>
            <div v-else>
              Unknown report date.
            </div>
            <div v-if="report.setting" class="LayoutReport__setting_text">
              {{ report.setting }}
            </div>
          </div>
        </div>

        <div class="LayoutReport__events">
          <div v-for="(event, index) in report.report_events" :key="index">
            <report-event-box :class="getClasses(index, report.report_events)" :event="event"
                :showSection="showCurrentSection(index, report.report_events)"/>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { inject } from "vue";
import { handleMessageError, setMessage } from "@/assets/lib/message_util";
import { useReportStore } from "@/assets/lib/reportstore";
import TimestampText from "@/components/TimestampText.vue";
import ReportEventBox from "@/components/ReportEventBox.vue";
import DrugSummaryBox from "@/components/DrugSummaryBox.vue";
import HeaderRowBox from "@/components/HeaderRowBox.vue";
import Timestamp from "@/assets/lib/timestamp";
import log from "@/assets/lib/logger";

const store = useReportStore();
let state = new Map();

export default {
  name: "ReportBox",
  computed: {
    Timestamp() {
      return Timestamp
    }
  },
  components: {
    HeaderRowBox,
    DrugSummaryBox,
    ReportEventBox,
    TimestampText
  },
  props: {
    id: String
  },
  methods: {
    getStore() {
      return store.m.get(this.id)
    },
    isLoaded() {
      return this.getStore() !== undefined && store.isLoaded(this.id);
    },
    showCurrentSection(index, events) {
      if (index === 0) {
        return true;
      }

      return !!(events[index - 1] && events[index].section !== events[index - 1].section);
    },
    getClasses(index, sections) {
      const boxClass = index % 2 === 0 ? 'LayoutReportBox' : 'LayoutReportBox_alt'
      const classes = {}
      classes[boxClass] = true
      classes["LayoutReportBox_last"] = index === sections.length - 1

      return classes
    }
  },
  async setup(props) {
    let report = state.get(props.id)
    if (report === undefined) {
      report = {
        ranSetup: Boolean
      }
    }

    if (report.ranSetup === true) {
      return
    }
    report.ranSetup = true
    state.set(props.id, report)

    const axios = inject('axios')
    await axios.get('/report/' + props.id).then(function (response) {
      log("axios updateData", store)
      store.updateData(props.id, response.status, response.data)
      setMessage(response.data.msg, "", store.m.get(props.id) ? store.m.get(props.id).apiSuccess : false);
    }).catch(function (error) {
      store.updateData(props.id, error.response.status, error.response.data)
      setMessage(error.response.data.msg, "", store.m.get(props.id) ? store.m.get(props.id).apiSuccess : false);
      handleMessageError(error);
    })
  }
}
</script>

<style scoped>
@import url(@/assets/css/message_util.css);

.LayoutReport__main {
    text-align: left;
}

.LayoutReport__main h1 {
    text-align: center;
}

.LayoutReport__report {
    max-width: 90%;
    margin: auto;
}

/* override LayoutReport__report for desktop browsers */
@media only screen and (min-width: 680px) {
    .LayoutReport__report {
        max-width: 75vw;
    }
}


.LayoutReport__report_summary {
    display: flex;
    flex-direction: row;
    align-items: baseline;
    justify-content: center;
    margin-bottom: 1em;
}

.LayoutReport__report_summary_entry {
    flex-grow: 1;
    margin-right: 1em;
}

.LayoutReport__report_summary_entry:last-child {
    margin-right: 0;
}

.LayoutReport__setting, .LayoutReport__events {
    background-color: var(--tr-background-raised-light);
    border-radius: 10px;
    border: 1px solid var(--tr-border);
    position: relative;
    margin-bottom: 1em;
}

.LayoutReport__setting_date {
    font-weight: bold;
}

.LayoutReport__setting_text {
    margin: 0.5em;
}

.LayoutReportBox, .LayoutReportBox_alt {
    padding: 10px;
    border-bottom: 1px solid var(--tr-border);
}

.LayoutReportBox_alt {
    background-color: var(--tr-background-raised);
}

.LayoutReportBox_last {
    border-bottom: none;
}
</style>

<template>
  <div class="LayoutReport__main" v-if="getStore().isLoaded()">
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
                  'Name': 'Josie Kins',
                  'Date': new Timestamp({date: report.report_date, longFormat: true}).get(),
                  'Gender': 'Female',
                  'Height': '5\'9',
                  'Weight': '~150 lbs'
              }"
            />
          </div>
          <div class="LayoutReport__report_summary_entry">
            <drug-summary-box style="margin-right: 0;" :events="report.report_events"/>
          </div>
        </div>

        <div class="LayoutReport__setting">
          <div :class="{'LayoutReportBox': true, 'LayoutReportBox_last': true}">
            Experienced on
            <timestamp-text :date="report.report_date" :long-format="true"/>
            <br>
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
import { useReportsStore } from "@/assets/lib/reportsstore";
import TimestampText from "@/components/TimestampText.vue";
import ReportEventBox from "@/components/ReportEventBox.vue";
import DrugSummaryBox from "@/components/DrugSummaryBox.vue";
import HeaderRowBox from "@/components/HeaderRowBox.vue";
import Timestamp from "@/assets/lib/timestamp";

const store = useReportsStore();
let ranSetup = false

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
      return store
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
      classes['LayoutReportBox_last'] = index === sections.length - 1

      return classes
    }
  },
  async setup(props) {
    if (ranSetup) {
      return
    }
    ranSetup = true

    const axios = inject('axios')

    await axios.get('/report/' + props.id).then(function (response) {
      store.updateData(response.status, response.data)
      setMessage(response.data.msg, "", store.apiSuccess);
    }).catch(function (error) {
      store.updateData(error.response.status, error.response.data)
      setMessage(error.response.data.msg, "", store.apiSuccess);
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
    max-width: 75%;
    margin: auto;
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
}

.LayoutReport__report_summary_entry:first-child {
    margin-right: 1em;
}

.LayoutReport__setting, .LayoutReport__events {
    background-color: #fbfbfb;
    border-radius: 10px;
    border: 1px solid hsla(0, 0%, 66.7%, .35);
    position: relative;
    margin-bottom: 1em;
}

.LayoutReport__setting_text {
    margin: 0.5em;
}

.LayoutReportBox, .LayoutReportBox_alt {
    padding: 10px;
    border-bottom: 1px solid hsla(0, 0%, 66.7%, .35);
}

.LayoutReportBox_alt {
    background-color: #f6f6f6;
}

.LayoutReportBox_last {
    border-bottom: none;
}
</style>

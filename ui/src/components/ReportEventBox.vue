<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <div class="ReportEventBoxSection__wrapper">
    <div class="ReportEventBoxSection">
      <div :class="getSectionClass(event.section)">{{ showSection ? getSection(event.section) : "" }}</div>
    </div>

    <div v-if="event.timestamp !== '0001-01-01T00:00:00Z'" class="ReportEventBoxTime">
      <timestamp-text :date="event.timestamp" :hide-date="true"/>
    </div>
    <div v-if="event['type'] === 1">
      {{ event.content }}
    </div>
    <div v-else>
      <drug-box :drug="new DrugData({obj: event.drug})"/>
    </div>
  </div>
</template>

<script>
import TimestampText from "@/components/TimestampText.vue";
import DrugBox from "@/components/DrugBox.vue";
import ReportEvent from "@/assets/lib/report-event";
import DrugData from "@/assets/lib/drug-data";

export default {
  name: "ReportEventBox",
  computed: {
    DrugData() {
      return DrugData
    }
  },
  components: { DrugBox, TimestampText },
  props: {
    event: ReportEvent,
    showSection: Boolean
  },
  methods: {
    getSection(section) {
      switch (section) {
        case 2:
          return "Onset"
        case 3:
          return "Peak"
        case 4:
          return "Offset"
        default:
          return "Other"
      }
    },
    getSectionClass(section) {
      let textClass = "ReportEventBox__section_other";

      switch (section) {
        case 1:
          textClass = "ReportEventBox__section_description"
          break
        case 2:
          textClass = "ReportEventBox__section_onset"
          break
        case 3:
          textClass = "ReportEventBox__section_peak"
          break
        case 4:
          textClass = "ReportEventBox__section_offset"
          break
      }

      return `{ 'ReportEventBoxSectionText': true, '${textClass}': true`
    }
  }
}
</script>

<style scoped>
.ReportEventBoxSection {
    position: absolute;
    right: 0;
}

.ReportEventBoxSection {
    color: var(--tr-default-alt-lighter-text);
    transform: translateX(-0.5em) translateY(-0.75em);
    /*TODO: transform: rotate(90deg);*/
    text-transform: uppercase;
    font-size: 12px;
}

.ReportEventBoxTime {
    color: var(--tr-default-accent-text);
    font-size: 12px;
    margin-bottom: 2px;
}

.ReportEventBox__section_onset {
    color: rgb(221, 255, 221)
}
</style>
<template>
  <div :class="getSectionClass(event.section)">
    <div class="ReportEventBoxSection">
      <div class="ReportEventBoxSectionText">{{ showSection ? getSection(event.section) : "" }}</div>
    </div>

    <div v-if="event.timestamp !== '0001-01-01T00:00:00Z'" class="ReportEventBoxTime">
      <timestamp-text :data="event.timestamp" :hide-date="true"/>
    </div>
    <div v-if="event.type === 1">
      {{ event.content }}
    </div>
    <div v-else>
      <drug-box :drug="event.drug"/>
    </div>
  </div>
</template>

<script>
import TimestampText from "@/components/TimestampText.vue";
import DrugBox from "@/components/DrugBox.vue";

export default {
  name: "ReportEventBox",
  components: {DrugBox, TimestampText},
  props: {
    event: undefined,
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
      switch (section) {
        case 1:
          return "ReportEventBox__section_description"
        case 2:
          return "ReportEventBox__section_onset"
        case 3:
          return "ReportEventBox__section_peak"
        case 4:
          return "ReportEventBox__section_offset"
        default:
          return "ReportEventBox__section_other"
      }
    },
  }
}
</script>

<style scoped>
.ReportEventBoxSection {
  position: absolute;
  right: 0;
}

.ReportEventBoxSectionText {
  color: #ccc;
  transform: translateX(-0.5em) translateY(-0.75em);
  /*TODO: transform: rotate(90deg);*/
  text-transform: uppercase;
  font-size: 12px;
}

.ReportEventBoxTime {
  color: #8a8a8a;
  font-size: 12px;
  margin-bottom: 2px;
}
</style>
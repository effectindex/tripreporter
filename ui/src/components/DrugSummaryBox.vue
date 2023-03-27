<template>
  <header-column-box
      header="Substances"
      icon="pills"
      :columns="['Name', 'Dosage', 'RoA']"
      :rows="getRows()"
  />
</template>

<script>
import {getDose, getRoA} from "@/assets/lib/drugs";
import HeaderColumnBox from "@/components/HeaderColumnBox.vue";

export default {
  name: "DrugSummaryBox",
  components: {HeaderColumnBox},
  methods: {
    getDose, getRoA,
    getRows() {
      let rows = []
      this.events.forEach(e => {
        if (e.type === 2) {
          rows.push({'Name': e.drug.name, 'Dosage': getDose(e.drug), 'RoA': getRoA(e.drug)})
        }
      })
      return rows
    }
  },
  props: {
    events: undefined
  },
}
</script>

<style scoped>

</style>

<template>
  <header-column-box
      header="Substances"
      icon="pills"
      :columns="['Name', 'Dosage', 'RoA']"
      :rows="getRows()"
  />
</template>

<script>
import HeaderColumnBox from "@/components/HeaderColumnBox.vue";
import DrugData from "@/assets/lib/drug-data";

export default {
  name: "DrugSummaryBox",
  components: { HeaderColumnBox },
  methods: {
    getRows() {
      let rows = []
      this.events.forEach(e => {
        const drug = new DrugData({ obj: e.drug })

        if (e['type'] === 2) {
          rows.push({ 'Name': drug.name, 'Dosage': drug.getDose(), 'RoA': drug.getRoA() })
        }
      })
      return rows
    }
  },
  props: {
    // TODO: Switch to proper shared type that reflects the Go variation / #106
    events: Array
  },
}
</script>

<style scoped>

</style>

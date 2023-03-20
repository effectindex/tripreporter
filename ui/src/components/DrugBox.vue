<template>
  <div class="DrugBox">
    <img class="DrugBox__pill" src="../assets/svg/pills.svg" alt="Pills icon" width="32" height="32">
    <span class="DrugBox__text">
      {{ getDose(drug) }}<span v-if="getDose(drug)"> of </span>{{ drug.name }}{{ getRoAJoined(drug) }}
    </span>
  </div>
</template>

<script>
export default {
  name: "DrugBox",
  props: {
    drug: {
      name: String,
      dosage: Number,
      dosage_unit: String,
      roa: Number,
      frequency: Number,
      Prescribed: Number
    }
  },
  methods: {
    getRoA(drug) {
      switch (drug.roa) {
        case 1:
          return "Other"
        case 2:
          return "Oral"
        case 3:
          return "Buccal"
        case 4:
          return "Rectal"
        case 5:
          return "Inhaled"
        case 6:
          return "Sublabial"
        case 7:
          return "Intranasal"
        case 8:
          return "Sublingual"
        case 9:
          return "Injection"
        case 10:
          return "Buccal Injection"
        case 11:
          return "Intravenous Injection"
        case 12:
          return "Subcutaneous Injection"
        case 13:
          return "Intramuscular Injection"
        default:
          return ""
      }
    },
    getRoAJoined(drug) {
      const roa = this.getRoA(drug);
      if (roa) {
        return `, ${roa}`
      }

      return roa
    },
    getPrescribed(drug) {
      switch (drug.prescribed) {
        case 1:
          return "Over the counter"
        case 2:
          return "Prescribed by a doctor"
      }
    },
    getDose(drug) {
      if (drug.dosage === 0 && drug.dosage_unit === "") {
        return ""
      }

      if (drug.dosage === 0) {
        return drug.dosage_unit
      }

      return `${drug.dosage}${drug.dosage_unit}`
    }
  }
}
</script>

<style scoped>
.DrugBox__pill {
  vertical-align: middle;
}

.DrugBox__text {
  font-weight: bold;
  padding-left: 0.5em;
}
</style>
<template>
  <div class="DrugBox">
    <img class="DrugBox__pill" src="../assets/svg/pills.svg" alt="Pills icon" width="32" height="32">
    <span class="DrugBox__text">
      {{ getDose(drug) }}<span v-if="getDose(drug)"> of </span>{{ drug.name }}{{ getRoAJoined(drug) }}
    </span>
  </div>
</template>

<script>
import {getDose, getRoAJoined} from "@/assets/lib/drugs";

export default {
  name: "DrugBox",
  methods: {getRoAJoined, getDose},
  props: {
    drug: {
      name: String,
      dosage: Number,
      dosage_unit: String,
      roa: Number,
      frequency: Number,
      Prescribed: Number
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
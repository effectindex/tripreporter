<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <div class="DrugBox">
    <img class="DrugBox__pill" src="../assets/svg/pills.svg" alt="Pills icon" width="32" height="32">
    <span class="DrugBox__text">
      {{ drug.getDose() }}<span v-if="drug.getDose()"> of </span>{{ drug.name }}{{ drug.getRoA() ? `, ${drug.getRoA()}` : '' }}
    </span>
  </div>
</template>

<script>
import DrugData from "@/assets/lib/drug-data";

export default {
  name: "DrugBox",
  props: {
    drug: DrugData
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
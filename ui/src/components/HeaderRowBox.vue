<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <div class="HeaderRowBox">
    <section class="HeaderRowBox__header_wrapper">
      <h3 class="HeaderRowBox__header">
        <!-- TODO: Make icon dynamic -->
        <img v-if="icon === 'pills'" class="HeaderRowBox__header_icon" src="../assets/svg/pills.svg" alt="Pills icon"
            width="20" height="20">
        <img v-else-if="icon === 'user'" class="HeaderRowBox__header_icon" src="../assets/svg/user.svg" alt="User icon"
            width="20" height="20">
        <img v-else class="HeaderRowBox__header_icon" src="../assets/svg/report.svg" alt="Report icon"
            width="20" height="20">
        <span class="HeaderRowBox__header_text">
        {{ header }}
      </span>
      </h3>
    </section>
    <div v-if="columns" class="HeaderRowBox__row_wrapper">
      <!--suppress JSUnusedLocalSymbols -->
      <div v-for="(label, index) in columns.filter((label) => rows[label])" :key="index" class="HeaderRowBox__row">
        <div class="HeaderRowBox__row_label">
          {{ label }}
        </div>
        <div class="HeaderRowBox__row_entry">
          <router-link v-if="links && links[label]" :to="links[label]" class="--tr-link">{{ rows[label] }}</router-link>
          <div v-else>{{ rows[label] }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "HeaderRowBox",
  props: {
    icon: String,
    header: String,
    columns: Array,
    rows: undefined,
    links: undefined
  },
}
</script>

<style scoped>
.HeaderRowBox {
    flex: 1;
    background-color: var(--tr-background-raised-light);
    border-radius: 10px;
    border: 1px solid var(--tr-border);
}

.HeaderRowBox__header_wrapper h3 {
    margin: 0;
}

.HeaderRowBox__header {
    border-top-left-radius: 10px;
    border-top-right-radius: 10px;
    background-color: var(--tr-background-raised-dark);
    padding: 0.25em 0.5em 0.5em;
}

.HeaderRowBox__header_icon {
    vertical-align: bottom;
    margin-bottom: 3px;
}

.HeaderRowBox__header_text {
    font-weight: bold;
    padding-left: 0.5em;
}

.HeaderRowBox__row_wrapper {
    padding: 0.25em;
}

.HeaderRowBox__row {
    padding: 5px 10px;
    display: flex;
    flex-direction: row;
}

.HeaderRowBox__row_label {
    flex: 1em;
    color: var(--tr-text-alt-dark);
    font-weight: 700;
}

.HeaderRowBox__row_entry {
    flex: 5;
}

/* override HeaderRowBox__row_label for desktop browsers */
@media only screen and (min-width: 680px) {
    .HeaderRowBox__row_label {
        flex: 2;
    }
}
</style>
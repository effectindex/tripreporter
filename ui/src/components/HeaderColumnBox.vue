<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <div class="HeaderColumnBox">
    <section class="HeaderColumnBox__header_wrapper">
      <h3 class="HeaderColumnBox__header">
        <!-- TODO: Make icon dynamic -->
        <img v-if="icon === 'pills'" class="HeaderColumnBox__header_icon" src="../assets/svg/pills.svg" alt="Pills icon"
            width="20" height="20">
        <img v-else-if="icon === 'user'" class="HeaderRowBox__header_icon" src="../assets/svg/user.svg" alt="User icon"
            width="20" height="20">
        <img v-else class="HeaderRowBox__header_icon" src="../assets/svg/report.svg" alt="Report icon"
            width="20" height="20">
        <span class="HeaderColumnBox__header_text">
        {{ header }}
      </span>
      </h3>
    </section>
    <div class="HeaderColumnBox__table_wrapper">
      <table class="HeaderColumnBox__table">
        <thead v-if="columns" class="HeaderColumnBox__table_columns">
        <tr>
          <td v-for="(label, index) in columns" :key="index">
            {{ label }}
          </td>
        </tr>
        </thead>
        <tbody v-if="rows" class="HeaderColumnBox__table_rows">
        <tr v-for="(row, index) in rows" :key="index"
            :class="index % 2 === 0 ? 'HeaderColumnBox__table_row' : 'HeaderColumnBox__table_row_alt'">
          <td v-for="(column, index1) in columns" :key="index1">
            <router-link v-if="links && links[index] && links[index][column]" :to="links[index][column]" class="--tr-link">{{ row[column] }}</router-link>
            <div v-else>{{ row[column] }}</div>
          </td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: "HeaderColumnBox",
  props: {
    icon: String,
    header: String,
    columns: Array,
    rows: undefined,
    links: undefined
  },
  data() {
    return {
      cssProps: {
        'mask-image': `url(../assets/svg/pills.svg)`
      }
    }
  },
  computed: {
    imgPath() {
      return require(`${this.icon}`);
    }
  }
}
</script>

<style scoped>
.HeaderColumnBox {
    flex: 1;
    background-color: var(--tr-background-raised-light);
    border-radius: 10px;
    border: 1px solid var(--tr-border);
}

.HeaderColumnBox__header_wrapper h3 {
    margin: 0;
}

.HeaderColumnBox__header {
    border-top-left-radius: 10px;
    border-top-right-radius: 10px;
    background-color: var(--tr-background-raised-dark);
    padding: 0.25em 0.5em 0.5em;
}

.HeaderColumnBox__header_icon {
    vertical-align: bottom;
    margin-bottom: 3px;
}

.HeaderColumnBox__header_text {
    font-weight: bold;
    padding-left: 0.5em;
}

.HeaderColumnBox__table_wrapper {
    padding: 0.75em;
    border-collapse: collapse;
}

.HeaderColumnBox__table {
    width: 100%;
    border-collapse: collapse;
}

.HeaderColumnBox__table_columns {
    font-weight: bold;
    padding: 5px 10px;
}

.HeaderColumnBox__table_rows {
    padding: 5px 10px;
}

.HeaderColumnBox__table_row td, .HeaderColumnBox__table_row_alt td {
    padding: 3px 0;
}
</style>
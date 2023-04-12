<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <span>{{ timestamp ? timestamp.get() : new Timestamp({ date: date, showTime: showTime, hideDate: hideDate, longFormat: longFormat }).get() }}</span>
</template>

<script>
import Timestamp from "../assets/lib/timestamp";

export default {
  name: "TimestampText",
  computed: {
    Timestamp() {
      return Timestamp
    }
  },
  props: {
    timestamp: Timestamp,
    date: String,
    showTime: Boolean,
    hideDate: Boolean,
    longFormat: Boolean
  }
}
</script>

<style scoped>

</style>

<template>
  <span>{{ new Timestamp({date: date, showTime: showTime, hideDate: hideDate, longFormat: longFormat}).get() }}</span>
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
    date: String,
    showTime: Boolean,
    hideDate: Boolean,
    longFormat: Boolean
  }
}
</script>

<style scoped>

</style>

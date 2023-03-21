<template>
  <span>{{ getFormattedTimestamp({data, showTime, hideDate, longFormat}) }}</span>
</template>

<script>
import getFormattedTimestamp from "../assets/lib/timestamp";

export default {
  name: "TimestampText",
  methods: {getFormattedTimestamp},
  props: {
    data: String,
    showTime: Boolean,
    hideDate: Boolean,
    longFormat: Boolean
  }
}
</script>

<style scoped>

</style>
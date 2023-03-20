<template>
  <span>{{ getFormatted() }}</span>
</template>

<script>
export default {
  name: "TimestampText",
  props: {
    data: String,
    showTime: Boolean,
    hideDate: Boolean,
    longFormat: Boolean
  },
  methods: {
    getFormatted() {
      const date = new Date(this.data);

      // Golang zero time
      if (date.getTime() === -62135596800000) {
        return ""
      }

      let options = this.longFormat ? {weekday: "long", year: "numeric", month: "long", day: "numeric"} : {year: "numeric", month: "numeric", day: "numeric"};

      if (this.showTime) {
        options.hour = "numeric";
        options.minute = "numeric";
      }

      if (this.hideDate) {
        return date.toLocaleTimeString(undefined, { hour: "2-digit", minute: "2-digit" });
      }

      return date.toLocaleString(undefined, options);
    }
  }
}
</script>

<style scoped>

</style>
// SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

// A formatted Date with easy to use options.
export default class Timestamp { // TODO: Rewrite in TS for #106
  constructor({ date, showTime, hideDate, longFormat }) {
    this.date = new Date(date);
    this.showTime = showTime;
    this.hideDate = hideDate;
    this.longFormat = longFormat;
    return this;
  }

  raw() { // Return the raw Timestamp data
    return this.date.toJSON()
  }

  epoch() {
    return -62135596800000
  }

  get() { // Return the formatted Timestamp based on the provided options
    if (this.date === undefined) {
      return ""
    }

    // Golang zero time in Unix epoch milliseconds
    if (this.date.getTime() === this.epoch()) {
      return ""
    }

    let options = this.longFormat ? { weekday: "long", year: "numeric", month: "long", day: "numeric" } : {
      year: "numeric",
      month: "numeric",
      day: "numeric"
    };

    if (this.showTime) {
      options.hour = "numeric";
      options.minute = "numeric";
    }

    if (this.hideDate) {
      return this.date.toLocaleTimeString(undefined, { hour: "2-digit", minute: "2-digit" });
    }

    return this.date.toLocaleString(undefined, options);
  }
}

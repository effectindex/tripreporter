// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

export default class ReportSubject { // TODO: Rewrite in TS for #106
  constructor({ obj, age, gender, display_unit, height_cm, weight_kg }) {
    this.age = age;
    this.gender = gender;
    this.display_unit = display_unit;
    this.height_cm = height_cm;
    this.weight_kg = weight_kg;
    obj && Object.assign(this, obj);
    return this;
  }

  height() {
    if (this.display_unit !== 2) {
      return `${this.display_unit.toFixed(1)}cm`
    }

    let inches = (this.height_cm * 0.393700787).toFixed(0)
    const feet = Math.floor(inches / 12);
    inches %= 12;

    return `${feet}'${inches}"`
  }

  weight() {
    if (this.display_unit !== 2) {
      return `${this.weight_kg}kg`
    }

    return `${+(this.weight_kg * 2.205).toFixed(1)}lbs`
  }
}
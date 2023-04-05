// SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

export default class DrugData { // TODO: Rewrite in TS for #106
  constructor({ obj, name, dosage, dosage_unit, roa, frequency, prescribed }) {
    this.name = name;
    this.dosage = dosage;
    this.dosage_unit = dosage_unit;
    this.roa = roa;
    this.frequency = frequency;
    this.prescribed = prescribed;

    this.getRoA = () => {
      switch (this.roa) {
        case 1:
          return "Other"
        case 2:
          return "Oral"
        case 3:
          return "Buccal"
        case 4:
          return "Rectal"
        case 5:
          return "Inhaled"
        case 6:
          return "Sublabial"
        case 7:
          return "Intranasal"
        case 8:
          return "Sublingual"
        case 9:
          return "Injection"
        case 10:
          return "Buccal Injection"
        case 11:
          return "Intravenous Injection"
        case 12:
          return "Subcutaneous Injection"
        case 13:
          return "Intramuscular Injection"
        default:
          return undefined
      }
    }

    this.getPrescribed = () => {
      switch (this.prescribed) {
        case 1:
          return "Over the counter"
        case 2:
          return "Not prescribed by a doctor"
        case 3:
          return "Prescribed by a doctor"
        default:
          return undefined
      }
    }

    this.getDose = () => {
      if (this.dosage === 0) {
        return this.dosage_unit !== "" ? `${this.dosage_unit}` : undefined
      }

      let joiner = " "
      if (this.dosage_unit.length < 5) {
        joiner = ""
      }

      return `${this.dosage}${joiner}${this.dosage_unit}`
    }

    obj && Object.assign(this, obj);
    return this;
  }
}

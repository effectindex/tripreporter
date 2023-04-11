<template>
  <FormKit
      type="text"
      id="drug_name"
      name="drug_name"
      :label="`${labelPrefix} name`"
      validation="length:0,4096"
      validation-visibility="live"
      :placeholder="placeholderName"
      help="(optional) Leave blank if unknown."
  />
  <FormKit
      type="text"
      id="drug_dosage"
      name="drug_dosage"
      :label="`${labelPrefix} dosage`"
      validation="length:0,4096"
      validation-visibility="live"
      :placeholder="placeholderDosage"
      help="(optional) Leave blank if unknown."
  />
  <FormKit
      type="select"
      :label="`${labelPrefix} Route of Administration`"
      name="roa"
      id="roa"
      :options="optionsRoa"
      :placeholder="placeholderRoa"
      help="(optional)"
  />
  <FormKit
      type="select"
      :label="`${labelPrefix} OTC / Prescription`"
      name="prescribed"
      id="prescribed"
      :options="optionsPrescribed"
      :placeholder="placeholderPrescribed"
      help="(optional)"
  />
</template>

<script>
export default {
  name: "FormKitSubstance",
  props: {
    labelPrefix: String,
    placeholderName: String,
    placeholderDosage: String,
    placeholderRoa: String,
    placeholderPrescribed: String
  }
}
</script>

<script setup>
const optionsRoa = [
  { label: '', value: '0' },
  { label: 'Other', value: '1' },
  { label: 'Oral (swallowed)', value: '2' },
  { label: 'Sublingual (under tongue)', value: '8' },
  { label: 'Inhaled', value: '5' },
  { label: 'Intranasal (snorted)', value: '7' },
  { label: 'Rectal (boofed)', value: '4' },
  { label: 'Buccal (held in gums)', value: '3' },
  { label: 'Intravenous (IV injection)', value: '11' },
  { label: 'Intramuscular (injection into muscle)', value: '13' },
  { label: 'Intrabuccal (injection into gums)', value: '10' },
  { label: 'Subcutaneous (injection into fat)', value: '12' },
  { label: 'Sublabial (under the lip)', value: '6' },
  { label: 'Injection (other method)', value: '9' },
];
const optionsPrescribed = [
  { label: '', value: '0' },
  { label: 'It\'s over the counter', value: '1' },
  { label: 'It\'s prescribed by a doctor', value: '3' },
  { label: 'It\'s not prescribed by a doctor', value: '2' },
];
</script>

<style scoped>

</style>
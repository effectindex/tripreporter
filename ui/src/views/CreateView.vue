<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <AuthWrapper>
  <div class="create">
      <h1 class="--tr-header-h1">Create a Subjective Experience Report</h1>

      <div class="DefaultView__message" id="DefaultView__message">
        <div class="DefaultView__message_text" id="DefaultView__message_text"></div>
      </div>

      <div class="DefaultView__form_wide">
        <FormKit type="form" @submit="handleNext" #default="{ value, state: { errors } }" :actions="false">
          <FormKit type="multi-step" name="report_form" tab-style="progress" :hide-progress-labels="true" :allow-incomplete="false" :classes="{ wrapper: 'formkit-wrapper-wide' }">
            <FormKit type="step" name="report_info" v-model="createStore.reportInfo" :stepActions-class="{ 'formkit-justify-right': true }">
              <FormKit
                  :classes="{ wrapper: 'formkit-wrapper-wide' }"
                  type="text"
                  id="title"
                  name="title"
                  label="Report Title"
                  validation="required|length:0,4096"
                  placeholder="My subjective experience with LSD"
                  :help="getTextLength(value.title, 4096)"
                  :delay="60"
              />

              <FormKit
                  :classes="{ wrapper: 'formkit-wrapper-wide' }"
                  type="textarea"
                  id="setting"
                  name="setting"
                  label="Setting"
                  rows="5"
                  validation="length:0,4096"
                  placeholder="Describe the setting / place the experience started in."
                  :help="'(optional) ' + getTextLength(value.setting, 4096)"
                  :delay="60"
              />

              <FormKit
                  :classes="{ outer: 'formkit-outer-narrow' }"
                  type="date"
                  id="date"
                  name="report_date"
                  label="Date"
                  help="When did the experience occur?"
              />
              <!--suppress VueUnrecognizedSlot -->
              <template #stepNext="{ handlers, node }">
                <FormKit
                    :classes="{ input: { 'formkit-input-no-margin': true } }"
                    type="button"
                    @click="handlers.incrementStep(1, node.context)()"
                    label="Next"
                    data-next="true"
                />
              </template>
            </FormKit>
            <FormKit type="step" name="subject_info" v-model="createStore.reportSubject">
              <FormKit
                  type="number"
                  min="0"
                  id="subject_age"
                  name="subject_age"
                  label="Age"
                  help="(optional)"
                  v-model="createStore.subjectInfo.age"
              />

              <FormKit
                  type="taglist"
                  id="subject_gender"
                  name="subject_gender"
                  label="Gender"
                  :options="optionsGender"
                  :allow-new-values="true"
                  max="1"
                  help="(optional) Custom values are supported."
                  v-model="createStore.subjectInfo.gender"
              />

              <FormKit
                  type="toggle"
                  name="use_imperial"
                  label="Use imperial for units?"
                  v-model="createStore.useImperial"
              />

              <div v-show="!createStore.useImperial">
                <FormKit
                    type="number"
                    min="0"
                    step="0.5"
                    id="subject_height_cm"
                    name="subject_height_cm"
                    label="Height (cm)"
                    help="(optional)"
                    v-model="createStore.subjectInfo.heightCm"
                />
                <FormKit
                    type="number"
                    min="0"
                    step="0.5"
                    id="subject_weight_kg"
                    name="subject_weight_kg"
                    label="Weight (kg)"
                    help="(optional)"
                    v-model="createStore.subjectInfo.weightKg"
                />
              </div>
              <div v-show="createStore.useImperial">
                <FormKit
                    type="number"
                    min="0"
                    id="subject_height_ft"
                    name="subject_height_ft"
                    label="Height (ft)"
                    help="(optional)"
                    v-model="createStore.subjectInfo.heightFt"
                />
                <FormKit
                    type="number"
                    min="0"
                    step="0.5"
                    id="subject_height_in"
                    name="subject_height_in"
                    label="Height (in)"
                    help="(optional)"
                    v-model="createStore.subjectInfo.heightIn"
                />
                <FormKit
                    type="number"
                    min="0"
                    step="0.5"
                    id="subject_weight_lbs"
                    name="subject_weight_lbs"
                    label="Weight (lbs)"
                    help="(optional)"
                    v-model="createStore.subjectInfo.weightLbs"
                />
              </div>
              <!--suppress VueUnrecognizedSlot -->
              <template #stepNext="{ handlers, node }">
                <FormKit
                    :classes="{ input: { 'formkit-input-no-margin': true, 'formkit-input-justify-right': false } }"
                    type="button"
                    @click="handlers.incrementStep(1, node.context)()"
                    label="Next"
                    data-next="true"
                />
              </template>
            </FormKit>
            <FormKit type="step" name="medication_info" v-model="createStore.reportMedication">
              <FormKit
                  type="repeater"
                  id="subject_medications"
                  name="subject_medications"
                  label="Are you on any medications relevant to the report?"
                  help="(optional)"
                  :add-button="false"
                  :insert-control="true"
              >
                <FormKitDrug
                    label-prefix="Medication"
                    placeholder-name="Omeprazole"
                    placeholder-dosage="10mg"
                    placeholder-roa="How do you take this medication?"
                    placeholder-prescribed="Is this medication prescribed?"
                />
              </FormKit>
              <!--suppress VueUnrecognizedSlot -->
              <template #stepNext="{ handlers, node }">
                <FormKit
                    :classes="{ input: { 'formkit-input-no-margin': true } }"
                    type="button"
                    @click="handlers.incrementStep(1, node.context)()"
                    label="Next"
                    data-next="true"
                />
              </template>
            </FormKit>
            <FormKit type="step" name="report_events" v-model="createStore.reportEvents">
              <FormKit
                  type="repeater"
                  id="report_sections"
                  name="report_sections"
                  label="Add doses and content describing the experience, as timestamped sections."
                  help="(optional)"
                  :add-button="false"
                  :insert-control="true"
                  #default="{ value, index }"
              >
                <FormKit
                    type="time"
                    name="timestamp"
                    label="Time"
                    :help="'(optional) ' + getEventTimestampText(value)"
                />

                <FormKit
                    type="radio"
                    label="During what part of the experience is this?"
                    name="section"
                    id="section"
                    :value="getRadioDefault(createStore.reportEvents, index)"
                    :options="[
                      { label: 'Other', value: '1', help: 'This description is not during the experience itself.' },
                      { label: 'Onset', value: '2' },
                      { label: 'Peak', value: '3' },
                      { label: 'Offset', value: '4' }
                    ]"
                    help="(optional)"
                />

                <FormKit
                    type="toggle"
                    name="is_drug"
                    label="Adding substance dose?"
                    :value="index === 0"
                />

                <FormKit
                    v-if="!getEventType(value)"
                    :classes="{ wrapper: 'formkit-wrapper-wide' }"
                    type="textarea"
                    id="content"
                    name="content"
                    label="Description"
                    rows="5"
                    placeholder="Describe this part of the subjective experience."
                    :help="'(optional) ' + getTextLength(value ? value.content : '', 10485760)"
                    :delay="60"
                />
                <div v-else>
                  <FormKitDrug
                      label-prefix="Substance"
                      placeholder-name="LSD"
                      placeholder-dosage="100μg"
                      placeholder-roa="How did you take the substance?"
                      placeholder-prescribed="Is this substance prescribed?"
                  />
                </div>
              </FormKit>

              <!--suppress VueUnrecognizedSlot -->
              <template #stepNext>
                <FormKit
                    :classes="{ input: { 'formkit-input-no-margin': true } }"
                    type="submit"
                    @submit="submitForm"
                    label="Create Report!"
                    data-next="true"
                    :disabled="errors && submitting"
                />
              </template>
            </FormKit>
          </FormKit>
        </FormKit>
      </div>
    </div>
  </AuthWrapper>
</template>

<script>
import LayoutDefault from "@/layouts/LayoutDefault.vue";

export default {
  name: "CreateView",
  created() {
    this.$emit('update:layout', LayoutDefault);
  },
  methods: {
    getEventTimestampText(value) {
      if (this.getEventType(value)) {
        return "What time was this substance dosed?"
      }

      return "What time did this description occur?"
    },
    getEventType(value) {
      return value.is_drug === true;
    },
    getRadioDefault(value, index) {
      if (!value.report_sections || index === 0) {
        return '0'
      }
      return value.report_sections[index - 1] ? value.report_sections[index - 1].section : '0'
    }
  }
}
</script>

<script setup>
import { inject, ref } from "vue";
import { useCreateStore } from "@/assets/lib/createstore";
import { handleMessageError, setMessage } from "@/assets/lib/message_util";
import { getTextLength } from "@/assets/lib/form";
import log from "@/assets/lib/logger";
import FormKitDrug from "@/components/FormKitDrug.vue";
import AuthWrapper from "@/components/AuthWrapper.vue";

const router = inject('router')
const axios = inject('axios')
const createStore = useCreateStore();

const optionsGender = ["Male", "Female", "Nonbinary"];

const messageSuccess = "Successfully created report!";
let success = ref(false);
let submitting = ref(false);

const submitForm = async (fields) => {
  submitting.value = true;

  // Transform form steps into a more sane data format, for the server to handle
  fields = fields.report_form;
  fields = {
    ...fields.report_info,
    ...fields.report_events,
    "report_subject": {
      ...fields.subject_info,
      "medications": {
        ...fields.medication_info.subject_medications
      }
    }
  };

  log("submitForm", fields)
  await axios.post('/report', fields).then(function (response) {
    success.value = response.status === 201;
    submitting.value = false;
    setMessage(response.data.msg, messageSuccess, success.value, router, `/report/${response.data.id}`);
  }).catch(function (error) {
    success.value = error.response.status === 201;
    submitting.value = false;
    setMessage(error.response.data.msg, messageSuccess, success.value, router, `/report/${error.response.data.id}`);
    handleMessageError(error);
  })
}

// TODO: Workaround for https://github.com/formkit/formkit/issues/641
const handleNext = async (fields, handlers) => {
  let incrementedStep = false;
  let focusedStep = false;

  const activateStep = async(child, step, nextStep) => {
    if (child.name === step && child.context.isActiveStep && !incrementedStep) {
      if (nextStep === "submit_form") {
        return submitForm(fields)
      }

      incrementedStep = true
      child.context.handlers.next()
    }

    if (child.name === nextStep && incrementedStep && !focusedStep) {
      focusedStep = true

      let elementId = child.context.node.children[0].props.id
      if (child.context.node.children[0].children.length > 0 && child.context.node.children[0].children[0].children.length > 0) {
        elementId = child.context.node.children[0].children[0].children[0].props.id
      }

      setTimeout(function () {
        const el = document.getElementById(elementId)
        el.focus()
        el.click()
      }, 20)
    }
  }

  // report_info, subject_info, medication_info, report_events
  handlers.children[0].walk(child => {
    activateStep(child, "report_info", "subject_info")
    activateStep(child, "subject_info", "medication_info")
    activateStep(child, "medication_info", "report_events")
    activateStep(child, "report_events", "submit_form")
  })
}
</script>

<style>
@import url(@/assets/css/forms.css);
@import url(@/assets/css/forms-repeater.css);
@import url(@/assets/css/forms-multi-step.css);
</style>

<style scoped>
@import url(@/assets/css/message_util.css);
</style>

<!--
SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>

SPDX-License-Identifier: OSL-3.0
-->

<template>
  <div class="create">
    <div v-if="!store.activeSession" class="no-session">
      <div v-if="store.updatedPreviously">
        <not-found/>
      </div>
    </div>
    <div v-else>
      <h1 class="--tr-header-h1">Create a Subjective Experience Report</h1>

      <div class="DefaultView__message" id="DefaultView__message">
        <div class="DefaultView__message_text" id="DefaultView__message_text"></div>
      </div>

      <div class="DefaultView__form_wide">
        <FormKit type="form" :actions="false" @submit="submitForm" #default="{ value }">

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
              type="date"
              id="date"
              name="report_date"
              label="Date"
              validation="required"
              help="When did the experience occur?"
          />

          <FormKit
              id="repeater"
              name="report_sections"
              type="repeater"
              label="Report Content"
              :add-button="false"
              :insert-control="true"
              #default="{ index }"
          >
            <FormKit
                type="time"
                name="timestamp"
                label="Time"
                :help="'(optional) ' + getEventTimestampText(value, index)"
            />

            <FormKit
                type="radio"
                label="During what part of the experience is this?"
                name="section"
                id="section"
                :value="getRadioDefault(value, index)"
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
            />

            <FormKit
                v-if="!getEventType(value, index)"
                :classes="{ wrapper: 'formkit-wrapper-wide' }"
                type="textarea"
                id="content"
                name="content"
                label="Description"
                rows="5"
                placeholder="Describe this part of the subjective experience."
                :help="'(optional) ' + getTextLength(value.report_sections ? value.report_sections[index] ? value.report_sections[index].content : '' : '', 10485760)"
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
          <!--<pre wrap>{{ value }}</pre>-->

          <div class="DefaultView__inline_box">
            <div class="DefaultView__inline">
              <FormKit
                  type="submit"
                  @submit="submitForm"
              >
                Create Report!
              </FormKit>
            </div>
          </div>
        </FormKit>

      </div>
    </div>
  </div>
</template>

<script>
import LayoutDefault from "@/layouts/LayoutDefault.vue";

export default {
  name: "CreateView",
  created() {
    this.$emit('update:layout', LayoutDefault);
  },
  methods: {
    getEventTimestampText(value, index) {
      const event = this.getEventType(value, index)
      if (event) {
        return "What time was this substance dosed?"
      }

      return "What time did this description occur?"
    },
    getEventType(value, index) {
      if (!value.report_sections) {
        return false
      }
      const event = value.report_sections[index] ? value.report_sections[index].is_drug : false
      return event === true;
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
import { inject } from "vue";
import { useSessionStore } from '@/assets/lib/sessionstore'
import NotFound from "@/views/NotFound.vue";
import { handleMessageError, setMessage } from "@/assets/lib/message_util";
import { getTextLength } from "@/assets/lib/form";
import FormKitDrug from "@/components/FormKitDrug.vue";

const router = inject('router')
const axios = inject('axios')
const store = useSessionStore();

const messageSuccess = "Successfully created report!";
let success = false;

const submitForm = async (fields) => {
  // don't do anything if the user presses the button again, for example, while waiting for a redirect
  if (success) {
    return
  }

  await axios.post('/report', fields).then(function (response) {
    success = response.status === 201;
    setMessage(response.data.msg, messageSuccess, success, router, `/reports?id=${response.data.id}`);
  }).catch(function (error) {
    success = error.response.status === 201;
    setMessage(error.response.data.msg, messageSuccess, success, router, `/reports?id=${error.response.data.id}`);
    handleMessageError(error);
  })
}
</script>

<style>
@import url(@/assets/css/forms.css);
</style>

<style scoped>
@import url(@/assets/css/message_util.css);
</style>

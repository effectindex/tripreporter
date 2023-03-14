<template>
  <div class="create">
    <div v-if="!store.activeSession" class="no-session">
      <div v-if="store.updatedPreviously">
        <not-found/>
      </div>
    </div>
    <div v-else>
      <h1>Create a Subjective Experience Report</h1>

      <div class="DefaultView__message" id="DefaultView__message">
        <div class="DefaultView__message_text" id="DefaultView__message_text"></div>
      </div>

      <div class="DefaultView__form_wide">
        <!--        :submit-attrs="{inputClass: 'formkit-input-hidden'}"-->
        <FormKit type="form" :actions="false" @submit="submitForm" #default="{ value }">

          <FormKit
              :classes="{ outer: createStore.submitClass, wrapper: 'formkit-wrapper-wide' }"
              type="text"
              id="title"
              name="title"
              label="Report Title"
              validation="required|length:0,4096"
              placeholder="My subjective experience with LSD"
          />

          <FormKit
              :classes="{ outer: createStore.submitClass, wrapper: 'formkit-wrapper-wide' }"
              v-show="createStore.page > 0"
              type="textarea"
              id="setting"
              name="setting"
              label="Setting"
              rows="5"
              validation="length:0,4096"
              validation-visibility="live"
              placeholder="Briefly describe the setting / place the experience started in."
              help="(optional)"
          />

          <FormKit
              :classes="{ outer: createStore.submitClass }"
              v-show="createStore.page > 0"
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
                :placeholder="getSectionPlaceholder(value, index)"
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
                :classes="{ outer: createStore.submitClass, wrapper: 'formkit-wrapper-wide' }"
                type="textarea"
                id="content"
                name="content"
                label="Description"
                rows="5"
                placeholder="Describe this part of the subjective experience."
                help="(optional)"
            />
            <!-- TODO: Refactor substance dosing into separate component -->
            <div v-else>
              <FormKit
                  :classes="{ outer: createStore.submitClass }"
                  type="text"
                  id="drug_name"
                  name="drug_name"
                  label="Substance name"
                  validation="length:0,4096"
                  placeholder="LSD"
                  help="(optional) Leave blank if unknown."
              />
              <FormKit
                  :classes="{ outer: createStore.submitClass }"
                  type="text"
                  id="drug_dosage"
                  name="drug_dosage"
                  label="Substance dosage"
                  validation="length:0,4096"
                  placeholder="100μg"
                  help="(optional) Leave blank if unknown."
              />
              <FormKit
                  type="select"
                  label="Substance Route of Administration"
                  name="roa"
                  id="roa"
                  :options="[
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
                  ]"
                  placeholder="How did you take the substance?"
                  help="(optional)"
              />
              <FormKit
                  type="select"
                  label="Substance OTC / Prescription"
                  name="prescribed"
                  id="prescribed"
                  :options="[
                    { label: '', value: '0' },
                    { label: 'It\'s over the counter', value: '1' },
                    { label: 'It\'s prescribed by a doctor', value: '3' },
                    { label: 'It\'s not prescribed by a doctor', value: '2' },
                  ]"
                  placeholder="Is this substance prescribed?"
                  help="(optional)"
              />
            </div>
          </FormKit>
          <pre wrap>{{ value }}</pre>
          <pre wrap>{{ events }}</pre>
          <!--                    <FormKit-->
          <!--                        type="submit"-->
          <!--                        label="Register"-->
          <!--                    />-->
          <div class="DefaultView__inline_box">
            <div class="DefaultView__inline">
              <FormKit
                  type="button"
                  @click="updateBackProgress"
                  :disabled="createStore.page === 0"
              >
                Back
              </FormKit>
              <FormKit
                  type="button"
                  @click="updateNextProgress"
              >
                Next
              </FormKit>
              <FormKit
                  type="submit"
                  @submit="submitForm"
                  :submit-attrs="{inputClass: 'formkit-input-hidden'}"
              >
                SubmitReal
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
    getSectionPlaceholder(value, index) {
      console.log(`HERE: ${index}`)
      if (!value.report_sections) {
        return ""
      }
    }
  }
}
</script>

<script setup>
import {ref} from 'vue'
import {useSessionStore} from '@/assets/lib/sessionstore'
import {useCreateStore} from "@/assets/lib/createstore";
import NotFound from "@/views/NotFound.vue";

// const event = ref([])
const events = ref([])
let report = ref({})

const store = useSessionStore();
const createStore = useCreateStore();

const submitForm = async (fields) => {
  // Let's pretend this is an ajax request:
  // await new Promise((r) => setTimeout(r, 1000))
  // createStore.page += 1
  report.value = fields;
  console.log(report.value.title)
}

createStore.page = 0;

const updateBackProgress = async (fields) => {
  createStore.page -= 1
  if (createStore.page === 0) {
    createStore.submitClass = 'formkit-outer'
  }
}

const updateNextProgress = async (fields) => {
  createStore.page += 1
  createStore.submitClass = 'formkit-outer-hidden'
}
</script>

<style>
@import url(@/assets/css/forms.css);
</style>

<style scoped>
@import url(@/assets/css/message_util.css);
</style>

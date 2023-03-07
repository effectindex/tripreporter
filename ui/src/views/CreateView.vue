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
              v-model="events"
              type="list"
              name="events"
          >
            <FormKit
                type="select"
                label="What type of report section is this?"
                name="select"
                id="select"
                :options="[
                    { label: 'Description', value: '1' },
                    { label: 'Substance Dose', value: '2' },
                  ]"
                validation="required"
            />

            <FormKit
                type="time"
                name="timestamp"
                label="Time"
                :help="getEventTimestampText(value.events ? value.events[0] : '0')"
            />
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
    getEventTimestampText(value) {
      if (value === "1") {
        return "What time did this description occur?"
      }
      if (value === "2") {
        return "What time was this substance dosed?"
      }

      return "Select a report section type!"
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

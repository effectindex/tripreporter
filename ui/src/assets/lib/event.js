// SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import { ref } from "vue";

const bus = ref(new Map());

export default function useEventsBus() {
  function emit(event, ...args) {
    bus.value.set(event, args);
  }

  return {
    emit,
    bus
  }
}

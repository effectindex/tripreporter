// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import log from "@/assets/lib/logger";

export function setMessage(message, messageSuccess, status, router, location, routerTimeout) {
  if (!message && !messageSuccess) {
    log("message_util: message and messageSuccess are null!", arguments)
    return
  }

  const elemText = document.getElementById("DefaultView__message_text");
  if (elemText === null) {
    log("message_util: elemText is null!", arguments)
    return
  }

  elemText.textContent = message;

  const elem = document.getElementById("DefaultView__message")
  elem.style.display = 'block';

  if (status === true) {
    elemText.style.background = 'var(--tr-accent)'
    elemText.innerHTML = messageSuccess;

    if (router && location) {
      const timeout = routerTimeout ? routerTimeout : 0
      window.setTimeout(function () {
        router.push(location)
      }, timeout);
    }
  } else {
    elemText.style.background = 'var(--tr-error)'
  }
}

export function clearMessage(expectStatus, status, clearTimeout) {
  if (expectStatus !== status) {
    return
  }

  const elem = document.getElementById("DefaultView__message")
  if (elem === null) {
    log("message_util: elem is null!")
    return
  }

  const timeout = clearTimeout ? clearTimeout : 0
  window.setTimeout(function () {
    elem.style.display = 'none';
  }, timeout);
}

export function handleMessageError(error) {
  if (error.response) {
    log(error.response)
  } else if (error.request) {
    log(error.request);
  } else {
    log("Unexpected error when handling message", error);
  }
}

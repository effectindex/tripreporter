// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

import log from "@/assets/lib/logger";

export function setMessage(message, messageSuccess, status, router, location, routerTimeout) {
  if (!message && !messageSuccess) {
    log("message_util: message and messageSuccess are null!", arguments)
    return
  }

  const timeout = routerTimeout ? routerTimeout : 0
  const elemText = document.getElementById("DefaultView__message_text");
  if (elemText === null) {
    log("message_util: elemText is null!", arguments)
    return
  }

  elemText.textContent = message;

  const elem = document.getElementById("DefaultView__message")
  // Show message if it's an error, or a success that won't redirect or redirect immediately
  if (status !== true || (!location || (location && timeout !== 0))) {
    elem.style.display = 'block';
  }

  if (status === true) {
    elemText.style.background = 'var(--tr-accent)'
    elemText.innerHTML = messageSuccess;

    if (router && location) { // Message is cleared by the router on redirect
      window.setTimeout(function () {
        router.push(location)
      }, timeout);
    } else { // Clear message if we're not redirecting elsewhere
      elem.style.display = 'none';
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

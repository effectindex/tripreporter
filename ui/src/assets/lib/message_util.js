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
        elem.style.background = 'var(--tr-default-primary)'
        elemText.innerHTML = messageSuccess;

        if (router && location) {
            const timeout = routerTimeout ? routerTimeout : 0
            window.setTimeout(function () {
                router.push(location)
            }, timeout);
        }
    } else {
        elem.style.background = 'var(--tr-default-message-error)'
    }
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

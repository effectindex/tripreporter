export function setMessage(message, messageSuccess, status, location) {
    const elemText = document.getElementById("DefaultView__message_text");
    elemText.textContent = message;

    const elem = document.getElementById("DefaultView__message")
    elem.style.display = 'block';

    if (status === true) {
        elem.style.background = 'var(--tr-default-primary)'
        elemText.innerHTML = messageSuccess;

        if (location) {
            window.setTimeout(function () {
                window.location.href = location; // faster than window.location.replace()
            }, 3000);
        }
    } else {
        elem.style.background = 'var(--tr-default-message-error)'
    }
}

export function handleMessageError(error) {
    if (error.response) {
        console.log(error.response)
    } else if (error.request) {
        console.log(error.request);
    } else {
        console.log(`Error: ${error.message} - ${error}`);
    }
}

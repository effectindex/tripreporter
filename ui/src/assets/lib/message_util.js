export function setMessage(message, messageSuccess, status) {
    const elemText = document.getElementById("DefaultView__message_text");
    elemText.textContent = message;

    const elem = document.getElementById("DefaultView__message")
    elem.style.display = 'block';

    if (status === true) {
        elem.style.background = '#3d9991'
        elemText.innerHTML = messageSuccess;
        window.setTimeout(function () {
            window.location.href = "/login";
        }, 3000);
    } else {
        elem.style.background = '#a83232'
    }
}

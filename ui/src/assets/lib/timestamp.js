export default function getFormattedTimestamp({data, showTime, hideDate, longFormat}) {
    const date = new Date(data);

    // Golang zero time
    if (date.getTime() === -62135596800000) {
        return ""
    }

    let options = longFormat ? {weekday: "long", year: "numeric", month: "long", day: "numeric"} : {
        year: "numeric",
        month: "numeric",
        day: "numeric"
    };

    if (showTime) {
        options.hour = "numeric";
        options.minute = "numeric";
    }

    if (hideDate) {
        return date.toLocaleTimeString(undefined, {hour: "2-digit", minute: "2-digit"});
    }

    return date.toLocaleString(undefined, options);
}

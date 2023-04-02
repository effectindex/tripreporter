export function getTextLength(text, max) {
  if (!text) {
    return ""
  }

  return `(${text.length.toLocaleString()} / ${max.toLocaleString()})`
}

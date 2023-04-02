/* eslint-disable no-console */

export default function log(message, ...optionalParams) {
  if (process.env.NODE_ENV !== "production") {
    console.log(message, ...optionalParams)
  }
}

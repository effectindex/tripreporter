/* eslint-disable no-console */

export default function log(message, ...optionalParams) {
    console.log(`BASE_URL: ${process.env.BASE_URL}`)
    console.log(`NODE_ENV: ${process.env.NODE_ENV}`)

    if (process.env.NODE_ENV !== "production") {
        console.log(message, ...optionalParams)
    }
}

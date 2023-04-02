import axios from 'axios'

let url = process.env.VUE_APP_PROD_URL
if (process.env.NODE_ENV !== "production") {
  url = process.env.VUE_APP_DEV_URL
}

const apiClient = axios.create({
  baseURL: `${url}/api/v1`,
  headers: {
    "Access-Control-Allow-Origin": url,
    "Content-Type": "application/json",
  },
})

export default apiClient

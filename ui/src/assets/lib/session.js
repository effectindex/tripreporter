import log from "@/assets/lib/logger";

export async function validateSession(axios) {
  let res = await axios.get('/session/validate').catch(function (error) {
    log("Caught exception in validateSession", error)
  })
  return res ? res.status === 200 : false;
}

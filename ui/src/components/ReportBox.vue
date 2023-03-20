<template>
<h1>meow</h1>
</template>

<script>
import {inject} from "vue";
import {handleMessageError, setMessage} from "@/assets/lib/message_util";
import {useReportsStore} from "@/assets/lib/reportsstore";

const store = useReportsStore();
let ranSetup = false

export default {
  name: "ReportBox",
  props: {
    id: String
  },
  async setup(props) {
    if (ranSetup) {
      return
    }
    ranSetup = true

    const axios = inject('axios')

    await axios.get('/report/'+props.id).then(function (response) {
      store.apiSuccess = response.status === 200;
      store.data = response.data;
      console.log(response.data)
      setMessage(response.data.msg, `<wrap pre>${JSON.stringify(response.data, null, 4)}</wrap>`, store.apiSuccess);
    }).catch(function (error) {
      console.log(error.response)
      console.log(error.response.data.msg)
      store.apiSuccess = error.response.status === 200;
      setMessage(error.response.data.msg, "", store.apiSuccess);
      handleMessageError(error);
    })
  }
}
</script>

<style scoped>
@import url(@/assets/css/message_util.css);
</style>

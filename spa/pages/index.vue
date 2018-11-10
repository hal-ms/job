<template>
<div>
  <qrcode-vue :value="url" size="200" level="H"></qrcode-vue>
</div>
</template>

<script>
import io from "socket.io-client";
import QrcodeVue from "qrcode.vue";
import axios from "axios";

export default {
  data() {
    return {
      url: ""
    };
  },
  components: {
    QrcodeVue
  },
  mounted() {
    this.GetToken();
    this.socket = io("https://hal-iot.net");
    // this.socket = io("localhost:8000");
    this.socket.on("open", msg => {
      this.GetToken();
    });
  },
  methods: {
    GetToken: function() {
      axios.get("https://hal-iot.net/api/token").then(res => {
        this.url = "https://hal-iot.net/create/" + res.data.id;
        console.log(res.data.id);
      });
    }
  }
};
</script>

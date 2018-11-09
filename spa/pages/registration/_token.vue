<template>
  <section>
    <h1 class="title">ジョブ登録</h1>
    <b-field label="ユーザー名">
      <b-input v-model="userName"></b-input>
    </b-field>

    <b-table :data="jobCategory" :columns="jobCategoryClm" :selected.sync="selected" focusable></b-table>
    <a class="button" @click="registerJob">登録</a>

    <!-- <h2 class="title">DEBUG</h2>
    <a class="button" @click="showLocalStorage">ローカルストレージを表示</a>
    <a class="button" @click="clearLocalStorage">ローカルストレージを削除</a>
    <p>local storage: {{ jobIDList }}</p> -->
  </section>
</template>

<script>
import axios from 'axios'

export default {
  // retのに値が仕込まれるのを待つためにasync, awaitを用いている
  async validate ({ params }) {
    let ret = null
    await axios.get("https://hal-iot.net/api/alive_token/" + params.token)
      .then(res => {
        ret = true
      })
      .catch(err => {
        ret = false
      })
    return ret
  },

  data() {
    return {
      jobIDList: null,  // for debug
      userName: "",

      jobCategory: [],
      jobCategoryClm: [{field: "display_name", label: "ジョブ名"}],
      selected: null,
    }
  },

  mounted() {
    this.userName = localStorage.getItem("userName")
    axios.get("https://hal-iot.net/api/job_categorys")
      .then(res => {
        this.jobCategory = res.data
      })
  },

  methods: {
    registerJob() {
      let token = this.parseToken()

      if (this.selected == null) {
        this.danger("ジョブが選択されていません")
        return
      }

      axios.post("https://hal-iot.net/api/job/" + token, {name: this.selected.name, user_name: this.userName})
        .then(res => {
          if (this.userName == "") {
            this.danger("ユーザー名が入力されていません")
            return
          }

          console.log("hello")

          let jobIDStr = localStorage.getItem("jobIDList")
          if (jobIDStr) {
            let jobIDList = jobIDStr.split(",")
            jobIDList.push(res.data.id)
            localStorage.setItem("jobIDList", jobIDList)
          } else {
            localStorage.setItem("jobIDList", res.data.id)
          }
          localStorage.setItem("userName", this.userName)
          this.$router.push("/jobs")
        })
    },

    danger(msg) {
      this.$toast.open({
        message: msg,
        type: 'is-danger'
      })
    },

    showLocalStorage() {
      this.jobIDList = localStorage.getItem("jobIDList")
    },

    clearLocalStorage() {
      localStorage.removeItem("jobIDList")
      localStorage.removeItem("userName")
      console.log("cleared local storage")
    },

    parseToken() {
      return location.pathname.split('/')[2]  // 決め打ち
    },

  },
};
</script>
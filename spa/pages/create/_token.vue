<template>
  <section class="create-job-page-container">
    <div class="input-container">
      <b-field label="name">
        <input type="text" v-model="userName">
      </b-field>
    </div>

    <div class="job-categories">
      <ul>
        <li v-for="(job, index) in jobCategory" :key="index" @click="selectJob(job)">
          <div
            class="job"
            :class="{'selected': currentSelectedJob && currentSelectedJob.name === job.name}"
          >
            <div class="job-icon">
              <img :src="job.image_icon" alt>
            </div>
            <div class="job-name">{{job.display_name}}</div>
          </div>
        </li>
      </ul>
    </div>

    <div class="button-container">
      <button class="create-button" @click="registerJob"></button>
    </div>

    <!-- <h2 class="title">DEBUG</h2>
    <a class="button" @click="showLocalStorage">ローカルストレージを表示</a>
    <a class="button" @click="clearLocalStorage">ローカルストレージを削除</a>
    <p>local storage: {{ jobIDList }}</p> -->
  </section>
</template>


<script>
import axios from "axios";

export default {
  // retのに値が仕込まれるのを待つためにasync, awaitを用いている
  async validate({ params }) {
    let ret = null;
    await axios
      .get("https://hal-iot.net/api/alive_token/" + params.token)
      .then(res => {
        ret = true;
      })
      .catch(err => {
        ret = false;
      });
    return ret;
  },

  data() {
    return {
      jobIDList: null, // for debug
      userName: "",

      jobCategory: [],
      selected: null,

      currentSelectedJob: null
    };
  },

  mounted() {
    this.userName = localStorage.getItem("userName");
    axios.get("https://hal-iot.net/api/job_categorys").then(res => {
      this.jobCategory = res.data;
    });
  },

  methods: {
    selectJob(job) {
      this.currentSelectedJob = job;
    },
    registerJob() {
      let token = this.parseToken();
      if (this.userName == null) {
        this.danger("名前が入力されていません");
        return;
      }
      if (this.currentSelectedJob == null) {
        this.danger("ジョブが選択されていません");
        return;
      }

      axios
        .post("https://hal-iot.net/api/job/" + token, {
          name: this.currentSelectedJob.name,
          user_name: this.userName
        })
        .then(res => {
          if (this.userName == "") {
            this.danger("ユーザー名が入力されていません");
            return;
          }

          let jobIDStr = localStorage.getItem("jobIDList");
          if (jobIDStr) {
            let jobIDList = jobIDStr.split(",");
            jobIDList.push(res.data.id);
            localStorage.setItem("jobIDList", jobIDList);
          } else {
            localStorage.setItem("jobIDList", res.data.id);
          }
          localStorage.setItem("userName", this.userName);

          // this.load()
          // this.$router.push("/jobs")
          this.load();
        });
    },

    danger(msg) {
      this.$toast.open({
        message: msg,
        type: "is-danger"
      });
    },

    sleepByPromise(sec) {
      return new Promise(resolve => setTimeout(resolve, sec * 1000));
    },

    async load() {
      const loadingComponent = this.$loading.open({
        container: this.isFullPage
      });
      await this.sleepByPromise(1);
      loadingComponent.close();
      this.$router.push("/jobs");
    },

    showLocalStorage() {
      this.jobIDList = localStorage.getItem("jobIDList");
    },

    clearLocalStorage() {
      localStorage.removeItem("jobIDList");
      localStorage.removeItem("userName");
      localStorage.removeItem("doneIds");
      console.log("cleared local storage");
    },

    parseToken() {
      return location.pathname.split("/")[2]; // 決め打ち
    }
  }
};
</script>


<style scoped>
.create-job-page-container {
  width: 100vw;
  min-height: 100vh;
  padding-top: 50vw;
  background-image: url("~assets/create_back.svg");
  background-size: 100%;
}
.input-container {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.input-container input {
  outline: none;
  background: transparent;
  border: 2px solid #000;
  border-radius: 5px;
  font-size: 1.5rem;
  text-indent: 0.5rem;
}
.job-categories ul {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  padding: 2rem 0;
}
.job {
  margin: 0.5rem;
  width: 25vw;
  height: 25vw;
  border-radius: 50%;
  background-image: url("~assets/circle.svg");
  background-size: cover;
  background-repeat: no-repeat;
  background-color: #fff;
}
.job.selected {
  width: 30vw;
  height: 30vw;
}
.job .job-icon {
  width: 100%;
  margin: 0 auto;
}
.job .job-icon img {
  display: block;
  width: 100%;
}
.job .job-name {
  text-align: center;
  font-weight: bold;
  font-size: 0.9rem;
}
.button-container {
  display: flex;
  justify-content: center;
}
.create-button {
  border: 0;
  outline: none;
  width: 60vw;
  height: 13vw;
  margin: 0 auto;
  background-image: url("~assets/create_button.svg");
  background-size: 100%;
  background-repeat: no-repeat;
  background-position: center;
  text-align: center;
  font-size: 2rem;
  color: #fff;
  border-radius: 10px;
}
</style>

<template>
  <section>
    <h1 class="title">ジョブ進捗確認</h1>
    <ul >
    <li v-for="(job, index) in jobs" :key="index">
    {{ job }}
  </li>
</ul>
  </section>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      jobs: []
    };
  },

  mounted() {
    // for test
    //this.setAlwaysDoneJobID()

    this.getJobs();
    setInterval(this.getJobs, 500);
  },

  methods: {
    getJobs() {
      let jobIDListStr = localStorage.getItem("jobIDList");

      if (jobIDListStr) {
        let dat = [];
        axios.get("https://hal-iot.net/api/jobs/" + jobIDListStr).then(res => {
          for (let job of res.data) {
            if (job.done) {
              let nflg = true;
              for (let id of this.getDoneIds()) {
                if (job.id == id) {
                  nflg = false;
                }
              }
              if (nflg) {
                this.doneToast(job.display_name);
                this.setDoneIds(job.id);
              }
            }
          }
          this.jobs = res.data;
        });
      }
    },
    setDoneIds(id) {
      let jobIDStr = localStorage.getItem("doneIds");
      if (jobIDStr) {
        let jobIDList = jobIDStr.split(",");
        jobIDList.push(id);
        localStorage.setItem("doneIds", jobIDList);
      } else {
        localStorage.setItem("doneIds", id);
      }
    },
    getDoneIds() {
      let jobIDList;
      let jobIDStr = localStorage.getItem("doneIds");
      if (jobIDStr) {
        jobIDList = jobIDStr.split(",");
      } else {
        jobIDList = new Array();
      }
      return jobIDList;
    },
    doneToast(jobName) {
      this.$toast.open({
        message: "Congratulations！ " + jobName + "が達成されました！",
        type: "is-success"
      });
    },

    // for test
    setAlwaysDoneJobID() {
      let jobIDList = localStorage.getItem("jobIDList").split(",");
      jobIDList.push("5be522641473860001e87a8f");
      localStorage.setItem("jobIDList", jobIDList);
    }
  }
};
</script>
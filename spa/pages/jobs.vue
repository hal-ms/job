<template>
  <div class="jobs-page-wrapper">
    <section class="jobs-page-container">
      <div class="input-container">
        <div class="border-container">
          <p class="from-title">FROM</p>
          <p class="from-name">{{from}}</p>
        </div>
      </div>
      <div class="jobs-container">
        <ul>
          <li v-for="(job, index) in jobs" :key="index">
            <div class="job" :class="{'is-complete': job.done}">
              <div class="icon" :style="{'border-color': job.image_color}">
                <img :src="job.image_icon">
              </div>
              <div class="job-right">
                <p class="date">{{job.created}}</p>
                <p class="display-name">{{job.display_name}}</p>
              </div>
            </div>
          </li>
        </ul>
      </div>
    </section>
  </div>
</template>


<script>
import axios from "axios";

export default {
  data() {
    return {
      jobs: [],
      from: "たなか"
    };
  },

  mounted() {
    // for test
    //this.setAlwaysDoneJobID()
    this.from = localStorage.getItem("userName");
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
          this.jobs = res.data.reverse();
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


<style scoped>
.jobs-page-wrapper {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100%;
  background-color: rgb(255, 192, 179);
  overflow: scroll;
  -webkit-overflow-scrolling: touch;
  /* scroll-behavior: smooth; */
  /* -webkit-overflow-scrolling: touch; */
}
.jobs-page-container {
  width: 100vw;
  padding-top: 50vw;
  padding-bottom: 5rem;
  background-image: url("~assets/jobs_back.svg");
  background-size: 100%;
  background-repeat: no-repeat;
}
.jobs-page-container .input-container {
  display: flex;
  justify-content: flex-end;
  padding-right: 10vw;
}
.jobs-page-container .border-container {
  display: flex;
  align-items: flex-end;
  border-bottom: 1px solid #000;
}
.jobs-page-container .input-container .from-title {
  padding-bottom: 0.2rem;
  font-weight: bold;
}
.jobs-page-container .input-container .from-name {
  font-weight: bold;
  font-size: 2rem;
  padding-left: 1rem;
}
.jobs-container {
  display: flex;
  justify-content: center;
  padding-top: 1.4rem;
}
.jobs-container li + li {
  margin-top: 1.5rem;
}
.jobs-container .job {
  position: relative;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-radius: 5px;
  box-shadow: 0 3px 6px rgba(0, 0, 0, 0.16), 0 3px 6px rgba(0, 0, 0, 0.23);
  background-color: #fff;
  padding: 0 10px;
}
.jobs-container .job.is-complete::after {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 90vw;
  height: 50vw;
  background-image: url("~assets/complete.svg");
  background-size: contain;
  background-repeat: no-repeat;
}
.jobs-container .job .icon {
  width: 20vw;
  height: 20vw;
  border-radius: 50%;
  border: 2px solid red;
  overflow: hidden;
}
.jobs-container .job .icon img {
  display: block;
  width: 100%;
}
.jobs-container .job .display-name {
  font-weight: bold;
}
.jobs-container .job .display-name {
  font-size: 2rem;
  font-weight: bold;
}
.job .job-right {
  padding: 1rem 0;
  padding-left: 1rem;
}
</style>

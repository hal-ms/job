<template>
  <section>
    <h1 class="title">ジョブ進捗確認</h1>
    <b-table :data="jobs" :columns="jobsClm"></b-table>
  </section>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      jobs: [],
      jobsClm: [{field: "name", label: "ジョブ名"}, {field: "progress", label: "進捗"}]
    }
  },

  mounted() {
    // for test
    //this.setAlwaysDoneJobID()

    this.getJobs()
    setInterval(this.getJobs, 3000)
  },

  methods: {
    getJobs() {
      console.log("polling")
      let jobIDListStr = localStorage.getItem("jobIDList")

      if (jobIDListStr) {
        let dat = []
        axios.get("https://hal-iot.net/api/jobs/" + jobIDListStr)
          .then(res => {
            for(let job of res.data) {
              if (job.done) {
                if (job.notification) {
                  this.doneToast(job.display_name)
                }
                dat.push({name: job.display_name, progress: "達成"})
              } else {
                dat.push({name: job.display_name, progress: "未達成"})
              }
            }
          })
        this.jobs = dat
      }
    },

    doneToast(jobName) {
      this.$toast.open({
          message: 'Congratulations！ ' + jobName + 'が達成されました！',
          type: 'is-success'
      })
    },

    // for test
    setAlwaysDoneJobID() {
      let jobIDList = localStorage.getItem("jobIDList").split(',')
      jobIDList.push("5be522641473860001e87a8f")
      localStorage.setItem("jobIDList", jobIDList)
    },

  }
}
</script>
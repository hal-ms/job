<template>
  <section>
    <div v-if="error.statusCode === 404">
      <b-message type="is-danger" has-icon>
        <p>404 ERROR</p>
        <p>In most cases, a person who encounters "404" looks like the following.</p>
        <p>訳: 大抵の場合、404に出くわした人間は以下のような表情をします。</p>
      </b-message>
      <p><img src="../assets/masui.jpg" width="256" height="256"></p>
    </div>
    <div v-else>
      <b-message type="is-danger" has-icon>
        <p>ERROR!!</p>
      </b-message>
    </div>
  </section>
</template>
<script>
export default {
  props: ["error"],
  mounted() {
    setTimeout(() => {
      this.$router.push("/jobs");
    }, 2000);
  }
};
</script>
(window.webpackJsonp=window.webpackJsonp||[]).push([[4],{202:function(t,e){},207:function(t,e,n){"use strict";n.r(e);var o=n(175),i=n.n(o),s=n(205),a=n(135),u=n.n(a),c={data:function(){return{url:""}},components:{QrcodeVue:s.a},mounted:function(){var t=this;this.GetToken(),this.socket=i()("https://hal-iot.net"),this.socket.on("open",function(e){t.GetToken()})},methods:{GetToken:function(){var t=this;u.a.get("https://hal-iot.net/api/token").then(function(e){t.url="https://hal-iot.net/create/"+e.data.id,console.log(e.data.id)})}}},l=n(17),r=Object(l.a)(c,function(){var t=this.$createElement,e=this._self._c||t;return e("div",[e("qrcode-vue",{attrs:{value:this.url,size:"200",level:"H"}})],1)},[],!1,null,null,null);r.options.__file="index.vue";e.default=r.exports}}]);
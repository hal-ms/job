(window.webpackJsonp=window.webpackJsonp||[]).push([[3],{206:function(e,t,a){"use strict";a.r(t);a(77),a(35),a(36);var o,n,s=a(3),r=a.n(s),i=a(135),l=a.n(i),c={validate:(n=r()(regeneratorRuntime.mark(function e(t){var a,o;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return a=t.params,o=null,e.next=4,l.a.get("https://hal-iot.net/api/alive_token/"+a.token).then(function(e){o=!0}).catch(function(e){o=!1});case 4:return e.abrupt("return",o);case 5:case"end":return e.stop()}},e,this)})),function(e){return n.apply(this,arguments)}),data:function(){return{jobIDList:null,userName:"",jobCategory:[],jobCategoryClm:[{field:"display_name",label:"ジョブ名"}],selected:null}},mounted:function(){var e=this;this.userName=localStorage.getItem("userName"),l.a.get("https://hal-iot.net/api/job_categorys").then(function(t){e.jobCategory=t.data})},methods:{registerJob:function(){var e=this,t=this.parseToken();null!=this.selected?l.a.post("https://hal-iot.net/api/job/"+t,{name:this.selected.name,user_name:this.userName}).then(function(t){if(""!=e.userName){var a=localStorage.getItem("jobIDList");if(a){var o=a.split(",");o.push(t.data.id),localStorage.setItem("jobIDList",o)}else localStorage.setItem("jobIDList",t.data.id);localStorage.setItem("userName",e.userName),e.load()}else e.danger("ユーザー名が入力されていません")}):this.danger("ジョブが選択されていません")},danger:function(e){this.$toast.open({message:e,type:"is-danger"})},sleepByPromise:function(e){return new Promise(function(t){return setTimeout(t,1e3*e)})},load:(o=r()(regeneratorRuntime.mark(function e(){var t;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return t=this.$loading.open({container:this.isFullPage}),e.next=3,this.sleepByPromise(1);case 3:t.close(),console.log("konnichiwa"),this.$router.push("/jobs");case 6:case"end":return e.stop()}},e,this)})),function(){return o.apply(this,arguments)}),showLocalStorage:function(){this.jobIDList=localStorage.getItem("jobIDList")},clearLocalStorage:function(){localStorage.removeItem("jobIDList"),localStorage.removeItem("userName"),localStorage.removeItem("doneIds"),console.log("cleared local storage")},parseToken:function(){return location.pathname.split("/")[2]}}},u=a(17),m=Object(u.a)(c,function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("section",[a("h1",{staticClass:"title"},[e._v("ジョブ登録")]),e._v(" "),a("b-field",{attrs:{label:"ユーザー名"}},[a("b-input",{model:{value:e.userName,callback:function(t){e.userName=t},expression:"userName"}})],1),e._v(" "),a("b-table",{attrs:{data:e.jobCategory,columns:e.jobCategoryClm,selected:e.selected,focusable:""},on:{"update:selected":function(t){e.selected=t}}}),e._v(" "),a("a",{staticClass:"button",on:{click:e.registerJob}},[e._v("登録")]),e._v(" "),a("h2",{staticClass:"title"},[e._v("DEBUG")]),e._v(" "),a("a",{staticClass:"button",on:{click:e.showLocalStorage}},[e._v("ローカルストレージを表示")]),e._v(" "),a("a",{staticClass:"button",on:{click:e.clearLocalStorage}},[e._v("ローカルストレージを削除")]),e._v(" "),a("p",[e._v("local storage: "+e._s(e.jobIDList))])],1)},[],!1,null,null,null);m.options.__file="_token.vue";t.default=m.exports}}]);
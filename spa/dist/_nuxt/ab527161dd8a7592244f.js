(window.webpackJsonp=window.webpackJsonp||[]).push([[2],{130:function(n,t,a){var e=a(158);"string"==typeof e&&(e=[[n.i,e,""]]),e.locals&&(n.exports=e.locals);(0,a(35).default)("7a40ee17",e,!0,{sourceMap:!1})},131:function(n,t){n.exports=function(n){return"string"!=typeof n?n:(/^['"].*['"]$/.test(n)&&(n=n.slice(1,-1)),/["'() \t\n]/.test(n)?'"'+n.replace(/"/g,'\\"').replace(/\n/g,"\\n")+'"':n)}},157:function(n,t,a){"use strict";var e=a(130);a.n(e).a},158:function(n,t,a){var e=a(131);(n.exports=a(34)(!1)).push([n.i,"\n.jobs-page-wrapper[data-v-c3d0a26a] {\n  position: fixed;\n  top: 0;\n  left: 0;\n  width: 100vw;\n  height: 100%;\n  background-color: rgb(255, 192, 179);\n  overflow: scroll;\n  -webkit-overflow-scrolling: touch;\n  /* scroll-behavior: smooth; */\n  /* -webkit-overflow-scrolling: touch; */\n}\n.jobs-page-container[data-v-c3d0a26a] {\n  width: 100vw;\n  padding-top: 50vw;\n  padding-bottom: 5rem;\n  background-image: url("+e(a(159))+');\n  background-size: 100%;\n  background-repeat: no-repeat;\n}\n.jobs-page-container .input-container[data-v-c3d0a26a] {\n  display: flex;\n  justify-content: flex-end;\n  padding-right: 10vw;\n}\n.jobs-page-container .border-container[data-v-c3d0a26a] {\n  display: flex;\n  align-items: flex-end;\n  border-bottom: 1px solid #000;\n}\n.jobs-page-container .input-container .from-title[data-v-c3d0a26a] {\n  padding-bottom: 0.2rem;\n  font-weight: bold;\n}\n.jobs-page-container .input-container .from-name[data-v-c3d0a26a] {\n  font-weight: bold;\n  font-size: 2rem;\n  padding-left: 1rem;\n}\n.jobs-container[data-v-c3d0a26a] {\n  display: flex;\n  justify-content: center;\n  padding-top: 1.4rem;\n}\n.jobs-container li + li[data-v-c3d0a26a] {\n  margin-top: 1.5rem;\n}\n.jobs-container .job[data-v-c3d0a26a] {\n  position: relative;\n  display: flex;\n  justify-content: space-between;\n  align-items: center;\n  border-radius: 5px;\n  box-shadow: 0 3px 6px rgba(0, 0, 0, 0.16), 0 3px 6px rgba(0, 0, 0, 0.23);\n  background-color: #fff;\n  padding: 0 10px;\n}\n.jobs-container .job.is-complete[data-v-c3d0a26a]::after {\n  content: "";\n  position: absolute;\n  top: 0;\n  left: 0;\n  width: 90vw;\n  height: 50vw;\n  background-image: url('+e(a(160))+");\n  background-size: contain;\n  background-repeat: no-repeat;\n}\n.jobs-container .job .icon[data-v-c3d0a26a] {\n  width: 20vw;\n  height: 20vw;\n  border-radius: 50%;\n  border: 2px solid red;\n  overflow: hidden;\n}\n.jobs-container .job .icon img[data-v-c3d0a26a] {\n  display: block;\n  width: 100%;\n}\n.jobs-container .job .display-name[data-v-c3d0a26a] {\n  font-weight: bold;\n}\n.jobs-container .job .display-name[data-v-c3d0a26a] {\n  font-size: 2rem;\n  font-weight: bold;\n}\n.job .job-right[data-v-c3d0a26a] {\n  padding: 1rem 0;\n  padding-left: 1rem;\n}\n",""])},159:function(n,t,a){n.exports=a.p+"img/ead96c4.svg"},160:function(n,t,a){n.exports=a.p+"img/8285fcd.svg"},167:function(n,t,a){"use strict";a.r(t);a(76),a(49),a(50),a(25);var e=a(133),o=a.n(e),i={data:function(){return{jobs:[],from:"たなか"}},mounted:function(){this.from=localStorage.getItem("userName"),this.getJobs(),setInterval(this.getJobs,500)},methods:{getJobs:function(){var n=this,t=localStorage.getItem("jobIDList");if(t){o.a.get("https://hal-iot.net/api/jobs/"+t).then(function(t){var a=!0,e=!1,o=void 0;try{for(var i,s=t.data[Symbol.iterator]();!(a=(i=s.next()).done);a=!0){var r=i.value;if(r.done){var d=!0,c=!0,l=!1,p=void 0;try{for(var g,b=n.getDoneIds()[Symbol.iterator]();!(c=(g=b.next()).done);c=!0){var f=g.value;r.id==f&&(d=!1)}}catch(n){l=!0,p=n}finally{try{c||null==b.return||b.return()}finally{if(l)throw p}}d&&(n.doneToast(r.display_name),n.setDoneIds(r.id))}}}catch(n){e=!0,o=n}finally{try{a||null==s.return||s.return()}finally{if(e)throw o}}n.jobs=t.data.reverse()})}},setDoneIds:function(n){var t=localStorage.getItem("doneIds");if(t){var a=t.split(",");a.push(n),localStorage.setItem("doneIds",a)}else localStorage.setItem("doneIds",n)},getDoneIds:function(){var n=localStorage.getItem("doneIds");return n?n.split(","):new Array},doneToast:function(n){this.$toast.open({message:"Congratulations！ "+n+"が達成されました！",type:"is-success"})},setAlwaysDoneJobID:function(){var n=localStorage.getItem("jobIDList").split(",");n.push("5be522641473860001e87a8f"),localStorage.setItem("jobIDList",n)}}},s=(a(157),a(17)),r=Object(s.a)(i,function(){var n=this,t=n.$createElement,a=n._self._c||t;return a("div",{staticClass:"jobs-page-wrapper"},[a("section",{staticClass:"jobs-page-container"},[a("div",{staticClass:"input-container"},[a("div",{staticClass:"border-container"},[a("p",{staticClass:"from-title"},[n._v("FROM")]),n._v(" "),a("p",{staticClass:"from-name"},[n._v(n._s(n.from))])])]),n._v(" "),a("div",{staticClass:"jobs-container"},[a("ul",n._l(n.jobs,function(t,e){return a("li",{key:e},[a("div",{staticClass:"job",class:{"is-complete":t.done}},[a("div",{staticClass:"icon",style:{"border-color":t.image_color}},[a("img",{attrs:{src:t.image_icon}})]),n._v(" "),a("div",{staticClass:"job-right"},[a("p",{staticClass:"date"},[n._v(n._s(t.created))]),n._v(" "),a("p",{staticClass:"display-name"},[n._v(n._s(t.display_name))])])])])}))])])])},[],!1,null,"c3d0a26a",null);r.options.__file="jobs.vue";t.default=r.exports}}]);
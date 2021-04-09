// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from "vue";
import App from "./App";
import router from "./router";
import ElementUI from "element-ui"; //element-ui的全部组件
import "element-ui/lib/theme-chalk/index.css"; //element-ui的css
import Store from "@/store/store";
import "vue-awesome/icons";
import Icon from "vue-awesome/components/Icon";
import api from "@/api/api";

import layer from 'vue-layer'
import 'vue-layer/lib/vue-layer.css';

Vue.prototype.$layer = layer(Vue);

Vue.use(ElementUI); //使用elementUI
Vue.config.productionTip = false;
Vue.component("icon", Icon);

/* eslint-disable no-new */

router.beforeEach((to, from, next) => {
    if (to.meta.requireAuth) {
        if (!api.isEmptyObject(Store.state.user)) {
            next();
        } else {
            next({
                path: "/login",
                query: { redirect: to.fullPath } // 将跳转的路由path作为参数，登录成功后跳转到该路由
            });
        }
    } else {
        next();
    }
});

new Vue({
    el: "#app",
    router,
    components: { App },
    template: "<App/>"
});

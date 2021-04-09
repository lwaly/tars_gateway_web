<template>
    <span class="leables" :class="router == names.path ? 'active' : ''">
        <router-link :to="names.path"> {{ names.name }} </router-link>
        <i class="el-icon-close" @click="handleClose(names.path)"></i>
    </span>
</template>

<script>
export default {
    data() {
        return {
            nameTab: "用户列表",
        };
    },
    props: ["names"],
    methods: {
        handleClose(tag) {
            this.$emit("removeTable", tag);
        },
    },
    computed: {
        router() {
            return this.$route.path;
        },
    },
    watch: {
        router(news) {},
    },
};
</script>

<style style lang='scss' scoped>
.leables {
    padding: 5px 10px;
    border: 1px solid #f8c56e;
    border-radius: 2px;
    cursor: pointer;
    margin-right: 5px;
}
.leables a {
    color: #f8c56e;
    text-decoration: none;
}
.active {
    background: #f8c56e;
}
.active a {
    color: #ffffff !important;
}
</style>
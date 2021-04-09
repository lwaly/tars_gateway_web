<template>
    <div class="headers">
        <div class="headers-left">
            <div>
                {{ router.meta.title }}
            </div>
            <div class="headers-left-bottom">
                <Leables
                    v-for="(item, index) in tableList"
                    :key="index"
                    :names="item"
                    @removeTable="removeTable"
                />
            </div>
        </div>
    </div>
</template>

<script>
import Leables from "@/components/Leables";
import Store from "@/store/store";

export default {
    data() {
        return {
            squareUrl:
                "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png",
            tableList: [],
            icons: true,
        };
    },
    methods: {
        quies: function () {
            //   console.log(15455)
            Store.commit("removeToken");
        },
        goIndex: function () {
            this.$router.push("/index");
        },
        removeTable(tag) {
            for (let i = 0; i < this.tableList.length; i++) {
                if (this.tableList[i].path == tag) {
                    this.tableList.splice(i, 1);
                }
            }
            if (this.$route.path == tag) {
                if (this.tableList.length > 1) {
                    this.$router.push(
                        this.tableList[this.tableList.length - 1]
                    );
                } else {
                    this.$router.push("/index");
                }
            }
        },
        iconsBtns() {
            this.icons = !this.icons;
        },
    },
    components: {
        Leables,
    },
    computed: {
        userList() {
            return Store.state.load.userList;
        },
        router: function () {
            return this.$route;
        },
    },
    watch: {
        router: {
            handler(news, olds) {
                console.log("news", news);
                console.log("olds", olds);

                if (news.meta.table) {
                    let offs = this.tableList.every((item, index) => {
                        console.log("every", item.path, news.path);
                        return item.path != news.path;
                    });
                    console.log("this.tableList", offs, this.tableList);
                    if (offs) {
                        this.tableList.push(news);
                    }
                }
            },
            immediate: true,
        },
    },
};
</script>

<style style lang='scss' scoped>
.headers {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: space-between;
    overflow-y: hidden;
    box-sizing: border-box;
}

.el-dropdown-link {
    cursor: pointer;
    color: #409eff;
    line-height: 50px;
    display: inline-block;
    height: 50px;
    width: auto;
    display: flex;
    justify-content: flex-start;
}
.el-icon-arrow-down {
    font-size: 12px;
}
.demonstration {
    display: block;
    color: #8492a6;
    font-size: 14px;
    margin-bottom: 20px;
}
.el-dropdown-link-names {
    display: inline-block;
    height: 50px;
    line-height: 50px;
    margin-left: 5px;
}
.headers-left {
    width: calc(100% - 1px);
    background: #e9e5df;
}
.headers-left-top {
    width: 100%;
    height: 30px;
    font-size: 12px;
    line-height: 30px;
    text-align: left;
}
.headers-left-top i {
    margin-right: 10px;
    cursor: pointer;
    font-size: 16px;
    text-align: left;
}
.headers-left-bottom {
    width: 100%;
    height: 34px;
    font-size: 12px;
    padding-top: 5px;
    text-align: left;
}
</style>
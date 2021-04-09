<template>
    <div id="index" style="height: 100%">
        <el-container style="height: 100%">
            <el-header class="header">
                <Header />
            </el-header>
            <el-container style="height: 100%">
                <el-aside width="200px" style="height: 100%">
                    <Home></Home>
                </el-aside>
                <el-main>
                    <el-aside width="100%" style="height: 30px">
                        <NavigationTab></NavigationTab>
                    </el-aside>
                    <div style="padding-top: 25px">
                        <router-view></router-view>
                    </div>
                </el-main>
            </el-container>
        </el-container>
    </div>
</template>

<script>
import Home from "@/components/Home";
import NavigationTab from "@/components/NavigationTab";
import Header from "@/components/Header";

export default {
    name: "index",
    components: {
        Home: Home,
        Header: Header,
        NavigationTab: NavigationTab,
    },
    data() {
        return {
            squareUrl:
                "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png",
            tableList: [],
        };
    },
};
</script>

<style>
#index {
    font-family: "Avenir", Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
}
.header {
    background-color: #409eff;
    color: #fff;
    line-height: 60px;
}
.headers-left-bottom {
    width: 100%;
    height: 34px;
    font-size: 12px;
    padding-top: 5px;
}
</style>

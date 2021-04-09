<template>
    <div class="editmanage">
        <form v-on:submit="editManage">
            <div class="input-box">
                <span class="text-box"> 账号：</span>
                <el-input
                    clearable
                    size="small"
                    style="width: 160px"
                    type="text"
                    placeholder="账号"
                    v-model="manage.account"
                />
            </div>
            <div class="input-box">
                <span class="text-box">用户名：</span>
                <el-input
                    clearable
                    size="small"
                    style="width: 160px"
                    type="text"
                    placeholder="用户名"
                    v-model="manage.username"
                />
            </div>
            <div class="input-box">
                <span class="text-box">密码：</span>
                <el-input
                    clearable
                    size="small"
                    style="width: 160px"
                    type="password"
                    placeholder="密码"
                    v-model="manage.password"
                />
            </div>
            <div class="input-box">
                <span class="text-box">权限：</span>
                <el-select
                    placeholder="请选择人员角色"
                    v-model="manage.authority"
                >
                    <el-option label="超级管理员" value="admin"
                        >超级管理员</el-option
                    >
                    <el-option label="管理员" value="admin1">管理员</el-option>
                </el-select>
            </div>
            <div class="input-box">
                <el-button class="button-box" type="submit">立即提交</el-button>
            </div>
        </form>
    </div>
</template>
 
<script>
export default {
    name: "addmanage",
    data() {
        return {
            manage: {},
            form: {},
        };
    },
    props: {
        //接收父组件传来的id值
        info: {
            type: Object,
            default: () => {
                return {};
            },
        },
        layerid: {
            type: String,
            default: "",
        },
        lydata: {
            type: Object,
            default: () => {
                return {};
            },
        },
    },
    methods: {
        //用来显示对应id的编辑内容，在created中调用
        fetchManage(id) {
            this.manage = {
                account: "account",
                username: "username",
                password: "password",
                authority: "",
            };
            console.log("fetchManage", this.manage);
        },
        editManage(e) {
            if (
                !this.manage.account ||
                !this.manage.username ||
                !this.manage.password
            ) {
                this.$layer.msg("请添加对应信息！");
            } else {
                let updateManage = {
                    account: this.manage.account,
                    username: this.manage.username,
                    password: this.manage.password,
                    authority: this.manage.authority,
                };
                this.$http
                    .put(
                        "http://localhost:3000/manage/" + this.info.id,
                        updateManage
                    )
                    .then(function (response) {
                        //关闭父组件中的编辑弹框
                        this.$layer.close(this.layerid);
                        //弹出提示信息，默认时间为3秒
                        this.$layer.msg("修改管理员信息成功！");
                    });
                e.preventDefault();
            }
            e.preventDefault();
        },
    },
    created() {
        // this.info.id父组件传给子组件的id值
        this.fetchManage(this.info.id);
    },
};
</script>
 
<style style lang='scss' scoped>
.editmanage {
    width: 100%;
    height: 100%;
    padding: vw(50);
    top: 100px;
    .input-box {
        position: relative;
        top: 20%;
        left: 23%;
        top: 30px;
        box-sizing: border-box;
        padding: vw(50);
        margin: 10px;
        .text-box {
            width: 70px;
            display: inline-block;
            text-align: left;
        }
        .button-box {
            position: relative;
            bottom: 50%;
            left: 15%;
            margin: 20px;
        }
    }
}
</style>

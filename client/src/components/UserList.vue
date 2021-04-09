<template>
    <div class="customzie">
        <div class="search-box">
            <div>
                <span>
                    ip：
                    <el-input
                        clearable
                        placeholder="请输入"
                        size="small"
                        style="width: 160px"
                        v-model.number="ip"
                        type="number"
                    >
                    </el-input>
                </span>
                <el-select v-model="area" size="small" placeholder="请选择模式">
                    <el-option
                        v-for="item in selectmode"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    ></el-option>
                </el-select>

                <el-button
                    type="primary"
                    plain
                    size="small"
                    @click="UserListGet"
                    >查询</el-button
                >
                <el-button type="primary" plain size="small" @click="Reset"
                    >重置</el-button
                >
            </div>
        </div>
        <el-table :data="tableData" style="width: 100%">
            <el-table-column label="区域" align="center">
                <template slot-scope="scope">
                    <div slot="reference" class="nickname-wrapper">
                        <el-tag size="medium">{{ scope.row.nickname }}</el-tag>
                    </div>
                </template>
            </el-table-column>
            <el-table-column label="ip" align="center" width="180">
                <template slot-scope="scope">
                    <i class="el-icon-time"></i>
                    <span style="margin-left: 10px">{{
                        scope.row.create_time
                    }}</span>
                </template>
            </el-table-column>
            <el-table-column label="创建时间" align="center" width="180">
                <template slot-scope="scope">
                    <i class="el-icon-time"></i>
                    <span style="margin-left: 10px">{{
                        scope.row.create_time
                    }}</span>
                </template>
            </el-table-column>
            <el-table-column label="操作" align="center">
                <template slot-scope="scope">
                    <el-button
                        size="mini"
                        @click="handleEdit(scope.$index, scope.row)"
                        >编辑</el-button
                    >
                    <el-button
                        size="mini"
                        type="danger"
                        @click="handleDelete(scope.$index, scope.row)"
                        >删除</el-button
                    >
                </template>
            </el-table-column>
        </el-table>
        <el-pagination
            @current-change="handleCurrentChange"
            :page-size="10"
            layout="total, sizes, prev, pager, next, jumper"
            :total="totalCount"
        >
        </el-pagination>
    </div>
</template>

<script>
import Store from "@/store/store";
import api from "@/api/api";
import UserEdit from "@/components/UserEdit";

export default {
    components: {},
    data() {
        return {
            // 显示提示框
            showMsg: true,
            totalCount: 1,
            tableData: [],
            area: "",
            selectmode: [
                { value: "repair", label: "修复" },
                { value: "append", label: "追加" },
                { value: "overwirtte", label: "覆盖" },
            ],
            ip: "",
            sex: "",
            loginTypeWx: false,
            loginTypeWb: false,
            loginTypeQq: false,
            startTime: "",
            endTime: "",
        };
    },
    computed: {},
    watch: {},
    created() {},
    mounted() {},
    destroyed() {},
    methods: {
        handleEdit(index, row) {
            this.$layer.iframe({
                type: 2,
                title: "编辑",
                area: ["600px", "450px"],
                shade: true,
                offset: "auto",
                content: {
                    content: UserEdit, //传递的编辑组件主线
                    parent: this,
                    data: {
                        info: { ip: this.ip }, // 传递的要编辑内容的id值
                    },
                },
            });
        },
        getSex(val) {
            this.sex = val;
            console.log("获得的单选框值是：", val, typeof val);
        },
        handleDelete(index, row) {
            this.showMsg = false;
            // 提示是否确认删除
            console.log(index);
            console.log(row);
        },
        Reset() {
            this.selectmode.push({ value: "add", label: "add" });
        },
        UserListGet() {
            var parm = { page_index: 1, page_size: 10 };
            if (0 != this.ip) {
                parm.ip = this.ip;
            }
            if (0 != this.sex) {
                parm.sex = this.sex;
            }
            if (this.loginTypeWx) {
                parm.login_type = 1;
            } else if (this.loginTypeWb) {
                parm.login_type = 2;
            } else if (this.loginTypeQq) {
                parm.login_type = 3;
            }
            if (0 < new Date(this.startTime[0]).getTime() / 1000) {
                parm.start_time = new Date(this.startTime[0]).getTime() / 1000;
            }
            if (0 < new Date(this.startTime[1]).getTime() / 1000) {
                parm.end_time = new Date(this.startTime[1]).getTime() / 1000;
            }
            console.log(
                this.startTime,
                new Date(this.startTime[0]).getTime() / 1000
            );
            api.postJSON("/MuzenBAS/User/UserFind", parm).then((res) => {
                console.log(res);
                if (0 == res.data.code) {
                    this.tableData = res.data.data.data.slice(0, 10);
                    this.totalCount = res.data.data.count;
                    console.log(res.data.data);
                }
            });
        },
        /**
         *@description: 点击确定执行的方法
         */
        confirm() {
            // this.$refs.dialog.visible = false;
            console.log("执行确认方法");
        },
        handleCurrentChange(pageIndex) {
            api.postJSON("/MuzenBAS/User/UserFind", {
                page_index: pageIndex,
                page_size: 10,
            }).then((res) => {
                console.log(res);
                if (0 == res.data.code) {
                    this.tableData = res.data.data.data.slice(0, 10);
                    this.totalCount = res.data.data.count;
                    console.log(res.data.data);
                }
            });
        },
    },
};
</script>
<style lang='scss' scoped>
.customzie {
    width: 100%;
    height: vh(550);
    .search-box {
        position: relative;
        .btn-right {
            position: absolute;
            top: 0;
            right: 0;
        }
    }
}
.el-pagination {
    text-align: right;
}
</style>

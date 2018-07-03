var lof = console.log.bind(this)
var app = new Vue({
    el: '#app',
    data: {
        username: "",
        password: "",
        tologin: true,
        toregister: false,
       
    },
    methods: {
        showTime() {
            lof("i am show Time ")
        },
        toLogin() {
            this.tologin = true,
                this.toregister = false
               
        },
       
        register() {
            $.ajax({
                type: "post",
                url: "/api/register",
                data: {
                    username: this.username,
                    password: this.password
                },
                dataType: "json",
                timeout: 10000,
                success: function (d) {
                    console.log(d, "d");
                },
                error: function (XMLHttpRequest, textStatus, errorThrown) {
                    console.log("获取用户信息失败");
                }
            });
        },
        toRegister() {
            this.tologin = false,
            this.toregister = true
               
        },
        login() {
            var that = this;
            $.ajax({
                type: "post",
                url: "/api/login",
                data: {
                    username: this.username,
                    password: this.password
                },
                dataType: "json",
                timeout: 10000,
                success: function (d) {
                    console.log(d, "d");
                    if (d.Code == 200) {
                      location.href = "http://localhost:8082/game.html"
                       //window.location.href = "localhost:8082/game.html"
                    } else {
                        alert("用户名或密码错误")
                    }
                },
                error: function (XMLHttpRequest, textStatus, errorThrown) {
                    console.log("获取用户信息失败");
                }
            });
        }
    },
    mounted() {
        // this.showTime()
        // this.login()
    },
    created() {

    }
})
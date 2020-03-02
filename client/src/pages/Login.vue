<template>
    <div>
        <h1>login</h1>
	<form @submit.prevent="login">
		<label for="">ID：</label>
		<input type="text" v-model="loginform.id">
		<label for="">Password：</label>
		<input type="text" v-model="loginform.password">
		<input type="submit" value="ログイン">
	</form>
	<a href="register">新規登録</a>
    </div>
</template>

<script>
export default {
    data() {
        return {
            loginform : {
                id : '',
                password : '',
            },
            hoge : ''
        };
    },
    methods : {
        async login() {
            await this.axios
            .post('/login', {"id" : this.loginform.id,"password" : this.loginform.password})
            .then(response => (this.hoge = response))
            console.log(this.loginform)
            console.log(this.hoge)
        }
    }
}
</script>
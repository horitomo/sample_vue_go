<template>
    <div class="container">
		<ul class="tab">
			<li class="tab__item" @click="tab = 1">
				Login
			</li>
			<li class="tab__item" @click="tab = 2">
				Register
			</li>
		</ul>
		<div class="panel" v-show="tab === 1">
			<form class="form" @submit.prevent="login">
				<label for="login-email">Email</label>
				<input type="text" class="form__item" id="login-email"  v-model="loginForm.email">
				<label for="login-password">Password</label>
				<input type="password" class="form__item" id="login-password"  v-model="loginForm.password">
				<div class="form__button">
					<button type="submit" class="button button--inverse">login</button>
				</div>
			</form>
		</div>
		<div class="panel" v-show="tab === 2">
			{{ this.msg }}
			<form class="form" @submit.prevent="register">
			<label for="username">Name</label>
			<input type="text" class="form__item" id="username" v-model="registerForm.name">
			<label for="email">Email</label>
			<input type="text" class="form__item" id="email" v-model="registerForm.email">
			<label for="password">Password</label>
			<input type="password" class="form__item" id="password" v-model="registerForm.password">
			<div class="form__button">
				<button type="submit" class="button button--inverse">register</button>
			</div>
			</form>
		</div>
	</div>
</template>

<script>
export default {
    data() {
        return {
            tab: 1,
			loginForm: {
				email: '',
				password: ''
			},
			registerForm: {
				name: '',
				email: '',
				password: '',
            },
			user : '',
			msg : ''
        };
    },
    methods : {
        async login() {
            await this.axios
            .post('/login', {"email" : this.loginForm.email,"password" : this.loginForm.password})
            .then(response => (this.user = response))
			console.log(this.user)
			await this.$store.dispatch('auth/login', this.user.data.request)
			this.$router.push("/")
        },
        async register() {
			this.msg = ''
			if(!this.registerForm.name || !this.registerForm.password || !this.registerForm.email){
				this.msg = "入力が足りません"
				console.log(this.msg)
			}else {
				console.log(this.registerForm)
				await this.axios
				.post('/register', {"email" : this.registerForm.email,"password" : this.registerForm.password, "name" : this.registerForm.name})
				.then(response => (this.user = response))
				if(!this.user.data.error){
					await this.$store.dispatch('auth/register', this.user.data.request)
					this.$router.push("/")
				}
				console.log(this.user)
			}
        }
    }
}
</script>
<template>
    <div>
        <p>{{ hoge }}</p>
        <button>hoge</button>
        <form @submit.prevent="test">
            <button type="submit">get</button>
        </form>
    </div>
</template>
<script>
export default {
    data() {
        return {
            msg : 'msg',
            hoge: ''
        }
    },
    // methods : {
    //     async test(){
    //         var a = this.axios.post("http://localhost:8888/image")
    //         console.log(a)
    //         this.msg = a
    //     }
    // },
    methods: {
        test() {
            console.log(this.msg)
        this.axios
        .post('/api', {"msg" : this.msg})
        .then(response => (this.hoge = response))
        }
    }
}
</script>
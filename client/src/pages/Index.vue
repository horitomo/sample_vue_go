<template>
    <div>
        <p>{{ hoge }}</p>
        <button>hoge</button>
        <form @submit.prevent="test">
            <button type="submit">get</button>
        </form>
        <h2>追加</h2>
		<form @submit.prevent="todoRegister">
			<p>内容<input type="text" name="text" size="30" placeholder="入力してください" ></p>
			<p>状態
			<select name="status">
				<option value="未実行">未実行</option>
				<option value="実行中">実行中</option>
				<option value="終了">終了</option>
			</select>
			</p>
			<p><button type="submit">Send</button></p>
		</form>
		
		<ul>
            <Todos
                v-for="todo in todos"
                :key="todo.id"
                :item="todo"
            />
		</ul>
    </div>
</template>
<script>
import Todos from '../components/Todos'
export default {
    data() {
        return {
            todos : [],
            msg : 'msg',
            hoge: '',
        }
    },
    methods: {
        test() {
            console.log(this.msg)
            this.axios
            .post('/api', {"msg" : this.msg})
            .then(response => (this.hoge = response))
            
            this.todos[1] = {"id" : 1, "Text" : "aaa" , "status" : "実行済"}
        },
        fetchTodos() {
            this.todos[0] = {"id" : 1, "Text" : "aaa" , "status" : "実行済"}
            console.log(this.todos)
        }
    },
    components : {
        Todos
    },
    watch: {
        $route: {
        async handler () {
            await this.fetchTodos()
        },
        immediate: true
        }
    }
}
</script>
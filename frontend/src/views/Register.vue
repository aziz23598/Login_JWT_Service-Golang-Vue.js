<template> 
<form @submit.prevent="submit">
    <h1 class="h3 mb-3 fw-normal">Please register</h1>
      <input v-model="data.name" type="name" class="form-control" placeholder="Name">

      <input v-model="data.email" type="email" class="form-control" placeholder="Email">
    
      <input v-model="data.password" type="password" class="form-control" placeholder="Password">

    <button class="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
  </form>
</template>

<script lang="ts">
import {reactive} from 'vue';
import {useRouter} from 'vue-router';

export default {
    name: "Register",
    setup() {
        const data = reactive( {
            name: '',
            email: '',
            passwod: ''
        });
        const router = useRouter();

        const submit = async () => {
            await fetch('http://localhost:8000/api/register', {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify(data)
            });

            await router.push('/login');
        }

        return {
            data,
            submit
        }
    }
}
</script>

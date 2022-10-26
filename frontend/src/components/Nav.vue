<template>
<nav class="navbar navbar-expand-md navbar-dark bg-dark mb-4">
  <div class="container-fluid">
    <router-link to="/" class="navbar-brand" href="#">Home</router-link>
    
    <div>
      <ul class="navbar-nav me-auto mb-2 mb-md-0" v-if="!auth">
        <li class="nav-item">
          <router-link to="/login" class="nav-link active" href="#">Login</router-link>
        </li>
        <li class="nav-item">
          <router-link to="/register" class="nav-link active" href="#">Register</router-link>
        </li>
      </ul>
      <ul class="navbar-nav me-auto mb-2 mb-md-0" v-if="auth">
        <li class="nav-item">
          <a href="#" class="nav-link active" @click="logout">Logout</a>
        </li>
      </ul>
    </div>
  </div>
</nav>
</template>

<script lang="ts"> 
import {computed} from 'vue';
import {useStore} from "vuex";
import {useRouter} from 'vue-router';
export default {
    name: "Nav",
    setup() {
        const store = useStore();

        const auth = computed(() => store.state.authenticated)

        const router = useRouter();

        const logout = async () => {
            await fetch('http://localhost:8000/api/logout', {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                credentials: 'include',
            });
            await router.push('/login')
        }
        return {
            auth,
            logout
        }
    }
}
</script>
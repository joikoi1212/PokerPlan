<script setup>
import { useRouter } from 'vue-router';

const router = useRouter(); 
const handleLogin = () => {
  try {
    // Redirect the user to the backend's login handler
    window.location.href = 'http://localhost:8080/login';
  } catch (error) {
    console.error('Error during login:', error);
  }
};
</script>

<template>
  <nav class="navbar navbar-expand-lg fixed-top" >
    <div class="container-fluid d-flex">
      <img src="../../assets/logo.png" class="me-2"/>
      <span class="navbar-brand mb-0 h1">PokerPlan</span>
      <button class="btn btn-outline-primary ms-auto me-3" @click="handleLogin"  type="button">Login</button>
    </div>
  </nav>
</template>

<style scoped>
.navbar {
  background-color: #eff3f5;
  height: 10%;
}

.navbar div {
  height: 100%;
}

.navbar div img {
  height: 100%;
}

</style>
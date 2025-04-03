<template>
    <Nav></Nav>
    <Main1></Main1>
    <router-view></router-view> 
</template>

<script setup>
import Nav from './components/main-page/Nav.vue'
import Main1 from './components/main-page/main/Main1.vue'
import { ref, computed } from 'vue'



</script>

<style scoped>
.div {
  height: 40vh; 
  width: 100%; 
  background-color: #d3d3d3;
  margin-top: 5%; 
}
</style>
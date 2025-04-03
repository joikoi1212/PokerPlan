<script setup>
</script>

<template>
    <div class="info-container">
      <div class="info-text">
        <ul>
          <li>
            <img src="../../../assets/checkmark-xxl.png" alt="Checkmark" class="checkmark-icon" />
            Easy to use for planning poker sessions
          </li>
          <li>
            <img src="../../../assets/checkmark-xxl.png" alt="Checkmark" class="checkmark-icon" />
            Collaborate with your team in real-time
          </li>
          <li>
            <img src="../../../assets/checkmark-xxl.png" alt="Checkmark" class="checkmark-icon" />
            Track and manage your estimates effectively
          </li>
        </ul>
      </div>
    </div>
  </template>
  
  <style scoped>
  .info-container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-top: 20px;
    padding: 20px;
    border-radius: 8px;
    width: 100%;
  }
  
  .info-text {
    flex: 1;
    font-size: 1.2rem;
    line-height: 1.8;
  }
  
  .info-text ul {
    list-style: none; 
    padding: 0;
  }
  
  .info-text li {
    display: flex;
    align-items: center;
    margin-bottom: 10px;
  }
  
  .checkmark-icon {
    width: 24px;
    height: 24px;
    margin-right: 10px; 
  }
  
  .info-image {
    flex: 1;
    display: flex;
    justify-content: center;
  }
  
  .info-image img {
    max-width: 80%;
    height: auto;
    border-radius: 8px;
  }
  </style>
<template>
  <v-app>
    <AppHeader />
    <GamesList />
    
    <v-main class=" fill-height">
      <GameField />
    </v-main>
  </v-app>
</template>

<script setup>
import AppHeader from './components/AppHeader.vue'
import GamesList from './components/GamesList.vue'
import GameField from './components/GameField.vue'
</script>

<style>
html, body {
  overflow: hidden;
  height: 100%;
}

.v-application {
  height: 100vh;
}
</style>

<style scoped>
.border-r {
  border-right: 1px solid rgba(255, 255, 255, 0.12);
}

.game-field {
  display: flex;
  justify-content: center;
  align-items: center;
}

.game-grid {
  display: grid;
  grid-template-columns: repeat(15, 1fr);
  gap: 1px;
  background: rgba(255, 255, 255, 0.12);
  padding: 1px;
}

.grid-cell {
  aspect-ratio: 1;
  background: #1a237e;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style> 
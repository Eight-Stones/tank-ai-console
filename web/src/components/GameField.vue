<template>
  <v-container fluid class="d-flex flex-column fill-height pa-4">
    <!-- Game content -->
    <div class="d-flex flex-grow-1">
      <GameGrid :players="players" :is-running="isRunning" @start="startGame" @stop="stopGame" @add-bot="addBot" />
      <PlayersPanel :players="players" :grid-size="gridSize" />
    </div>
  </v-container>
</template>

<script setup>
import { ref, computed, onUnmounted } from 'vue'
import GameGrid from './GameGrid.vue'
import PlayersPanel from './PlayersPanel.vue'
import { updateTanks } from '../utils/tankMovement'

const gameInterval = ref(null)
const isRunning = ref(false)

const players = ref([])

// Compute grid size based on viewport height
const gridSize = computed(() => {
  const vh = Math.min(window.innerHeight - 200, window.innerWidth - 400)
  return Math.floor(vh)
})

// Методы управления игрой
const startGame = () => {
  if (!isRunning.value) {
    isRunning.value = true
    gameInterval.value = setInterval(() => {
      players.value = updateTanks(players.value)
    }, 1000) // Обновление каждую секунду
  }
}

const stopGame = () => {
  if (isRunning.value) {
    clearInterval(gameInterval.value)
    gameInterval.value = null
    isRunning.value = false
  }
}

const addBot = () => {
  // Пытаемся найти свободную позицию
  let x, y
  do {
    x = Math.floor(Math.random() * 20)
    y = Math.floor(Math.random() * 20)
  } while (players.value.some(p => p.x === x && p.y === y))

  const newBot = {
    name: `Bot${players.value.length + 1}`,
    hp: 20,
    ammo: 10,
    x,
    y,
    direction: ['up', 'down', 'left', 'right'][Math.floor(Math.random() * 4)]
  }
  players.value.push(newBot)
}

// Очистка интервала при размонтировании компонента
onUnmounted(() => {
  if (gameInterval.value) {
    clearInterval(gameInterval.value)
  }
})
</script>

<style scoped>
.game-container {
  min-height: 0;
  margin-right: 16px;
}

.game-field {
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
}

.game-grid {
  display: grid;
  grid-template-columns: repeat(20, 1fr);
  gap: 1px;
  background: rgba(255, 255, 255, 0.12);
  padding: 1px;
  aspect-ratio: 1;
  width: v-bind(gridSize + 'px');
  height: v-bind(gridSize + 'px');
}

.grid-cell {
  aspect-ratio: 1;
  background: #1a237e;
  display: flex;
  align-items: center;
  justify-content: center;
  min-width: 0;
  min-height: 0;
}

.game-controls {
  margin: 0 auto;
}

.players-container {
  height: v-bind(gridSize + 'px');
}

.players-panel {
  width: 100%;
  display: flex;
  flex-direction: column;
}

:deep(.v-table) {
  background-color: rgba(0, 0, 0, 0.2) !important;
  height: 100%;
}

:deep(.v-table__wrapper) {
  height: calc(100% - 49px); /* Вычитаем высоту заголовка */
}

:deep(.v-card-title) {
  padding: 8px 16px;
  font-size: 1rem;
  border-bottom: 1px solid rgba(255, 255, 255, 0.12);
  background-color: rgba(0, 0, 0, 0.2);
}

:deep(.v-table > .v-table__wrapper > table) {
  width: 100%;
  table-layout: fixed;
}

:deep(.v-table th) {
  white-space: nowrap;
  font-weight: 600 !important;
  color: rgb(255, 255, 255, 0.9) !important;
  font-size: 0.875rem;
}

:deep(.v-table td) {
  font-size: 0.875rem;
}

.h-100 {
  height: 100%;
}
</style> 
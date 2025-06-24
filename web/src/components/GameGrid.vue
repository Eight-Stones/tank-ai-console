<template>
  <div class="game-container flex-grow-1 d-flex flex-column">
    <div class="game-field flex-grow-1 d-flex align-center justify-center">
      <div class="game-grid">
        <template v-for="row in 20" :key="row">
          <template v-for="col in 20" :key="col">
            <div class="grid-cell">
              <v-icon v-if="getTankAtPosition(col-1, 20-row)"
                      icon="mdi-tank"
                      :color="getTankColor(col-1, 20-row)"
                      :style="getTankRotation(col-1, 20-row)"
                      size="x-large" />
            </div>
          </template>
        </template>
      </div>
    </div>

    <!-- Controls centered under game field -->
    <div class="game-controls" :style="{ width: gridSize + 'px' }">
      <v-card flat class="mt-4 bg-transparent">
        <v-card-actions class="justify-center">
          <v-btn 
            prepend-icon="mdi-play" 
            color="success" 
            :disabled="isRunning"
            @click="$emit('start')"
          >
            Run
          </v-btn>
          <v-btn 
            prepend-icon="mdi-stop" 
            color="error" 
            class="mx-2"
            :disabled="!isRunning"
            @click="$emit('stop')"
          >
            Stop
          </v-btn>
          <v-btn 
            prepend-icon="mdi-plus"
            @click="$emit('add-bot')"
          >
            Add Bot
          </v-btn>
        </v-card-actions>
      </v-card>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  players: {
    type: Array,
    required: true
  },
  isRunning: {
    type: Boolean,
    required: true
  }
})

defineEmits(['start', 'stop', 'add-bot'])

const getTankAtPosition = (x, y) => {
  return props.players.find(p => p.x === x && p.y === y)
}

const getTankColor = (x, y) => {
  const tank = getTankAtPosition(x, y)
  return tank ? 'primary' : ''
}

const getTankRotation = (x, y) => {
  const tank = getTankAtPosition(x, y)
  if (!tank) return ''
  
  const rotations = {
    'up': 'rotate(0deg)',
    'right': 'rotate(90deg)',
    'down': 'rotate(180deg)',
    'left': 'rotate(270deg)'
  }
  
  return `transform: ${rotations[tank.direction]}`
}

// Compute grid size based on viewport height
const gridSize = computed(() => {
  const vh = Math.min(window.innerHeight - 200, window.innerWidth - 400)
  return Math.floor(vh)
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
</style> 
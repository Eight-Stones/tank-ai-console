<template>
  <div class="players-container d-flex align-center" style="width: 350px">
    <v-card class="players-panel h-100" style="margin-top: 35px">
      <v-card-title class="text-subtitle-1 py-2">Players</v-card-title>
      <v-table density="compact" fixed-header height="100%">
        <thead>
          <tr>
            <th class="text-center" style="width: 25%">name</th>
            <th class="text-center" style="width: 15%">hp</th>
            <th class="text-center" style="width: 15%">ammo</th>
            <th class="text-center" style="width: 15%">x</th>
            <th class="text-center" style="width: 15%">y</th>
            <th class="text-center" style="width: 20%">dir</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="player in players" :key="player.name">
            <td class="text-center">{{ player.name }}</td>
            <td class="text-center">{{ player.hp }}</td>
            <td class="text-center">{{ player.ammo }}</td>
            <td class="text-center">{{ player.x }}</td>
            <td class="text-center">{{ player.y }}</td>
            <td class="text-center">{{ player.direction }}</td>
          </tr>
        </tbody>
      </v-table>
    </v-card>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  players: {
    type: Array,
    required: true
  },
  gridSize: {
    type: Number,
    required: true
  }
})
</script>

<style scoped>
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
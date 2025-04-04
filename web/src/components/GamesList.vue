<template>
  <v-navigation-drawer permanent width="300" class="fill-height">
    <v-list class="fill-height">
      <v-list-subheader>Games</v-list-subheader>
      <v-list-item
        v-for="game in games"
        :key="game.id"
        :value="game"
        active-color="primary"
        class="text-truncate"
      >
        <v-list-item-title class="text-truncate">{{ game.id }}</v-list-item-title>
        <v-list-item-subtitle>Created</v-list-item-subtitle>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>

<script setup>
import { ref } from 'vue'

const games = ref([
  { id: '5742a9d6-9296-4f1d-97AB-2bdf56103468' },
  { id: '1b227f18-b882-4aa0-88c1-265329992058' },
  { id: 'c54eb552-4ab0-47fb-8d90-959550c68f57' }
])
</script>

<style scoped>
.v-navigation-drawer {
  top: 64px !important; /* Отступ под AppBar */
  max-height: calc(100vh - 64px);
}
</style> 
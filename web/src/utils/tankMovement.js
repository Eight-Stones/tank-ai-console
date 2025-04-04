const GRID_SIZE = 20
const DIRECTIONS = ['up', 'right', 'down', 'left']

// Получение следующей позиции танка на основе его текущего положения и направления
export const getNextPosition = (x, y, direction) => {
  switch (direction) {
    case 'up':
      return { x, y: y - 1 }  // Движение вверх уменьшает y
    case 'down':
      return { x, y: y + 1 }  // Движение вниз увеличивает y
    case 'left':
      return { x: x - 1, y }
    case 'right':
      return { x: x + 1, y }
    default:
      return { x, y }
  }
}

// Проверка, находится ли позиция в пределах карты
export const isValidPosition = (x, y) => {
  return x >= 0 && x < GRID_SIZE && y >= 0 && y < GRID_SIZE
}

// Проверка на столкновение с другими танками
export const checkCollision = (x, y, players) => {
  return players.some(player => player.x === x && player.y === y)
}

// Получение случайного направления
export const getRandomDirection = () => {
  return DIRECTIONS[Math.floor(Math.random() * DIRECTIONS.length)]
}

// Основная функция движения танка
export const moveTank = (tank, players) => {
  const nextPos = getNextPosition(tank.x, tank.y, tank.direction)
  
  // Если следующая позиция выходит за пределы карты или там есть другой танк
  if (!isValidPosition(nextPos.x, nextPos.y) || checkCollision(nextPos.x, nextPos.y, players)) {
    // Меняем направление на случайное
    return {
      ...tank,
      direction: getRandomDirection()
    }
  }
  
  // Иначе двигаемся в выбранном направлении
  return {
    ...tank,
    x: nextPos.x,
    y: nextPos.y
  }
}

// Обновление состояния всех танков
export const updateTanks = (players) => {
  return players.map(tank => moveTank(tank, players.filter(p => p !== tank)))
} 
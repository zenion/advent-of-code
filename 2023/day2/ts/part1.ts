const file = Bun.file('../input.txt')
const lines = (await file.text()).split('\n')

interface Set {
  red: number
  green: number
  blue: number
}

interface Game {
  id: number
  sets: Set[]
}

const games: Game[] = []

for (const line of lines) {
  const [idRaw, setsRaw] = line.split(':')

  const id = /\d+/.exec(idRaw)![0]
  const setsRaw2 = setsRaw.split(';')
  const sets: Set[] = []

  for (const setRaw of setsRaw2) {
    const redMatch = setRaw.match(/(\d+) red/)
    const greenMatch = setRaw.match(/(\d+) green/)
    const blueMatch = setRaw.match(/(\d+) blue/)
    let red = 0
    let green = 0
    let blue = 0

    if (redMatch) red = Number(redMatch[1])
    if (greenMatch) green = Number(greenMatch[1])
    if (blueMatch) blue = Number(blueMatch[1])

    sets.push({ red, green, blue })
  }

  games.push({ id: Number(id), sets })
}

const maxRed = 12
const maxGreen = 13
const maxBlue = 14

let gameIdSum = 0

for (const game of games) {
  let badGame = false
  for (const set of game.sets) {
    if (set.red > maxRed || set.green > maxGreen || set.blue > maxBlue) {
      badGame = true
      break
    }
  }
  if (!badGame) {
    gameIdSum += game.id
  }
}

console.log(gameIdSum)

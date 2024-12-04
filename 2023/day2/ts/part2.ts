const file = Bun.file('../input.txt')
const lines = (await file.text()).split('\n')

interface Set {
  red: number
  green: number
  blue: number
}

interface Game {
  id: number
  power: number
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

  games.push({ id: Number(id), sets, power: 0 })
}

for (const game of games) {
  let red = 0
  let green = 0
  let blue = 0
  for (const set of game.sets) {
    // find highest red green and blue values
    if (set.red > red) red = set.red
    if (set.green > green) green = set.green
    if (set.blue > blue) blue = set.blue
  }
  game.power = red * green * blue
}

console.log(games.reduce((acc, game) => acc + game.power, 0))

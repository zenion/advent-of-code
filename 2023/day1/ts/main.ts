const file = Deno.readTextFileSync('../input.txt')
const lines = file.split('\n').filter((line) => line.trim() !== '')

interface Match {
  index: number
  digit: string
  type: 'digit' | 'word'
}

const lookup: { [key: string]: string } = {
  one: '1',
  two: '2',
  three: '3',
  four: '4',
  five: '5',
  six: '6',
  seven: '7',
  eight: '8',
  nine: '9',
}

let sum = 0

for (const line of lines) {
  const digits: Match[] = []

  // find all digits in the line
  for (let i = 0; i < line.length; i++) {
    const char = line[i]
    if (char >= '0' && char <= '9') {
      digits.push({ index: i, digit: char, type: 'digit' })
    }
  }

  // find first and last unique spelled numbers in the line
  for (const key in lookup) {
    const firstMatch = line.indexOf(key)
    if (firstMatch !== -1) {
      digits.push({ index: firstMatch, digit: lookup[key], type: 'word' })
    }

    const lastMatch = line.lastIndexOf(key)
    if (lastMatch !== -1 && lastMatch !== firstMatch) {
      digits.push({ index: lastMatch, digit: lookup[key], type: 'word' })
    }
  }

  // sort digits by index
  digits.sort((a, b) => a.index - b.index)

  // concat first and last digits
  const num = Number(digits[0].digit + digits[digits.length - 1].digit)

  sum += num
}

console.log(sum)

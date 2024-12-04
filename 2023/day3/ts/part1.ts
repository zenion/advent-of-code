const file = Bun.file('../input.txt')
const lines = (await file.text()).split('\n')

function charIsSymbol(char: string) {
  return /[`!@#$%^&*()_+\-=\[\]{};':"\\|,<>\/?~]/.test(char)
}

let nums: number[] = []

for (const i in lines) {
  const line = lines[i]
  const match = line.matchAll(/\d+/g)

  for (const m of match) {
    const startIdx = m.index!
    const nextCharIdx = startIdx + m[0].length

    const prevChar = line[startIdx - 1]
    const nextChar = line[nextCharIdx]

    if (charIsSymbol(prevChar) || charIsSymbol(nextChar)) {
      nums = [...nums, Number(m[0])]
    }

    const prevLine = lines[Number(i) - 1]
    if (prevLine) {
      for (let j = startIdx - 1; j <= nextCharIdx; j++) {
        const char = prevLine[j]
        if (charIsSymbol(char)) {
          nums = [...nums, Number(m[0])]
        }
      }
    }

    const nextLine = lines[Number(i) + 1]
    if (nextLine) {
      for (let j = startIdx - 1; j <= nextCharIdx; j++) {
        const char = nextLine[j]
        if (charIsSymbol(char)) {
          nums = [...nums, Number(m[0])]
        }
      }
    }
  }
}

console.log(nums.reduce((a, b) => a + b, 0))

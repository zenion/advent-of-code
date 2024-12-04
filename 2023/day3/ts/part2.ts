const file = Bun.file('../input.txt')
const lines = (await file.text()).split('\n')

function charIsSymbol(char: string) {
  return /[`!@#$%^&*()_+\-=\[\]{};':"\\|,<>\/?~]/.test(char)
}

interface Num {
  value: number
  lineIdx: number
  startIdx: number
  endIdx: number
}

interface StarSym {
  line: number
  idx: number
  adjNums?: Num[]
}

let nums: Num[] = []
let starSyms: StarSym[] = []

for (const i in lines) {
  const line = lines[i]

  const matchStar = line.matchAll(/\*/g)

  for (const m of matchStar) {
    starSyms = [
      ...starSyms,
      {
        line: Number(i),
        idx: m.index!,
      },
    ]
  }

  const match = line.matchAll(/\d+/g)

  for (const m of match) {
    const startIdx = m.index!
    const nextCharIdx = startIdx + m[0].length

    const prevChar = line[startIdx - 1]
    const nextChar = line[nextCharIdx]

    if (charIsSymbol(prevChar) || charIsSymbol(nextChar)) {
      nums = [
        ...nums,
        {
          value: Number(m[0]),
          lineIdx: Number(i),
          startIdx,
          endIdx: nextCharIdx,
        },
      ]
    }

    const prevLine = lines[Number(i) - 1]
    if (prevLine) {
      for (let j = startIdx - 1; j <= nextCharIdx; j++) {
        const char = prevLine[j]
        if (charIsSymbol(char)) {
          nums = [
            ...nums,
            {
              value: Number(m[0]),
              lineIdx: Number(i),
              startIdx,
              endIdx: nextCharIdx,
            },
          ]
        }
      }
    }

    const nextLine = lines[Number(i) + 1]
    if (nextLine) {
      for (let j = startIdx - 1; j <= nextCharIdx; j++) {
        const char = nextLine[j]
        if (charIsSymbol(char)) {
          nums = [
            ...nums,
            {
              value: Number(m[0]),
              lineIdx: Number(i),
              startIdx,
              endIdx: nextCharIdx,
            },
          ]
        }
      }
    }
  }
}

for (const star of starSyms) {
  let adjNums: Num[] = []

  for (const num of nums) {
    if (num.lineIdx === star.line && (num.startIdx - 1 === star.idx || num.endIdx === star.idx)) {
      adjNums = [...adjNums, num]
    } else if (star.line - 1 === num.lineIdx && star.idx >= num.startIdx - 1 && star.idx <= num.endIdx) {
      adjNums = [...adjNums, num]
    } else if (star.line + 1 === num.lineIdx && star.idx >= num.startIdx - 1 && star.idx <= num.endIdx) {
      adjNums = [...adjNums, num]
    }
  }

  star.adjNums = adjNums
}

const filteredStarSyms: StarSym[] = starSyms.filter((star) => star.adjNums!.length === 2)

let sum = 0
for (const star of filteredStarSyms) {
  sum += star.adjNums![0].value * star.adjNums![1].value
}

console.log(sum)

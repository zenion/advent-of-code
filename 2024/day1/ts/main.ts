import { assertEquals } from '@std/assert'

type Columns = {
  left: number[]
  right: number[]
}

function parseColumns(fileContents: string): Columns {
  const [left, right] = fileContents
    .trim()
    .split('\n')
    .filter((line) => line.trim() !== '')
    .reduce<[number[], number[]]>(
      ([leftCol, rightCol], line) => {
        const [leftNum, rightNum] = line.split(/\s+/).map(Number)
        return [leftCol.concat(leftNum), rightCol.concat(rightNum)]
      },
      [[], []],
    )
    .map((arr) => arr.toSorted())

  return { left, right }
}

function solvePart1(fileContents: string): number {
  const { left, right } = parseColumns(fileContents)

  return left.reduce((acc, curr, i) => {
    return acc + Math.abs(curr - right[i])
  }, 0)
}

function solvePart2(fileContents: string): number {
  const { left, right } = parseColumns(fileContents)

  return left.reduce((acc, leftNum) => {
    const countInRight = right.filter((rightNum) => rightNum === leftNum).length
    return acc + leftNum * countInRight
  }, 0)
}

if (import.meta.main) {
  const fileContents = Deno.readTextFileSync('../input.txt')
  console.log(`Part 1: ${solvePart1(fileContents)}`)
  console.log(`Part 2: ${solvePart2(fileContents)}`)
}

Deno.test('test', async (t) => {
  const contents = `3   4
4   3
2   5
1   3
3   9
3   3
`
  await t.step('part 1', () => {
    assertEquals(solvePart1(contents), 11)
  })
  await t.step('part 2', () => {
    assertEquals(solvePart2(contents), 31)
  })
})

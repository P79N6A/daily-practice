const println = console.log

class FixedSizeArray {
  constructor(m) {
    this.capacity = m
    this.array = new Array(m)
  }

  store(elem) {
    const nullElemIndex = this.array.findIndex(elem => !!!elem)
    if (nullElemIndex === -1) {
      for (let i = 0; i < this.capacity - 1; i++) {
        this.array[i] = this.array[i+1]
      }
      this.array[this.capacity-1] = elem
    } else {
      this.array[nullElemIndex] = elem
    }
  }
}

function minSubArray(array, m, sizedArray) {
  let i, l = array.length, initSum = 0, minSum
  for (let i = 0; i < m && i < l; i++) {
    sizedArray.store(array[i])
    initSum += array[i]
  }
  if (m > l) {
    return initSum
  } else {
    minSum = initSum
    for (let i = m; i < l; i++) {
      sizedArray.store(array[i])
      initSum = initSum + array[i] - array[i-m]
      if (initSum < minSum) {
        minSum = initSum
      }
    }
    return minSum
  }
}

const array = [61, 12, 13, 15, 22, -1, 4]
const m = 3
const sizedArray = new FixedSizeArray(m)
println("Array:", array)
println(minSubArray(array, m, sizedArray))
println(sizedArray)

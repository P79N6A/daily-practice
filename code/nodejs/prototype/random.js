const println = console.log

function RandomWord(chars) {
  if (!(this instanceof RandomWord)) {
    println('NEW')
    return new RandomWord(chars)
  }
  this.chars = chars
}

const r = RandomWord('hello')
println(r)
println(r instanceof RandomWord)

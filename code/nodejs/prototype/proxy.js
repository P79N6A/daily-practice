const println = console.log

function tryGet(obj, pth) {
  const pths = pth.split('.')
  const curr = pths.shift()
  return tryGetProp(obj, curr, pths)
}

function tryGetProp(obj, curr, pths) {
  const prop = obj[curr]
  if (prop) {
     if (pths.length === 0) {
       return prop
     } else {
       const curr = pths.shift()
       return tryGetProp(prop, curr, pths)
     }
  }
  return null
}

const obj = { a: { b : { c: 'd'}}}
println(tryGet(obj, 'a.b.c'))

const println = console.log

class Person {
  constructor(name) {
    this.name = name
  }

  set Name(value) {
    this.name = value
  }

  get Name() {
    return this.name
  }
}

const p = new Person("Jack")
println(p.Name)
p.Name = "Tom"
println(p.Name)

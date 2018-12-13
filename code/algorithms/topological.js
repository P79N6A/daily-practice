const dc = require('./directedCycle')
const dfo = require('./depthFirstOrder')

class Topological {
    constructor(g) {
        const cyclefinder = new dc.DirectedCycle(g)
        if (cyclefinder.hasCycle()) {
            const dfs = new dfo.DepthFirstOrder(g)
            this.order = dfs.ReversePost()
        }
    }

    Order () { return this.order }

    IsDAC() { return !!this.order }
}

module.exports = {
    Topological,
}

global.println = console.log
global.error = console.error

const graph = require('./graph')
const diGraph = require('./digraph')
const depthSearchFirst = require('./depthSearchFirst')
const breadthFirstSearch = require('./breadthSearchFirst')
const directedDFS = require('./directedDFS')
const topological = require('./topological')
const cc = require('./CC')
const kosarajuSCC = require('./kosarajuSCC')

const g = graph.constructGraph([
    [0, 5],
    [4, 3],
    [0, 1],
    [9, 12],
    [6, 4],
    [5, 4],
    [0, 2],
    [11, 12],
    [9, 10],
    [0, 6],
    [7, 8],
    [9, 11],
    [5, 3],
], 13)
const dfs = new depthSearchFirst.DepthFirstSearch(g)
const CC = new cc.CC(g).components()
println(g.toString())
println(`${dfs.connected(0) ? '是' : '不是'}连通图`)
println(`连通分量：${CC.length}`)
println(`是否有环：${dfs.cycle() ? '是' : '否'}`)

let v = 0, w = 6
const pathway = dfs.pathway(v, w)
println(`深度遍历：${v} 与 ${w} 的通路： [${pathway  || '无'}]`)

let x = 0, y = 3
const bfs = new breadthFirstSearch.BreadthFirstSearch(g)
const path = bfs.pathTo(x, y)
println(`广度遍历：${x} 与 ${y} 的通路： [${path  || '无'}]`)

const dg = diGraph.constructDiGraph([
    [4, 2],
    [2, 3],
    [3, 2],
    [6, 0],
    [0, 1],
    [2, 0],
    [11, 12],
    [12, 9],
    [9, 10],
    [9, 11],
    [8, 9],
    [10, 12],
    [11, 4],
    [4, 3],
    [3, 5],
    [7, 8],
    [8, 7],
    [5, 4],
    [0, 5],
    [6, 4],
    [6, 9],
    [7, 6],
], 13)
println(`有向图：\n${dg.toString()}`)
const ddfs = new directedDFS.DirectedDFS(dg)
println(`是否有环：${ddfs.cycle() ? '是' : '否'}`)
println(`有向图反向：\n${dg.reverse().toString()}`)

const tp = new topological.Topological(dg)
println(`逆后序排列：\n${tp.Order().toString()}`)
const kcc = new kosarajuSCC.KosarajuSCC(dg)
v = 11, w = 9
println(`${v} 和 ${w} 是否为强连接：${kcc.stronglyConnected(v, w) ? '是' : '否'}`)
println(`强连通分量：${kcc.components().length}`)

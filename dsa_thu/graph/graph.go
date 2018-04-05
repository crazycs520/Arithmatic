package graph

import (
	"math"
	"time"
)

type VertexStatu int
type EdgeStatus int

const (
	UNDISCOVERED VertexStatu = 0 //未发现状态
	DISCOVERED   VertexStatu = 1 //发现状态
	VISITED      VertexStatu = 2 //访问过状态

	UNDETERMINED EdgeStatus = 0
	TREE         EdgeStatus = 1
	CROSS        EdgeStatus = 2
	FORWAED      EdgeStatus = 3
	BACKWAED     EdgeStatus = 4
)

//定点对象
type Vertex struct {
	data         interface{}
	inDegree     int
	outDegree    int
	status       VertexStatu
	dTime, fTime time.Duration //时间标签
	parent       int           //遍历树中的树节点
	priority     int           //遍历树中的优先级
}

//边对象
type Edge struct {
	data   interface{} //数据
	weight int         //权重
	status EdgeStatus  //类型
}

//邻节矩阵实现
type GraphMatrix struct {
	v []Vertex //顶点数组
	e [][]Edge //边(二维)数组
}

func NewVertex(data interface{}) Vertex {
	v := Vertex{}
	v.data = data
	v.inDegree = 0
	v.outDegree = 0
	v.status = UNDISCOVERED //未发现
	v.parent = -1
	v.priority = math.MaxInt64
	return v
}

func NewEdge(data interface{}, weight int) Edge {
	e := Edge{}
	e.data = data
	e.weight = weight
	e.status = UNDETERMINED
	return e
}

//判断顶点到另一顶点的边是否存在
func (g *GraphMatrix) exists(i, j int) bool {
	return g.e[i][j] != Edge{}
}

//返回下一个邻居
func (g *GraphMatrix) nextNbr(i, j int) int {
	for -1 < j && !g.exists(i, j) {
		j--
	}
	return j
}

//返回第一个邻居
func (g *GraphMatrix) firstNbr(i, j int) int {
	return g.nextNbr(i, len(g.v)-1)
}

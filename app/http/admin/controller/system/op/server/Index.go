package server

import (
	"github.com/gin-gonic/gin"
	"go_admin/app/res"
	"runtime"
)

func Index(c *gin.Context) {
	numCPUs := runtime.NumCPU()
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)
	res.Json(c, res.Data(gin.H{
		"num_cpus":     numCPUs,
		"memory_Total": (m.TotalAlloc) / 1024 / 1024,
		"memory_Usage": (m.Alloc) / 1024 / 1024,
	}))
	return
}

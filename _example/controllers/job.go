package controllers

import (
	"encoding/json"
	"net/http"

	dbpkg "github.com/vharish836/apig/_example/db"
	"github.com/vharish836/apig/_example/helper"
	"github.com/vharish836/apig/_example/models"
	"github.com/vharish836/apig/_example/version"

	"github.com/gin-gonic/gin"
)

// GetJobs ...
func GetJobs(c *gin.Context) {
	ver, err := verpkg.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.Instance(c)
	parameter, err := dbpkg.NewParameter(c, models.Job{})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db, err = parameter.Paginate(db)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db = parameter.SetPreloads(db)
	db = parameter.SortRecords(db)
	db = parameter.FilterFields(db)
	jobs := []models.Job{}
	fields := helper.ParseFields(c.DefaultQuery("fields", "*"))
	queryFields := helper.QueryFields(models.Job{}, fields)

	if err := db.Select(queryFields).Find(&jobs).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	index := 0

	if len(jobs) > 0 {
		index = int(jobs[len(jobs)-1].ID)
	}

	if err := parameter.SetHeaderLink(c, index); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if verpkg.Range("1.0.0", "<=", ver) && verpkg.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	if _, ok := c.GetQuery("stream"); ok {
		enc := json.NewEncoder(c.Writer)
		c.Status(200)

		for _, job := range jobs {
			fieldMap, err := helper.FieldToMap(job, fields)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}

			if err := enc.Encode(fieldMap); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
		}
	} else {
		fieldMaps := []map[string]interface{}{}

		for _, job := range jobs {
			fieldMap, err := helper.FieldToMap(job, fields)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}

			fieldMaps = append(fieldMaps, fieldMap)
		}

		if _, ok := c.GetQuery("pretty"); ok {
			c.IndentedJSON(200, fieldMaps)
		} else {
			c.JSON(200, fieldMaps)
		}
	}
}

// GetJob ...
func GetJob(c *gin.Context) {
	ver, err := verpkg.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.Instance(c)
	parameter, err := dbpkg.NewParameter(c, models.Job{})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db = parameter.SetPreloads(db)
	job := models.Job{}
	id := c.Params.ByName("id")
	fields := helper.ParseFields(c.DefaultQuery("fields", "*"))
	queryFields := helper.QueryFields(models.Job{}, fields)

	if err := db.Select(queryFields).First(&job, id).Error; err != nil {
		content := gin.H{"error": "job with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	fieldMap, err := helper.FieldToMap(job, fields)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if verpkg.Range("1.0.0", "<=", ver) && verpkg.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	if _, ok := c.GetQuery("pretty"); ok {
		c.IndentedJSON(200, fieldMap)
	} else {
		c.JSON(200, fieldMap)
	}
}

// CreateJob ...
func CreateJob(c *gin.Context) {
	ver, err := verpkg.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.Instance(c)
	job := models.Job{}

	if err := c.Bind(&job); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&job).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if verpkg.Range("1.0.0", "<=", ver) && verpkg.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(201, job)
}

// UpdateJob ...
func UpdateJob(c *gin.Context) {
	ver, err := verpkg.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.Instance(c)
	id := c.Params.ByName("id")
	job := models.Job{}

	if db.First(&job, id).Error != nil {
		content := gin.H{"error": "job with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	if err := c.Bind(&job); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&job).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if verpkg.Range("1.0.0", "<=", ver) && verpkg.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(200, job)
}

// DeleteJob ...
func DeleteJob(c *gin.Context) {
	ver, err := verpkg.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.Instance(c)
	id := c.Params.ByName("id")
	job := models.Job{}

	if db.First(&job, id).Error != nil {
		content := gin.H{"error": "job with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	if err := db.Delete(&job).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if verpkg.Range("1.0.0", "<=", ver) && verpkg.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}

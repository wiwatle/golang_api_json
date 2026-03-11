package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Task โครงสร้างข้อมูลของเรา
type Group struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status_use"`
}

const fileName = "data.json"

// ฟังก์ชันช่วยอ่านข้อมูลจาก JSON File
func readGroups() ([]Group, error) {
	var groups []Group
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(file, &groups)
	return groups, nil
}

// ฟังก์ชันช่วยเขียนข้อมูลลง JSON File
func writeGroups(groups []Group) error {
	data, err := json.MarshalIndent(groups, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func main() {
	r := gin.Default()

	// GET: ดึงข้อมูลทั้งหมด
	r.GET("/groups", func(c *gin.Context) {
		groups, _ := readGroups()
		c.JSON(http.StatusOK, groups)
	})

	// POST: เพิ่มข้อมูลใหม่
	r.POST("/groups", func(c *gin.Context) {
		var newGroup Group
		if err := c.ShouldBindJSON(&newGroup); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		groups, _ := readGroups()
		newGroup.ID = len(groups) + 1 // แบบง่าย: ใช้ความยาว slice กำหนด ID
		groups = append(groups, newGroup)

		writeGroups(groups)
		c.JSON(http.StatusCreated, newGroup)
	})

	// PUT: แก้ไขข้อมูลตาม ID
	r.PUT("/groups/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var updatedGroup Group
		c.ShouldBindJSON(&updatedGroup)

		groups, _ := readGroups()
		for i, t := range groups {
			if t.ID == id {
				groups[i].Name = updatedGroup.Name
				groups[i].Status = updatedGroup.Status
				writeGroups(groups)
				c.JSON(http.StatusOK, groups[i])
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Group not found"})
	})

	// DELETE: ลบข้อมูล
	r.DELETE("/groups/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		groups, _ := readGroups()

		newGroups := []Group{}
		for _, t := range groups {
			if t.ID != id {
				newGroups = append(newGroups, t)
			}
		}

		writeGroups(newGroups)
		c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
	})

	r.Run(":8080")
}

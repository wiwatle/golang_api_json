package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

// นี่คือการประกาศใช้งาน fmt

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

	// เพิ่ม Route หน้าแรก
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Go API is live on Azure!",
		})
	})

	r.Run(":8080") // หรือระบุ Port ที่ต้องการ
}

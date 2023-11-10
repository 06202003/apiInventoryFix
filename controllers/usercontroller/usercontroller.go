package usercontroller

import (
    "net/http"
    "github.com/06202003/apiInventory/models"
    "github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
    var users []models.User
    models.DB.Find(&users)
    c.JSON(http.StatusOK, gin.H{"users": users})
}

func Show(c *gin.Context) {
    var user models.User
    id := c.Param("id")

    if err := models.DB.First(&user, id).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
}

func Create(c *gin.Context) {
    var user models.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    models.DB.Create(&user)
    c.JSON(http.StatusOK, gin.H{"user": user})
}

func Update(c *gin.Context) {
    var user models.User
    id := c.Param("id")

    if err := c.ShouldBindJSON(&user); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    if models.DB.Model(&user).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Diperbaharui"})
}

func Delete(c *gin.Context) {
    var user models.User
    id := c.Param("id")

    if err := models.DB.First(&user, id).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User tidak ditemukan"})
        return
    }

    if err := models.DB.Delete(&user).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}

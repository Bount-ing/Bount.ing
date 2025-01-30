package controllers

import (
	"github.com/bount-ing/bount.ing/api/db"
	"github.com/bount-ing/bount.ing/api/models"
)

func CreateHost(host models.Host) error {
	err := db.DB.Create(&host)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func GetHost(hostID string) (models.Host, error) {
	var host models.Host
	err := db.DB.First(&host, hostID)
	if err.Error != nil {
		return host, err.Error
	}
	return host, nil
}

func GetHostFromAddress(hostAddress string) (models.Host, error) {
	var host models.Host
	err := db.DB.Where("address = ?", hostAddress).First(&host)
	if err.Error != nil {
		return host, err.Error
	}
	return host, nil
}

func GetHosts() ([]models.Host, error) {
	var hosts []models.Host
	err := db.DB.Find(&hosts)
	if err.Error != nil {
		return hosts, err.Error
	}
	return hosts, nil
}

func UpdateHost(host models.Host) error {
	err := db.DB.Save(&host)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func DeleteHost(hostID string) error {
	var host models.Host
	err := db.DB.First(&host, hostID)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

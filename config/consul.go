package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/hashicorp/consul/api"
)

func RegisterServiceWithConsul() {
	fmt.Println("Registering service with Consul...")
	// Konfigurasi Consul
	consulAddress := os.Getenv("CONSUL_ADDRESS") // Contoh: "127.0.0.1:8500"
	serviceID := os.Getenv("SERVICE_ID")         // Contoh: "go-api-1"
	serviceName := os.Getenv("SERVICE_NAME")     // Contoh: "go-api"
	servicePort := os.Getenv("SERVICE_PORT")     // Contoh: "8080"
	serviceHost := os.Getenv("SERVICE_HOST")     // Contoh: "127.0.0.1"

	config := api.DefaultConfig()
	config.Address = consulAddress

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Gagal membuat client Consul: %v", err)
	}

	// Konfigurasi Service Registrasi
	registration := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Port:    atoi(servicePort),
		Address: serviceHost,
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%s/health", serviceHost, servicePort),
			Interval: "10s",
			Timeout:  "5s",
		},
	}

	// Daftarkan service ke Consul
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("Gagal mendaftarkan service ke Consul: %v", err)
	}

	fmt.Println("Service berhasil terdaftar di Consul!")
}

// Fungsi untuk konversi string ke int
func atoi(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Gagal mengkonversi string ke int: %v", err)
	}
	return value
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func StorePost(post string) error {
	if err := os.MkdirAll("posts", 0744); err != nil {
		return err
	}

	// Obtener la fecha actual en formato YYYY-MM-DD
	dateStr := time.Now().Format("2006-01-02")

	// Crear el nombre del directorio con la fecha
	dirPath := fmt.Sprintf("posts/%s", dateStr)

	// Crear el directorio con la fecha
	if err := os.MkdirAll(dirPath, 0744); err != nil {
		return err
	}

	// Encontrar el siguiente número de archivo disponible
	nextId, err := getNextAvailableId(dirPath)
	if err != nil {
		return err
	}

	// Crear el nombre del archivo con el siguiente ID
	filePath := fmt.Sprintf("%s/%d.txt", dirPath, nextId)

	// Escribir el contenido del post en el archivo
	if err := os.WriteFile(filePath, []byte(post), 0644); err != nil {
		return err
	}

	return nil
}

func getNextAvailableId(dirPath string) (int, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		// Si el directorio está vacío o no existe, devolver 1
		if os.IsNotExist(err) {
			return 1, nil
		}
		return 0, err
	}

	maxId := 0
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".txt") {
			filename := strings.TrimSuffix(entry.Name(), ".txt")
			id, err := strconv.Atoi(filename)
			if err == nil && id > maxId {
				maxId = id
			}
		}
	}

	return maxId + 1, nil
}

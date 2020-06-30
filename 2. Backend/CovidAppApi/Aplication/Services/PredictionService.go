package services

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	model "../../domain/entities"
	service "../interfaces"
)

type predict struct {
}

type distances struct {
	distance float64
	sick     int
}

var arrayOfData = make([]model.Prediction, 0, 1000)

// PredictionService Implemantation
func PredictionService() service.IPredictionService {
	return &predict{}
}

func (*predict) Predict(prediction model.Prediction) (bool, error) {
	dataBuffer := make(chan model.Prediction, 1000)
	csvfile, err := os.Open("covidPeruDataSet.csv")
	if err != nil {
		log.Fatalln("No se pudo abrir el archivo", err)
	}
	r := csv.NewReader(csvfile)
	for i := 0; i < 1000; i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		noCommasRow := strings.Split(record[0], ";")
		go load(noCommasRow, dataBuffer)
	}
	for i := 0; i < 1000; i++ {
		arrayOfData = append(arrayOfData, <-dataBuffer)
	}
	fmt.Println("Antes")
	result := predictGo(prediction)
	fmt.Println("Despues")
	isSick := true
	if result == 1 {
		isSick = true
	} else {
		isSick = false
	}

	return isSick, nil
}

func load(noCommasRow []string, dataBuffer chan model.Prediction) {
	edad, _ := strconv.Atoi(noCommasRow[0])
	peso, _ := strconv.ParseFloat(noCommasRow[1], 2)
	tos, _ := strconv.Atoi(noCommasRow[3])
	fiebre, _ := strconv.Atoi(noCommasRow[4])
	dificultadRespirar, _ := strconv.Atoi(noCommasRow[5])
	perdidaOlfato, _ := strconv.Atoi(noCommasRow[6])
	enfermo, _ := strconv.Atoi(noCommasRow[7])
	distrito := 0
	switch noCommasRow[2] {
	case "Callao":
		distrito = 1
	case "Ventanilla":
		distrito = 2
	case "Ate":
		distrito = 3
	case "Barranco":
		distrito = 4
	case "Chorrillos":
		distrito = 5
	case "Comas":
		distrito = 6
	case "Jesus Maria":
		distrito = 7
	case "La Molina":
		distrito = 8
	case "La Victoria":
		distrito = 9
	case "Lince":
		distrito = 10
	case "Los Olivos":
		distrito = 11
	case "Lurin":
		distrito = 12
	case "Magdalena del Mar":
		distrito = 13
	case "Miraflores":
		distrito = 14
	case "Pueblo Libre":
		distrito = 15
	case "Puente Piedra":
		distrito = 16
	case "Rimac":
		distrito = 17
	case "San Borja":
		distrito = 18
	case "San Isidro":
		distrito = 19
	case "San Juan de Lurigancho":
		distrito = 20
	case "San Martin de Porres":
		distrito = 21
	case "San Miguel":
		distrito = 22
	case "Santiago de Surco":
		distrito = 23
	case "Surquillo":
		distrito = 24
	case "Villa El Salvador":
		distrito = 25
	default:
		distrito = 0
	}
	data := model.Prediction{edad, peso, distrito, tos, fiebre, dificultadRespirar, perdidaOlfato, enfermo}
	dataBuffer <- data
}

func predictGo(prediction model.Prediction) int {
	distArray := getAllDistances(prediction)
	minDists := findClosestGroups(distArray, 5)
	fmt.Println(minDists)
	score := make([]int, 6, 6)
	for i := 0; i < len(score); i++ {
		score[i] = 0
	}
	fmt.Println(score)
	for i := 0; i < len(minDists); i++ {
		score[minDists[i].sick] = score[minDists[i].sick] + 1
	}
	fmt.Println(score)
	quality := 1
	maxi := score[1]
	for i := 1; i < len(score); i++ {
		if score[i] >= maxi {
			maxi = score[i]
			quality = i
		}
	}
	return quality
}

func getAllDistances(predict model.Prediction) []distances {
	distArray := make([]distances, 0, len(arrayOfData))
	distanceBuffer := make(chan distances)
	for i := 0; i < len(arrayOfData); i++ {
		go calculateDistance(predict, arrayOfData[i], distanceBuffer)
	}
	for i := 0; i < len(arrayOfData); i++ {
		fmt.Println(i)
		distArray = append(distArray, <-distanceBuffer)
	}
	return distArray
}

func calculateDistance(predict model.Prediction, train model.Prediction, distanceBuffer chan distances) {
	result1 := math.Pow((float64(predict.Edad) - float64(train.Edad)), 2)
	result2 := math.Pow((predict.Peso - train.Peso), 2)
	result3 := math.Pow((float64(predict.Distrito) - float64(train.Distrito)), 2)
	result4 := math.Pow((float64(predict.Tos) - float64(train.Tos)), 2)
	result5 := math.Pow((float64(predict.Fiebre) - float64(train.Fiebre)), 2)
	result6 := math.Pow((float64(predict.DificultadRespirar) - float64(train.DificultadRespirar)), 2)
	result7 := math.Pow((float64(predict.PerdidaOlfato) - float64(train.PerdidaOlfato)), 2)
	result := math.Sqrt(result1 + result2 + result3 + +result4 + result5 + result6 + result7)

	distStruct := distances{result, train.Enfermo}
	distanceBuffer <- distStruct
}

func bubbleSort(dists []distances) []distances {
	for i := 0; i < len(dists)-1; i++ {
		for j := 0; j < len(dists)-i-1; j++ {
			if dists[j].distance > dists[j+1].distance {
				aux := dists[j]
				dists[j] = dists[j+1]
				dists[j+1] = aux
			}
		}
	}
	return dists
}

func findClosestGroups(dists []distances, k int) []distances {
	sortedDists := bubbleSort(dists)
	minDists := make([]distances, 0, k)
	for j := 0; j < k; j++ {
		minDists = append(minDists, sortedDists[j])
	}
	return minDists
}

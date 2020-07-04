package services

import (
	"encoding/csv"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"

	model "../../domain/entities"
	service "../interfaces"
)

type group struct {
}

var arrayOfGroupData = make([]model.GroupData, 0, 496)
var arrayOfClusters = make([]model.Clusters, 0, 10)

// GroupsService function
func GroupsService() service.IGroupsService {
	return &group{}
}

func (*group) GetGroups() (model.ClustersString, error) {
	trainKmeans()
	groups := findClustersNtimes()
	return groups, nil
}

func loadSick(i int, noCommasRow []string, dataBuffer chan model.GroupData) {
	enfermo, _ := strconv.Atoi(noCommasRow[7])
	if enfermo == 1 {
		edad, _ := strconv.Atoi(noCommasRow[0])
		peso, _ := strconv.ParseFloat(noCommasRow[1], 2)
		tos, _ := strconv.Atoi(noCommasRow[3])
		fiebre, _ := strconv.Atoi(noCommasRow[4])
		dificultadRespirar, _ := strconv.Atoi(noCommasRow[5])
		perdidaOlfato, _ := strconv.Atoi(noCommasRow[6])
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
		data := model.GroupData{Edad: edad, Peso: peso, Distrito: distrito, Tos: tos, Fiebre: fiebre, DificultadRespirar: dificultadRespirar, PerdidaOlfato: perdidaOlfato, Enfermo: enfermo}
		dataBuffer <- data
	}
}

func trainKmeans() {
	dataBuffer := make(chan model.GroupData, 496)
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
		go loadSick(i, noCommasRow, dataBuffer)
	}
	for i := 0; i < 496; i++ {
		arrayOfGroupData = append(arrayOfGroupData, <-dataBuffer)
	}
}

func calculateDistance2(predict model.GroupData, train model.GroupData) float64 {
	result1 := math.Pow((float64(predict.Edad) - float64(train.Edad)), 2)
	result2 := math.Pow((predict.Peso - train.Peso), 2)
	result3 := math.Pow((float64(predict.Distrito) - float64(train.Distrito)), 2)
	result4 := math.Pow((float64(predict.Tos) - float64(train.Tos)), 2)
	result5 := math.Pow((float64(predict.Fiebre) - float64(train.Fiebre)), 2)
	result6 := math.Pow((float64(predict.DificultadRespirar) - float64(train.DificultadRespirar)), 2)
	result7 := math.Pow((float64(predict.PerdidaOlfato) - float64(train.PerdidaOlfato)), 2)
	result := math.Sqrt(result1 + result2 + result3 + +result4 + result5 + result6 + result7)

	return result
}

func createClusters(k int) {
	arrayOfRootClusters := make([]model.GroupData, 0, 4)
	arrayOfGroups := make([]model.ClusteredGroup, 0, 4)

	variation := make([]float64, 0, 4)
	contOfDistances := make([]int, 0, 4)

	for i := 0; i < k; i++ {
		name := "Group " + strconv.Itoa(i)
		contOfDistances = append(contOfDistances, 0)
		variation = append(variation, 0.0)
		auxArray := make([]model.GroupData, 0, 150)
		auxGroup := model.ClusteredGroup{Name: name, Data: auxArray}
		arrayOfGroups = append(arrayOfGroups, auxGroup)
	}

	for i := 0; i < k; i++ {
		x := rand.Intn(495)
		rootCluster := arrayOfGroupData[x]
		arrayOfRootClusters = append(arrayOfRootClusters, rootCluster)
	}

	for i := 0; i < 496; i++ {
		arrayOfDistances := make([]float64, 0, 4)
		for j := 0; j < 4; j++ {
			distance := calculateDistance2(arrayOfGroupData[i], arrayOfRootClusters[j])
			arrayOfDistances = append(arrayOfDistances, distance)
		}
		indexOfMin := 0
		min := arrayOfDistances[0]
		for j := 0; j < 4; j++ {
			if min > arrayOfDistances[j] {
				min = arrayOfDistances[j]
				indexOfMin = j
			}
		}
		arrayOfGroups[indexOfMin].Data = append(arrayOfGroups[indexOfMin].Data, arrayOfGroupData[i])
		variation[indexOfMin] = variation[indexOfMin] + min
		contOfDistances[indexOfMin] = contOfDistances[indexOfMin] + 1
	}

	score := 0.0
	for i := 0; i < k; i++ {
		variation[i] = variation[i] / float64(contOfDistances[i])
		score = score + variation[i]
	}
	score = score / float64(k)

	cluster := model.Clusters{Score: score, Clusters: arrayOfGroups}
	arrayOfClusters = append(arrayOfClusters, cluster)
}

func findClustersNtimes() model.ClustersString {
	N := 100
	k := 4
	for i := 0; i < N; i++ {
		createClusters(k)
	}

	min := arrayOfClusters[0]
	for i := 0; i < len(arrayOfClusters); i++ {
		if min.Score > arrayOfClusters[i].Score {
			min = arrayOfClusters[i]
		}
	}

	clusters := make([]model.ClusteredGroupString, 0, 4)
	for i := 0; i < len(min.Clusters); i++ {
		auxGroup := make([]model.GroupDataString, 0, 400)
		for j := 0; j < len(min.Clusters[i].Data); j++ {

			distrito := ""
			switch min.Clusters[i].Data[j].Distrito {
			case 1:
				distrito = "Callao"
			case 2:
				distrito = "Ventanilla"
			case 3:
				distrito = "Ate"
			case 4:
				distrito = "Barranco"
			case 5:
				distrito = "Chorrillos"
			case 6:
				distrito = "Comas"
			case 7:
				distrito = "Jesus Maria"
			case 8:
				distrito = "La Molina"
			case 9:
				distrito = "La Victoria"
			case 10:
				distrito = "Lince"
			case 11:
				distrito = "Los Olivos"
			case 12:
				distrito = "Lurin"
			case 13:
				distrito = "Magdalena del Mar"
			case 14:
				distrito = "Miraflores"
			case 15:
				distrito = "Pueblo Libre"
			case 16:
				distrito = "Puente Piedra"
			case 17:
				distrito = "Rimac"
			case 18:
				distrito = "San Borja"
			case 19:
				distrito = "San Isidro"
			case 20:
				distrito = "San Juan de Lurigancho"
			case 21:
				distrito = "San Martin de Porres"
			case 22:
				distrito = "San Miguel"
			case 23:
				distrito = "Santiago de Surco"
			case 24:
				distrito = "Surquillo"
			case 25:
				distrito = "Villa El Salvador"
			default:
				distrito = "Nothing"
			}

			respirar := ""
			switch min.Clusters[i].Data[j].DificultadRespirar {
			case 0:
				respirar = "Falso"
			case 1:
				respirar = "Verdad"
			}

			perdidaOlfato := ""
			switch min.Clusters[i].Data[j].PerdidaOlfato {
			case 0:
				perdidaOlfato = "Falso"
			case 1:
				perdidaOlfato = "Verdad"
			}

			aux := model.GroupDataString{Edad: min.Clusters[i].Data[j].Edad, Peso: min.Clusters[i].Data[j].Peso, Distrito: distrito, DificultadRespirar: respirar, PerdidaOlfato: perdidaOlfato}
			auxGroup = append(auxGroup, aux)
		}
		name := "Group " + strconv.Itoa(i)
		auxCluster := model.ClusteredGroupString{Name: name, Data: auxGroup}
		clusters = append(clusters, auxCluster)
	}

	response := model.ClustersString{Score: min.Score, Clusters: clusters}
	return response
}

import csv
import random

sickPeople = 0
random.seed()
corpus = ""
rows = []
distritos = [	"Callao", "Callao", "Callao", "Callao", "Callao",
              "Ventanilla", "Ventanilla", "Ventanilla",
              "Ate", "Ate", "Ate",
              "Barranco", "Barranco", "Barranco", "Barranco",
              "Chorrillos", "Chorrillos", "Chorrillos",
              "Comas", "Comas", "Comas", "Comas", "Comas", "Comas",
              "Jesus Maria", "Jesus Maria", "Jesus Maria",
              "La Molina",
              "La Victoria", "La Victoria", "La Victoria", "La Victoria",
              "Lince", "Lince", "Lince",
              "Los Olivos", "Los Olivos", "Los Olivos", "Los Olivos", "Los Olivos", "Los Olivos",
              "Lurin", "Lurin", "Lurin", "Lurin", "Lurin", "Lurin", "Lurin", "Lurin", "Lurin",
              "Magdalena del Mar", "Magdalena del Mar",
              "Miraflores", "Miraflores",  "Miraflores",
              "Pueblo Libre", "Pueblo Libre",
              "Puente Piedra", "Puente Piedra", "Puente Piedra", "Puente Piedra",
              "Rimac", "Rimac", "Rimac", "Rimac", "Rimac", "Rimac",
              "San Borja", "San Borja", "San Borja", "San Borja", "San Borja", "San Borja",
              "San Isidro", "San Isidro",
              "San Juan de Lurigancho", "San Juan de Lurigancho", "San Juan de Lurigancho", "San Juan de Lurigancho", "San Juan de Lurigancho", "San Juan de Lurigancho"
              "San Martin de Porres", "San Martin de Porres", "San Martin de Porres", "San Martin de Porres", "San Martin de Porres", "San Martin de Porres",
              "San Miguel",  "San Miguel",  "San Miguel",
              "Santiago de Surco", "Santiago de Surco", "Santiago de Surco", "Santiago de Surco", "Santiago de Surco", "Santiago de Surco",
              "Surquillo", "Surquillo",
              "Villa El Salvador",  "Villa El Salvador", "Villa El Salvador"]

for i in range(1000):
    row = []
    # age
    age = random.randint(10, 75)
    if age < 65:
        age = random.randint(10, 75)
        if age < 45:
            age = random.randint(10, 75)

    weight = random.uniform(30.0, 120.0)
    if weight < 75:
        weight = random.uniform(30.0, 80.0)
        if weight < 45:
            weight = random.uniform(30.0, 80.0)

    weight = round(weight, 2)
    distritoIndex = random.randint(0, len(distritos)-1)
    distrito = distritos[distritoIndex]

    tos = random.randint(0, 1)
    fiebre = random.randint(0, 1)
    respirar = random.randint(0, 1)
    gusto = random.randint(0, 1)

    row.append(age)
    row.append(weight)
    row.append(distrito)
    row.append(tos)
    row.append(fiebre)
    row.append(respirar)
    row.append(gusto)

    chances = []
    cont = 100
    if tos:
        for i in range(10):
            chances.append(1)
            cont -= 1
    if fiebre:
        for i in range(20):
            chances.append(1)
            cont -= 1
    if respirar:
        for i in range(30):
            chances.append(1)
            cont -= 1
    if gusto:
        for i in range(40):
            chances.append(1)
            cont -= 1

    for i in range(cont):
        chances.append(0)

    sickIndex = random.randint(0, len(chances)-1)
    sick = chances[sickIndex]
    if sick == 1:
        sickPeople += 1
    row.append(sick)
    rows.append(row)

print(sickPeople)

with open('covidPeruDataSet.csv', 'w', newline='') as csvfile:
    writer = csv.writer(csvfile, delimiter=';', quoting=csv.QUOTE_MINIMAL)
    writer.writerow(
        ["edad", "peso", "distrito", "tos", "fiebre", "respirar", "gusto", "enfermo"])
    for line in rows:
        writer.writerow([line[0], line[1], line[2], line[3],
                         line[4], line[5], line[6], line[7]])

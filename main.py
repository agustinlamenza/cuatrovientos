def main():
    program = input("""
                    Seleccione el programa a ejecutar:
                    Filtro por edades (1)
                    Filtro por nombres (2)
                    Mostrar tipos de datos (3)
                    """)
    if program == "1":
        filtrar_edades()
    if program == "2":
        filtrar_nombres()
    elif program == "3":
        mostrar_tipos()
    else:
        print(f'"{program}" no es una opción valida')

def filtrar_edades():
    param = input("Ingrese la edad mínima: ")
    min = int(param)

    param = input("Ingrese la edad máxima: ")
    max = int(param)

    param = input("Ingrese el sexo: (1 femenimo, 2 masculino) ")
    sex = param

    people = []

    with open("personas.csv", "r") as f:
        for line in f.readlines():
            data = line.split(";")
            person = {
                "id": data[0],
                "first_name": data[1],
                "last_name": data[2],
                "sex" : data[3]
            }

            try:
                person["age"] = int(data[4])
            except ValueError:
                person["age"] = 0

            people.append(person)

    filtered = list(p for p in people if p["sex"] == sex and p["age"] >= min and p["age"] <= max)

    for p in filtered:
        print(f'{p["age"]} - {p["first_name"]} - {p["last_name"]}')

    print(f"Cantidad total: {len(filtered)}")

def filtrar_nombres():
    print("El programa no esta disponible aun")

def mostrar_tipos():
    cadena = "hola como estas"
    numerico = 123
    flotante = 123.654
    print(type(cadena))
    print(type(numerico))
    print(type(flotante))

if __name__ == "__main__":
   main()
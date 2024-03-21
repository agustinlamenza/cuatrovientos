namespace cuatrovientos
{
    public class Persona
    {
        public string Nombre { get; set; }
        public string Apellido { get; set; }
        public int Edad { get; set; }
    }

    class Program
    {
        static void Main(string[] args)
        {
            var personas = new List<Persona>();

            var reader = new StreamReader($"{Environment.CurrentDirectory}/personas.csv");
            while (!reader.EndOfStream)
            {
                var valor = reader.ReadLine() ?? string.Empty;
                var elementos = valor.Split(';');

                var persona = new Persona
                {
                    Nombre = elementos[1] ?? string.Empty,
                    Apellido = elementos[2] ?? string.Empty,
                    Edad = $"{elementos[1]}{elementos[2]}".Length
                };

                personas.Add(persona);
            }

            foreach (var item in personas.OrderByDescending(p => p.Edad).Take(10))
            {
                Console.WriteLine($"El nombre es: {item.Nombre}, {item.Apellido}");
            }
        }
    }
}


// while (reader.EndOfStream)
// {
//     var valor = reader.ReadLine();

//     Console.WriteLine($"{valor}");
// }



// bool bandera = true;
// var personas = new List<Persona>();

// while (bandera)
// {
//     Console.WriteLine("Ingrese el nombre: ");
//     var nombre = Console.ReadLine();

//     Console.WriteLine("Ingrese la edad: ");
//     var edad = Console.ReadLine();

//     var persona = new Persona
//     {
//         Nombre = nombre,
//         Edad = Convert.ToInt32(edad)
//     };

//     personas.Add(persona);

//     Console.WriteLine("Quiere ingresar otra persona? si/no");
//     var cortar = Console.ReadLine();
//     if (cortar == "no")
//     {
//         bandera = false;
//     }
// }

// var resultado = personas.OrderBy(p => p.Edad);

// foreach (var item in resultado)
// {
//     Console.WriteLine($"El nombre es: {item.Nombre}, y la edad es: {item.Edad}");
// }



package mypackage

// CarPublic es una estructura con campos públicos
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// Función pública
func PrintMessage() {
	println("Hola desde mypackage!")
}

// Función privada
func printPrivateMessage() {
	println("Este es un mensaje privado desde mypackage.")
}

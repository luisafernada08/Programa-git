package main

import (
    "fmt"
    "github.com/jung-kurt/gofpdf"
    "log"
)

type Producto struct {
    Nombre     string
    PrecioUnit float64
    Cantidad   int
}

func factura() {
    // Solicitar información al usuario
    var nombreCliente, fecha string
    var numProductos int

    fmt.Print("Ingrese el nombre del cliente: ")
    fmt.Scanln(&nombreCliente)
    fmt.Print("Ingrese la fecha de la factura: ")
    fmt.Scanln(&fecha)
    fmt.Print("Ingrese el número de productos: ")
    fmt.Scanln(&numProductos)

    // Crear una lista de productos
    productos := make([]Producto, numProductos)
    for i := 0; i < numProductos; i++ {
        fmt.Printf("Ingrese el nombre del producto %d: ", i+1)
        fmt.Scanln(&productos[i].Nombre)
        fmt.Print("Ingrese el precio unitario: ")
        fmt.Scanln(&productos[i].PrecioUnit)
        fmt.Print("Ingrese la cantidad: ")
        fmt.Scanln(&productos[i].Cantidad)
    }

    // Crear el PDF
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()

    // Agregar información al PDF
    pdf.SetFont("Arial", "B", 16)
    pdf.CellFormat(0, 10, "Factura","", 1, "C", false,0,"")
    pdf.SetFont("Arial", "", 12)

	// Agregar información cliente
    pdf.Cell(0, 10, fmt.Sprintf("Nombre del cliente: %s", nombreCliente), "", 1, "",false, 0, "")
    pdf.Cell(0, 10, fmt.Sprintf("Fecha: %s", fecha), "",1,"",false,0,"")
    pdf.Cell(0, 10, "",1,"",false,0,"") // Espacio en blanco

    // Encabezado de la tabla de productos
    
pdf.Cell(100, 10, "Producto", "B", 0, "", false, 0, "")
pdf.Cell(30, 10, "Precio Unit.", "B", 0, "R", false, 0, "")
pdf.Cell(30, 10, "Cantidad", "B", 0, "R", false, 0, "")
pdf.Cell(30, 10, "Total", "B", 1, "R", false, 0, "")

    // Calcular total y agregar productos
    var totalGeneral float64
    for _, producto := range productos {
        total := producto.PrecioUnit * float64(producto.Cantidad)
        totalGeneral += total

		// Imprimir productos en la tabla	
        pdf.CellFormat(100, 10, producto.Nombre,"",0,"",false,0,"")
        pdf.CellFormat(30, 10, fmt.Sprintf("%.2f", producto.PrecioUnit), "",0,"R",false, 0,"")
        pdf.Cell(30, 10, fmt.Sprintf("%d", producto.Cantidad), "",0,"R",false,0,"")
        pdf.Cell(30, 10, fmt.Sprintf("%.2f", total), "",1,"R",false,0,"")
    
    }

    // Total general
    pdf.Cell(100, 10, "Total General:", "T",0,"",false,0,"")
    pdf.Cell(30, 10,"", "T", 0, "", false, 0, "")
    pdf.Cell(30, 10,"", "T", 0, "", false, 0, "")
    pdf.Cell(30, 10, fmt.Sprintf("%.2f", totalGeneral),"", "T", 1, "", false, 0, "" )
    
    // Guardar el PDF
    err := pdf.OutputFileAndClose("factura.pdf")
    if err != nil {
        log.Fatalf("Error al guardar el PDF: %v", err)
    }

    fmt.Println("Factura generada: factura.pdf")

pdf.Cell( 0 , 10 , fmt.Sprintf( "Customer name: %s" , customerName), "" , 1 , "" , false , 0 , "" ) 
pdf.Cell( 0 , 10 , fmt .Sprintf( "Date: %s" , date), "" , 1 , "" , false , 0 , "" ) 
pdf.Cell( 0 , 10 , "" , "" , 1 , "" , false , 0 , "" ) // Blank // Product table header 
pdf.Cell( 100 , 10 , "Product" , "B" , 0 , "" , false , 0 , "" )"10"-10249pdf.Cell( 30 , 10 , "Unit Price." , "B" , 0 , "R" , false , 0 , "" ) 
pdf.Cell( 30 , 10 , "Quantity" , "B" , 0 , "R" , false , 0 , "" ) 
pdf.Cell( 30 , 10 , "Total" , "B" , 1 , "R" , false , 0 , "" ) // Calculate total and add products var totalGeneral float64 for _, product := range products { 
total := product.PriceUnit * float64 (product.Quantity) totalGeneral += total // Print products in table pdf.CellFormat( "" " , 0 , "R" , false , 0 , "" ) 
pdf.Cell( 30 , 10 , fmt.Sprintf( "%d" , product.Quantity), "" , 0 , "R" , false , 0 , "" ) 
pdf.Cell( 30 , 10 ) 

}

func main() {
	//factura()
	}
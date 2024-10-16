from reportlab.lib.pagesizes import letter
from reportlab.pdfgen import canvas
from datetime import date

def generar_factura_pdf(nombre_cliente, fecha, productos, archivo_salida="factura.pdf"):
    # Crear un canvas para el PDF
    c = canvas.Canvas(archivo_salida, pagesize=letter)
    ancho, alto = letter

    # Título
    c.setFont("Helvetica-Bold", 20)
    c.drawString(200, alto - 50, "Factura")

    # Información del cliente y fecha
    c.setFont("Helvetica", 12)
    c.drawString(50, alto - 100, f"Nombre del cliente: {nombre_cliente}")
    c.drawString(50, alto - 120, f"Fecha de la factura: {fecha}")

    # Encabezado de la tabla de productos
    c.setFont("Helvetica-Bold", 12)
    c.drawString(50, alto - 160, "Producto")
    c.drawString(250, alto - 160, "Precio Unitario")
    c.drawString(350, alto - 160, "Cantidad")
    c.drawString(450, alto - 160, "Costo Total")

    # Inicializar el total general
    total_general = 0
    y = alto - 180  # posición inicial para los productos

     # Imprimir en consola y en el PDF
    print("\nProductos y precios:")
    
    # Imprimir en la terminal
    #print(f"{nombre} - \nPrecio Unitario: ${precio:.2f},\n Cantidad: {cantidad},\n Costo Total: ${costo_total:.2f}")    

    # Listar productos
    for producto in productos:
        nombre, precio, cantidad = producto
        costo_total = precio * cantidad
        total_general += costo_total

        # Imprimir cada fila de productos
        c.setFont("Helvetica", 12)
        c.drawString(50, y, nombre)
        c.drawString(250, y, f"${precio:.2f}")
        c.drawString(350, y, str(cantidad))
        c.drawString(450, y, f"${costo_total:.2f}")
        y -= 20  # mover hacia abajo para el siguiente producto

    # Total general
    c.setFont("Helvetica-Bold", 12)
    c.drawString(50, y - 20, f"Total General: ${total_general:.2f}")

    # Guardar el PDF
    c.save()
    print("PDF generado con éxito: Factura.pdf")
    
    # Abrir el PDF después de generarlo
    import os
    os.startfile(archivo_salida)
    #lo que vayan a usar para funciones del PDF se menciona con pdf. y la funcio u opcion de la libreria
    

# Función para ingresar datos de la factura
def ingresar_datos_factura():
    nombre_cliente = input("Ingrese el nombre del cliente: ")
    fecha = input("Ingrese la fecha (YYYY-MM-DD) o presione enter para usar la fecha de hoy: ")
    if not fecha:
        fecha = str(date.today())

    productos = []
    while True:
        nombre_producto = input("Ingrese el nombre del producto (o 'fin' para terminar): ")
        if nombre_producto.lower() == 'fin':
            break
        precio_unitario = float(input(f"Ingrese el precio unitario de {nombre_producto}: "))
        cantidad = int(input(f"Ingrese la cantidad de {nombre_producto}: "))
        productos.append((nombre_producto, precio_unitario, cantidad))

    return nombre_cliente, fecha, productos


# Programa principal
if __name__ == "__main__":
    nombre_cliente, fecha, productos = ingresar_datos_factura()
    generar_factura_pdf(nombre_cliente, fecha, productos)
    print("Factura.pdf generada exitosamente.")


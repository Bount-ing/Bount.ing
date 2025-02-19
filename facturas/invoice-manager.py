#!/usr/bin/env python3

from reportlab.lib.pagesizes import A4
from reportlab.pdfgen import canvas
from reportlab.lib.utils import ImageReader
from reportlab.platypus import Table, TableStyle
from reportlab.lib import colors
from datetime import datetime
from PIL import Image
import io

def generar_factura(nombre, direccion, nif, comision, iva=21, num_factura="0001"):
    """
    Genera una factura en PDF para un freelance (solucionador) con un diseño mejorado, incluyendo logo y formato ordenado.
    """
    fecha = datetime.today().strftime('%d/%m/%Y')
    total_iva = (comision * iva) / 100
    total_factura = comision + total_iva

    nombre_archivo = f"factura_{num_factura}.pdf"
    c = canvas.Canvas(nombre_archivo, pagesize=A4)
    width, height = A4
    c.setFont("Helvetica", 12)

    # Process logo with white background
    logo_path = "logo.png"
    try:
        # Open the image
        img = Image.open(logo_path)

        # Convert to RGBA if it's not already
        if img.mode != 'RGBA':
            img = img.convert('RGBA')

        # Create a pure white background
        white_bg = Image.new("RGB", img.size, (255, 255, 255))

        # Paste the image onto the white background using the alpha channel as mask
        white_bg.paste(img, (0, 0), img.split()[3])

        # Save to a bytes buffer
        img_buffer = io.BytesIO()
        white_bg.save(img_buffer, format='PNG')
        img_buffer.seek(0)

        # Use the buffer with ImageReader
        c.drawImage(ImageReader(img_buffer), 50, 773, width=100, height=25)
    except Exception as e:
        print(f"Error processing logo: {str(e)}")
        # Continue without logo

    # Header
    c.setFont("Helvetica-Bold", 14)
    c.drawString(width/2 - 30, height - 50, "FACTURA")
    c.setFont("Helvetica", 12)
    c.drawString(width - 150, height - 50, f"Nº: {num_factura}")
    c.drawString(width - 150, height - 70, f"Fecha: {fecha}")

    # Create data for emisor and client in table format
    data = [
        ["EMISOR:", "CLIENTE:"],
        ["Bount.ing S.L.", nombre],
        ["Dirección: [Tu Dirección]", f"Dirección: {direccion}"],
        ["NIF: [Tu NIF]", f"NIF: {nif}"]
    ]

    # Create the table
    table = Table(data, colWidths=[width/2 - 75, width/2 - 75])

    # Style the table
    style = TableStyle([
        ('FONT', (0, 0), (1, 0), 'Helvetica-Bold', 12),
        ('FONT', (0, 1), (1, 3), 'Helvetica', 12),
        ('TOPPADDING', (0, 0), (-1, -1), 6),
        ('BOTTOMPADDING', (0, 0), (-1, -1), 6),
        ('ALIGNMENT', (0, 0), (0, -1), 'LEFT'),
        ('ALIGNMENT', (1, 0), (1, -1), 'RIGHT'),
    ])

    table.setStyle(style)

    # Draw the table
    table.wrapOn(c, width, height)
    table.drawOn(c, 50, height - 230)

    # Línea separadora
    c.line(50, height - 250, width - 50, height - 250)

    # Detalles de la factura
    c.setFont("Helvetica-Bold", 12)
    c.drawString(50, height - 280, "Concepto")
    c.drawString(width - 150, height - 280, "Importe (€)")
    c.setFont("Helvetica", 12)
    c.drawString(50, height - 300, "Comisión por uso de la plataforma Bount.ing")
    c.drawString(width - 150, height - 300, f"{comision:.2f}")

    # Subtotal e IVA
    c.drawString(50, height - 330, "Base Imponible:")
    c.drawString(width - 150, height - 330, f"{comision:.2f} €")
    c.drawString(50, height - 350, f"IVA ({iva}%):")
    c.drawString(width - 150, height - 350, f"{total_iva:.2f} €")

    # Total Factura
    c.setFont("Helvetica-Bold", 12)
    c.drawString(50, height - 370, "Total Factura:")
    c.drawString(width - 150, height - 370, f"{total_factura:.2f} €")

    # Línea separadora
    c.line(50, height - 390, width - 50, height - 390)

    # Pie de página
    c.setFont("Helvetica", 10)
    c.drawString(50, height - 410, "Método de pago: Procesado a través de Stripe Connect")
    c.drawString(50, height - 430, "Gracias por usar Bount.ing!")

    c.save()
    print(f"Factura generada: {nombre_archivo}")

# Ejemplo de uso
if __name__ == "__main__":
    generar_factura(
        nombre="Juan Pérez",
        direccion="Calle Falsa 123, Madrid",
        nif="12345678X",
        comision=10.00  # Comisión en euros
    )

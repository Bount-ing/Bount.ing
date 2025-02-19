#!/usr/bin/env python3

from reportlab.lib.pagesizes import A4
from reportlab.pdfgen import canvas
from datetime import datetime
import os

def test_step1():
    """Basic PDF with just invoice header"""
    try:
        output_file = "step1_test.pdf"
        c = canvas.Canvas(output_file, pagesize=A4)

        # Basic header
        c.setFont("Helvetica-Bold", 14)
        c.drawString(200, 780, "FACTURA")
        c.setFont("Helvetica", 12)

        print("About to save PDF...")
        c.save()
        print(f"Attempted to save to: {os.path.abspath(output_file)}")

        if os.path.exists(output_file):
            print(f"SUCCESS: File created at {os.path.abspath(output_file)}")
        else:
            print(f"FAILURE: File not created at {os.path.abspath(output_file)}")
    except Exception as e:
        print(f"EXCEPTION: {str(e)}")
        import traceback
        print(traceback.format_exc())

def test_step2():
    """Add simple text fields (no logo yet)"""
    try:
        output_file = "step2_test.pdf"
        c = canvas.Canvas(output_file, pagesize=A4)

        # Header
        c.setFont("Helvetica-Bold", 14)
        c.drawString(200, 780, "FACTURA")
        c.setFont("Helvetica", 12)
        c.drawString(400, 780, "Nº: 0001")

        # Basic client info
        c.drawString(50, 700, "Cliente: Test Client")
        c.drawString(50, 680, "Dirección: Test Address")

        print("About to save PDF...")
        c.save()
        print(f"Attempted to save to: {os.path.abspath(output_file)}")

        if os.path.exists(output_file):
            print(f"SUCCESS: File created at {os.path.abspath(output_file)}")
        else:
            print(f"FAILURE: File not created at {os.path.abspath(output_file)}")
    except Exception as e:
        print(f"EXCEPTION: {str(e)}")
        import traceback
        print(traceback.format_exc())

def test_step3():
    """Test with improved transparent PNG handling"""
    try:
        from reportlab.lib.utils import ImageReader
        from PIL import Image
        import io

        output_file = "improved_transparency_test.pdf"
        c = canvas.Canvas(output_file, pagesize=A4)

        logo_path = "logo.png"
        if os.path.exists(logo_path):
            try:
                # Process logo with transparency
                img = Image.open(logo_path)
                print(f"Opened image: {logo_path}, mode: {img.mode}, size: {img.size}")

                if img.mode == 'RGBA':
                    # Create a white RGB background
                    background = Image.new('RGB', img.size, (255, 255, 255))

                    # Extract alpha channel
                    alpha = img.split()[3]

                    # Remove alpha channel from image
                    rgb_img = img.convert('RGB')

                    # Paste using alpha as mask
                    background.paste(rgb_img, (0, 0), mask=alpha)

                    # Save to a bytes buffer
                    img_buffer = io.BytesIO()
                    background.save(img_buffer, format='PNG')
                    img_buffer.seek(0)

                    # Use the buffer with ImageReader
                    c.drawImage(ImageReader(img_buffer), 50, 750, width=100, height=50)
                    print("Processed and added transparent PNG to canvas with improved method")
                else:
                    c.drawImage(ImageReader(logo_path), 50, 750, width=100, height=50)
                    print("Added regular image to canvas")
            except Exception as e:
                print(f"Error processing logo: {str(e)}")
                import traceback
                print(traceback.format_exc())

        # Header
        c.setFont("Helvetica-Bold", 14)
        c.drawString(200, 780, "FACTURA")

        print("About to save PDF...")
        c.save()
        print(f"Attempted to save to: {os.path.abspath(output_file)}")

        if os.path.exists(output_file):
            print(f"SUCCESS: File created at {os.path.abspath(output_file)}")
        else:
            print(f"FAILURE: File not created at {os.path.abspath(output_file)}")
    except Exception as e:
        print(f"EXCEPTION: {str(e)}")
        import traceback
        print(traceback.format_exc())

if __name__ == "__main__":
    print("Starting step 1 test")
    test_step1()
    print("\nStarting step 2 test")
    test_step2()
    print("\nStarting step 3 test")
    test_step3()
    print("All tests completed")

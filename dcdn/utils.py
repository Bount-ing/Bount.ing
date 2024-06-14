import base64
import requests

def interpolate_color(start_color, end_color, factor):
    """
    Interpolates between two colors.
    start_color and end_color are in hex format (e.g., "000000" for black).
    factor is a float between 0 and 1.
    """
    start_color = int(start_color, 16)
    end_color = int(end_color, 16)

    start_r = (start_color >> 16) & 0xFF
    start_g = (start_color >> 8) & 0xFF
    start_b = (start_color & 0xFF)

    end_r = (end_color >> 16) & 0xFF
    end_g = (end_color >> 8) & 0xFF
    end_b = (end_color & 0xFF)

    r = round(start_r + (end_r - start_r) * factor)
    g = round(start_g + (end_g - start_g) * factor)
    b = round(start_b + (end_b - start_b) * factor)

    return f'{r:02x}{g:02x}{b:02x}'

def image_to_base64(url):
    response = requests.get(url)
    return base64.b64encode(response.content).decode('utf-8')

def local_image_to_base64(image_path):
    with open(image_path, 'rb') as image_file:
        return base64.b64encode(image_file.read()).decode('utf-8')
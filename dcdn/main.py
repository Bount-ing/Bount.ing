from flask import Flask, request, jsonify, send_file, Response
import psycopg2
import requests
import base64
import os
import cairosvg
from PIL import Image
from io import BytesIO
import imageio.v2 as imageio
from utils import image_to_base64, local_image_to_base64, interpolate_color

app = Flask(__name__)

# Database configuration
DATABASE = {
    'dbname': os.environ.get('POSTGRES_DB', 'postgres'),
    'user': os.environ.get('POSTGRES_USER', 'postgres'),
    'password': os.environ.get('POSTGRES_PASSWORD', 'postgres'),
    'host': os.environ.get('POSTGRES_HOST', 'db'),
    'port': os.environ.get('POSTGRES_PORT', '5432')
}

base_url = os.environ.get('BASE_URL', 'http://0.0.0.0:3000')




# Connect to the database
def get_db_connection():
    conn = psycopg2.connect(**DATABASE)
    return conn

# Function to retrieve total bounties for an issue
def get_total_bounties(issue_id):
    conn = get_db_connection()
    cursor = conn.cursor()
    query = """
    SELECT COALESCE(SUM(amount), 0) as total_bounty
    FROM bounties
    WHERE issue_id = %s;
    """
    cursor.execute(query, (issue_id,))
    total_bounty = cursor.fetchone()[0]
    cursor.close()
    conn.close()
    return total_bounty

# Function to retrieve the issue image URL from the bounties table
def get_issue_image_url(issue_id):
    conn = get_db_connection()
    cursor = conn.cursor()
    query = """
    SELECT issue_image_url
    FROM bounties
    WHERE issue_id = %s
    LIMIT 1;
    """
    cursor.execute(query, (issue_id,))
    result = cursor.fetchone()
    cursor.close()
    conn.close()
    return result[0] if result else None

def create_issue_card(issue, total_bounty, issue_image_url, local_image_path):
    owner, repo = issue['url'].split('/')[-4], issue['url'].split('/')[-3]
    repo_url = f"https://github.com/{owner}/{repo}"
    issue_title = issue['title']
    
    # Convert images to base64
    external_image_base64 = image_to_base64(issue_image_url)
    local_image_base64 = local_image_to_base64(local_image_path)

    svg_content = f'''
    <svg width="338" height="213" viewBox="0 0 338 213" xmlns="http://www.w3.org/2000/svg">
        <!-- Background with a soft dark matrix effect -->
        <defs>
            <linearGradient id="bgGradient" x1="0%" y1="0%" x2="100%">
                <stop offset="0%" style="stop-color:#000000;stop-opacity:1" />
                <stop offset="100%" style="stop-color:#0c0c0c;stop-opacity:1" />
            </linearGradient>
            <filter id="softGlow" x="-50%" y="-50%" width="200%" height="200%">
                <feGaussianBlur stdDeviation="2" result="coloredBlur"/>
                <feMerge>
                    <feMergeNode in="coloredBlur"/>
                    <feMergeNode in="SourceGraphic"/>
                </feMerge>
            </filter>
        </defs>
        <rect width="100%" height="100%" fill="url(#bgGradient)" rx="15" ry="15" />

        <!-- Title with soft glow effect -->
        <text x="50%" y="40" font-family="Nimbus Mono L" font-size="28" fill="#1abc9c" text-anchor="middle" letter-spacing="2" filter="url(#softGlow)">
            <tspan id="wanted">WANTED</tspan>
            <animate attributeName="opacity" values="1;0;0;0;0;1" dur="6s" repeatCount="indefinite" />
        </text>
        <text x="50%" y="40" font-family="Nimbus Mono L" font-size="28" fill="#1abc9c" text-anchor="middle" letter-spacing="2" filter="url(#softGlow)">
            <tspan id="solved">SOLVED</tspan>
            <animate attributeName="opacity" values="0;0;1;1;0;0" dur="6s" repeatCount="indefinite" />
        </text>

        <!-- Repo and Owner -->
        <text x="50%" y="70" font-family="Nimbus Mono L" font-size="10" fill="#1abc9c" text-anchor="middle" filter="url(#softGlow)">
            {owner}/{repo}
        </text>

        <!-- Issue Title -->
        <text id="issueTitle" x="50%" y="150" font-family="Nimbus Mono L" font-size="16" fill="#1abc9c" text-anchor="middle" filter="url(#softGlow)">
            <tspan x="50%" dy="-1em">{issue_title[:35]}</tspan>
            <tspan x="50%" dy="1.4em">{issue_title[35:70]}</tspan>
            <animate attributeName="opacity" values="1;0;0;0;0;1" dur="6s" repeatCount="indefinite" />
        </text>
        
        <!-- Bounty Amount -->
        <text id="bounty" x="50%" y="150" font-family="Nimbus Mono L" font-size="42" fill="#1abc9c" text-anchor="middle" filter="url(#softGlow)">
            <tspan id="total" >{total_bounty} €</tspan>
            <animate attributeName="opacity" values="0;0;1;1;0;0" dur="6s" repeatCount="indefinite" />
        </text>

        <!-- Issue Image with circular clipping and link to repo -->
        <a href="{repo_url}" target="_blank">
            <clipPath id="clipCircleIssue">
                <circle cx="60" cy="60" r="30" />
            </clipPath>
            <image x="30" y="30" width="60" height="60" href="data:image/png;base64,{external_image_base64}" clip-path="url(#clipCircleIssue)" />
        </a>

        <!-- Logo with circular clipping and link to bounty page -->
        <a href="https://bount.ing" target="_blank">
            <clipPath id="clipCircleLogo">
                <circle cx="278" cy="60" r="45" />
            </clipPath>
            <image x="233" y="15" width="90" height="90" href="data:image/png;base64,{local_image_base64}" clip-path="url(#clipCircleLogo)" />
        </a>

        <!-- Futuristic Border with a soft glow and blinking -->
        <rect x="5" y="5" width="328" height="203" rx="15" ry="15" fill="none" stroke="#1abc9c" stroke-width="1" stroke-dasharray="5,3" filter="url(#softGlow)">
            <animate attributeName="stroke-dashoffset" values="0;30;0" dur="6s" repeatCount="indefinite" />
        </rect>

        <rect x="2" y="2" width="334" height="209" rx="15" ry="15" fill="none" stroke="#1abc9c" stroke-width="1" stroke-dasharray="5,3" filter="url(#softGlow)">
            <animate attributeName="stroke-dashoffset" values="30;0;30" dur="6s" repeatCount="indefinite" />
        </rect>
    </svg>
    '''
    return svg_content



@app.route('/issues/<int:issue_id>/card.svg', methods=['GET'])
def get_issue_card(issue_id):
    conn = get_db_connection()
    cursor = conn.cursor()
    cursor.execute("SELECT github_url, title, description FROM issues WHERE id = %s", (issue_id,))
    issue = cursor.fetchone()
    if issue:
        issue_data = {
            'id': issue_id,
            'url': issue[0],
            'title': issue[1],
            'body': issue[2]
        }
        total_bounty = get_total_bounties(issue_id)
        issue_image_url = get_issue_image_url(issue_id)  # Fetch the image URL from bounties table
        svg_content = create_issue_card(issue_data, total_bounty, issue_image_url, 'logo.png')
        return Response(svg_content, mimetype='image/svg+xml')
    return jsonify({'error': 'Issue not found'}), 404

def create_bounty_card(bounty, issue, local_image_path, options):
    owner, repo = issue['url'].split('/')[-4], issue['url'].split('/')[-3]
    repo_url = f"https://github.com/{owner}/{repo}"
    external_image_base64 = image_to_base64(issue["image_url"])
    local_image_base64 = local_image_to_base64(local_image_path)
    opacity_logo_value = (float(options["opacity_1"]) + float(options["opacity_2"]))
    opacity_logos = f"%.1f" % opacity_logo_value
    
    dark_red = "DB0000"
    
    svg_content = f'''
    <svg width="338" height="213" viewBox="0 0 338 213" xmlns="http://www.w3.org/2000/svg">
        <!-- Background with a soft dark matrix effect -->
                <defs>
            <linearGradient id="bgGradient" x1="0%" y1="0%" y2="100%" x2="0%">
                <stop offset="99%" style="stop-color:#000000;stop-opacity:1" />
                <stop offset="100%" style="stop-color:#{options["color"]};stop-opacity:1" />
            </linearGradient>
            <filter id="softGlow" x="-50%" y="-50%" width="200%" height="200%">
                <feGaussianBlur stdDeviation="2" result="coloredBlur"/>
                <feMerge>
                    <feMergeNode in="coloredBlur"/>
                    <feMergeNode in="SourceGraphic"/>
                </feMerge>
            </filter>
        </defs>
        <rect width="100%" height="100%" fill="url(#bgGradient)" rx="15" ry="15" />

        <!-- Title with soft glow effect -->
        <text x="50%" y="40" font-family="Nimbus Mono L" font-size="28" fill="#{dark_red}" text-anchor="middle" letter-spacing="2" filter="url(#softGlow)" style="opacity: {options["opacity_1"]}">
            <tspan id="wanted">WANTED</tspan>
        </text>
        <text x="50%" y="40" font-family="Nimbus Mono L" font-size="28" fill="#{options["color"]}" text-anchor="middle" letter-spacing="2" filter="url(#softGlow)" style="opacity: {options["opacity_2"]}">
            <tspan id="solved">SOLVED</tspan>
            <animate attributeName="opacity" values="0;0;1;1;0;0" dur="6s" repeatCount="indefinite" />
        </text>

        <!-- Repo and Owner -->
        <text x="50%" y="70" font-family="Nimbus Mono L" font-size="10" fill="#1abc9c" text-anchor="middle" filter="url(#softGlow)" style="opacity: {opacity_logos}">
            {owner}
        </text>
        <text x="50%" y="85" font-family="Nimbus Mono L" font-size="10" fill="#1abc9c" text-anchor="middle" filter="url(#softGlow)" style="opacity: {opacity_logos}">
            {repo}
        </text>

        <!-- Issue Title -->
        <text id="issueTitle" x="50%" y="130" font-family="Nimbus Mono L" font-size="14" fill="#1abc9c" text-anchor="middle" filter="url(#softGlow)" style="opacity: {opacity_logos}">
            <tspan x="50%" dy="-1em">{issue["title"][:35]}</tspan>
            <tspan x="50%" dy="1.4em">{issue["title"][35:70]}</tspan>
        </text>
        
        <!-- Bounty Amount -->
        <text id="bounty" x="50%" y="175" font-family="Nimbus Mono L" font-size="21" fill="#{dark_red}" text-anchor="middle" filter="url(#softGlow)" style="opacity: {options["opacity_2"]}">
            <tspan id="total" >Up to {bounty["amount"]} €</tspan>
        </text>

        <!-- Issue Image with circular clipping and link to repo -->
        <a href="{repo_url}" target="_blank" style="opacity: {opacity_logos}">
            <clipPath id="clipCircleIssue">
                <circle cx="60" cy="60" r="30" />
            </clipPath>
            <image x="30" y="30" width="60" height="60" href="data:image/png;base64,{external_image_base64}" clip-path="url(#clipCircleIssue)" />
        </a>

        <!-- Logo with circular clipping and link to bounty page -->
        <a href="https://bount.ing" target="_blank" style="opacity: {opacity_logos}">
            <clipPath id="clipCircleLogo">
                <circle cx="278" cy="60" r="45" />
            </clipPath>
            <image x="233" y="15" width="90" height="90" href="data:image/png;base64,{local_image_base64}" clip-path="url(#clipCircleLogo)" />
        </a>

        <!-- Futuristic Border with a soft glow and blinking -->
        <rect x="5" y="5" width="328" height="203" rx="15" ry="15" fill="none" stroke="#{dark_red}" stroke-width="1" stroke-dasharray="5,3" filter="url(#softGlow)" style="opacity: {opacity_logos}">
            <animate attributeName="stroke-dashoffset" values="0;30;0" dur="6s" repeatCount="indefinite" />
        </rect>

        <rect x="2" y="2" width="334" height="209" rx="15" ry="15" fill="none" stroke="#1abc9c" stroke-width="1" stroke-dasharray="5,3" filter="url(#softGlow)" style="opacity: {opacity_logos}">
            <animate attributeName="stroke-dashoffset" values="30;0;30" dur="6s" repeatCount="indefinite" />
        </rect>
    </svg>
    '''
    # SVG saving and conversion to PNG
    svg_file_path = f'bounty_card_{options["name"]}.svg'
    png_file_path = f'bounty_card_{options["name"]}.png'
    with open(svg_file_path, 'w', encoding='utf-8') as f:
        f.write(svg_content)
    cairosvg.svg2png(url=svg_file_path, write_to=png_file_path)
    os.remove(svg_file_path)  # Remove the SVG to avoid confusion
    return png_file_path

def generate_frames(bounty, issue, local_image_path, frame_modifiers, path):
    frames = []
    for i, modifier in enumerate(frame_modifiers):
        frame_path = create_bounty_card(bounty, issue, local_image_path, modifier)
        print(f"Creating frame {i} at path {frame_path}")
        try:
            frame = imageio.imread(frame_path)
            frames.append(frame)
        except Exception as e:
            print(f"Failed to read image {frame_path}: {str(e)}", flush=True)
    
    print(f"Total frames created: {len(frames)}")
    # Copy frames array and reverse it to create the reverse animation
    reverse_frames = frames.copy()
    reverse_frames.reverse()
    frames.extend(reverse_frames)
    try:
        imageio.mimsave(path, frames, format='GIF', duration=len(frames)*2.4, loop=0)
        print(f"Saved GIF to {path}", flush=True)
    except Exception as e:
        print(f"Failed to save GIF: {str(e)}", flush=True)



@app.route('/bounties/<int:bounty_id>/card.gif', methods=['GET'])
def get_bounty_card(bounty_id):
    conn = get_db_connection()
    cursor = conn.cursor()
    
    # Retrieve the bounty details along with the related issue from the database
    cursor.execute("""
        SELECT b.amount, b.issue_id, i.github_url, i.title, i.description
        FROM bounties b
        JOIN issues i ON b.issue_id = i.id
        WHERE b.id = %s
    """, (bounty_id,))
    
    bounty = cursor.fetchone()
    if bounty:
        bounty_amount, issue_id, issue_github_url, issue_title, issue_description = bounty
        
        # Prepare the data structure for the issue
        issue_image_url = get_issue_image_url(issue_id)

        issue = {
            'id': issue_id,
            'url': issue_github_url,
            'image_url': issue_image_url,
            'title': issue_title,
            'body': issue_description
        }

        bounty_data = {
            'id': bounty_id,
            'amount': bounty_amount
        }

        # Assume get_issue_image_url fetches the image URL based on the issue ID
        
        # Create the SVG content using the issue data and bounty details
        file_path = f'bounty_card_{bounty_id}.gif'  # Ensure this path is correctly formed
        # progressive black green red
        # Define the colors to transition between
        color_steps = [
            "1abc9c",  # logo
            "000", # Black
            "1abc9c",  # logo
        ]

        # Number of frames
        num_frames = 20
        frames_per_segment = num_frames // (len(color_steps) - 1)
        frame_modifiers = []

        # Create the frame modifiers
        
        for segment in range(len(color_steps) - 1):
            start_color = color_steps[segment]
            end_color = color_steps[segment + 1]
            for i in range(frames_per_segment):
                factor = i / (frames_per_segment - 1)
                color = interpolate_color(start_color, end_color, factor)
                total_index = segment * frames_per_segment + i
                opacity_1 = 0.91 * (1 - ((total_index*1.25) / num_frames))
                opacity_2 = 0.91 * ((total_index * 0.85) / num_frames)
                frame_modifiers.append({
                    "name": f"frame{total_index + 1}",
                    "color": color,
                    "opacity_1": f"{opacity_1:.1f}",
                    "opacity_2": f"{opacity_2:.1f}"
                })



        generate_frames(bounty_data, issue, 'logo.png', frame_modifiers, file_path)
        return send_file(file_path, mimetype='image/png')
    
    return jsonify({'error': 'Bounty not found'}), 404

@app.route('/logo.png', methods=['GET'])
def get_logo():
    logo_path = 'logo.png'
    return send_file(logo_path, mimetype='image/png')


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=3001)

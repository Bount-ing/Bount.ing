from flask import Flask, request, jsonify, send_file, Response
import psycopg2
import requests
import base64
import io
import os

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

def image_to_base64(url):
    response = requests.get(url)
    return base64.b64encode(response.content).decode('utf-8')

def local_image_to_base64(image_path):
    with open(image_path, 'rb') as image_file:
        return base64.b64encode(image_file.read()).decode('utf-8')


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

def create_wanted_poster(issue, total_bounty, issue_image_url, local_image_path):
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
        <text x="50%" y="40" font-family="ui-sans-serif, sans-serif" font-size="28" fill="#1abc9c" text-anchor="middle" letter-spacing="2" filter="url(#softGlow)">
            <tspan id="wanted">WANTED</tspan>
            <animate attributeName="opacity" values="1;0;0;1" keyTimes="0;0.5;0.5;1" dur="6s" repeatCount="indefinite" />
        </text>
        <text x="50%" y="40" font-family="ui-sans-serif, sans-serif" font-size="28" fill="#1abc9c" text-anchor="middle" letter-spacing="2" filter="url(#softGlow)">
            <tspan id="solved">SOLVED</tspan>
            <animate attributeName="opacity" values="0;0;1;0" keyTimes="0;0.5;0.5;1" dur="6s" repeatCount="indefinite" />
        </text>

        <!-- Repo and Owner -->
        <text x="50%" y="70" font-family="ui-sans-serif, sans-serif" font-size="10" fill="#1abc9c" text-anchor="middle" filter="url(#softGlow)">
            {owner}/{repo}
        </text>

        <!-- Issue Title -->
        <text id="issueTitle" x="50%" y="180" font-family="ui-sans-serif, sans-serif" font-size="16" fill="#1abc9c" text-anchor="middle" filter="url(#softGlow)">
            <tspan x="50%" dy="-1em">{issue_title[:35]}</tspan>
            <tspan x="50%" dy="1.4em">{issue_title[35:70]}</tspan>
        </text>
        
        <!-- Bounty Amount -->
        <text id="bounty" x="50%" y="120" font-family="ui-sans-serif, sans-serif" font-size="24" fill="#1abc9c" text-anchor="middle" filter="url(#softGlow)" display="none">
            {total_bounty} €
            <set attributeName="display" to="inline" begin="3s" dur="6s" repeatCount="indefinite" />
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
        <rect x="5" y="5" width="328" height="203" rx="15" ry="15" fill="none" stroke="#1abc9c" stroke-width="3" stroke-dasharray="5,3" filter="url(#softGlow)">
            <animate attributeName="stroke-dashoffset" from="0" to="30" dur="6s" repeatCount="indefinite" />
        </rect>
    </svg>
    '''
    return svg_content





@app.route('/issues/<int:issue_id>/wanted_card.svg', methods=['GET'])
def get_wanted_card(issue_id):
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
        svg_content = create_wanted_poster(issue_data, total_bounty, issue_image_url, 'logo.png')
        return Response(svg_content, mimetype='image/svg+xml')
    return jsonify({'error': 'Issue not found'}), 404

@app.route('/logo.png', methods=['GET'])
def get_logo():
    logo_path = 'logo.png'
    return send_file(logo_path, mimetype='image/png')


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=3001)

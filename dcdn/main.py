from flask import Flask, request, jsonify, send_file, Response
import psycopg2
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

def create_wanted_poster(issue, total_bounty, issue_image_url):
    owner, repo = issue['url'].split('/')[-4], issue['url'].split('/')[-3]
    repo_url = f"https://github.com/{owner}/{repo}"
    svg_content = f'''
    <svg width="338" height="213" viewBox="0 0 338 213" xmlns="http://www.w3.org/2000/svg">
        <!-- Background black with subtle gold gradient -->
        <defs>
            <linearGradient id="bgGradient" x1="0%" y1="0%" x2="100%" y2="100%">
                <stop offset="0%" style="stop-color:#0a0a0a;stop-opacity:1" />
                <stop offset="100%" style="stop-color:#1a1a1a;stop-opacity:1" />
            </linearGradient>
            <linearGradient id="goldGradient" x1="0%" y1="0%" x2="100%" y2="100%">
                <stop offset="0%" style="stop-color:#FFD700;stop-opacity:1" />
                <stop offset="100%" style="stop-color:#FF8C00;stop-opacity:1" />
            </linearGradient>
            <filter id="neonGlow" x="-50%" y="-50%" width="200%" height="200%">
                <feGaussianBlur stdDeviation="4" result="coloredBlur"/>
                <feMerge>
                    <feMergeNode in="coloredBlur"/>
                    <feMergeNode in="SourceGraphic"/>
                </feMerge>
            </filter>
        </defs>
        <rect width="100%" height="100%" fill="url(#bgGradient)" rx="15" ry="15" />

        <!-- Futuristic Title with Neon Glow -->
        <text x="50%" y="40" font-family="Orbitron, sans-serif" font-size="28" fill="url(#goldGradient)" text-anchor="middle" letter-spacing="2" filter="url(#neonGlow)">
            WANTED
        </text>
        <text x="50%" y="70" font-family="Orbitron, sans-serif" font-size="10" fill="url(#goldGradient)" text-anchor="middle">{owner}/{repo}</text>

        <!-- Multiline Issue Title -->
        <text x="50%" y="90" font-family="Orbitron, sans-serif" font-size="12" fill="#FFFFFF" text-anchor="middle">
            <tspan x="50%" dy="1.2em">{issue['title'][:30]}</tspan>
            <tspan x="50%" dy="1.2em">{issue['title'][30:60]}</tspan>
            <tspan x="50%" dy="1.2em">{issue['title'][60:90]}</tspan>
        </text>

        <!-- Issue Image with circular clipping and link to repo -->
        <a href="{repo_url}" target="_blank">
            <clipPath id="clipCircleIssue">
                <circle cx="60" cy="60" r="30" />
            </clipPath>
            <image x="30" y="30" width="60" height="60" href="{issue_image_url}" clip-path="url(#clipCircleIssue)" />
        </a>

        <!-- Logo with circular clipping and link to bounty page -->
        <a href="https://bount.ing" target="_blank">
            <clipPath id="clipCircleLogo">
                <circle cx="278" cy="60" r="45" />
            </clipPath>
            <image x="233" y="15" width="90" height="90" href="/logo.png" clip-path="url(#clipCircleLogo)" />
        </a>

        <!-- Reward Section -->
        <text x="50%" y="160" font-family="Orbitron, sans-serif" font-size="18" fill="url(#goldGradient)" text-anchor="middle" filter="url(#neonGlow)">
            {total_bounty} €
        </text>

        <!-- Futuristic Border with glowing effect -->
        <rect x="5" y="5" width="328" height="203" rx="15" ry="15" fill="none" stroke="url(#goldGradient)" stroke-width="4" stroke-dasharray="10,5" filter="url(#neonGlow)" />
    </svg>
    '''
    html_content = f'''
    <html>
    <body style="background-color: #0a0a0a; display: flex; justify-content: center; align-items: center; height: 100vh;">
        {svg_content}
    </body>
    </html>
    '''
    return html_content




@app.route('/issues/<int:issue_id>/wanted_card.svg', methods=['GET'])
def get_wanted_card(issue_id):
    conn = get_db_connection()
    cursor = conn.cursor()
    cursor.execute("SELECT github_url, title, description FROM issues WHERE id = %s", (issue_id,))
    issue = cursor.fetchone()
    if issue:
        issue_data = {
            'url': issue[0],
            'title': issue[1],
            'body': issue[2]
        }
        total_bounty = get_total_bounties(issue_id)
        issue_image_url = get_issue_image_url(issue_id)  # Fetch the image URL from bounties table
        svg_content = create_wanted_poster(issue_data, total_bounty, issue_image_url)
        return Response(svg_content, mimetype='image/svg+xml')
    return jsonify({'error': 'Issue not found'}), 404

@app.route('/logo.png', methods=['GET'])
def get_logo():
    logo_path = 'logo.png'
    return send_file(logo_path, mimetype='image/png')


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=3001)

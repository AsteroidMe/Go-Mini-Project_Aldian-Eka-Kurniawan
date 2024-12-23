# **Eco Journal**

## **Apa itu Eco Journal?**

Eco Journal yang saya buat diharapkan bisa untuk menjadi wadah untuk karya tulis yang berkaitan dengan lingkungan hijau, memudahkan para pengguna dalam mengaksesnya.

## **Persiapan**
* **Go:** 1.23 or higher
* **Database:** MySQL 8.0 or higher
* **API Keys:**
    * `GEMINI_API_KEY` for use of Gen AI Gemini 1.5 Flash. [Get your API key](https://ai.google.dev/gemini-api/docs?gad_source=1&gclid=CjwKCAiAxea5BhBeEiwAh4t5K-uLKpnHMmmUmfdgAgRQG-WXsX2AP1N9CETrOASezuTErrrTuhGiHBoCTaUQAvD_BwE&hl=id)
* **Alat:**
    * `Docker` (opsional)

## **Instalasi**

Ikuti langkah berikut:

1. Clone repository: 
    ```bash
    git clone https://github.com/AsteroidMe/Go-Mini-Project_Aldian-Eka-Kurniawan
2. Navigate to the project directory:
    ```bash
    cd Go-Mini-Project_Aldian-Eka-Kurniawan
3. Install dependencies:
    ```bash
    go get .
3. Add ENV file:
    ```bash
    copy .env.example .env
4. Modify ENV file:
    ```bash
    # DB CONFIG
    DB_HOST="localhost"
    DB_PORT="3306"
    DB_USERNAME="root"
    DB_PASSWORD=""
    DB_NAME="ecojournal"
    
    # JWT
    JWT_SECRET_KEY="your_jwt_secret"
    
    # API KEYS
    GEMINI_API_KEY="your_gemini_api_key"
6. Run tests:
    ```bash
    go test ./...
7. Start the project:
    ```bash
    go run main.go

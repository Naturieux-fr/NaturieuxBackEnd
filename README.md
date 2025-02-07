# NaturieuxBackEnd

Backend pour l'application de quiz naturaliste Naturieux, développé en Go.

## 🌿 À propos

NaturieuxBackEnd est une API REST qui gère la logique métier des quiz naturalistes. Elle permet de créer, gérer et passer des quiz sur la faune et la flore.

## 🚀 Technologies

- Go
- PostgreSQL (base de données principale)
- Redis (cache)
- Docker & Docker Compose
- Chi Router

## 🛠 Installation

### Prérequis

- Go 1.21+
- Docker & Docker Compose
- Make (optionnel)

### Installation pour le développement

1. Cloner le repository
```bash
git clone https://github.com/Naturieux-fr/NaturieuxBackEnd.git
```

2. Installer les dépendances
```bash
go mod download
```

3. Copier le fichier d'environnement exemple
```bash
cp .env.example .env
```

4. Lancer les services avec Docker Compose
```bash
docker-compose up -d
```

5. Lancer le serveur de développement
```bash
go run cmd/api/main.go
```

## 📚 Documentation API

La documentation de l'API est disponible via Swagger à l'adresse `/docs/swagger` une fois le serveur lancé.

## 🧪 Tests

Pour lancer les tests :
```bash
go test ./...
```

## 🤝 Contribution

Les contributions sont les bienvenues ! N'hésitez pas à ouvrir une issue ou une pull request.

## 📝 License

[MIT](LICENSE)
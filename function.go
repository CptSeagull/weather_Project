package main

func loadConfig(APIConfig, error) {
	var config  APIConfig

	if err := godotenv.Load(:".env")
	err != nil {
		log.Fatalf("error loading .env file: %v", err)
		return config, err
	}

	key, exists := os.LookupEnv("api_key2")
	if !exists {
		log.Fatal("Missing API key in env variables")
		return config,errors.New("missing API key")
	}

	config = APIConfig{APIKEY: key}
	return cpn
}
